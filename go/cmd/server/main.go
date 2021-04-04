package main

import (
	"net"
	"os"

	internal "github.com/kashifsoofi/bygfoot-go/internal/api"
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

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.WithFields(log.Fields{
			"Error": err.Error(),
		}).Fatal("Failed to listen")
	}

	s := internal.Server{}

	grpcServer := grpc.NewServer()
	log.Info("gRPC server started at ", port)

	api.RegisterRegionServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.WithFields(log.Fields{
			"Error": err.Error(),
		}).Fatal("Failed to serve")
	}
}
