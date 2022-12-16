package main

import (
	"log"

	"github.com/ferch5003/ferch5003/readme-maker/internal/platform/nasa"
	"github.com/ferch5003/ferch5003/readme-maker/internal/platform/nasa/dto"
	"github.com/joho/godotenv"
)

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}

func run() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}

	nasaClient := nasa.NewClient()

	apodParams := dto.APODRequestParams{}
	response, err := nasaClient.GetAPOD(apodParams)

	log.Println(response)

	return err
}
