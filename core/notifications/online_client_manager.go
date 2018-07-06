package notifications

import (
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

// OnlineClient is the central class in the domain model.
type OnlineClient struct {
	ws           *websocket.Conn
	send         chan UserNotificationDto
	ConnectionID string
	UserID       int64
	ConnectTime  int64
}

func (o *OnlineClient) Sending() {
	defer func() {
		o.ws.Close()
	}()

	for {
		select {
		case message, ok := <-o.send:
			if !ok {
				o.ws.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			o.ws.WriteJSON(message)

		}
	}
}

func (o *OnlineClient) Closing(closed chan *OnlineClient) {
	defer func() {
		closed <- o
		o.ws.Close()
	}()

	for {
		_, _, err := o.ws.ReadMessage()

		if err != nil {
			closed <- o

			o.ws.Close()
			close(o.send)
			break
		}

	}
}

// NewOnlineClient create a new instance of OnlineClient.
func NewOnlineClient(ws *websocket.Conn, userID int64) *OnlineClient {
	return &OnlineClient{
		ws:           ws,
		ConnectionID: string(NextID()),
		UserID:       userID,
		ConnectTime:  time.Now().Unix(),
		send:         make(chan UserNotificationDto),
	}
}

// OnlineClientManager used to manage connected clients.
type OnlineClientManager interface {
	Add(OnlineClient)
	Remove(connectionID string)
	GetAllClients() []OnlineClient
	GetAllByUserID(userID int64) []OnlineClient
}

// NewOnlineClientManager create a new instance of OnlineClientManager.
func NewOnlineClientManager() OnlineClientManager {
	return &onlineClientManager{
		clients: make(map[string]*OnlineClient),
	}
}

type onlineClientManager struct {
	mtx     sync.RWMutex
	clients map[string]*OnlineClient
}

func (o *onlineClientManager) Add(oc OnlineClient) {
	o.mtx.Lock()
	defer o.mtx.Unlock()

	o.clients[oc.ConnectionID] = &oc
}

func (o *onlineClientManager) Remove(connectionID string) {
	o.mtx.Lock()
	defer o.mtx.Unlock()

	delete(o.clients, connectionID)
}

func (o *onlineClientManager) GetAllClients() []OnlineClient {
	clients := []OnlineClient{}
	for _, v := range o.clients {
		clients = append(clients, *v)
	}

	return clients
}

func (o *onlineClientManager) GetAllByUserID(userID int64) []OnlineClient {
	clients := []OnlineClient{}
	for _, v := range o.clients {
		if v.UserID == userID {
			clients = append(clients, *v)
		}
	}

	return clients
}
