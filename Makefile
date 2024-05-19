.PHONY: run
run:
	cd vue-app && npm run build
	open http://localhost:8080
	go run ./cmd/http