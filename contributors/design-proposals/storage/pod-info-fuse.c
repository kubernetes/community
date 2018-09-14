/*
  FUSE flexVolume which provides pod.uid, pod.nodename,
  and pid (in host PID namespace).
  Author: Jan Pazdziora

  gcc -Wall pod-info-fuse.c $( pkg-config fuse json-c --cflags --libs ) \
      -D LOG_FILE=/tmp/pod-info.log \
      -o /usr/libexec/kubernetes/kubelet-plugins/volume/exec/example.com~pod-info-fuse/pod-info-fuse
*/

#define _GNU_SOURCE
#include <stdio.h>
#include <string.h>
#include <errno.h>
#include <fcntl.h>
#include <unistd.h>
#include <sys/wait.h>
#include <sys/utsname.h>
#include <time.h>

#define FUSE_USE_VERSION 29

#include <fuse.h>

#include <json.h>
#include <json_tokener.h>
#include <json_object.h>

#define _QUOTE1(x) #x
#define _QUOTE2(x) _QUOTE1(x)
#ifdef LOG_FILE
#define _LOG_FILE _QUOTE2(LOG_FILE)
#else
#define _LOG_FILE "/dev/null"
#endif

static FILE * log = NULL;

static const char * pod_uid = NULL;
static const char * pod_uid_path = "/pod.uid";
static const char * pod_nodename_path = "/pod.nodename";
static const char * host_pid_path = "/pid";

static int pod_info_readdir(const char * path, void * dir,
		fuse_fill_dir_t filler,
		off_t offset, struct fuse_file_info * fi) {
	if (strcmp(path, "/") != 0) {
		return -ENOENT;
	}
	filler(dir, ".", NULL, 0);
	filler(dir, "..", NULL, 0);
	filler(dir, pod_uid_path + 1, NULL, 0);
	filler(dir, pod_nodename_path + 1, NULL, 0);
	filler(dir, host_pid_path + 1, NULL, 0);
	return 0;
}

static int pod_info_open(const char * path, struct fuse_file_info * fi) {
	if (strcmp(path, pod_uid_path)
		&& strcmp(path, pod_nodename_path)
		&& strcmp(path, host_pid_path)) {
		return -ENOENT;
	}
	if ((fi->flags & 3) != O_RDONLY) {
		return -EACCES;
	}
	return 0;
}

static int pod_info_getattr(const char * path, struct stat * stinfo) {
	memset(stinfo, 0, sizeof(struct stat));
	if (strcmp(path, "/") == 0) {
		stinfo->st_mode = S_IFDIR | 0755;
		stinfo->st_nlink = 2;
	} else if (strcmp(path, pod_uid_path) == 0) {
		stinfo->st_mode = S_IFREG | 0444;
		stinfo->st_nlink = 1;
		if (pod_uid) {
			stinfo->st_size = strlen(pod_uid) + 1;
		}
	} else if (strcmp(path, pod_nodename_path) == 0) {
		stinfo->st_mode = S_IFREG | 0444;
		stinfo->st_nlink = 1;
		struct utsname buf;
		if (uname(&buf) == 0) {
			stinfo->st_size = strlen(buf.nodename) + 1;
		}
	} else if (strcmp(path, host_pid_path) == 0) {
		stinfo->st_mode = S_IFREG | 0444;
		stinfo->st_nlink = 1;
		struct fuse_context * info = fuse_get_context();
		if (info) {
			char buf[3];
			stinfo->st_size = snprintf(buf, 3, "%d\n", info->pid);
		}
	} else {
		return -ENOENT;
	}
	return 0;
}

static int pod_info_read(const char * path, char * buf,
		size_t size, off_t offset, struct fuse_file_info * fi) {
	size_t len;
	char * value = NULL;
	int need_free = 0;
	if (strcmp(path, pod_uid_path) == 0) {
		value = (char *)pod_uid;
	} else if (strcmp(path, pod_nodename_path) == 0) {
		struct utsname buf;
		if (uname(&buf) == 0) {
			value = buf.nodename;
		}
	} else if (strcmp(path, host_pid_path) == 0) {
		struct fuse_context * info = fuse_get_context();
		if (info) {
			if (asprintf(&value, "%d", info->pid) >= 0) {
				need_free = 1;
			}
		}
	}
	if (value) {
		len = strlen(value) + 1;
		if (offset < len) {
			if (offset + size > len)
				size = len - offset;
			memcpy(buf, value + offset, size - 1);
			if (size > strlen(value + offset)) {
				buf[size - 1] = '\n';
			}
			if (need_free) {
				free(value);
			}
			return size;
		}
		if (need_free) {
			free(value);
		}
		return 0;
	}
	return -ENOENT;
}

static struct fuse_operations pod_info_operations = {
	.getattr	= pod_info_getattr,
	.readdir	= pod_info_readdir,
	.open		= pod_info_open,
	.read		= pod_info_read,
};

int not_supported(const char * msg) {
	printf("{ \"status\": \"Not supported\", \"message\": \"%s\" }\n", msg);
	exit(1);
}

int do_mount(int argc, char * argv[]) {
	if (argc != 4) {
		not_supported("mount expects two additional parameters");
	}

	json_object * obj;
	obj = json_tokener_parse(argv[3]);
	if (!obj || !json_object_is_type(obj, json_type_object)) {
		not_supported("second parameter is not JSON object");
	}
	json_object * pod_uid_obj;
	if (!json_object_object_get_ex(obj, "kubernetes.io/pod.uid", &pod_uid_obj)) {
		not_supported("second parameter does not have kubernetes.io/pod.uid");
	}
	if (!json_object_is_type(pod_uid_obj, json_type_string)) {
		not_supported("second parameter does not have kubernetes.io/pod.uid string");
	}
	pod_uid = json_object_get_string(pod_uid_obj);

	char * new_argv[] = {
		argv[0],
		"-o", "allow_other",
		argv[2],
		NULL
	};

	int wstatus;
	int stderr_pipe[2];
	char stderr_buffer[1024];
	pipe(stderr_pipe);
	if (fork()) {
		close(stderr_pipe[1]);
		int count = read(stderr_pipe[0], stderr_buffer, sizeof(stderr_buffer) - 1);
		if (count >= 0) {
			stderr_buffer[count] = '\0';
		}
		close(stderr_pipe[0]);
		wait(&wstatus);
	} else {
		close(STDIN_FILENO);
		close(STDOUT_FILENO);
		close(STDERR_FILENO);
		dup2(stderr_pipe[1], STDERR_FILENO);
		close(stderr_pipe[0]);
		close(stderr_pipe[1]);
		int ret = fuse_main(4, new_argv, &pod_info_operations, NULL);
		exit(ret);
	}

	if (!WIFEXITED(wstatus) || WEXITSTATUS(wstatus)) {
		json_object * obj;
		obj = json_object_new_string(stderr_buffer);
		printf("%s%s%s\n", "{ \"status\": \"Failure\", \"message\": ",
			json_object_to_json_string(obj),
			" }");
		exit(1);
	}
	puts("{ \"status\": \"Success\" }");
	exit(0);
}

int do_umount(int argc, char * argv[]) {
	if (argc != 3) {
		not_supported("umount expects one additional parameter");
	}
	int wstatus;
	if (fork()) {
		wait(&wstatus);
	} else {
		execl("/usr/bin/umount", "-f",  argv[2], (char *) 0);
		exit(0);
	}

	if (!WIFEXITED(wstatus) || WEXITSTATUS(wstatus)) {
		puts("{ \"status\": \"Failure\", \"message\": \"umount failed\" }");
		exit(1);
	}

	unlink(argv[3]);
	puts("{ \"status\": \"Success\" }");
	exit(0);
}

int main(int argc, char * argv[]) {
	log = fopen(_LOG_FILE, "a");
	time_t now = time(NULL);
	char now_str[64];
	strftime(now_str, sizeof(now_str), "%F %T", localtime(&now));
	fputs(now_str, log);
	int i;
	for (i = 1; i < argc; i++) {
		fputs(" ", log);
		fputs(argv[i], log);
	}
	fputs("\n", log);
	fclose(log);

	if (argc < 2) {
		not_supported("at least one parameter expected");
	}

	if (strcmp(argv[1], "init") == 0) {
		puts("{ \"status\": \"Success\", \"capabilities\": {\"attach\": false, \"selinuxRelabel\": false }}");
		exit(0);
	} else if (strcmp(argv[1], "umount") == 0) {
		do_umount(argc, argv);
	} else if (strcmp(argv[1], "mount") == 0) {
		do_mount(argc, argv);
	}
	not_supported("unknown command");
	exit(2);
}
