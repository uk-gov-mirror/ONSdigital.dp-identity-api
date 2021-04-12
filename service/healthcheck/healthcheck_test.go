package healthcheck_test

import (
	"context"
	"errors"
	"net/http"
	"testing"

	health "github.com/ONSdigital/dp-healthcheck/healthcheck"
	healthcheck "github.com/ONSdigital/dp-identity-api/service/healthcheck/healthcheckmock"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider/cognitoidentityprovideriface"

	. "github.com/smartystreets/goconvey/convey"
)

// Cognito Mocking - start..
type mockCognitoIdentityProviderClient struct {
    cognitoidentityprovideriface.CognitoIdentityProviderAPI
}
func (m *mockCognitoIdentityProviderClient) DescribeUserPool(poolInputData *cognitoidentityprovider.DescribeUserPoolInput) (*cognitoidentityprovider.DescribeUserPoolOutput, error) {
	exsitingPoolID := "us-west-2_aaaaaaaaa"
    if *poolInputData.UserPoolId != exsitingPoolID {
		return nil, errors.New("Failed to load user pool data")
	}
	// just don't return an error here which satisfies health checker
	return nil, nil
}
// Cognito Mocking - end.

func TestGetHealthCheck(t *testing.T) {
	ctx := context.Background()

	Convey("dp-identity-api healthchecker reports healthy", t, func() {
		healthcheck.CognitoClient = &mockCognitoIdentityProviderClient{}
		healthcheck.AWSUserPoolID = "us-west-2_aaaaaaaaa"
	
		checkState := health.NewCheckState("dp-identity-api-test")
	
		err := healthcheck.Checker(ctx, checkState)
		Convey("When GetHealthCheck is called", func() {
			Convey("Then the HealthCheck flag is set to true and HealthCheck is returned", func() {
				So(checkState.StatusCode(), ShouldEqual, http.StatusOK)
				So(checkState.Status(), ShouldEqual, health.StatusOK)
				So(err, ShouldBeNil)
			})
		})
	})

	Convey("dp-identity-api healthchecker reports critical", t, func() {
		healthcheck.CognitoClient = &mockCognitoIdentityProviderClient{}
		healthcheck.AWSUserPoolID = "us-west-2_aaaaaaaab"
	
		checkState := health.NewCheckState("dp-identity-api-test")
	
		err := healthcheck.Checker(ctx, checkState)
		Convey("When GetHealthCheck is called", func() {
			Convey("Then the HealthCheck flag is set to true and HealthCheck is returned", func() {
				So(checkState.StatusCode(), ShouldEqual, http.StatusTooManyRequests)
				So(checkState.Status(), ShouldEqual, health.StatusCritical)
				So(err, ShouldNotBeNil)
			})
		})
	})
}