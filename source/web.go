package source

import (
	"strconv"
	"ccms/gee"
	//"ccms/hc/accessory"
)

func NewWeb(){
	r := gee.NewUse(DebugMode, true)
	
	OpsRouter(r)
	
	Info("WebHook Runing Port:",ServerConfig.HTTPPort)
	
	r.RunSrv(":"+ServerConfig.HTTPPort)
}

func OpsRouter(r *gee.Engine){
	r.RESTFUL("/ops/:ops_id/:gpio_id/:status", func (c *gee.Context){
		ops_id, ops_id_err := strconv.Atoi(c.Param("ops_id"))
		gpio_id, gpio_id_err := strconv.Atoi(c.Param("gpio_id"))
		if ops_id_err == nil && gpio_id_err == nil  {
			if Outlet, ok := Outlets[int(NewUUID(ops_id, gpio_id))]; ok == true {
				on := (c.Param("status") == "true")
				Outlet.Outlet.On.SetValue(on)
				if OpGpios(ops_id, gpio_id, on) == false {
					c.JSON(200, gee.H{
						"message":"op gpio fail",
						"status":false,
					})
					Outlet.Outlet.On.SetValue(!on)
					return
				}
				c.JSON(200, gee.H{
					"message":"success",
					"status":true,
				})
				return 
			}
			c.JSON(200, gee.H{
				"message":"not the op",
				"status":false,
			})
			return
		}
		Debug(":(", ops_id_err, gpio_id_err)
		c.JSON(200, gee.H{
			"message":"gpio_id er ops_id parse int error!",
			"status":false,
		})
		return 
	})
}
