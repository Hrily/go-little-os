# Go Little OS

This is the go based version of [The little book about OS
development](http://littleosbook.github.io/).

## Requirements

The following are required for building this project:

+ `gccgo`: For compiling go files (without go runtime)
+ `nasm`: For compiling assembly based loader

## Usage

### Build

You can build the image using:

```
make all
```

In case you want to use docker for build (if you are not on linux)

```
make build-docker
```

### Run

Run the image using bochs
```
make run-bochs
```
