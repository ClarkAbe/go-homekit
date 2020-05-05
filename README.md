# Golang Homekit

#### Run / 运行

```
go build -o homekit.bin main.go
chmod +x ./homekit.bin
#auto mkdir all dir
#自动创建所有文件夹
```


#### Config / 配置文件

```
{
    "http_port": "9999", // Web Hook Port / Web Hook 端口
    "time_zone": "CST:+8", // Time Zone / 时区
    "full_cpu": false, // Multi CPU Core / 多核CPU
    "debug": false, // Debug Mode / 调试模式
    "web_hook": true, // Web Hook Switch / Web Hook 开关
    "homekit": true // Homekit Switch / Homekit 开关
}



```