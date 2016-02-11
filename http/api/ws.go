package api

//
// import (
// 	"io"
// 	"log"
//
// 	"github.com/pdxjohnny/microsocket/random"
// 	"github.com/spf13/viper"
//
// 	"golang.org/x/net/websocket"
// )
//
// const (
// 	// StateNotReady -
// 	StateNotReady = "not_ready"
// 	// StateHasLoginInfo -
// 	StateHasLoginInfo = "has_login_info"
// 	// StateLogin -
// 	StateLogin = "login"
// 	// StateReady -
// 	StateReady = "ready"
// 	// StateMakeCall -
// 	StateMakeCall = "make_call"
// 	// StateInCall -
// 	StateInCall = "in_call"
// 	// StateEndCall -
// 	StateEndCall = "end_call"
// 	// StateError -
// 	StateError = "error"
// )
//
// // WsMessage holds messages like state updates that get sent to the wss
// type WsMessage struct {
// 	// For changing state
// 	State string `json:"state,omitempty"`
// 	// For updating properties like gmail_username
// 	Set   string `json:"set,omitempty"`
// 	Value string `json:"value,omitempty"`
// 	// For making a call
// 	Number string `json:"number,omitempty"`
// }
//
// // WsReadyMessage is used to pass an id and websocket to the operator so that the
// // ws can be removed if it disconnects
// type WsReadyMessage struct {
// 	ID string
// 	Ws *websocket.Conn
// }
//
// // WsHandler handles wss connecting via websocket
// func WsHandler(ws *websocket.Conn) {
// 	// Generate a random wsID so that it can be removed if it disconnects
// 	wsID := random.Letters(10)
// 	log.Println("Ws connected", wsID)
// 	// Receive message from the ws
// 	for {
// 		var message WsMessage
// 		err := websocket.JSON.Receive(ws, &message)
// 		if err == io.EOF {
// 			ws.Close()
// 			// Make sure the operator knows that the ws has disconnected
// 			Op.WsDisconnected <- wsID
// 			log.Println("Ws disconnected", wsID)
// 			return
// 		} else if err != nil {
// 			log.Println("Error receiving from wsID", wsID, ":", err)
// 			return
// 		} else {
// 			// Echo back
// 			err := websocket.JSON.Send(ws, message)
// 			if err != nil {
// 				log.Println("Error sending to wsID", wsID, ":", err)
// 				return
// 			}
// 			// Preform actions based on state
// 			switch message.State {
// 			case StateNotReady:
// 				// Send it the login info
// 				websocket.JSON.Send(ws, map[string]string{
// 					"set":   "gmail_username",
// 					"value": viper.GetString("gmail_username"),
// 				})
// 				websocket.JSON.Send(ws, map[string]string{
// 					"set":   "gmail_password",
// 					"value": viper.GetString("gmail_password"),
// 				})
// 				// Tell it that it has the login info
// 				websocket.JSON.Send(ws, map[string]string{
// 					"state": StateHasLoginInfo,
// 				})
// 			case StateReady:
// 				// If its ready then the operator needs to know that
// 				wsReadyMessage := &WsReadyMessage{
// 					ID: wsID,
// 					Ws: ws,
// 				}
// 				Op.WsReady <- wsReadyMessage
// 			}
// 		}
// 	}
// }
