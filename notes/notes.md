* https://elephantsql.com/

```
go mod init github.com/E-Health/goscar-fhir

```

There are four directives: module, require, exclude, replace.

```
module github.com/my/thing

require (
    github.com/some/dependency v1.2.3
    github.com/another/dependency/v4 v4.0.0
)
```

```
project/
├── go.mod
├── bar
│   └── bar.go
└── foo
    └── foo.go
```

* import "example.com/my/module/v2/pkg/foo" to import foo from the v2 version of module example.com/my/module.



* run go get -u to use the latest minor or patch releases
* run go get -u=patch to use the latest patch releases
