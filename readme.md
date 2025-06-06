病假管理系统
注意需要配置setting.yaml
默认连接的是postgres数据库
1.安装依赖
```bash
go mod tidy
```

2.编译成二进制文件
```bash
go build -o ./bin/leave main.go
```

3.运行
```bash
./bin/leave
```