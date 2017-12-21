# Client Tool Release Publishing Guidelines

Projects should publish releases for client side tools.

## Go Projects

### Static Linking

See [Go executables are statically linked, except when they are not](http://matthewkwilliams.com/index.php/2014/09/28/go-executables-are-statically-linked-except-when-they-are-not/). 


- How to compile a statically linked binary: `go` file must be compiled without cgo support. 

```sh
# Disable cgo
export CGO_ENABLED=0
```

- How to check if a binary is statically linked

```sh
# List dynamic dependencies (shared libraries):
# 1. if it's dynamically linked, you'll see
$ ldd <your_tool> 
    linux-vdso.so.1 =>  (0x00007ffe937ea000)
    libpthread.so.0 => /lib/x86_64-linux-gnu/libpthread.so.0 (0x00007f0a7dae5000)
    libc.so.6 => /lib/x86_64-linux-gnu/libc.so.6 (0x00007f0a7d720000)
    /lib64/ld-linux-x86-64.so.2 (0x00007f0a7dd03000)
# 2. if it's statically linked, you'll see
$ ldd <your_tool>  
	not a dynamic executable

# Recognize the type of data in a file
# 1. if it's dynamically linked, you'll see
$ file <your_tool> 
/usr/local/your_tool: ELF 64-bit LSB  executable, x86-64, version 1 (SYSV), dynamically linked (uses shared libs), for GNU/Linux 2.6.32, BuildID[sha1]=86c6d2ff21297a06cc7319244f35e2671612beae, not stripped
# 2. if it's statically linked, you'll see
$ file  <your_tool>  
/usr/local/your_tool: ELF 64-bit LSB  executable, x86-64, version 1 (SYSV), statically linked, not stripped
```

### Targets
Build your release binary for the following targets:

- darwin-amd64
- linux-386
- linux-amd64
- linux-armv6l
- linux-ppc64le
- windows-amd64

### Packaging

Package binaries into a tar.gz file and make available on GitHub releases page.

# Service Side Release Publishing Guidelines

### Packaging

Server side programs should be packaged into container images.  Stateless services should be run as Deployments (as opposed to Replication Controllers).

# Documentation Guidelines

TODO: Write this
