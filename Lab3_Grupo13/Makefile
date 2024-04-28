.PHONY: all build run clean

all: build run

build:
	go build -o earth-server ./earth/earth.go
	go build -o team-client ./team/team.go

docker-build:
	docker-compose build

run: build
	./earth-server & sleep 2 && ./team-client

clean:
	rm -f earth-server team-client
	docker-compose down
