.PHONY: \
  goose-status \
  goose-up \
  goose-down \
  goose-create \
	goose-reset \
	golangci-lint \
	fmtcheck

# データベース接続情報
GOOSE_DRIVER=postgres
GOOSE_DBSTRING=postgres://yondeco:yondeco@localhost:5432/yondeco?sslmode=disable
MIGRATION_DIR=migration
SRCS = $(shell git ls-files '*.go')

define format
  go fmt ./... && goimports -w ./ && go mod tidy
endef

format:
  $(call format)

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
golangci-lint:
  golangci-lint run --timeout=1h
fmtcheck:
  @ $(foreach file,$(SRCS),gofmt -s -l $(file);)
wire-gen-usecase:
	wire gen ./core/usecase
wire-gen-injector:
	wire gen ./core/injector
