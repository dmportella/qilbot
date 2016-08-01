package edsm

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/dmportella/qilbot/logging"
	"github.com/dmportella/qilbot/utilities"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// Helper variables for url creation
var (
	urlSphereSystems = func(systemName string, radius float64) string {
		return fmt.Sprintf("%s?coords=1&showid=1&radius=%s&systemName=%s", EndpointSphereSystems, url.QueryEscape(strconv.FormatFloat(radius, 'f', 2, 64)), url.QueryEscape(systemName))
	}

	urlSystem = func(systemName string) string {
		return fmt.Sprintf("%s?coords=1&systemName=%s", EndpointSystem, url.QueryEscape(systemName))
	}
)

// NewAPIClient created an instance of APIClient
// debug: tells APIClient if it should be running in debug mode.
func NewAPIClient(debug bool) APIClient {
	return APIClient{Debug: debug}
}

func (client *APIClient) request(method string, url string, b []byte) (response []byte, err error) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(b))

	req.Header.Set("accept", "application/json; charset=utf-8")
	req.Header.Set("User-Agent", "Discord Bot (https://github.com/dmportella/qilbot, 0.0.0)")

	httpClient := &http.Client{Timeout: (120 * time.Second)}

	res, err := httpClient.Do(req)

	defer res.Body.Close()

	if err != nil {
		logging.Warning.Println("Request error", err)
		err = errors.New("Http request returned an error")
		return
	}

	response, err = ioutil.ReadAll(res.Body)

	if client.Debug {
		logging.Trace.Printf("API REQUEST\tURL :: %s\n", url)
		logging.Trace.Printf("API RESPONSE\tSTATUS :: %s\n", res.Status)
		for k, v := range res.Header {
			logging.Trace.Printf("API RESPONSE\tHEADER :: [%s] = %+v\n", k, v)
		}
		logging.Trace.Printf("API RESPONSE\tBODY :: [%s]\n", response)
	}
	return
}

// GetAPIStatus gets the API status for EDSM.
func (client *APIClient) GetAPIStatus() (status APIStatus, err error) {
	response, err := client.request("GET", EndpointStatus, nil)
	if err != nil {
		return
	}

	err = utilities.FromJSON(response, &status)
	return
}

// GetSphereSystems gets all the systems within a specified radius of the system provided.
// systemName: the name of the system to use as original.
// radius: float64 radius of the search.
func (client *APIClient) GetSphereSystems(systemName string, radius float64) (systems []System, err error) {
	response, err := client.request("GET", urlSphereSystems(systemName, radius), nil)
	if err != nil {
		return
	}

	err = utilities.FromJSON(response, &systems)
	return
}

// GetSystem gets the the specified system information.
// systemName: the name of the system to fetch
func (client *APIClient) GetSystem(systemName string) (system System, err error) {
	response, err := client.request("GET", urlSystem(systemName), nil)
	if err != nil {
		return
	}

	err = utilities.FromJSON(response, &system)
	return
}
