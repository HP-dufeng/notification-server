package notifications

type Hub struct {
	AddClient           chan *OnlineClient
	RemoveClient        chan *OnlineClient
	OnlineClientManager OnlineClientManager
}

func NewHub(o OnlineClientManager) *Hub {
	return &Hub{
		AddClient:           make(chan *OnlineClient),
		RemoveClient:        make(chan *OnlineClient),
		OnlineClientManager: o,
	}
}

func (h *Hub) Start() {
	defer func() {
		close(h.AddClient)
		close(h.RemoveClient)
	}()

	for {
		select {
		case c := <-h.AddClient:
			h.OnlineClientManager.Add(*c)
		case c := <-h.RemoveClient:
			h.OnlineClientManager.Remove(c.ConnectionID)
		}
	}
}

// RealTimeNotifier to send real time notifications to users.
type RealTimeNotifier interface {
	SendNotifications(notifications []*UserNotificationDto) error
}

type wsRealTimeNotifier struct {
	onlineClientManager OnlineClientManager
}

// NewWebSocketRealTimeNotifier create a new instance of RealTimeNotifier interface implemented by WebSocket.
func NewWebSocketRealTimeNotifier(onlineClientManager OnlineClientManager) RealTimeNotifier {
	return &wsRealTimeNotifier{onlineClientManager}
}

func (r *wsRealTimeNotifier) SendNotifications(notifications []*UserNotificationDto) error {
	for _, v := range notifications {
		for _, c := range r.onlineClientManager.GetAllByUserID(v.UserID) {
			c.send <- *v
		}
	}

	return nil
}

type nullRealTimeNotifier struct{}

// NewNullRealTimeNotifier null pattern implementation of RealTimeNotifier interface.
func NewNullRealTimeNotifier() RealTimeNotifier {
	return nullRealTimeNotifier{}
}

func (nullRealTimeNotifier) SendNotifications(notifications []*UserNotificationDto) error {
	return nil
}
