package api

import (
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"

	dbAPI "github.com/pdxjohnny/s-db/api"
	"github.com/pdxjohnny/s/restQuickReply"
	"github.com/pdxjohnny/s/token"

	"github.com/pdxjohnny/eco/variables"
)

// PostNear retives locations near the point specified
func PostNear(w rest.ResponseWriter, r *rest.Request) {
	collection := r.PathParam("collection")

	var recvDoc map[string]interface{}
	err := r.DecodeJsonPayload(&recvDoc)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Preform the near query
	doc, err := dbAPI.Near(variables.ServiceDBURL, token.BackendToken, collection, recvDoc)
	// doc, err := dbAPI.Save(variables.ServiceDBURL, r.Env["JWT_RAW"].(string), collection, id, recvDoc)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	if doc == nil {
		w.(http.ResponseWriter).Write(restQuickReply.BlankResponse)
		return
	}
	w.WriteJson(doc)
}
