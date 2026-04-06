migrate:
	migrate -path /Users/mozhno/Documents/golang/finance/db -database 'postgres://postgres:qwerty@0.0.0.0:5442/postgres?sslmode=disable' up
	
tidy:
	go mod tidy && go fmt ./...