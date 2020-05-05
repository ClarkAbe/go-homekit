package config
import(
	"os"
	"bytes"
	"io/ioutil"
	"encoding/json"
	"path/filepath"
)
//读取过后会自动格式化然后重写一次
func ReadConfig(path string,out interface{})(interface{},bool){
	//var out tp
	if path, err := filepath.Abs(path); err == nil {
		if data, err := ioutil.ReadFile(path); err == nil {
			if json.Unmarshal(data, &out) == nil {
				WriteConfig(path,out)
				return out,true
			}
		}
	}
	return out,false
}

func WriteConfig(path string,data interface{})(bool){
	if path, err := filepath.Abs(path); err == nil {
		if byte_json,err := json.Marshal(data);err == nil {
			var str bytes.Buffer
			if json.Indent(&str, byte_json, "", "    ") == nil {
				if ioutil.WriteFile(path, str.Bytes(), os.ModePerm) == nil {
					return true
				}
			}
		}
	}
	return false
}

