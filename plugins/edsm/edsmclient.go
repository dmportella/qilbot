package edsm

import (
	"time"
)

// Known endpoints for EDSM System and Distance API
const (
	EndpointEDSM = "https://www.edsm.net/"

	EndpointAPI = EndpointEDSM + "api-v1/"

	EndpointSystem        = EndpointAPI + "system"
	EndpointSystems       = EndpointAPI + "systems"
	EndpointDistances     = EndpointAPI + "distances"
	EndpointSphereSystems = EndpointAPI + "sphere-systems"

	EndpointSubmitDistances = EndpointAPI + "submit-distances"
)

// Known endpoints for EDSM Status API
const (
	EndpointStatus = EndpointEDSM + "api-status-v1/elite-server"
)

// Know endpoints for EDSM Log API
const (
	EndpointLogAPI = EndpointEDSM + "api-logs-v1/"

	EndpointLogGetPosition = EndpointLogAPI + "get-position"

	EndpointLogGetLogs     = EndpointLogAPI + "get-logs"
	EndpointLogSetLogs     = EndpointLogAPI + "set-logs"
	EndpointLogDeleteLogs  = EndpointLogAPI + "delete-logs"
	EndpointLogGetComments = EndpointLogAPI + "get-comments"
	EndpointLogGetComment  = EndpointLogAPI + "get-comment"
	EndpointLogSetComment  = EndpointLogAPI + "set-comment"
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
