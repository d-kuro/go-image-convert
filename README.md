# go-image-convert

The Golang cli tool to convert image extension.

## Build

<!-- markdownlint-disable MD014 -->

```bash
$ go build main.go
```

<!-- markdownlint-enable MD014 -->

## Usage

<!-- markdownlint-disable MD014 -->

```bash
$ ./main ~/Desktop/hoge/
```

<!-- markdownlint-enable MD014 -->

## Option

```text
Usage of ./main:
  -from string
        input file extension (support: jpg/png/gif, default: jpg)
  -to string
        output file extension (support: jpg/png/gif, default: png)
```

## GoDoc

<!-- markdownlint-disable MD014 -->

```bash
$ godoc -http=:6060
```

<!-- markdownlint-enable MD014 -->

You can access to read the documentation. See this link:
[http://localhost:6060/pkg/github.com/d-kuro/](http://localhost:6060/pkg/github.com/d-kuro/)
