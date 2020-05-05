 package source

import (
	"os"
	"strings"
	"strconv"
	"os/signal" 
	"ccms/gee"
)


var (
	ServerConfig = &ServerConfigStruct{"9999", "CST:+8", false, false, true, true}
	HomekitConfig = &HomekitConfigStruct{"./db", "11223", "12344321", Bridge{"Bridge", "wDDWW-dWDWD"}, []Ops{{"ExampleOps", "socket", "1.2.3.4:1234", Ont{"on(%v)", "off(%v)", "ok"}, []Gpios{{"ExampleGpio", "1"}}}}}
)

const (
	ServerJsonPath = "./config/server.json"
	HomekitJsonPath = "./config/homekit.json"
)
//Example
type (
	ServerConfigStruct struct {
		HTTPPort string `json:"http_port"`
		TimeZone string `json:"time_zone"`
		FullCPU bool `json:"full_cpu"`
		Debug bool `json:"debug"`
		WebHook bool `json:"web_hook"`
		Homekit bool `json:"homekit"`
	}
	HomekitConfigStruct struct {
		Storage string `json:"storage"`
		Port string `json:"port"`
		Pin string `json:"pin"`
		Bridge Bridge `json:"bridge"`
		Ops []Ops `json:"ops"`
	}
	Bridge struct {
		Name string `json:"name"`
		Serial string `json:"serial"`
	}
	Ont struct {
		On string `json:"on"`
		Off string `json:"off"`
		Success string `json:"success"`
	}
	Gpios struct {
		Name string `json:"name"`
		Gpio string `json:"gpio"`
	}
	Ops struct {
		Name string `json:"name"`
		Type string `json:"type"`
		Addr string `json:"addr"`
		Ont Ont `json:"ont"`
		Gpios []Gpios `json:"gpios"`
	}
)

func Init(){
	MkdirAppDir()
	if InitConfigs() == false {
		Error("Init Config Error")
		Exit(0)
	}
	if ServerConfig.Homekit == true {
		go HomeKit()
	}
	if ServerConfig.WebHook == true {
		go NewWeb()
	}
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	for _ = range c {
		return
	}
}

func InitConfigs()(bool){
	if InitConfig(ServerJsonPath, ServerConfig) == false {
		return false
	}
	
	if InitConfig(HomekitJsonPath, HomekitConfig) == false {
		return false
	}
	timeSplit := strings.Split(ServerConfig.TimeZone, ":")
	if len(timeSplit) != 2 {
		Error("TimeZone Split Error")
		Exit(0)
	}
	if timeIn, err := strconv.Atoi(timeSplit[1]); err == nil {
		LoadTimeZone(timeSplit[0], timeIn)
		LogTimeZone(timeSplit[0], timeIn)
		gee.LoadTimeZone(timeSplit[0], timeIn)
		Info("Load TimeZone....")
	}
	HomekitLog(ServerConfig.Debug)
	DebugMode = ServerConfig.Debug
	FullCPU(ServerConfig.FullCPU)
	Debug("Server Config : ", ServerConfig)
	return true
}














