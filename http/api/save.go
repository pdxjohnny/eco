package api

import (
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"

	dbAPI "github.com/pdxjohnny/s-db/api"
	"github.com/pdxjohnny/s/restQuickReply"

	"github.com/pdxjohnny/eco/variables"
)

// GetDoc returns the accounts for an id
func GetDoc(w rest.ResponseWriter, r *rest.Request) {
	collection := r.PathParam("collection")
	id := r.PathParam("id")
	doc, err := dbAPI.Get(variables.ServiceDBURL, r.Env["JWT_RAW"].(string), collection, id)
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

// GetSaveDoc returns the accounts for an id
func GetSaveDoc(w rest.ResponseWriter, r *rest.Request) {
	collection := r.PathParam("collection")
	id := r.PathParam("id")
	value := r.PathParam("value")
	doc, err := dbAPI.GetSave(variables.ServiceDBURL, r.Env["JWT_RAW"].(string), collection, id, value)
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

// PostSaveDoc uses get to retrive a document
func PostSaveDoc(w rest.ResponseWriter, r *rest.Request) {
	var recvDoc map[string]interface{}
	collection := r.PathParam("collection")
	id := r.PathParam("id")
	err := r.DecodeJsonPayload(&recvDoc)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	doc, err := dbAPI.Save(variables.ServiceDBURL, r.Env["JWT_RAW"].(string), collection, id, recvDoc)
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
