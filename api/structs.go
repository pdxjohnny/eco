package api

import "golang.org/x/net/websocket"

// LoginData provides username and password to authenticate with
type LoginData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// CallData contains the number to call
type CallData struct {
	Number string `json:"number"`
}

// ClientInCall contains the lock that gets sent back to the client once the
// ws is in a call for the requested number
type ClientInCall struct {
	Number string `json:"number"`
	Lock   string `json:"lock"`
}

// EndData the the lock on the call which allows a client to end that call
type EndData struct {
	Lock string `json:"lock"`
}

// WsData contains a recv channel that will be filled with messages
// received via websocket from the call center as well as a send channel
// which is used to send messages to the call center
type WsData struct {
	Ws    *websocket.Conn
	Close chan bool
	Err   chan error
	Recv  chan interface{}
	Send  chan interface{}
}

// StringData is used for the write method of WsData so that it can be used
// as an io.Writer
type StringData struct {
	Data string `json:"data"`
}
