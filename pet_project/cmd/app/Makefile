DB_DSN := "postgres://postgres:yourpassword@localhost:5432/main?sslmode=disable"
MIGRATE := migrate -path ./migrations -database $(DB_DSN)
 
migrate-new:
	migrate create -ext sql -dir ./migrations $(NAME)

migrate:
	$(MIGRATE) up

migrate-down:
	$(MIGRATE) down

run:
	go run D:/pet_project/cmd/app/main.go

gen:
	oapi-codegen -config "openapi\.openapi" -include-tags tasks -package tasks "openapi\openapi.yaml" > "D:\pet_project\internal\web\tasks\api.gen.go"

lint:
	golangci-lint run --out-format=colored-line-number
