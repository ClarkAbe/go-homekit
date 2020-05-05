// THIS FILE IS AUTO-GENERATED
package service

import (
	"ccms/hc/characteristic"
)

const TypeStatefulProgrammableSwitch = "88"

type StatefulProgrammableSwitch struct {
	*Service

	ProgrammableSwitchEvent       *characteristic.ProgrammableSwitchEvent
	ProgrammableSwitchOutputState *characteristic.ProgrammableSwitchOutputState
}

func NewStatefulProgrammableSwitch() *StatefulProgrammableSwitch {
	svc := StatefulProgrammableSwitch{}
	svc.Service = New(TypeStatefulProgrammableSwitch)

	svc.ProgrammableSwitchEvent = characteristic.NewProgrammableSwitchEvent()
	svc.AddCharacteristic(svc.ProgrammableSwitchEvent.Characteristic)

	svc.ProgrammableSwitchOutputState = characteristic.NewProgrammableSwitchOutputState()
	svc.AddCharacteristic(svc.ProgrammableSwitchOutputState.Characteristic)

	return &svc
}
