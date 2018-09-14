/*
  FUSE flexVolume which calls helper program

    <helper> get <the-path> <the-mount-parameters>

  and presents its standard output as file contents.

  Author: Jan Pazdziora

  For a flexVolume plugin example.com/custodia-cli-fuse, compile with

  gcc -Wall flexvolume-fuse-external.c $( pkg-config fuse json-c --cflags --libs ) \
      -D LOG_FILE=/tmp/custodia-cli.log \
      -o /usr/libexec/kubernetes/kubelet-plugins/volume/exec/example.com~custodia-cli-fuse/custodia-cli-fuse

  or place the output to appropriate location and symlink it from

  /usr/libexec/kubernetes/kubelet-plugins/volume/exec/example.com~custodia-cli-fuse/custodia-cli-fuse

  The matching helper (or a symlink to it) then has to be in "next door" location

  /usr/libexec/kubernetes/kubelet-plugins/volume/libexec/example.com~custodia-cli-fuse/custodia-cli-fuse
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

#ifdef HELPER_PREFIX
#define _HELPER_PREFIX _QUOTE2(HELPER_PREFIX)
#else
#define _HELPER_PREFIX "/usr/libexec/kubernetes/kubelet-plugins/volume/libexec"
#endif

#ifndef DEBUG
#define DEBUG 0
#endif

#define debug(x) if (DEBUG) { x }

static char * external_helper_path = NULL;
static char ** mount_params = NULL;
static char ** get_params = NULL;
static char ** enable_dirs = (char *[]){ "/", NULL };

static FILE * log = NULL;

static time_t mount_time = 0;

char * file_path = NULL;
long file_length;
void * file_content = NULL;
time_t file_mtime;

#define READ_CHUNK_SIZE 2048
static int external_helper_get_stdout(char * const * params, const char * debug_caller, void ** data) {
	if (external_helper_path == NULL) {
		return -1;
	}
	int wstatus;
	int stdout_pipe[2];
	ssize_t stdout_length = 0;
	char * stdout_buffer = malloc(READ_CHUNK_SIZE + 1);
	if (stdout_buffer == NULL) {
		return -1;
	}

	fflush(log);
	pipe(stdout_pipe);
	if (fork()) {
		debug( fprintf(log, "%s/external_helper_get_stdout forked for %s\n", debug_caller, params[0]); fflush(log); )
		close(stdout_pipe[1]);
		ssize_t chunk_length;
		while ((chunk_length = read(stdout_pipe[0], stdout_buffer + stdout_length, READ_CHUNK_SIZE)) > 0) {
			debug( fprintf(log, "%s/external_helper_get_stdout read %ld\n", debug_caller, chunk_length); fflush(log); )
			stdout_length += chunk_length;
			char * new_stdout_buffer = realloc(stdout_buffer, stdout_length + READ_CHUNK_SIZE + 1);
			if (new_stdout_buffer == NULL) {
				free(stdout_buffer);
				stdout_buffer = NULL;
				break;
			}
			stdout_buffer = new_stdout_buffer;
		}
		debug( fprintf(log, "%s/external_helper_get_stdout length %ld\n", debug_caller, stdout_length); fflush(log); )
		close(stdout_pipe[0]);
		wait(&wstatus);
		if (!WIFEXITED(wstatus) || WEXITSTATUS(wstatus)) {
			debug( fprintf(log, "%s/external_helper_get_stdout wstatus indicates error\n", debug_caller); fflush(log); )
			if (stdout_buffer) {
				free(stdout_buffer);
			}
			return -1;
		}
		stdout_buffer[stdout_length] = '\0';
		*data = stdout_buffer;
		return stdout_length;
	} else {
		debug ( fprintf(log, "%sexternal_helper_get_stdout in child for %s\n", debug_caller, params[0]); fflush(log); )
		close(STDIN_FILENO);
		close(STDOUT_FILENO);
		if (DEBUG) {
			dup2(fileno(log), STDERR_FILENO);
		} else {
			close(STDERR_FILENO);
		}
		dup2(stdout_pipe[1], STDOUT_FILENO);
		close(stdout_pipe[0]);
		close(stdout_pipe[1]);

		debug( fprintf(log, "%s/external_helper_get_stdout before exec of %s\n", debug_caller, params[0]); fflush(log); )

		execv(external_helper_path, params);
		debug( fprintf(log, "%s/external_helper_get_stdout execv of %s failed\n", debug_caller, params[0]); fflush(log); )
		exit(127);
	}
}

static int custodia_cli_readdir(const char * path, void * dir,
		fuse_fill_dir_t filler,
		off_t offset, struct fuse_file_info * fi) {
	debug( fprintf(log, "custodia_cli_readdir(%s)\n", path); fflush(log); )
	if (strcmp(path, "/") != 0) {
		/* To avoid the flexVolume logic from being confused,
		   we need to show root (the mount point existing).
		   For the other dirs we just return -ENOENT, it seems to work. */
		return -ENOENT;
	}
	filler(dir, ".", NULL, 0);
	filler(dir, "..", NULL, 0);
	return 0;
}

static int custodia_cli_get_file_cached(const char * path) {
	debug( fprintf(log, "custodia_cli_get_file_cached(%s)\n", path); fflush(log); )
	if (path == NULL || path[0] != '/') {
		return -ENOENT;
	}
	/* Invalidate the last cache file if it was a different one. */
	if (file_path != NULL && strcmp(path, file_path) != 0) {
		free(file_content);
		file_content = NULL;
		free(file_path);
		file_path = NULL;
	}
	if (file_path) {
		/* We reuse the last-used file information. */
		return 0;
	}

	get_params[0] = external_helper_path;
	get_params[1] = "get";
	get_params[2] = (char *)path;
	file_length = external_helper_get_stdout(get_params, "custodia_cli_get_file_cached", &file_content);
	if (file_length < 0) {
		return -ENOENT;
	}
	file_mtime = time(NULL);
	file_path = strdup(path);
	return 0;
}

static int custodia_cli_get_dir_file(const char * path, struct stat * stinfo) {
	memset(stinfo, 0, sizeof(struct stat));
	if (file_path && strcmp(path, file_path) == 0 && file_content) {
		goto stat_file;
	}
	if (enable_dirs == NULL) {
		return -ENOENT;
	}
	size_t path_len = strlen(path);
	int i;
	for (i = 0; enable_dirs[i]; i++) {
		size_t dir_len = strlen(enable_dirs[i]);
		if (path_len > dir_len) {
			if (strncmp(path, enable_dirs[i], dir_len) == 0
				&& (path[dir_len] == '/'
					|| (dir_len > 0 && path[dir_len - 1] == '/'))
				&& strchr(path + dir_len + 1, '/') == NULL) {
				if (custodia_cli_get_file_cached(path) == 0) {
					goto stat_file;
				}
				return -ENOENT;
			}
		} else if (strncmp(path, enable_dirs[i], path_len) == 0) {
			if (path_len == dir_len
				|| enable_dirs[i][path_len] == '/'
				|| (path_len > 0 && path[path_len - 1] == '/')) {
				stinfo->st_mode = S_IFDIR | 0755;
				stinfo->st_nlink = 2;
				stinfo->st_mtime = mount_time;
				return 0;
			}
		}
	}
	return -ENOENT;
stat_file:
	stinfo->st_mode = S_IFREG | 0444;
	stinfo->st_nlink = 1;
	stinfo->st_size = file_length;
	stinfo->st_mtime = file_mtime;
	return 0;
}

static int custodia_cli_getattr(const char * path, struct stat * stinfo) {
	debug( fprintf(log, "custodia_cli_getattr(%s)\n", path); fflush(log); )
	return custodia_cli_get_dir_file(path, stinfo);
}

static int custodia_cli_open(const char * path, struct fuse_file_info * fi) {
	debug( fprintf(log, "custodia_cli_open(%s)\n", path); fflush(log); )
	if ((fi->flags & 3) != O_RDONLY) {
		return -EACCES;
	}
	struct stat stinfo;
	if (custodia_cli_get_dir_file(path, &stinfo) == 0
		&& stinfo.st_mode & S_IFREG
		&& file_content) {
		return 0;
	}
	return -ENOENT;
}

static int custodia_cli_read(const char * path, char * buf,
		size_t size, off_t offset, struct fuse_file_info * fi) {
	debug( fprintf(log, "custodia_cli_read(%s)\n", path); fflush(log); )
	struct stat stinfo;
	if (custodia_cli_get_dir_file(path, &stinfo) == 0
		&& stinfo.st_mode & S_IFREG
		&& file_content) {
		if (offset < file_length) {
			if (offset + size > file_length)
				size = file_length - offset;
			memcpy(buf, file_content + offset, size);
			return size;
		}
		return 0;
	}
	return -ENOENT;
}

static struct fuse_operations custodia_cli_operations = {
	.getattr	= custodia_cli_getattr,
	.readdir	= custodia_cli_readdir,
	.open		= custodia_cli_open,
	.read		= custodia_cli_read,
};

static int exit_not_supported(const char * msg) {
	printf("{ \"status\": \"Not supported\", \"message\": \"%s\" }\n", msg);
	exit(1);
}

static void set_external_helper_path(char * argv0) {
	char * double_basename = NULL;
	if (argv0) {
		double_basename = strchrnul(argv0, '\0');
		int found = 0;
		while (double_basename > argv0) {
			if (*double_basename == '/') {
				found++;
				if (found == 2) {
					break;
				}
			}
			double_basename--;
		}
		if (found != 2) {
			double_basename = NULL;
		}
	}
	if (double_basename == NULL) {
		char * error_message = "could not determine helper name from argv[0]";
		asprintf(&error_message, "could not determine helper name from %s", argv0);
		exit_not_supported(error_message);
	}
	asprintf(&external_helper_path, "%s%s", _HELPER_PREFIX, double_basename);
	if (external_helper_path == NULL) {
		exit_not_supported("could not allocate space for helper name");
	}
	if (access(external_helper_path, F_OK)) {
		char * error_message = "helper does not seem to exist";
		asprintf(&error_message, "helper %s does not seem to exist", external_helper_path);
		exit_not_supported(error_message);
	}
}

static char ** get_array_of_strings(json_object * json_obj, char * key, int * count_ptr) {
	json_object * json_value_array_obj;
	char ** ret = NULL;
	int param_count = 0;
	if (json_object_object_get_ex(json_obj, key, &json_value_array_obj)) {
		char * error_message = "value is not an array of strings";
		if (!json_object_is_type(json_value_array_obj, json_type_array)) {
			asprintf(&error_message, "value of %s is not an array", key);
			exit_not_supported(error_message);
		}
		param_count = json_object_array_length(json_value_array_obj);
		ret = calloc(param_count + 1, sizeof(char *));
		int i;
		for (i = 0; i < param_count; i++) {
			json_object * json_value_obj = json_object_array_get_idx(json_value_array_obj, i);
			if (!json_object_is_type(json_value_obj, json_type_string)) {
				asprintf(&error_message, "value of %s[%d] is not an array", key, i);
				exit_not_supported(error_message);
			}
			ret[i] = strdup(json_object_get_string(json_value_obj));
		}
	}
	if (count_ptr) {
		*count_ptr = param_count;
	}
	return ret;
}

static int do_mount(const int argc, char * argv[]) {
	if (argc != 4) {
		exit_not_supported("mount expects two additional parameters");
	}

	set_external_helper_path(argv[0]);

	void * mount_data = NULL;
	char ** params = (char *[]){ external_helper_path, "mount", NULL };
	int length = external_helper_get_stdout(params, "do_mount", &mount_data);
	if (length < 0) {
		char * error_message = "error running helper";
		asprintf(&error_message, "error running helper %s", external_helper_path);
		exit_not_supported(error_message);
	}

	json_tokener * json_tok = json_tokener_new();
	json_object * json_helper_obj = json_tokener_parse_ex(json_tok, mount_data, length);
	json_tokener_free(json_tok);
	if (json_helper_obj == NULL || !json_object_is_type(json_helper_obj, json_type_object)) {
		if (length == 0) {
			debug( fprintf(log, "zero output from helper %s", external_helper_path); )
		} else {
			debug( fwrite(mount_data, length, 1, log); )
		}
		exit_not_supported("unexpected mount response from the helper");
	}

	int mount_params_count;
	mount_params = get_array_of_strings(json_helper_obj, "mount-param", &mount_params_count);
	json_object * json_mount_obj = json_tokener_parse(argv[3]);
	if (!json_mount_obj || !json_object_is_type(json_mount_obj, json_type_object)) {
		exit_not_supported("second parameter is not JSON object");
	}
	get_params = calloc(3 + mount_params_count + 1, sizeof(char *));
	if (mount_params) {
		int i = 0;
		for (i = 0; mount_params[i]; i++) {
			json_object * json_value_obj;
			if (!json_object_object_get_ex(json_mount_obj, mount_params[i], &json_value_obj)) {
				char * error_message = "mount parameters to not match mount-param requested by helper";
				asprintf(&error_message, "mount parameter does not contain %s requested by helper", mount_params[i]);
				exit_not_supported(error_message);
			}
			if (!json_object_is_type(json_value_obj, json_type_string)) {
				char * error_message = "mount parameters are not string";
				asprintf(&error_message, "mount parameter %s is not string", mount_params[i]);
				exit_not_supported(error_message);
			}
			get_params[i + 3] = strdup(json_object_get_string(json_value_obj));
		}
	}

	enable_dirs = get_array_of_strings(json_helper_obj, "enable-dirs", NULL);

	mount_time = time(NULL);

	char * new_argv[] = {
		argv[0],
		"-o", "allow_other",
		argv[2],
		NULL
	};

	fflush(log);
	int wstatus;
	int stderr_pipe[2];
	char stderr_buffer[512];
	pipe(stderr_pipe);
	if (fork()) {
		close(stderr_pipe[1]);
		/* We only read a piece of stderr. */
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
		int ret = fuse_main(4, new_argv, &custodia_cli_operations, NULL);
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

static int do_umount(int argc, char * argv[]) {
	if (argc != 3) {
		exit_not_supported("umount expects one additional parameter");
	}
	fflush(log);
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
	fflush(log);

	if (argc < 2) {
		exit_not_supported("at least one parameter expected");
	}

	if (strcmp(argv[1], "init") == 0) {
		set_external_helper_path(argv[0]);
		puts("{ \"status\": \"Success\", \"capabilities\": {\"attach\": false, \"selinuxRelabel\": false }}");
		exit(0);
	} else if (strcmp(argv[1], "umount") == 0) {
		do_umount(argc, argv);
	} else if (strcmp(argv[1], "mount") == 0) {
		do_mount(argc, argv);
	}
	exit_not_supported("unknown command");
	exit(2);
}
