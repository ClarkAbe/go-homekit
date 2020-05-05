// THIS FILE IS AUTO-GENERATED
package service

import (
	"ccms/hc/characteristic"
)

const TypeAirQualitySensor = "8D"

type AirQualitySensor struct {
	*Service

	AirQuality *characteristic.AirQuality
}

func NewAirQualitySensor() *AirQualitySensor {
	svc := AirQualitySensor{}
	svc.Service = New(TypeAirQualitySensor)

	svc.AirQuality = characteristic.NewAirQuality()
	svc.AddCharacteristic(svc.AirQuality.Characteristic)

	return &svc
}
