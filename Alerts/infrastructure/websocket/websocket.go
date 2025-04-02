package websocket

import (
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}
var clients = make(map[*websocket.Conn]bool)

func HandleConnections(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("Error al actualizar conexión a WebSocket:", err)
		return
	}
	defer conn.Close()

	clients[conn] = true
	log.Println("Nuevo cliente WebSocket conectado")

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("Cliente desconectado:", err)
			delete(clients, conn)
			break
		}
		log.Printf("Mensaje recibido: %s\n", msg)
	}
}

func BroadcastMessage(message string) {
	for client := range clients {
		err := client.WriteMessage(websocket.TextMessage, []byte(message))
		if err != nil {
			log.Println("Error enviando mensaje:", err)
			client.Close()
			delete(clients, client)
		}
	}
}
