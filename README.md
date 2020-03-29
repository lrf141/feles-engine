[![GoVersion](https://img.shields.io/badge/Go-v1.12-brightgreen)]()
[![GccVersion](https://img.shields.io/badge/gcc-v7.5.0-brightgreen)]()
[![GoVersion](https://img.shields.io/badge/LICENSE-MIT-blue)]()
# feles-engine
A tiny container engine as a code execution environment


# Debug

## Build Feles Engine

```
$ make
```

## Start Daemon
```
// local mode
$ ./feles-engine
```

## Send Http Request
```
// if daemon launch as local mode
$ curl --unix-socket /tmp/feles.sock localhost
```