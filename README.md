# go-cat

Simple implementation of `cat`.

## Install

```sh
go install .
```

## Usage

- argument

```sh
$ go-cat -i /proc/version
Linux version 5.13.19-2-MANJARO (builduser@fv-az39-393) (gcc (GCC) 11.1.0, GNU ld (GNU Binutils) 2.36.1)
```

- stdin

```sh
$ cat /proc/version | go-cat
Linux version 5.13.19-2-MANJARO (builduser@fv-az39-393) (gcc (GCC) 11.1.0, GNU ld (GNU Binutils) 2.36.1)
```

```sh
$ go-cat
a
abc
[Ctrl-d]
a
abc
```

  - filename as `-` alias of stdin

  ```sh
  cat /proc/version | go-cat -i -
  Linux version 5.13.19-2-MANJARO (builduser@fv-az39-393) (gcc (GCC) 11.1.0, GNU ld (GNU Binutils) 2.36.1)
  ```
