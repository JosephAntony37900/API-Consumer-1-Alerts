package websocket

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // En producción, restringe esto a tus dominios
	},
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type ClientManager struct {
	clients map[*websocket.Conn]bool
	sync.RWMutex
	maxClients int
}

var manager = ClientManager{
	clients:    make(map[*websocket.Conn]bool),
	maxClients: 100, // Ajusta según tu instancia
}

func HandleConnections(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("Error al actualizar conexión a WebSocket:", err)
		return
	}

	manager.Lock()
	if len(manager.clients) >= manager.maxClients {
		conn.Close()
		manager.Unlock()
		log.Println("Conexión rechazada: límite de clientes alcanzado")
		return
	}
	manager.clients[conn] = true
	manager.Unlock()

	log.Println("Nuevo cliente WebSocket conectado. Total:", len(manager.clients))

	// Configurar ping/pong
	conn.SetReadDeadline(time.Now().Add(60 * time.Second))
	conn.SetPongHandler(func(string) error {
		conn.SetReadDeadline(time.Now().Add(60 * time.Second))
		return nil
	})

	defer func() {
		manager.Lock()
		delete(manager.clients, conn)
		manager.Unlock()
		conn.Close()
		log.Println("Cliente desconectado. Total restante:", len(manager.clients))
	}()

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("Error de lectura: %v", err)
			}
			break
		}
		log.Printf("Mensaje recibido: %s\n", msg)
	}
}

func BroadcastMessage(message []byte) {
	manager.RLock()
	defer manager.RUnlock()

	for client := range manager.clients {
		client.SetWriteDeadline(time.Now().Add(10 * time.Second))
		err := client.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			log.Println("Error enviando mensaje:", err)
			client.Close()
			manager.Lock()
			delete(manager.clients, client)
			manager.Unlock()
		}
	}
}