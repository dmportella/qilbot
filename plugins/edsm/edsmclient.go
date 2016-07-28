package edsm

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
