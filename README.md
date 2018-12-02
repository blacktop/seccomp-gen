# seccomp-gen

> Docker Secure Computing Profile Generator

---

## Install

### macOS

```bash
$ brew install blacktop/tap/seccomp-gen
```

### linux/windows

Download from [releases](https://github.com/blacktop/seccomp-gen/releases/latest)

## Getting Started

```bash
$ strace -ff curl github.com | scgen
```

## Credits

- https://blog.jessfraz.com/post/how-to-use-new-docker-seccomp-profiles/
- https://github.com/antitree/syscall2seccomp

## TODO

- [ ] add support for consuming sysdig output

## Issues

Find a bug? Want more features? Find something missing in the documentation? Let me know! Please don't hesitate to [file an issue](https://github.com/blacktop/seccomp-gen/issues/new)

## License

MIT Copyright (c) 2018 **blacktop**
