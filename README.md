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
server.js
{
    "http_port": "9999", // webhook port / Web Hook 端口
    "time_zone": "CST:+8", // time zone / 时区
    "full_cpu": false, // Multi CPU Core / 多核CPU
    "debug": false, // debug mode / 调试模式
    "web_hook": true, // web hook switch / Web Hook 开关
    "homekit": true // homekit switch / Homekit 开关
}

homekit.js
{
    "storage": "./db", // hc database / HC 数据文件
    "port": "11223", // homekit port / homekit端口
    "pin": "12344321", // homekit pin / homekit pin 密码
    "bridge": {
        "name": "Bridge", // btidge name / 中继节点名字
        "serial": "wDDWW-dWDWD" // btidge serial / 节点序列号,可瞎填
    },
    "ops": [ // switch group, now only switch group / 开关组,目前只有开关....
        {
            "name": "ExampleOps", // group name / 组名
            "type": "socket", // group net type (socket or http) / 通讯方式
            "addr": "1.2.3.4:1234", // net addr / 地址
            "ont": { 
                "on": "on(%v)", // "on" send string, %v will be replaced gpio num / "开" 发送的内容,%v会被替换为Gpio号
                "off": "off(%v)", // "off" send string, %v will be replaced gpio num / "关" 发送的内容,%v会被替换为Gpio号
                "success": "ok" // success callback string / 成功返回字符
            },
            "gpios": [
                {
                    "name": "ExampleGpio", // gpio name (homekit default name) / gpio 名字(homekit 默认名字)
                    "gpio": "1" //gpio num / gpio 号
                }
            ]
        }
    ]
}

```