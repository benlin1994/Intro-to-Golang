# Building your project


## Build commands

There are several ways you can build your Go projects:

* go build
* go run
* go install

### go build 

The `go build` command will compile the package that is in the current directory. If it finds a `main` package, it will create an executable in the current directory. If there is more than one package type in the directory, it will error. 

```shell
go build
```

You can specify a specific file or files to build as well:

```shell
go build filename.go
```

### go run

The `go run filename.go` command will compile and run a specified file. The specified file needs a `main` package for it to run. The executable will be deleted after it runs.

```shell
go run filename.go
```

### go install

The `go install` command works the same as the `go build` command, except that it creates the executable in the path specified by the GOBIN go environment variable. It will also create binaries of your imported packages in the `/pkg` folder of your workspace. 

```shell
go install
```

You can specify a specific file or files to build as well:

```shell
go install filename.go
```


## Other tools

In addition to the `go` command. Go comes with two other command line tools:

* godoc
* gofmt

### godoc

The `godoc` command is used to get the documentation of packages. It will print out the documentation in the terminal.

```shell
godoc fmt
```

### gofmt

The `gofmt` command will reformat any `.go` file to have standard spacing and alignment. 

```shell
gofmt filename.go
```

If you specificy a directory, it will format all `.go` files in the directory:

```
gofmt .
```