package api

import (
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"

	dbAPI "github.com/pdxjohnny/s-db/api"
	"github.com/pdxjohnny/s/restQuickReply"
	"github.com/pdxjohnny/s/token"

	"github.com/pdxjohnny/eco/variables"
)

// GetDoc returns the accounts for an id
func GetDoc(w rest.ResponseWriter, r *rest.Request) {
	collection := r.PathParam("collection")
	id := r.PathParam("id")
	doc, err := dbAPI.Get(variables.ServiceDBURL, token.BackendToken, collection, id)
	// doc, err := dbAPI.Get(variables.ServiceDBURL, r.Env["JWT_RAW"].(string), collection, id)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	if doc == nil {
		w.(http.ResponseWriter).Write(restQuickReply.BlankResponse)
		return
	}
	w.WriteJson(doc)
}
