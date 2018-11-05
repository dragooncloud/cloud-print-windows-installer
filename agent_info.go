package main

import (
	"os"
	"os/user"
	"strings"
	"time"

	"github.com/alexbrainman/printer"
)

type agentInfo struct {
	UUID         string    `firestore:"uuid"`
	PrinterNames []string  `firestore:"printerNames"`
	HostName     string    `firestore:"hostname"`
	GetInfoError string    `firestore:"errorGettingInfo"`
	SavedAt      time.Time `firestore:"savedAt"`
	SavedBy      string    `firestore:"savedBy"`
}

func getAgentInfo(uuid string) agentInfo {
	printerNames, err := printer.ReadNames()
	currentUser, _ := user.Current()
	result := agentInfo{
		UUID:         uuid,
		HostName:     strings.ToLower(os.Getenv("COMPUTERNAME")),
		PrinterNames: printerNames,
		SavedAt:      time.Now(),
		SavedBy:      currentUser.Username,
	}
	if err != nil {
		result.GetInfoError = err.Error()
	}
	return result
}
