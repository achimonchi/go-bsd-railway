running-test:
	go test -v ./test/...

run-app:
	DB_HOST=containers-us-west-107.railway.app \
	DB_PORT=5559 \
	DB_USER=postgres \
	DB_PASSWORD=0LMwCKrcqiLS7rMr111k \
	DB_NAME=railway \
	PORT=:8181 \
		go run main.go