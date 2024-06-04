.PHONY: run
run: build-ui
	open http://localhost:8080
	go run ./cmd/http

.PHONY: build-ui
build-ui:
	cd vue-app && npm run build

.PHONY: serve
serve:
	go run ./cmd/http

