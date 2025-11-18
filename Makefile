ARGS?=
dev:
	go run ./*.go $(ARGS)
build:
	go build -o ./out/ ./...


.PHONY: dev build

# go run ./*.go add --tag deploy --cmd 'docker compose build && docker compose push && kubectl rollout restart deploy/{{ .service }} -n {{ .ns }}'