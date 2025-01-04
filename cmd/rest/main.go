package rest

import (
	"github.com/rs/zerolog"

	"projeto-cnpj-go/internal/infra"
	"projeto-cnpj-go/internal/infra/engine"
	"projeto-cnpj-go/internal/repository"
	"projeto-cnpj-go/internal/routes"
	"projeto-cnpj-go/internal/services"
)

func Serve() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	client := infra.Connect()

	repository.New(client.Client())

	svc := services.Service{}
	r := routes.SetupRouter(svc)

	go engine.Start(svc)

	r.Run("localhost:8080")
}
