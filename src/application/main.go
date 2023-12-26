package main

import (
	"context"
	"healthy-mind/src/data"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
)

/*
	- Reflexions
	- Verses
	- Devotional
*/

var ginLambda *ginadapter.GinLambdaV2
var ginEngine *gin.Engine

func init() {
	r := gin.Default()

	r.GET("/prayer", data.GetPrayer())

	r.GET("/reflection", data.GetReflection())

	ginEngine = r
}

func main() {
	lambdaTaskRoot := os.Getenv("LAMBDA_TASK_ROOT")
	if lambdaTaskRoot != "" {
		// If LAMBDA_TASK_ROOT is set, we are running inside Lambda.
		ginLambda = ginadapter.NewV2(ginEngine)

		lambda.Start(func(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
			return ginLambda.ProxyWithContext(ctx, req)
		})
	} else {
		// Else, we are running in a local or other environment.
		ginEngine.Run()
	}
}
