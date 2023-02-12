package main

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/diegolemospadilha/api-products-go/internal/infra/akafka"
	"github.com/diegolemospadilha/api-products-go/internal/infra/repository"
	"github.com/diegolemospadilha/api-products-go/internal/infra/web"

	"github.com/diegolemospadilha/api-products-go/internal/usecase"
	"github.com/go-chi/chi/v5"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(host.docker.internal:3306)/products")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	repository := repository.NewProductRepositoryMysql(db)
	createProductUseCase := usecase.NewCreateProductUseCase(repository)
	listProductsUseCase := usecase.NewListProductsUseCase(repository)

	productHandlers := web.NewProductHandlers(
		createProductUseCase,
		listProductsUseCase,
	)

	r := chi.NewRouter()
	r.Post("/products", productHandlers.CreateProductHandler)
	r.Get("/products", productHandlers.ListProductsHandler)

	go http.ListenAndServe(":8000", r)

	msgChannel := make(chan *kafka.Message)
	go akafka.Consume([]string{"products"}, "host.docker.internal:9094", msgChannel)
	for msg := range msgChannel {
		dto := usecase.CreateProductInputDto{}
		err := json.Unmarshal(msg.Value, &dto)
		if err != nil {
			// logar o erro
		}

		_, err = createProductUseCase.Execute(dto)
	}
}
