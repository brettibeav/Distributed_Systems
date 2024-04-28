package main

import (
	"fmt"
	"math/rand"
	"net"
	"sync"
	"time"
)

var (
	// asignacion inicial de botines a planetas
	asignaciones = map[string]int{
		"A": 0,
		"B": 0,
		"C": 0,
		"D": 0,
		"E": 0,
		"F": 0,
	}
	mutex sync.Mutex
)

func main() {
	// conectarse al servidor
	puerto := ":8080"
	l, err := net.Listen("tcp", puerto)
	if err != nil {
		fmt.Println("Error al escuchar:", err)
		return
	}
	defer l.Close()
	fmt.Println("Servidor esperando a los capitanes en el puerto", puerto)

	// asignaciones aleatorias de botines para cada planeta
	rand.Seed(time.Now().UnixNano())
	for planeta := range asignaciones {
		asignaciones[planeta] = rand.Intn(10)
	}
	fmt.Println(asignaciones)

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error al aceptar conexión:", err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	// recibir mensaje del capitan
	defer conn.Close()
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error al leer mensaje:", err)
		return
	}
	planeta := string(buffer[:n])

	// encontrar el planeta con menos botines
	minBotines := 1 << 31 // inicializar con un valor grande
	var planetaMenosBotin string
	for p, botines := range asignaciones {
		if botines < minBotines {
			minBotines = botines
			planetaMenosBotin = p
		}
	}

	// asignar botín al planeta con menos botines y restar 1 al planeta del capitán
	mutex.Lock()
	asignaciones[planetaMenosBotin]++
	asignaciones[planeta]--
	botines := asignaciones[planetaMenosBotin]
	mutex.Unlock()
	fmt.Println(asignaciones)

	// mostrar detalles de la asignación
	respuesta := fmt.Sprintf("Botín asignado al planeta %s, cantidad actual: %d\n", planetaMenosBotin, botines)
	_, err = conn.Write([]byte(respuesta))
	if err != nil {
		fmt.Println("Error al enviar respuesta al cliente:", err)
		return
	}
}
