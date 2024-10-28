package storage

import (
	"encoding/json"
	"os"
)

const jsonFilePath = "packages.json"

func ReadPackages() (map[string]string, error) {
	file, err := os.ReadFile(jsonFilePath)
	if err != nil {
		if os.IsNotExist(err) {
			return make(map[string]string), nil
		}
		return nil, err
	}

	var packages map[string]string
	err = json.Unmarshal(file, &packages)
	if err != nil {
		return nil, err
	}

	return packages, nil
}

func WritePackages(packages map[string]string) error {
	file, err := json.MarshalIndent(packages, "", " ")
	if err != nil {
		return err
	}

	err = os.WriteFile(jsonFilePath, file, 0644)
	if err != nil {
		return err
	}

	return nil
}