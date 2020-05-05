// THIS FILE IS AUTO-GENERATED
package service

import (
	"ccms/hc/characteristic"
)

const TypeFilterMaintenance = "BA"

type FilterMaintenance struct {
	*Service

	FilterChangeIndication *characteristic.FilterChangeIndication
}

func NewFilterMaintenance() *FilterMaintenance {
	svc := FilterMaintenance{}
	svc.Service = New(TypeFilterMaintenance)

	svc.FilterChangeIndication = characteristic.NewFilterChangeIndication()
	svc.AddCharacteristic(svc.FilterChangeIndication.Characteristic)

	return &svc
}
