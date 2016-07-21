package edsm

import (
	"errors"
	"github.com/dmportella/qilbot/plugins"
)

const (
	NAME        = "EDSM Plugin"
	DESCRIPTION = "Client plugin for Elite Dangerous Star Map web site."
)

func NewEDSMClient() EDSMClient {
	return EDSMClient{plugins.Plugin{Name: NAME, Description: DESCRIPTION}}
}

func (self *EDSMClient) GetDistanceBetweenTwoSystems(systemName1 string, systemName2 string) (distance float64, err error) {
	if sys1, ok1 := getSystem(systemName1); ok1 == nil {
		if sys2, ok2 := getSystem(systemName2); ok2 == nil {
			distance = calculateDistance(sys1.Coords, sys2.Coords)
			return
		}
	}

	err = errors.New("Unable to get distance between these systems.")

	return
}
