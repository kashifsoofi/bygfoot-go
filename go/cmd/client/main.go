package main

import (
	"context"
	"os"

	"github.com/kashifsoofi/bygfoot-go/pkg/api"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})

	var port string
	var ok bool
	if port, ok = os.LookupEnv("PORT"); ok {
		log.WithFields(log.Fields{
			"PORT": port,
		}).Info("PORT env var defined")
	} else {
		port = "9000"
		log.WithFields(log.Fields{
			"PORT": port,
		}).Warn("PORT env var not defined. Going with default")
	}

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":"+port, grpc.WithInsecure())
	if err != nil {
		log.WithFields(log.Fields{
			"Error": err.Error(),
		}).Fatal("Failed to connect")
	}
	defer conn.Close()

	c := api.NewRegionServiceClient(conn)

	resp, err := c.GetRegions(context.Background(), &api.GetRegionsRequest{})
	if err != nil {
		log.WithFields(log.Fields{
			"Error": err.Error(),
		}).Fatal("Error when calling GetRegions")
	}

	log.WithFields(log.Fields{
		"Regions": resp.Regions,
	}).Info("Regions from server")
}
