package main

import (
	"fmt"
)

func main() {
	// generate new / get existing machine uid
	uuid, err := getMachineUID()
	if err != nil {
		fmt.Printf("Error fetching machine UUID, %s\n", err.Error())
		return
	}
	// produce agent info
	agentInfo := getAgentInfo(uuid)
	// push this up to registration
	err = sendDetailsToCloud(uuid, agentInfo)
	if err != nil {
		fmt.Printf("Error pushing agent info to cloud, %s\n", err.Error())
		return
	}
	// start browser to URL
	openbrowser(fmt.Sprintf("https://example.org/claim-agent?id=%s", uuid))
}
