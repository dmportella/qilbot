package edsm

import (
	"errors"
	"fmt"
	"github.com/dmportella/qilbot/utilities"
	"io/ioutil"
	"math"
	"net/http"
	"net/url"
	"time"
)

func calculateDistance(sys1 *Coordinates, sys2 *Coordinates) float64 {
	deltaX := sys1.X - sys2.X
	deltaY := sys1.Y - sys2.Y
	deltaZ := sys1.Z - sys2.Z

	return math.Sqrt(deltaX*deltaX + deltaY*deltaY + deltaZ*deltaZ)
}

func getSystem(systemName string) (system System, err error) {
	url := fmt.Sprintf("https://www.edsm.net/api-v1/system?coords=1&systemName=%s", url.QueryEscape(systemName))

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("accept", "application/json; charset=utf-8")
	req.Header.Set("User-Agent", fmt.Sprintf("Qilbot a Discord Bot (https://github.com/dmportella/qilbot)"))

	client := &http.Client{Timeout: (20 * time.Second)}

	res, err := client.Do(req)

	defer res.Body.Close()

	if err != nil {
		err = errors.New("could not retrieve system information.")
		return
	}

	body, err := ioutil.ReadAll(res.Body)

	err = utilities.FromJson(body, &system)

	return
}