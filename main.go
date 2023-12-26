package main

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"math/rand"
)

/*
	- Prayers
	- Reflexions
	- Verses
	- Devotional
*/

var Prayers []string

type PrayerResponse struct {
	Message string
}

func main() {
	Prayers = append(Prayers,
		"Señor, gracias por darme el mayor regalo imaginable: la salvación a través de Jesucristo. Ayúdame a cultivar un corazón generoso y a ayudar a los neceistados. Abre mis ojos para ver oportunidades de compartir tu amor con aquellos que no te conocen. En el nombre de Jesús, Amen",
	)

	lambda.Start(handler)
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	body := PrayerResponse{
		Message: fmt.Sprintf(Prayers[rand.Intn(len(Prayers))]),
	}

	jsonBody, _ := json.Marshal(body)

	response := events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(jsonBody),
		Headers:    map[string]string{"Content-Type": "application/json"},
	}

	return response, nil
}
