package main

type hub struct {
	c map[*connection]bool
	b chan []byte
	r chan *connection
	u chan *connection
}

var h = hub{
	c : make(map[*connection]bool),
	u : make(chan *connection),
	b : make(chan []byte),
	r : make(chan *connection),
}

func (h *hub) run()  {
	for{
		select {
		case c := <- h.r
			h.c[c] = true
			c.data.Ip = c.ws.RemoteAddr().String()
			c.data.Type = "handshake"
		}
	}
}
