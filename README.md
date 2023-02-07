# todo
学习如何使用ent库

## 
初始化
```console
go mod init github.com/hewenyu/todo
```

## 表设计生成
```bash
# 表原始model生成
go run -mod=mod entgo.io/ent/cmd/ent init User
# temple 模板表,里面更多的模板用于开发
go run -mod=mod entgo.io/ent/cmd/ent init Temple
```

## 公共参数设计

```bash
# 在ent目录下创建一个公共参数专用文件夹
cd ent && mkdir mixins && cd mixins
```
