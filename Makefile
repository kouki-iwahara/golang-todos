# @をつけるとmake実行時にコマンド名を表示しない
mysql:
	@docker exec -it golang_todos_db sh -c 'mysql -u golang_todos_user -plocal'

up:
	@docker-compose up

down:
	@docker-compose down

# 見つけたファイルをフルパスで出力→.goのファイルをgrep→結果に対しgoimports。
fmt:
	@find . -print | grep --regex '.*\.go' | xargs goimports -w -local "github.com/kouki-iwahara/golang-todos/api"

verify:
	@cd api && staticcheck ./... && go vet ./...