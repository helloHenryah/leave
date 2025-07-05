病假管理系统
注意需要配置setting.yaml
默认连接的是postgres数据库
1.安装依赖
```bash
go mod init leave
go mod tidy
```

2.编译成二进制文件
```bash
go build -o ./bin/leave main.go
```

3.设置环境变量
```bash
# 设置数据库类型 目前支持sqlite和postgres以及mysql
export LEAVE_APP_DB_TYPE="sqlite"
# 设置数据库连接 默认使用sqlite和file.db
export LEAVE_APP_DB_DSN="file.db"
# 设置运行端口 默认8080
export LEAVE_APP_PORT="8080"
```

3.运行
```bash
./bin/leave
```