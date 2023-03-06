package idfy

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"bitbucket.org/junglee_games/getsetgo/instrumenting/newrelic"
	logger "bitbucket.org/junglee_games/getsetgo/logger_old"
)

type Config struct {
	AccountId string
	ApiKey    string
	EndPoint  string
}

func (config *Config) GetIdfyAccountId() string {
	return config.AccountId
}
func (config *Config) GetIdfyApiKey() string {
	return config.ApiKey
}
func (config *Config) GetIdfyEndpoint() string {
	return config.EndPoint
}
func main() {

	cfg := Config{
		AccountId: "accountId",
		ApiKey:    "apikey",
		EndPoint:  "localhost:8080",
	}
	
	ctx := context.Background()
	logger, err := logger.GetLoggerHandle(ctx, "logType")

	app, err := newrelic.New("NewRelicAppName", "NewRelicLicenseKey")
	if err != nil {
		logger.Error(ctx, err)
		panic(err)
	}
	httpclint := http.Client{}
	idfyClient := New(&cfg, *app, &httpclint)

	idfyR := IdfyRequest{
		TaskID:  "12",
		GroupID: "34",
		Data: Data{
			Document1: "doc1",
			Document2: "doc2",
			Consent:   "nien",
			AdvancedDetails: AdvancedDetails{
				ExtractQrInfo:     true,
				ExtractLast4Digit: true,
			},
		},
	}
	res1, err := idfyClient.ExtractAadhar("StringOfImageWithBase64", idfyR)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(res1)
	}
	res2, err := idfyClient.ExtractPan("StringOfImageWithBase64", idfyR)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(res2)
	}
	res3, err := idfyClient.ExtractDl("StringOfImageWithBase64", idfyR)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(res3)
	}
	res4, err := idfyClient.ExtractVoter("StringOfImageWithBase64", idfyR)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(res4)
	}
	res5, err := idfyClient.ExtractPassport("StringOfImageWithBase64", idfyR)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(res5)
	}
}
