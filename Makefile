.PHONY: run
run: build-ui
	go run ./cmd/godeviz show

.PHONY: build-ui
build-ui:
	cd vue-app && npm run build

.PHONY: show
show:
	go run ./cmd/godeviz show

.PHONY: svg
svg:
	go run ./cmd/godeviz svg