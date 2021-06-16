# Introduction

This is a simple demo repository to test out Azure Functions Custom Handler with Golang.  

## Run
```
 cd src
 go build server.go
 func start
    Azure Functions Core Tools
    Core Tools Version:       3.0.3568 Commit hash: e30a0ede85fd498199c28ad699ab2548593f759b  (64-bit)
    Function Runtime Version: 3.0.15828.0

    [2021-06-16T15:08:31.711Z] Worker process started and initialized.

Functions:
	os: [GET] http://localhost:7071/api/os
```

## Test
```
    curl http://localhost:7071/api/os
    {"Time":"Wednesday, 16-Jun-21 10:08:56 CDT","Host":"test.local","OSType":"darwin","Version":"v2"}
```