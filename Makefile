.PHONY: openapi
openapi: openapi_http

.PHONY: openapi_http
openapi_http:
	@./scripts/openapi-http.sh tts pkg/tts/handlers handlers

.PHONY: lint
lint:
	@go-cleanarch
	@./scripts/lint.sh common
	@./scripts/lint.sh tts

.PHONY: fmt
fmt:
	goimports -l -w pkg/

.PHONY: build
build:
	docker build --no-cache --progress=plain -t ritlab/tts:dev -f docker/app/Dockerfile \
	--build-arg PROFILE=dev \
	.

.PHONY: run
run:
	docker compose up -d --force-recreate --build tts

delete-old-image:
	docker rmi $$(docker images --filter dangling=true -q )
