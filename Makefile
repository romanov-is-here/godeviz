.PHONY: runhttp
runhttp:
	cd vue-app && npm run build
	go run ./cmd/http