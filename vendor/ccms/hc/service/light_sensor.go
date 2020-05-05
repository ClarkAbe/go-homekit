// THIS FILE IS AUTO-GENERATED
package service

import (
	"ccms/hc/characteristic"
)

const TypeLightSensor = "84"

type LightSensor struct {
	*Service

	CurrentAmbientLightLevel *characteristic.CurrentAmbientLightLevel
}

func NewLightSensor() *LightSensor {
	svc := LightSensor{}
	svc.Service = New(TypeLightSensor)

	svc.CurrentAmbientLightLevel = characteristic.NewCurrentAmbientLightLevel()
	svc.AddCharacteristic(svc.CurrentAmbientLightLevel.Characteristic)

	return &svc
}
