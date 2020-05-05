package source

import (
	"fmt"
	"net"
	"time"
	"net/http"
	"io/ioutil"
	"strconv"
	"ccms/hc"
	"ccms/hc/accessory"
)

var (
	Outlets = make(map[int]*accessory.Outlet)
	Accessorys []*accessory.Accessory
)

func SendSocket(ops_id int, sdata []byte)(bool) {
	op := HomekitConfig.Ops[ops_id]
	conn, err := net.DialTimeout("tcp", op.Addr, 8 * time.Second)
	if err != nil {
		return false
	}
	defer conn.Close()
	conn.SetWriteDeadline(time.Now().Add(6 * time.Second))
	conn.Write(sdata)
	conn.SetReadDeadline(time.Now().Add(6 * time.Second))
	buffer := make([]byte, 2)
	_, err = conn.Read(buffer)
	if err != nil {
		return false
	}
	//Info(string(buffer))
	return (string(buffer) == op.Ont.Success)
}

func SendHttp(ops_id int, sdata string)(bool){
	op := HomekitConfig.Ops[ops_id]
	client := http.Client{
		Timeout: 3 * time.Second,
	}
	resp, err := client.Get(op.Addr + sdata)
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false
	}
	return (string(body) == op.Ont.Success)
}

func OpGpios(ops_id int, gpios_id int, on bool)(bool) {
	op := HomekitConfig.Ops[ops_id]
	of := op.Ont.Off
	if on == true {
		of = op.Ont.On
	}
	send_string := fmt.Sprintf(of, op.Gpios[gpios_id].Gpio)
	switch (op.Type) {
		case "socket":
			return SendSocket(ops_id, []byte(send_string))
		case "http":
			return SendHttp(ops_id, send_string)
	}
	return false
}

func NewUUID(id int, gpios_id int)(int64) {
	if uuid, err := strconv.ParseInt(fmt.Sprintf("%v%v", (id + 1), (gpios_id + 1)), 10, 64); err == nil {
		return uuid
	}
	Warning("UUID ParseInt not work")
	return int64(id + 21 + gpios_id)
}

func NewOpGpio(ops_id int, gpios_id int)(*accessory.Outlet){
	gpio := HomekitConfig.Ops[ops_id].Gpios[gpios_id]
	Outlet := accessory.NewOutlet(accessory.Info{
		Name: gpio.Name,
		SerialNumber: fmt.Sprintf("051AC-OnGPIO%v", gpio.Gpio),
		Manufacturer: "ClarkQAQ",
		Model: "HC-SBX45",
		ID:uint64(NewUUID(ops_id, gpios_id)),
	})
	Outlet.Outlet.On.SetValue(false)
	Outlet.Outlet.On.OnValueRemoteUpdate(func(on bool) {
		if on == true {
			Outlet.Outlet.On.SetValue(true)
			if OpGpios(ops_id, gpios_id, true) == false {
				Outlet.Outlet.On.SetValue(false)
			}
		} else {
			Outlet.Outlet.On.SetValue(false)
			if OpGpios(ops_id, gpios_id, false) == false {
				Outlet.Outlet.On.SetValue(true)
			}
		}
	})
	return Outlet
}

func HomekitLog(x bool){
	if x {
		hc.EnableLog()
	}else{
		hc.DisableLog()
	}
}

func HomeKit() {
	bridge := accessory.NewBridge(accessory.Info{
		Name: HomekitConfig.Bridge.Name,
		SerialNumber: HomekitConfig.Bridge.Serial,
		Manufacturer: "ClarkQAQ",
		Model: "HC-SB250",
		ID: 1,
	})
	
	for i := 0; i < len(HomekitConfig.Ops); i++ {
		Op := HomekitConfig.Ops[i]
		for o := 0; o < len(Op.Gpios); o++ {
			Outlet := NewOpGpio(i, o)
			Outlets[int(NewUUID(i, o))] = Outlet
			Accessorys = append(Accessorys, Outlet.Accessory)
		}
	}
	
	t, err := hc.NewIPTransportArray(hc.Config{Pin: HomekitConfig.Pin, Port: HomekitConfig.Port, StoragePath: HomekitConfig.Storage}, bridge.Accessory, Accessorys)
	if err != nil {
		Error(err)
		return
	}
	
	hc.OnTermination(func() {
		<-t.Stop()
	})
	Info("HomeKit Runing Port:", HomekitConfig.Port, " Pin:", HomekitConfig.Pin)
	t.Start()
	return
}




