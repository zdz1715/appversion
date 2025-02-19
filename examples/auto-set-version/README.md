# 自动设置版本

在执行`go build`或`go install`命令时，将自动检索 VCS（版本控制系统）信息以及包的版本号。请注意，此功能仅在代码处于版本控制之下，并且使用模块的全名时有效。
## 打包
请确保在构建命令中使用你的模块全名。例如，如果你的模块全名是`example.com/myapp/cmd/myapp`，则需要将以下命令中的路径替换为你的模块路径。
```shell
go build -o test-auto-set-app github.com/zdz1715/go-app-version/examples/auto-set-version
```
执行`test-auto-set-app`后的输出信息:
```text
goappversion.Info{Name:"", Major:"0", Minor:"0", Patch:"0", Version:"v0.0.0-20250219071117-3d3b8b40641b", GitCommit:"3d3b8b40641b951dc7a9a61647cb852463867d44", GitTreeState:"dirty", BuildDate:"2025-02-19T07:11:17Z", GoVersion:"go1.23.6", Compiler:"gc", Platform:"darwin/arm64"}
```