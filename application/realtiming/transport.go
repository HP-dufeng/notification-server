package realtiming

import (
	"net/http"
	"time"

	"github.com/fengdu/notification-server/core/notifications"
	"github.com/gorilla/websocket"
)

type onlineClient struct {
	ws           *websocket.Conn
	send         chan notifications.UserNotificationDto
	connectionID string
	userID       int64
	ConnectTime  int64
}

func (o *onlineClient) ConnectionID() string {
	return o.connectionID
}

func (o *onlineClient) UserID() int64 {
	return o.userID
}

func (o *onlineClient) Send() chan notifications.UserNotificationDto {
	return o.send
}

func (o *onlineClient) Write() {
	defer o.ws.Close()

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

func (o *onlineClient) Read() {
	defer o.ws.Close()

	for {
		_, _, err := o.ws.ReadMessage()
		if err != nil {
			break
		}

	}
}

func newOnlineClient(ws *websocket.Conn, userID int64) *onlineClient {
	return &onlineClient{
		ws:           ws,
		connectionID: string(notifications.NextID()),
		userID:       userID,
		ConnectTime:  time.Now().Unix(),
		send:         make(chan notifications.UserNotificationDto),
	}
}

type Hub struct {
	Join                chan *onlineClient
	Leave               chan *onlineClient
	OnlineClientManager notifications.OnlineClientManager
}

func NewHub(o notifications.OnlineClientManager) *Hub {
	return &Hub{
		Join:                make(chan *onlineClient),
		Leave:               make(chan *onlineClient),
		OnlineClientManager: o,
	}
}

func (h *Hub) Start() {
	for {
		select {
		case c := <-h.Join:
			h.OnlineClientManager.Add(c)
		case c := <-h.Leave:
			h.OnlineClientManager.Remove(c.connectionID)
		}
	}
}

const (
	socketBufferSize = 1024
)

var upgrader = &websocket.Upgrader{
	CheckOrigin:     func(r *http.Request) bool { return true },
	ReadBufferSize:  socketBufferSize,
	WriteBufferSize: socketBufferSize,
}

func (h *Hub) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	socket, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	client := newOnlineClient(socket, 1)
	h.Join <- client
	defer func() { h.Leave <- client }()

	go client.Write()

	client.Read()
}
