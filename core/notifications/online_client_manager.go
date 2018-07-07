package notifications

import (
	"sync"
)

type OnlineClient interface {
	ConnectionID() string
	UserID() int64
	Send() chan UserNotificationDto
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

	o.clients[oc.ConnectionID()] = &oc
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
		if (*v).UserID() == userID {
			clients = append(clients, *v)
		}
	}

	return clients
}
