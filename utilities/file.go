package utilities

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

// DefaultPermissions Permissions for file and directory creations.
const (
	DefaultPermissions = 0666
)

// FileOrDirectoryExists Check is a file or directory exists.
func FileOrDirectoryExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

// GetCurrentFolder Returns the
func GetCurrentFolder() (string, error) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	return dir, err
}

// SaveToFile Saves a byte array to a path.
func SaveToFile(data []byte, path string) error {
	return ioutil.WriteFile(path, data, DefaultPermissions)
}

// ReadFromFile Safe reads file content.
func ReadFromFile(filepath string) ([]byte, error) {
	if ok, err := FileOrDirectoryExists(filepath); err != nil || !ok {
		return nil, err
	}

	return ioutil.ReadFile(filepath)
}
