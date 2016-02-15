package api

import (
	"net/http"
	"strconv"

	"github.com/ant0ine/go-json-rest/rest"

	dbAPI "github.com/pdxjohnny/s-db/api"
	"github.com/pdxjohnny/s/restQuickReply"
	"github.com/pdxjohnny/s/token"

	"github.com/pdxjohnny/eco/variables"
)

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
	// We need to format as a GeoJSON point
	// https://docs.mongodb.org/manual/reference/geojson/
	// { type: "Point", coordinates: [ 40, 5 ] }
	// Always list coordinates in longitude, latitude order.
	longitude, err := strconv.ParseFloat(recvDoc["longitude"].(string), 64)
	if err != nil {
		rest.Error(w, "longitude needs to be a float", http.StatusInternalServerError)
		return
	}
	latitude, err := strconv.ParseFloat(recvDoc["latitude"].(string), 64)
	if err != nil {
		rest.Error(w, "latitude needs to be a float", http.StatusInternalServerError)
		return
	}
	recvDoc["location"] = map[string]interface{}{
		"type": "Point",
		"coordinates": []float64{
			longitude,
			latitude,
		},
	}
	// Now save it
	doc, err := dbAPI.Save(variables.ServiceDBURL, token.BackendToken, collection, id, recvDoc)
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
