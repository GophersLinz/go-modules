# dep

This folder describes the upgrade process from `dep` to `modules`

```
$ go mod init github.com/GophersLinz/go-modules
  - initialize the module

$ go mod tidy
  - prune old dependencies and add any dependencies needed for
    other combinations of OS, architecture, and build tags

$ rm -rf Gopkg.* vendor/
  - remove dep things
```