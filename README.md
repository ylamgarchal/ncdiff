# Ncdiff

Ncdiff is a Golang implementation of the Normalized Compression Distance (NCD). The NCD indicates how much two files are similar by giving a distance between them (0 closest, 100 farther).

The major advantage of the NCD function is that it is agnostic to the files format and give a good first similarity value.

![NCD](https://wikimedia.org/api/rest_v1/media/math/render/svg/e00cc1837ede7799f7dbdb924d88330c4cbfa899)


More information on NCD [here](https://en.wikipedia.org/wiki/Normalized_compression_distance).


## Usage
```bash
$ git clone https://github.com/ylamgarchal/ncdiff.git
$ cd ncdiff
$ go build
$ ./ncdiff -h
Usage: ncdiff file-x file-y
```

## Installation

Use `go install` to install the binary in your GOPATH.

Note: Be sure that your `GOPATH/bin` directory is in your `PATH` environment variable to be able to call ncdiff like any other binary tool.

```bash
$ cd ncdiff
$ go install
$ ncdiff -h
Usage: ncdiff file-x file-y
```

## Example

The examples directory contains four text files `examples/tests{1,2,3,4}.txt`. The NCD result is a value between 0 and 100.

```bash
$ ncdiff examples/test1.txt examples/test1.txt 
8.62
```

We can see that `test1.txt` has a distance of `8.62` from itself, it's not `0` because of the compression algorithm used. Depending of the file format some compression works better than others. 

Let's compute the NCD between `test1.txt` and `test2.txt`:

```bash
$ ncdiff examples/test1.txt examples/test2.txt 
20.31
```

We can see that `test1.txt` has a bigger distance from `test2.txt` than itself. This is because `test2.txt` is based on `test1.txt` plus some differences added.

Let's continue with `test3.txt` and `test4.txt` on which i added differences:

```bash
$ ncdiff examples/test1.txt examples/test3.txt 
28.57
$ ncdiff examples/test1.txt examples/test4.txt 
37.68
```

Given the content of the files we can see that the NCD results are consistent.