package main

import (
	"context"
	"fmt"
	"log"
	"munition/grupo13/munition"
	"net"
	"sync"
	"time"

	"google.golang.org/grpc"
)

type EarthServer struct {
	munition.UnimplementedEarthServer
	inventoryAT int32
	inventoryMP int32
	mutex       sync.Mutex
}

func (e *EarthServer) SolicitudeM(ctx context.Context, req *munition.MunitionRequest) (*munition.MunitionResponse, error) {
	e.mutex.Lock()
	defer e.mutex.Unlock()

	response := &munition.MunitionResponse{}

	if req.AtCount <= e.inventoryAT && req.MpCount <= e.inventoryMP {
		e.inventoryAT -= req.AtCount
		e.inventoryMP -= req.MpCount
		response.Granted = true
		response.Message = fmt.Sprintf("EXITOSA: AT %d; MP %d", e.inventoryAT, e.inventoryMP)
		log.Printf("Recibe solicitud de team %d, %d AT and %d MP, EXITOSA", req.TeamId, req.AtCount, req.MpCount)
	} else {
		response.Granted = false
		response.Message = fmt.Sprintf("DENEGADA: AT %d; MP %d", e.inventoryAT, e.inventoryMP)
		log.Printf("Recibe solicitud de team %d, %d AT and %d MP, DENEGADA", req.TeamId, req.AtCount, req.MpCount)
	}

	return response, nil
}

func (e *EarthServer) produceMunition() {
	for {
		time.Sleep(5 * time.Second)
		e.mutex.Lock()

		newAT := int32(10)
		newMP := int32(5)

		if e.inventoryAT+newAT > 50 {
			newAT = 0
		}

		if e.inventoryMP+newMP > 20 {
			newMP = 0
		}

		e.inventoryAT += newAT
		e.inventoryMP += newMP

		log.Printf("Producido %d AT y %d MP. Inventario actual: %d AT, %d MP.", newAT, newMP, e.inventoryAT, e.inventoryMP)

		e.mutex.Unlock()
	}
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	server := grpc.NewServer()
	earthServer := &EarthServer{
		inventoryAT: 0,
		inventoryMP: 0,
	}

	munition.RegisterEarthServer(server, earthServer)

	go earthServer.produceMunition()

	log.Printf("Starting Earth gRPC server on :50051")

	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
