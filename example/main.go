package main

import (
	"context"
	"log"
	"os"

	"github.com/kokizzu/gotro/L"
	"github.com/rizalgowandy/synapse-go"
	"github.com/rizalgowandy/synapse-go/pkg/api"
	"github.com/rizalgowandy/synapse-go/pkg/entity"
)

func main() {
	client, err := synapse.NewClient(&api.Config{
		ID:      os.Getenv("ID"),
		Key:     os.Getenv("KEY"),
		Debug:   true,
		HostURL: "https://uat-api.synapsefi.com",
	})
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	resp, err := client.GenerateRefreshToken(ctx, &entity.GenerateRefreshTokenReq{
		UserID:        "5fd7acfd8677040053ad486d",
		UserIPAddress: "255.127.79.76",
		Email:         "test@synapsepay.com",
		Password:      "test1234",
	})
	if err != nil {
		log.Fatal(err)
	}
	L.Describe(resp.ID)
}
