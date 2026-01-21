package ws

type Hub struct {
	clients    map[string]map[*Client]struct{}
	register   chan *Client
	unregister chan *Client
	broadcast  chan Message
}

type Message struct {
	GameID string
	Data   []byte
}

type Client struct {
	gameID string
	send   chan []byte
}

func NewHub() *Hub {
	hub := &Hub{
		clients:    map[string]map[*Client]struct{}{},
		register:   make(chan *Client),
		unregister: make(chan *Client),
		broadcast:  make(chan Message),
	}
	go hub.run()
	return hub
}

func (h *Hub) RunOnce() {
	select {
	case client := <-h.register:
		if h.clients[client.gameID] == nil {
			h.clients[client.gameID] = map[*Client]struct{}{}
		}
		h.clients[client.gameID][client] = struct{}{}
	case client := <-h.unregister:
		if clients := h.clients[client.gameID]; clients != nil {
			delete(clients, client)
			if len(clients) == 0 {
				delete(h.clients, client.gameID)
			}
		}
	case message := <-h.broadcast:
		for client := range h.clients[message.GameID] {
			select {
			case client.send <- message.Data:
			default:
			}
		}
	default:
	}
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			if h.clients[client.gameID] == nil {
				h.clients[client.gameID] = map[*Client]struct{}{}
			}
			h.clients[client.gameID][client] = struct{}{}
		case client := <-h.unregister:
			if clients := h.clients[client.gameID]; clients != nil {
				delete(clients, client)
				if len(clients) == 0 {
					delete(h.clients, client.gameID)
				}
			}
		case message := <-h.broadcast:
			for client := range h.clients[message.GameID] {
				select {
				case client.send <- message.Data:
				default:
				}
			}
		}
	}
}
