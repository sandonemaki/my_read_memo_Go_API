# データベース接続情報
GOOSE_DRIVER=postgres
GOOSE_DBSTRING=postgres://yondeco:yondeco@localhost:5432/yondeco?sslmode=disable
MIGRATION_DIR=migration

# Goose コマンドのエイリアス
GOOSE=goose -dir $(MIGRATION_DIR) $(GOOSE_DRIVER) "$(GOOSE_DBSTRING)"

goose-status:
	$(GOOSE) status
goose-up:
	$(GOOSE) up
goose-down:
	$(GOOSE) down
goose-create:
	@if [ -z "$(NAME)" ]; then \
		echo "NAME is required"; \
		exit 1; \
	fi; \
	$(GOOSE) create $(NAME) $(TYPE:=sql)
goose-reset:
	$(GOOSE) reset
