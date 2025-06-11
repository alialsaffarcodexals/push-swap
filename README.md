# push-swap

This repository contains a Go implementation of the classic **push-swap**
exercise. Two executables are provided:

- **push-swap** – generates a list of stack instructions that sorts its
  numeric arguments.
- **checker** – validates a list of instructions by executing them on the
  given stack and reports whether it is sorted.

The available instructions are:

```
pa pb sa sb ss ra rb rr rra rrb rrr
```

See the project specification for a detailed description of these
operations.

## Building

The project uses only the Go standard library. To build the executables
run:

```
go build ./cmd/pushswap
go build ./cmd/checker
```

This will produce `pushswap` and `checker` binaries in the current
folder.  Rename `pushswap` to `push-swap` if desired.

## Usage

To generate instructions that sort a list of numbers:

```
$ ./pushswap "2 1 3 6 5"
```

To verify a sequence of instructions:

```
$ echo -e "pb\nra\npa" | ./checker "2 1"
```

`checker` prints `OK` if the given instructions actually sort the stack
and leave the auxiliary stack empty, otherwise `KO` is printed.

## Sample numbers file

A sample numbers file `numbers.txt` is provided. It contains:

```
4 67 3 87 23
```

You can try sorting it with:

```
$ ARG="$(cat numbers.txt)"
$ ./pushswap "$ARG" | ./checker "$ARG"
```

Both executables return nothing when called without arguments.

## Tests

Unit tests can be run with:

```
go test ./...
```

