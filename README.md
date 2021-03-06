Goadd: Add Folder To Gopath™
==========================
![image](https://farm5.staticflickr.com/4317/35198386374_1939af3de6_k_d.jpg)
## Requirements

- Go development environment: >= **go1.2**

### Install from source code

    go get -u github.com/apple-han/goadd


The executable will be produced under `$GOPATH/bin` in your file system; for global use purpose, we recommend you to add this path into your `PATH` environment variable.

## Features

- Add folder to gopath.
- Easy to find files under gopath 
```
$ goadd
NAME:
   Goadd - Add Folder To GoPath

USAGE:
   Goadd [arguments...]

VERSION:
   0.0.1 Beta

COMMANDS:
   run          add a new folder, and add to gopath
   list         list all files of gopath
   help, h      Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h           show help
   --version, -v        print the version
```
## License

This project is under the Apache License, Version 2.0. See the [LICENSE](LICENSE) file for the full license text.
