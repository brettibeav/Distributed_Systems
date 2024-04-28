package main

import (
	"fmt"
	"math/rand"
	"net"
	"time"
)

var planetas = []string{"A", "B", "C", "D", "E", "F"}

func main() {
	direccion := "localhost:8080"

	for {
		// conectarse al servidor
		conn, err := net.Dial("tcp", direccion)
		if err != nil {
			fmt.Println("Error al conectar al servidor:", err)
			return
		}

		capitan := rand.Intn(3) + 1                   // seleccionar un capitan aleatorio
		planeta := planetas[rand.Intn(len(planetas))] // seleccionar un planeta aleatorio
		time.Sleep(time.Second * time.Duration(3))    // simular tiempo de descubrimiento aleatorio

		// enviar mensaje al servidor
		_, err = conn.Write([]byte(planeta))
		if err != nil {
			fmt.Println("Error al enviar mensaje:", err)
			return
		}

		// recibir respuesta del servidor
		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error al leer respuesta del servidor:", err)
			return
		}

		fmt.Println("Respuesta del servidor:", string(buffer[:n]))
		// mostrar detalles de la asignación
		fmt.Printf("Capitan C%d encontro botin en Planeta P%s, enviando solicitud de asignacion\n", capitan, planeta)

		conn.Close()

		time.Sleep(time.Second)
	}
}
