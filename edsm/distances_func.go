package edsm

import (
	"encoding/json"
	"errors"
	"fmt"
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

	fmt.Println(sys2)
	fmt.Println(sys1)

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

	err = fromJson(body, &system)

	return
}

func fromJson(data []byte, v interface{}) error {
	err := json.Unmarshal(data, v)
	if err != nil {
		fmt.Println(err)
		return errors.New("json decoding error.")
	}

	return nil
}
