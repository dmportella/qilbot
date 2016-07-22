package utilities

import (
	"encoding/json"
	"errors"
	"github.com/dmportella/qilbot/logging"
)

func FromJson(data []byte, v interface{}) error {
	err := json.Unmarshal(data, v)
	if err != nil {
		logging.Error.Println(err)
		return errors.New("json decoding error.")
	}

	return nil
}
