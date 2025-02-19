# go-app-version
快速设置Go应用版本信息，如果你需要编写多个应用的版本设置代码，你可能需要这个。

## 安装

```shell
go get github.com/zdz1715/go-app-version
```

## 使用方式
- [在项目中直接设置](./examples/set-version/main.go)
- [打包时注入版本](./examples/build/README.md)

## 版本信息字段

| 字段           | 说明                           | 
|:-------------|:-----------------------------|
| Name         | 程序名称                         |
| Major        | 主要版本号                        |
| Minor        | 次要版本号                        |
| Patch        | 修订号                          |
| Version      | 版本号                          |
| GitCommit    | Git 提交hash                   |
| GitTreeState | Git 提交状态: 'clean' or 'dirty' |
| BuildDate    | 构建时间                         |
| GoVersion    | go 版本                        |
| Compiler     | 编译器名称                        |
| Platform     | 系统架构，format: os/arch         |
