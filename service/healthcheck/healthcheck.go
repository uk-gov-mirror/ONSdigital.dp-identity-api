package healthcheck

import (
	"context"
	"net/http"

	health "github.com/ONSdigital/dp-healthcheck/healthcheck"
	cognito "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

const CognitoHealthy  = "Cognito Healthy"

// package variables
var (
	CognitoClient *cognito.CognitoIdentityProvider
	AWSUserPoolID string
)

func Checker(ctx context.Context, state *health.CheckState) error {
	// attempt to retrieve list of users, if err returned we have a problem with user pool and need to set health check to critical
	_, err := CognitoClient.DescribeUserPool(&cognito.DescribeUserPoolInput{UserPoolId: &AWSUserPoolID})

	if err != nil {
        state.Update(health.StatusCritical, err.Error(), http.StatusTooManyRequests)
		return err
	}
	state.Update(health.StatusOK, CognitoHealthy, http.StatusOK)

	return nil
}