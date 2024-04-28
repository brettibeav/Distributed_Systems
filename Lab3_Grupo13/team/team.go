package main

import (
	"context"
	"log"
	"math/rand"
	"munition/grupo13/munition"
	"sync"
	"time"

	"google.golang.org/grpc"
)

func solicitMunition(client munition.EarthClient, teamID int32) {
	reqAT := rand.Intn(11) + 20
	reqMP := rand.Intn(6) + 10

	log.Printf("Team %d solicitando %d AT y %d MP.", teamID, reqAT, reqMP)

	for {
		resp, err := client.SolicitudeM(context.Background(), &munition.MunitionRequest{
			TeamId:  teamID,
			AtCount: int32(reqAT),
			MpCount: int32(reqMP),
		})

		if err != nil {
			log.Printf("Error en solicitud: %v", err)
			time.Sleep(3 * time.Second)
			continue
		}

		if resp.Granted {
			log.Printf("EXITOSA: equipo %d recibe municiones. %s", teamID, resp.Message)
			break
		} else {
			log.Printf("DENEGADO: %s. Reintentando en 3 segundos.", resp.Message)
			time.Sleep(3 * time.Second)
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to Earth: %v", err)
	}

	defer conn.Close()

	client := munition.NewEarthClient(conn)

	time.Sleep(10 * time.Second)

	var wg sync.WaitGroup
	wg.Add(4)

	for i := int32(1); i <= 4; i++ {
		go func(teamID int32) {
			defer wg.Done()
			solicitMunition(client, teamID)
		}(i)
	}

	wg.Wait()
}
