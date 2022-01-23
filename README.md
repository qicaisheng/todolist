# todolist

## todolist CLI的价值
* 快捷管理 todo list，通过CLI 快捷创建、跟进todo，不用关心当前所在的目录 

## 如何安装
在终端执行如下命令:
```shell
go install github.com/qicaisheng/todolist@latest
```
接下来就可以确认是否安装成功
```shell
todolist -h
```


## 常用命令
```shell
todolist list
todolist add <title>
todolist close <todoId>
todolist modify <todoId>
todolist show <todoId>
```