.PHONY: dev
dev:
	go build -o ./tmp/$(APP_NAME) ./cmd/$(APP_NAME)/main.go && air

.PHONY: templ-watch
templ-watch:
	templ generate --watch --proxy=http://localhost:8080 --open-browser=false

.PHONY: tailwind-watch
tailwind-watch:
	./tailwindcss -i ./public/css/input.css -o ./public/css/styles.css --watch