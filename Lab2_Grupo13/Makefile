.PHONY: docker capitanes clean

docker:
	docker-compose build
	docker-compose up

capitanes:
	go run cliente.go

clean:
	docker-compose down