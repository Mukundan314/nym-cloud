package main

import (
	"github.com/mukundan314/nym-cloud/server/nymclient"
	log "github.com/sirupsen/logrus"
	"github.com/vmihailenco/msgpack/v5"
)

type requestHandler struct {
	Client *nymclient.Client
}

func (r requestHandler) handle(rawRequest nymclient.ReceivedResponse) {
	var request map[string]interface{}
	err := msgpack.Unmarshal(rawRequest.Message, &request)
	if err != nil {
		log.WithError(err).WithField("rawRequest", rawRequest.Message).Error("Failed to parse request")
		return
	}

	switch requestType := request["type"].(string); requestType {
	default:
		log.WithField("requestType", requestType).Error("Unknown reqeustType")
	}
}
