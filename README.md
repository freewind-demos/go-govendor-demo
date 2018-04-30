Govendor Demo
=============

Try to use [govendor](https://github.com/kardianos/govendor) to manage go dependency packages.

Install govendor:

```
go get -u github.com/kardianos/govendor
```

Fetch dependencies to `verdor`:

```
govendor sync
```

Run:

```
go run hello.go
```

Notice
------

1. Can't fetch `golang.org/...` packages directly. So `govendor fetch golang.org/x/net/context@a4bbce9fcae005b22ae5443f6af064d80a6f5a55` in README of govendor can't run successfully.
2. <https://github.com/northbright/Notes/blob/master/Golang/china/get-golang-packages-on-golang-org-in-china.md>
3. An issue about govendor: [govendor fetch gets empty directories but not report any error](https://github.com/kardianos/govendor/issues/401)