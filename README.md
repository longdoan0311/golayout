# Introduction
**golayout** provides a feature to generate a Golang project layout.

This layout is inspired from [Ben Johnson's article](https://medium.com/@benbjohnson/standard-package-layout-7cdbc8391fc1). 

# How to use it
* Get it
```
go get github.com/3t-dev/golayout/app/golayout
```
* Generate your project layout
```
golayout gen -n <<YOUR_PROJECT_NAME>>
```

# Used packages
Many thanks to known packages:
* [github.com/spf13/cobra](https://github.com/spf13/cobra)
* [github.com/sirupsen/logrus](https://github.com/sirupsen/logrus)
* [github.com/gobuffalo/packr/v2](https:github.com/gobuffalo/packr/v2)
* and many standard Golang packages.