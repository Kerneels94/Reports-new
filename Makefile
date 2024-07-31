# run:
# 	@go run ./cmd/main.go

run:
	@cd cmd && air

templ:
	@cd view && templ generate