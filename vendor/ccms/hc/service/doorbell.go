// THIS FILE IS AUTO-GENERATED
package service

import (
	"ccms/hc/characteristic"
)

const TypeDoorbell = "121"

type Doorbell struct {
	*Service

	ProgrammableSwitchEvent *characteristic.ProgrammableSwitchEvent
}

func NewDoorbell() *Doorbell {
	svc := Doorbell{}
	svc.Service = New(TypeDoorbell)

	svc.ProgrammableSwitchEvent = characteristic.NewProgrammableSwitchEvent()
	svc.AddCharacteristic(svc.ProgrammableSwitchEvent.Characteristic)

	return &svc
}
