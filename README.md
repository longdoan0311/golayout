# Introduction

**golayout** provides a feature to generate a Golang project layout.

This layout is inspired from [Ben Johnson's article](https://medium.com/@benbjohnson/standard-package-layout-7cdbc8391fc1). 

## Project layout

```
+-- api
|   +-- home.go
|   +-- server.go
|   +-- user.go
+-- app
|   +-- <<YOUR_PROJECT_NAME>>
|   |   +-- cmd
+-- service
|   +-- main.go
+-- conf
|   +-- config.yml
+-- store
|   +-- imdbuser.go
|   +-- sqluser.go
|   +-- store.go
+-- user
|   +-- userinfo.go
+-- uuid
|   +-- gen.go
+-- config.go
+-- go.mod
+-- store.go
+-- user.go
```

# How to use it

* Get it

```
go get github.com/techcomsecurities/golayout/app/golayout
```

* Generate your project layout

```
golayout gen -n <<YOUR_PROJECT_NAME>> -m <<YOUR_MODULE_PATH>>
```

* Run your new generated project

```
cd <<YOUR_PROJECT_NAME>>
go run app/service/main.go -c conf/config.yml
```

# Used packages

Many thanks to known packages:

* [github.com/spf13/cobra](https://github.com/spf13/cobra)
* [github.com/sirupsen/logrus](https://github.com/sirupsen/logrus)
* [github.com/gobuffalo/packr/v2](https:github.com/gobuffalo/packr/v2)
* and many standard Golang packages.