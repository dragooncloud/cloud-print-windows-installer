package main

import (
	"io/ioutil"
	"os"
	"path"

	"github.com/kjk/betterguid"
)

func writeUUID(path string, uuid string) (string, error) {
	err := ioutil.WriteFile(path, []byte(uuid), 0644)
	return uuid, err
}

func readUUID(path string) (string, error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	uuid := string(bytes)
	return uuid, nil
}

func uuidNotAlreadyPresent(path string) bool {
	_, err := os.Stat(path)
	return os.IsNotExist(err)
}

func getMachineUID() (string, error) {
	pathToUUIDFile := path.Join(os.Getenv("PROGRAMDATA"), "dragoon-cloud_print_agent-uuid")
	// file needs to be created
	if uuidNotAlreadyPresent(pathToUUIDFile) {
		return writeUUID(pathToUUIDFile, betterguid.New())
	}
	// file already exists
	return readUUID(pathToUUIDFile)
}
