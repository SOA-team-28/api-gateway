package main

import (
	"context"
	"example/gateway/config"
	encounter_service "example/gateway/proto/encounter-service"
	follower_service "example/gateway/proto/follower-service"
	user "example/gateway/proto/stakeholders-service"
	tour_service "example/gateway/proto/tour-service"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	cfg := config.GetConfig()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(
		ctx,
		cfg.ToursServiceAdress,
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("Failed to dial tour service:", err)
	}
	defer conn.Close()

	stakeConn, err := grpc.DialContext(
		ctx,
		cfg.StakeholdersServiceAddress,
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("Failed to dial user service:", err)
	}
	defer stakeConn.Close()

	encounterConn, err := grpc.DialContext(
		ctx,
		cfg.EncounterServiceAdress,
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("Failed to dial encounter service:", err)
	}
	defer encounterConn.Close()

	followerConn, err := grpc.DialContext(
		ctx,
		cfg.FollowerServiceAddress,
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("Failed to dial follower service:", err)
	}
	defer encounterConn.Close()

	gwmux := runtime.NewServeMux()

	client := tour_service.NewTourServiceClient(conn)
	err = tour_service.RegisterTourServiceHandlerClient(
		context.Background(),
		gwmux,
		client,
	)
	if err != nil {
		log.Fatalln("Failed to register tour service gateway:", err)
	}

	userClient := user.NewUserServiceClient(stakeConn)
	err = user.RegisterUserServiceHandlerClient(context.Background(), gwmux, userClient)
	if err != nil {
		log.Fatalln("Failed to register user service gateway:", err)
	}

	encounterClient := encounter_service.NewEncounterServiceClient(encounterConn)
	err = encounter_service.RegisterEncounterServiceHandlerClient(context.Background(), gwmux, encounterClient)
	if err != nil {
		log.Fatalln("Failed to register encounter service gateway:", err)
	}

	followerClient := follower_service.NewFollowerServiceClient(followerConn)
	err = follower_service.RegisterFollowerServiceHandlerClient(context.Background(), gwmux, followerClient)
	if err != nil {
		log.Fatalln("Failed to register follower service gateway:", err)
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

	stopCh := make(chan os.Signal, 1)
	signal.Notify(stopCh, syscall.SIGTERM, syscall.SIGINT)

	<-stopCh

	if err = gwServer.Close(); err != nil {
		log.Fatalln("error while stopping server: ", err)
	}
}
