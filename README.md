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
go run -mod=mod entgo.io/ent/cmd/ent init User Temple
```

## 公共参数设计

```bash
# 在ent目录下创建一个公共参数专用文件夹
cd ent && mkdir mixins && cd mixins
```

## 创建模板

```bash
# 在ent目录下创建一个模板专用文件夹
cd ent && mkdir template && cd template
# 这个模板主要用于分页查询处理，如果有其他需要，请自定义模板
touch  mutation_input.tmpl
```
