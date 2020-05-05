 package source

import (
	"os"
	"log"
	"fmt"
	"math"
	"time"
	//"math/rand"
	"runtime"
	"strings"
	"errors"
	"ccms/config"
	//"ccms/bolt"
	"path/filepath"
	"encoding/json"
)

var (
	DebugMode = true
	TimeZone = time.FixedZone("UTC", int((0 * time.Hour).Seconds()))
	timeZone = time.FixedZone("UTC", int((0 * time.Hour).Seconds()))
)

func showGoNum() {
	Info(runtime.NumGoroutine())
}


func FullCPU(x bool){
	if x {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}else{
		runtime.GOMAXPROCS(1)
	}
}

func MkdirAppDir(){
	Mkdir("./config")
	Mkdir("./db")
}

func Exit(o int){
	os.Exit(o)
}

func InitConfig(path string,x interface{})(bool){
	path = AbsPath(path)
	if IsExist(path) == false {
		Info(path," ReWrite")
		if config.WriteConfig(path, x) == false {
			Error(path,"ReWrite Config Error")
			return false
		}
	}
	
	if _, status := config.ReadConfig(path, x); status == false { // 获取的配置信息
		Error("Read Config Error")
		return false
	}
	return true
}
/*
func SaveConfig()(bool){
	if config.WriteConfig(ServerJsonPath, ServerConfig) == false {
		Error(ServerJsonPath, "ReWrite Config Error")
		return false
	}
	if config.WriteConfig(PayJsonPath, PayConfig) == false {
		Error(PayJsonPath, "ReWrite Config Error")
		return false
	}
	return InitConfigs()
}
*/
func AbsPath(path string)(string) {
	if absPath, err := filepath.Abs(path); err == nil {
		return absPath
	}
	return path
}

func Mkdir(path string)(bool){
	if absPath, err := filepath.Abs(path); err == nil { // 多平台目录兼容
		if _, err := os.Stat(absPath); err != nil { // 检查路径是否存在
			if err := os.MkdirAll(absPath, os.ModePerm); err == nil { // 是否创建目录成功
				return true
			}
		}
	}
	return false
}

func IsExist(path string)(bool){
	if absPath, err := filepath.Abs(path); err == nil { // 多平台目录兼容
		if _, err := os.Stat(absPath); err == nil { // 检查路径是否存在
			return true
		}
	}
	return false
}

func LoadTimeZone(x string, i int)(*time.Location){
	TimeZone = time.FixedZone(x, int((time.Duration(i) * time.Hour).Seconds()))
	return timeZone
}


func LogTimeZone(x string, i int)(*time.Location){
	timeZone =time.FixedZone(x, int((time.Duration(i) * time.Hour).Seconds()))
	return timeZone
}

func CallerName() string {
	pc, _, _, _ := runtime.Caller(2)
	return runtime.FuncForPC(pc).Name()
}

func Info(x ...interface{}){
	log.SetPrefix(time.Now().In(timeZone).Format("[Info] 2006-01-02 15:04:05 "))
	log.SetFlags(0)
	var s string
	for _, o := range x {
		s = s + fmt.Sprintf("%v",o)
	}
	log.Println(s)
}

func NewError(x string)(error){
	return errors.New(x)
}

func Error(x ...interface{}){
	log.SetPrefix(time.Now().In(timeZone).Format("[Error] 2006-01-02 15:04:05 "))
	log.SetFlags(0)
	var s string
	for _, o := range x {
		s = s + fmt.Sprintf("%v",o)
	}
	log.Println(s)
}

func Warning(x ...interface{}){
	log.SetPrefix(time.Now().In(timeZone).Format("[Warning] 2006-01-02 15:04:05 "))
	log.SetFlags(0)
	var s string
	for _, o := range x {
		s = s + fmt.Sprintf("%v",o)
	}
	log.Println(s)
}

func Debug(x ...interface{}){
	if DebugMode == true {
		log.SetPrefix(time.Now().In(timeZone).Format("[Debug] 2006-01-02 15:04:05 "))
		log.SetFlags(0)
		var s string
		for _, o := range x {
			s = s + fmt.Sprintf("%v",o)
		}
		log.Println(s)
	}
}

func Map2Json(Map interface{})([]byte, error){
	return json.Marshal(Map)
}

func Json2Map(json_byte []byte,out interface{})(error){
	return json.Unmarshal(json_byte, &out)
}

func JsonToMap(json_byte []byte)(interface{}){
	var out interface{}
	if json.Unmarshal(json_byte, &out) == nil {
		return out
	}
	return nil
}

func Decimal(f float64, n int) float64 {
	n10 := math.Pow10(n)
	return math.Trunc((f+0.5/n10)*n10) / n10
}

func Split(x, v string)([]string){
	return strings.Split(x, v)
}

func DelStrArrValue(arr []string,vals []string)([]string){
	for i := 0; i < len(arr); i++ {
		for o := 0; o < len(vals); o++ {
			if arr[i] == vals[o] {
				arr = append(arr[:i], arr[i+1:]...)
			}
		}
	}
	return arr
}
