package variables

const (
	// APIPathLoginServer authenticates a client
	APIPathLoginServer = "/login"
	// APIPathLogin is the path to login
	APIPathLogin = "/api" + APIPathLoginServer
	// APIPathRefreshServer updates a clients token
	APIPathRefreshServer = "/refresh"
	// APIPathRefresh is the path to login
	APIPathRefresh = "/api" + APIPathRefreshServer
	// APIPathCallServer is the path to make a call
	APIPathCallServer = "/call/:number"
	// APIPathCall is the path to make a call
	APIPathCall = "/api" + APIPathCallServer
	// APIPathEndServer is the path to end a call
	APIPathEndServer = "/end/:lock"
	// APIPathEnd is the path to end a call
	APIPathEnd = "/api" + APIPathEndServer
	// APIPathWsServer is the path wss connect to via websocket
	APIPathWsServer = "/ws"
	// APIPathWs is the path wss connect to via websocket
	APIPathWs = "/api" + APIPathWsServer
)
