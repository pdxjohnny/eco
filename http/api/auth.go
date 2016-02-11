package api

//
// import (
// 	"net/http"
//
// 	"github.com/ant0ine/go-json-rest/rest"
//
// 	"github.com/pdxjohnny/numapp/api"
// 	"github.com/pdxjohnny/numapp/variables"
// )
//
// // PostLogiAuth logs in a user
// func PostLoginAuth(w rest.ResponseWriter, r *rest.Request) {
// 	var recvDoc map[string]interface{}
// 	err := r.DecodeJsonPayload(&recvDoc)
// 	if err != nil {
// 		rest.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	doc, err := api.LoginAuth(variables.ServiceAuthURL, variables.BackendToken, recvDoc)
// 	if err != nil {
// 		rest.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	w.WriteHeader(http.StatusOK)
// 	if doc == nil {
// 		w.(http.ResponseWriter).Write(variables.BlankResponse)
// 		return
// 	}
// 	w.WriteJson(doc)
// }
//
// // PostRefreshAuth logs in a user
// func PostRefreshAuth(w rest.ResponseWriter, r *rest.Request) {
// 	doc, err := api.RefreshAuth(variables.ServiceAuthURL, r.Env["JWT_RAW"].(string))
// 	if err != nil {
// 		rest.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	w.WriteHeader(http.StatusOK)
// 	if doc == nil {
// 		w.(http.ResponseWriter).Write(variables.BlankResponse)
// 		return
// 	}
// 	w.WriteJson(doc)
// }
//
// // PostRegisterAuth registers a new user
// func PostRegisterAuth(w rest.ResponseWriter, r *rest.Request) {
// 	var recvDoc map[string]interface{}
// 	err := r.DecodeJsonPayload(&recvDoc)
// 	if err != nil {
// 		rest.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	doc, err := api.RegisterAuth(variables.ServiceAuthURL, variables.BackendToken, recvDoc)
// 	if err != nil {
// 		rest.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	w.WriteHeader(http.StatusOK)
// 	if doc == nil {
// 		w.(http.ResponseWriter).Write(variables.BlankResponse)
// 		return
// 	}
// 	w.WriteJson(doc)
// }
//
// // GetAuth returns the accounts for an id
// func GetAuth(w rest.ResponseWriter, r *rest.Request) {
// 	id := r.PathParam("id")
// 	doc, err := api.GetAuth(variables.ServiceAuthURL, r.Env["JWT_RAW"].(string), id)
// 	if err != nil {
// 		rest.Error(w, err.Error(), http.StatusNotFound)
// 		return
// 	}
// 	if doc == nil {
// 		w.(http.ResponseWriter).Write(variables.BlankResponse)
// 		return
// 	}
// 	w.WriteJson(doc)
// }
//
// // PostAuth uses get to retrive a document
// func PostAuth(w rest.ResponseWriter, r *rest.Request) {
// 	var recvDoc map[string]interface{}
// 	id := r.PathParam("id")
// 	err := r.DecodeJsonPayload(&recvDoc)
// 	if err != nil {
// 		rest.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	doc, err := api.SaveAuth(variables.ServiceAuthURL, r.Env["JWT_RAW"].(string), id, recvDoc)
// 	if err != nil {
// 		rest.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	w.WriteHeader(http.StatusOK)
// 	if doc == nil {
// 		w.(http.ResponseWriter).Write(variables.BlankResponse)
// 		return
// 	}
// 	w.WriteJson(doc)
// }
