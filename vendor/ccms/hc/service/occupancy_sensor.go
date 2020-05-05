// THIS FILE IS AUTO-GENERATED
package service

import (
	"ccms/hc/characteristic"
)

const TypeOccupancySensor = "86"

type OccupancySensor struct {
	*Service

	OccupancyDetected *characteristic.OccupancyDetected
}

func NewOccupancySensor() *OccupancySensor {
	svc := OccupancySensor{}
	svc.Service = New(TypeOccupancySensor)

	svc.OccupancyDetected = characteristic.NewOccupancyDetected()
	svc.AddCharacteristic(svc.OccupancyDetected.Characteristic)

	return &svc
}
