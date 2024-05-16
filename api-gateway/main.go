package main

import (
	"context"
	"example/gateway/config"
	user "example/gateway/proto/stakeholders-service"
	tour_service "example/gateway/proto/tour-service"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	cfg := config.GetConfig()

	conn, err := grpc.DialContext(
		context.Background(),
		cfg.ToursServiceAdress,
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	stakeConn, err := grpc.DialContext(
		context.Background(),
		cfg.StakeholdersServiceAddress,
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("Failed to dial user service:", err)
	}

	gwmux := runtime.NewServeMux()
	// Register Greeter
	client := tour_service.NewTourServiceClient(conn)
	err = tour_service.RegisterTourServiceHandlerClient(
		context.Background(), 
		gwmux,
		client,
	)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}


	userClient := user.NewUserServiceClient(stakeConn)
	err = user.RegisterUserServiceHandlerClient(context.Background(), gwmux, userClient)
	if err != nil {
		log.Fatalln("Failed to register user service gateway:", err)
	}

	gwServer := &http.Server{
		Addr:    cfg.Address,
		Handler: gwmux,
	}

	go func() {
		if err := gwServer.ListenAndServe(); err != nil {
			log.Fatal("server error: ", err)
		}
	}()

	stopCh := make(chan os.Signal)
	signal.Notify(stopCh, syscall.SIGTERM)

	<-stopCh

	if err = gwServer.Close(); err != nil {
		log.Fatalln("error while stopping server: ", err)
	}
}
