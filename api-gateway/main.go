package main

import (
	"context"
	"example/gateway/config"
	encounter_service "example/gateway/proto/encounter-service"
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

	"strings"

	"github.com/dgrijalva/jwt-go"
	"google.golang.org/grpc/metadata"
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

	gwServer := &http.Server{
		Addr:    cfg.Address,
		Handler: jwtMiddleware(gwmux),
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

var jwtKey = []byte("secreet")

func jwtMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Preskoči autorizaciju za login i registraciju
		if r.URL.Path == "/login" || r.URL.Path == "/users" {
			next.ServeHTTP(w, r)
			return
		}

		// Uzimanje tokena iz Authorization header-a
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Nedostaje Authorization header", http.StatusUnauthorized)
			return
		}

		log.Println("Authorization header:", authHeader)
		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		// Parsiranje JWT tokena
		// Parsiranje JWT tokena
		claims := jwt.MapClaims{}
		token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err != nil || !token.Valid {
			http.Error(w, "Nevalidan token", http.StatusUnauthorized)
			return
		}

		log.Println("claims", claims)

		log.Println("claims", claims)
		// Dodavanje korisničkog ID-ja u kontekst zahteva
		ctx := context.WithValue(r.Context(), "userId", claims["userId"])

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func grpcMetadataAnnotator(ctx context.Context, req *http.Request) metadata.MD {
	userID, ok := ctx.Value("userId").(string)
	if !ok {
		return metadata.MD{}
	}
	return metadata.Pairs("userId", userID)
}
