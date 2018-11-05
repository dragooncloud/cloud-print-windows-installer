package main

import (
	"context"
	"errors"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
)

// TODO: push through a firebase function and use a secret signing tool
func sendDetailsToCloud(uuid string, data agentInfo) error {
	ctx := context.Background()
	clientInfo := option.WithCredentialsJSON([]byte(firestoreDbCreds))
	client, err := firestore.NewClient(ctx, "dragoon-cloud-print", clientInfo)
	if err != nil {
		return errors.New("failed to connect to firestore, reason: " + err.Error())
	}
	// push to database
	_, err = client.Collection("agents").Doc(uuid).Set(ctx, data)
	return err
}
