package utilities

import (
	"encoding/json"
	"errors"
	"github.com/dmportella/qilbot/logging"
)

// FromJSON converts json object representation from a byte array into golang struct.
func FromJSON(data []byte, v interface{}) error {
	err := json.Unmarshal(data, v)
	if err != nil {
		logging.Error.Println(err)
		return errors.New("json decoding error")
	}

	return nil
}

// ToJSON coverts a object into it's json representation and returns a byte array.
func ToJSON(v interface{}) (data []byte, err error) {
	data, err = json.Marshal(v)
	if err != nil {
		return
	}
	return
}
