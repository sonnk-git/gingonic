dev-server:
	go get github.com/pilu/fresh && go run github.com/pilu/fresh
gqlgen:
	go get github.com/99designs/gqlgen@v0.17.20 && go run github.com/99designs/gqlgen generate
.PHONY: scheduler
scheduler:
	cd scheduler && go run main.go
