package main

import (
	"flag"

	"github.com/mukundan314/nym-cloud/server/nymclient"
	log "github.com/sirupsen/logrus"
)

var (
	nymClientUri = flag.String("n", "ws://localhost:1977", "uri to a nym-client")
	logLevel     = flag.String("l", "info", "log level to use; available options are debug, info, warn and error")
)

func main() {
	flag.Parse()

	switch *logLevel {
	case "debug":
		log.SetLevel(log.DebugLevel)
	case "info":
		log.SetLevel(log.InfoLevel)
	case "warn":
		log.SetLevel(log.WarnLevel)
	case "error":
		log.SetLevel(log.ErrorLevel)
	default:
		log.WithField("logLevel", *logLevel).Error("Unknown log level")
	}

	client, err := nymclient.New(*nymClientUri)
	if err != nil {
		log.WithField("nymClient", *nymClientUri).WithError(err).Fatal("Failed to connect to the nym client")
	}
	defer client.Close()
	log.WithField("nymClient", *nymClientUri).Debug("Successfully connected to the nym client")

	handler := requestHandler{Client: client}

	for {
		resp, err := client.Recv()
		if err != nil {
			log.WithError(err).Error("An error occurred when receiving data from the nym client")
		}

		if resp.Type == "error" {
			log.WithFields(log.Fields{
				"code": resp.Error.Code,
				"msg":  string(resp.Error.Message), // TODO: Verify that this is guaranteed to be a printable string
			}).Error("Received an error response from the nym client")
			continue
		}

		if resp.Type == "selfAddress" {
			log.WithField("address", resp.SelfAddress).Warn("Unexpectedly received a selfAddress response from the nym client")
			continue
		}

		if resp.Type != "received" {
			log.WithField("type", resp.Type).Warn("Received unknown response type from the nym client")
			continue
		}

		go handler.handle(*resp.Received)
	}
}
