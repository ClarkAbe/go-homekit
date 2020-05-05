package accessory

import (
	"ccms/hc/service"
)

type Switch struct {
	*Accessory
	Switch *service.Switch
}

// NewSwitch returns a switch which implements model.Switch.
func NewSwitch(info Info) *Switch {
	acc := Switch{}
	acc.Accessory = New(info, TypeSwitch)
	acc.Switch = service.NewSwitch()
	acc.AddService(acc.Switch.Service)

	return &acc
}
