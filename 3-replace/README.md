# replace

This folder is for the `replace` cmd and how to develop sub modules

```
$ git clone https://github.com/russross/blackfriday ../blackfriday

[change dir blackfriday]
$ git checkout v2

[change dir 3-replace]
$ go mod edit -replace "github.com/russross/blackfriday/v2=../blackfriday"

```

Alternatively use `gohack` - https://github.com/rogpeppe/gohack

