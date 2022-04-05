DOCKER_COMPOSE_PATH := $(PACKAGE_ROOT_PATH)docker-compose.yml
PJ_PATH := ./

.PHONY: docker_config
docker_config:
	cd $(PJ_PATH) && \
	docker-compose -f $(DOCKER_COMPOSE_PATH) config

.PHONY: app_init
app_init:
	@cp .env.example .env
	@cp .envrc.example .envrc

.PHONY: docker_init 
	@make docker_build
	@make docker_up

.PHONY: docker_build
docker_build:
	@cd $(PJ_PATH) && \
	docker-compose -f $(DOCKER_COMPOSE_PATH) build --no-cache

.PHONY: docker_up
docker_up:
	@cd $(PJ_PATH) && \
	docker-compose -f $(DOCKER_COMPOSE_PATH) up

.PHONY: docker_down
docker_down:
	@cd $(PJ_PATH) && \
	docker-compose -f $(DOCKER_COMPOSE_PATH) down

# Depends on "docker_up" task
.PHONY: api_exec
api_exec:
	@docker exec -it golang_todos_api sh

# Depends on "docker_up" task
.PHONY: mysql_exec
mysql:
	@docker exec -it golang_todos_db sh -c 'mysql -u golang_todos_user -p'

# 見つけたファイルをフルパスで出力→.goのファイルをgrep→結果に対しgoimports。
.PHONY: go_fmt
go_fmt:
	@find . -print | grep --regex '.*\.go' | xargs goimports -w -local "github.com/kouki-iwahara/golang-todos/api"

.PHONY: go_verify
go_verify:
	@cd api && staticcheck ./... && go vet ./...