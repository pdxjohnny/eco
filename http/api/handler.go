package api

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/ant0ine/go-json-rest/rest"
	"github.com/pdxjohnny/go-json-rest-middleware-jwt"

	dbVariables "github.com/pdxjohnny/s-db/db/variables"
	"github.com/pdxjohnny/s/token"
	"github.com/pdxjohnny/s/variables"
)

// CreateAuthMiddleware creates the middleware for authtication
func CreateAuthMiddleware() (*jwt.Middleware, error) {
	err := token.LoadTokenKeys()
	if err != nil {
		return nil, err
	}

	authMiddleware := &jwt.Middleware{
		Realm:            "numapp",
		SigningAlgorithm: token.SigningAlgorithm,
		Key:              token.TokenSignKey,
		VerifyKey:        token.TokenVerifyKey,
		Timeout:          time.Hour,
		MaxRefresh:       time.Hour * 24,
		Authenticator: func(username string, password string) error {
			return errors.New("This message should never be seen")
		},
	}
	return authMiddleware, nil
}

// MakeHandler creates the api request handler
func MakeHandler() *http.Handler {
	api := rest.NewApi()

	// Make sure we want to enable auth
	if variables.EnableAuth != false {
		authMiddleware, err := CreateAuthMiddleware()
		if err != nil {
			panic(err)
		}

		api.Use(&rest.IfMiddleware{
			// Only authenticate non login or register requests
			Condition: func(request *rest.Request) bool {
				return true
			},
			IfTrue: authMiddleware,
		})
	}
	api.Use(rest.DefaultProdStack...)
	router, err := rest.MakeRouter(
		// For accounts, looking up and updating
		rest.Get(dbVariables.APIPathSaveServer, GetDoc),
		rest.Get(dbVariables.APIPathGetSaveServer, GetSaveDoc),
		rest.Post(dbVariables.APIPathSaveServer, PostSaveDoc),
		// // For user actions such as login
		// rest.Post(variables.APIPathLoginUserServer, PostLoginUser),
		// rest.Get(variables.APIPathRefreshUserServer, PostRefreshUser),
		// rest.Post(variables.APIPathRegisterUserServer, PostRegisterUser),
		// rest.Get(variables.APIPathUserServer, GetUser),
		// rest.Post(variables.APIPathUserServer, PostUser),
	)
	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)
	handler := api.MakeHandler()
	return &handler
}
