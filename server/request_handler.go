package main

import (
	"crypto/sha1"
	"encoding/hex"
	"io/ioutil"
	"os"
	"path"

	"github.com/mukundan314/nym-cloud/server/nymclient"
	log "github.com/sirupsen/logrus"
	"github.com/vmihailenco/msgpack/v5"
)

type requestHandler struct {
	Client     *nymclient.Client
	StorageDir string
}

func (r requestHandler) store(data []byte) {
	hash := sha1.Sum(data)
	hexHash := hex.EncodeToString(hash[:])

	directory := path.Join(r.StorageDir, hexHash[:2])
	if _, err := os.Stat(directory); err != nil {
		if os.IsNotExist(err) {
			os.Mkdir(directory, 0700)
		} else {
			log.WithError(err).Error("Failed to store data")
			return
		}
	}

	filePath := path.Join(directory, hexHash[2:])
	if _, err := os.Stat(filePath); err != nil {
		if os.IsNotExist(err) {
			if err := ioutil.WriteFile(filePath, data, 0400); err != nil {
				log.WithError(err).Error("Failed to store data")
			}
		} else {
			log.WithError(err).Error("Failed to store data")
			return
		}
	}
}

func (r requestHandler) Handle(rawRequest nymclient.ReceivedResponse) {
	var request map[string]interface{}
	err := msgpack.Unmarshal(rawRequest.Message, &request)
	if err != nil {
		log.WithError(err).WithField("rawRequest", rawRequest.Message).Error("Failed to parse request")
		return
	}

	switch requestType := request["type"].(string); requestType {
	case "store":
		r.store(request["data"].([]byte))
	default:
		log.WithField("requestType", requestType).Error("Unknown reqeustType")
	}
}
