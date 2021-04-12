package api

import (
	"context"

	"github.com/gorilla/mux"

	cognito "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

var CognitoClient *cognito.CognitoIdentityProvider

//API provides a struct to wrap the api around
type API struct {
	Router        *mux.Router
	CognitoClient *cognito.CognitoIdentityProvider
}

//Setup function sets up the api and returns an api
func Setup(ctx context.Context, r *mux.Router) *API {
	api := &API{
		Router:        r,
		CognitoClient: CognitoClient,
	}

	r.HandleFunc("/hello", HelloHandler(ctx)).Methods("GET")
    return api
}
