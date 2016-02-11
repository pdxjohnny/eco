package api

//
// import (
// 	"log"
// 	"time"
//
// 	"github.com/pdxjohnny/microsocket/random"
// 	"golang.org/x/net/websocket"
// )
//
// const (
// 	opBuffer = 20
// 	// DefaultCallTimeout is the default time to wait before ending a call
// 	DefaultCallTimeout = 2 * time.Minute
// )
//
// // Op is used if you don't want to create a NewOperator
// var Op *Operator
//
// // Operator connects calls that need to be made with wss that make them
// type Operator struct {
// 	Timeout time.Duration
// 	// Send requests to make a call to this channel
// 	MakeCall chan *ClientMakeCall
// 	// Send requests to end a call to this channel
// 	EndCall chan string
// 	// Wss who enter the ready state need to be delivered to the op here
// 	WsReady chan *WsReadyMessage
// 	// Wss that disconnect must be removed from the ready map
// 	WsDisconnected chan string
// 	// These are the calls waiting for a ws to make them
// 	CallsWaiting []*ClientMakeCall
// 	// These are the wss that are ready to make calls
// 	WssReady map[string]*websocket.Conn
// 	// Wss maped by locks so their calls can be ended
// 	WssInCall map[string]*websocket.Conn
// }
//
// func init() {
// 	Op = NewOperator()
// 	go Op.Route()
// }
//
// // NewOperator creates an operator
// func NewOperator() *Operator {
// 	op := Operator{
// 		Timeout:        DefaultCallTimeout,
// 		MakeCall:       make(chan *ClientMakeCall, opBuffer),
// 		EndCall:        make(chan string, opBuffer),
// 		WsReady:        make(chan *WsReadyMessage, opBuffer),
// 		WsDisconnected: make(chan string, opBuffer),
// 		CallsWaiting:   make([]*ClientMakeCall, 0),
// 		WssReady:       make(map[string]*websocket.Conn, opBuffer),
// 		WssInCall:      make(map[string]*websocket.Conn, opBuffer),
// 	}
// 	return &op
// }
//
// // Route starts the operator so that is listens for incoming call requests
// // and pairs them with available wss
// func (o *Operator) Route() {
// 	for {
// 		select {
// 		case ws := <-o.WsReady:
// 			log.Println(ws, "is ready")
// 			// Apped to the wss that are ready
// 			o.WssReady[ws.ID] = ws.Ws
// 			// Check to see if this new ws can make a call that is waiting to be
// 			// made
// 			o.CheckCallsWaiting()
// 		case wsID := <-o.WsDisconnected:
// 			// Take the ws out of the ready map because it disconnected
// 			delete(o.WssReady, wsID)
// 		case makeCall := <-o.MakeCall:
// 			log.Println("Got request to call", makeCall)
// 			// Add this number to the numbers that are waiting to be called
// 			o.CallsWaiting = append(o.CallsWaiting, makeCall)
// 			// Check to see if there is a ws available to call the number
// 			o.CheckCallsWaiting()
// 		case lock := <-o.EndCall:
// 			// Check to see if there is a ws in the in_call state for this lock
// 			ws, ok := o.WssInCall[lock]
// 			// If there is not ws for this lock then the ws has already gone
// 			// out of the in_call state
// 			if ok != true {
// 				continue
// 			}
// 			// End the call because the lock is associated with a ws in a call
// 			websocket.JSON.Send(ws, WsMessage{
// 				State: "end_call",
// 			})
// 			// Delete the call from the WssInCall because it has just been ended
// 			delete(o.WssInCall, lock)
// 			log.Println("Call with lock", lock, "has been ended")
// 		}
// 	}
// }
//
// // CheckCallsWaiting checks to see if there are any wss available to make
// // the calls waiting to be called
// func (o *Operator) CheckCallsWaiting() {
// 	log.Println("Checking for calls waiting...")
// 	log.Println(o.CallsWaiting, o.WssReady)
// 	// If there are no calls waiting to be made there is nothing to do
// 	if len(o.CallsWaiting) < 1 {
// 		return
// 	}
// 	// There are no wss ready so there is nothing we can do
// 	if len(o.WssReady) < 1 {
// 		return
// 	}
// 	// Grab a random ws
// 	var wsID string
// 	var ws *websocket.Conn
// 	for id, wsWs := range o.WssReady {
// 		wsID = id
// 		ws = wsWs
// 		// Grab one and exit the loop
// 		break
// 	}
// 	// The ws should be removed from the ready state becase it will now be
// 	// in the in_call state
// 	delete(o.WssReady, wsID)
// 	// The call we are about to make will no longer waiting to be made and
// 	// therefore needs to be removed from the waiting slice
// 	makeCall, callsWaiting := o.CallsWaiting[0], o.CallsWaiting[1:]
// 	o.CallsWaiting = callsWaiting
// 	// Take the ws we just removed from ready and put them in the make_call
// 	// state for the number we just removed from CallsWaiting
// 	log.Println("About to have", wsID, "call", makeCall.Number)
// 	websocket.JSON.Send(ws, WsMessage{
// 		State:  "make_call",
// 		Number: makeCall.Number,
// 	})
//
// 	// We need a way to end the call so take the ws and put them in the
// 	// WssInCall map, the key to access it will be a randomly generated
// 	// lock which the client will get back in case they want to end the call.
// 	// If the client never ends the call then EndCallTimeout will occor and the
// 	// call will be ended
//
// 	// Create the lock
// 	log.Println("Creating lock for", makeCall.Number, "...")
// 	lock := random.Letters(10)
// 	log.Println("Status make_call has been sent to", wsID,
// 		"for number", makeCall.Number,
// 		"with lock", lock,
// 	)
// 	// Put the ws in the WssInCall map by that lock
// 	o.WssInCall[lock] = ws
// 	// Start the timeout in case the client does not end the call
// 	go o.EndCallTimeout(lock, o.Timeout)
// 	// Return the lock to the client that requested this call in case they want
// 	// to end it before the timeout
// 	makeCall.Lock <- lock
// }
//
// // EndCallTimeout will send the end call request when the timeout occurs
// // it whoud be called as soon as the call is made
// func (o *Operator) EndCallTimeout(lock string, timeout time.Duration) {
// 	<-time.After(timeout)
// 	log.Println("Lock:", lock, "timed out")
// 	o.EndCall <- lock
// }
