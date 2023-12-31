DB_HOST = localhost
DB_PORT = 3306
DB_USER = root
DB_PASSWORD = pass
DB_NAME = ipfs
DB_TABLE = users
DB_CONTAINER_NAME = mysql-container

.PHONY: start-db stop-db

go.run:
	go run main.go

go.build:
	go build -o ./bin/ipfs-server main.go

go.test:
	go test -v ./...

go.mod:
	go mod tidy

start-db:
	docker run --name $(DB_CONTAINER_NAME) -p $(DB_PORT):$(DB_PORT) -e MYSQL_ROOT_PASSWORD=$(DB_PASSWORD) -e MYSQL_DATABASE=$(DB_NAME) -e MYSQL_PASSWORD=$(DB_PASSWORD) -d mysql:latest

stop-db:
	docker stop $(DB_CONTAINER_NAME)
	docker rm $(DB_CONTAINER_NAME)

start-ipfs:
	docker run -d --name ipfs \
      -e PRIVATE_PEER_ID=... \
      -e PRIVATE_PEER_IP_ADDR=... \
      -v ./001-test.sh:/container-init.d/001-test.sh \
      -p 4001:4001 \
      -p 127.0.0.1:8080:8080 \
      -p 127.0.0.1:5001:5001 \
      ipfs/kubo

stop-ipfs:
	docker stop ipfs
	docker rm ipfs

start-docker: start-db start-ipfs

stop-docker: stop-db stop-ipfs


start: go.mod start-docker wait  go.run

wait:
	sleep 30