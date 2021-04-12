package api

import (
	"context"
	"net/http/httptest"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/gorilla/mux"
	. "github.com/smartystreets/goconvey/convey"
)

func TestSetup(t *testing.T) {
	Convey("Given an API instance", t, func() {
		r := mux.NewRouter()
		ctx := context.Background()
		api := Setup(ctx, r)

		sess := session.Must(session.NewSessionWithOptions(session.Options{
			SharedConfigState: session.SharedConfigEnable,
		}))

		awsRegion := "eu-west-1"
		api.CognitoClient = cognitoidentityprovider.New(sess, &aws.Config{Region: &awsRegion})

		Convey("When created the following route(s) should have been added", func() {
			So(hasRoute(api.Router, "/hello", "GET"), ShouldBeTrue)
		})

		Convey("Ensure cognito client has been added to api", func() {
			So(api.CognitoClient, ShouldNotBeNil)
		})
	})
}

func hasRoute(r *mux.Router, path, method string) bool {
	req := httptest.NewRequest(method, path, nil)
	match := &mux.RouteMatch{}
	return r.Match(req, match)
}
