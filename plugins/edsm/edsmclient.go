package edsm

import (
	"time"
)

// Known endpoints for EDSM
const (
	EndpointEDSM = "https://www.edsm.net/"

	EndpointStatus = EndpointEDSM + "api-status-v1/elite-server"

	EndpointAPI = EndpointEDSM + "api-v1/"

	EndpointSystem        = EndpointAPI + "system"
	EndpointSystems       = EndpointAPI + "systems"
	EndpointDistances     = EndpointAPI + "distances"
	EndpointSphereSystems = EndpointAPI + "sphere-systems"

	EndpointSubmitDistances = EndpointAPI + "submit-distances"
)

// APIClient EDSM api client.
type APIClient struct {
	Debug bool
}

// APIStatus represents the status response of edsm api status call.
type APIStatus struct {
	LastUpdate *time.Time `json:"lastUpdate,omitempty"`
	Type       string     `json:"type,omitempty"`
	Message    string     `json:"message,omitempty"`
	Status     int32      `json:"status"`
}
