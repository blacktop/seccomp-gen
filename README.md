# seccomp-gen

[![Circle CI](https://circleci.com/gh/blacktop/seccomp-gen.png?style=shield)](https://circleci.com/gh/blacktop/seccomp-gen) [![Github All Releases](https://img.shields.io/github/downloads/blacktop/seccomp-gen/total.svg)](https://github.com/blacktop/seccomp-gen/releases/latest) [![GitHub release](https://img.shields.io/github/release/blacktop/seccomp-gen.svg)](https://github.com/blacktop/seccomp-gen/releases) [![License](http://img.shields.io/:license-mit-blue.svg)](http://doge.mit-license.org)

> Docker Secure Computing Profile Generator

---

## Why 🤔

This tool allows you to pipe the output of [strace](https://strace.io) through it and will auto-generate a docker seccomp profile that can be used to only whitelist the syscalls your container needs to run and blacklists everything else.

This adds a LOT of security by drastically limiting your attack surface to only what is needed.

## Install

### macOS

```bash
$ brew install blacktop/tap/seccomp-gen
```

### linux/windows

Download from [releases](https://github.com/blacktop/seccomp-gen/releases/latest)

## Getting Started

```bash
$ strace -ff curl github.com 2>&1 | scgen -verbose

   • found syscall: execve
   • found syscall: brk
   • found syscall: access
   • found syscall: access
   • found syscall: openat
   • found syscall: fstat
   • found syscall: mmap
   ...
```

```bash
$ ls -lah

-rw-r--r--   1 blacktop  staff   6.7K Dec  1 21:23 seccomp.json
```

### Inside Docker

Create a new Dockerfile

```dockerfile
FROM <your>/<image>:<tag>
RUN apt-get update && apt-get install -y strace
CMD ["strace","-ff","/your-entrypoint.sh"]
```

Build `scgen` image

```bash
$ docker build -t <your>/<image>:scgen .
```

Generate `seccomp` profile from docker logs output

```bash
docker run --rm --security-opt seccomp=unconfined <your>/<image>:scgen 2>&1 | scgen -verbose
```

Use your :new: `seccomp` profile

```bash
docker run --rm --security-opt seccomp=/path/to/seccomp.json <your>/<image>:<tag>
```

## Credits

- https://blog.jessfraz.com/post/how-to-use-new-docker-seccomp-profiles/
- https://github.com/antitree/syscall2seccomp

## TODO

- [x] filter strace through linux (32|64bit) [tbl](https://github.com/torvalds/linux/blob/master/arch/x86/entry/syscalls/syscall_64.tbl) files like Jess does
- [ ] add support for consuming sysdig output
- [ ] only add current arch to arches

## Issues

Find a bug? Want more features? Find something missing in the documentation? Let me know! Please don't hesitate to [file an issue](https://github.com/blacktop/seccomp-gen/issues/new)

## License

MIT Copyright (c) 2018 **blacktop**
