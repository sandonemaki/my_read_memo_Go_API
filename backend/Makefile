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

# スキーマ確認コマンド
schema-show:
	@echo "=== テーブル一覧 ==="
	@psql "$(GOOSE_DBSTRING)" -c "\dt"
	@echo ""
	@echo "=== 全テーブル詳細 ==="
	@for table in $$(psql "$(GOOSE_DBSTRING)" -t -c "SELECT tablename FROM pg_tables WHERE schemaname='public' AND tablename != 'goose_db_version';"); do \
		echo "--- $$table ---"; \
		psql "$(GOOSE_DBSTRING)" -c "\d $$table"; \
		echo ""; \
	done

schema-dump:
	@echo "=== 最新スキーマ情報をpkg/schema.txtに出力中 ===" > pkg/schema.txt
	@echo "Generated at: $$(date '+%Y-%m-%d %H:%M:%S')" >> pkg/schema.txt
	@echo "Latest migration: $$(goose -dir migration postgres "$(GOOSE_DBSTRING)" status | tail -1 | awk '{print $$NF}')" >> pkg/schema.txt
	@echo "" >> pkg/schema.txt
	@echo "=== テーブル一覧 ===" >> pkg/schema.txt
	@psql "$(GOOSE_DBSTRING)" -c "\dt" >> pkg/schema.txt 2>/dev/null
	@echo "" >> pkg/schema.txt
	@echo "=== 全テーブル詳細 ===" >> pkg/schema.txt
	@for table in $$(psql "$(GOOSE_DBSTRING)" -t -c "SELECT tablename FROM pg_tables WHERE schemaname='public' AND tablename != 'goose_db_version';"); do \
		echo "" >> pkg/schema.txt; \
		echo "--- $$table ---" >> pkg/schema.txt; \
		psql "$(GOOSE_DBSTRING)" -c "\d $$table" >> pkg/schema.txt 2>/dev/null; \
	done
	@echo ""
	@echo "✅ Schema saved to pkg/schema.txt"

schema-simple:
	@echo "=== 最新スキーマ一覧 ==="
	@for table in $$(psql "$(GOOSE_DBSTRING)" -t -c "SELECT tablename FROM pg_tables WHERE schemaname='public' AND tablename != 'goose_db_version' ORDER BY tablename;"); do \
		echo ""; \
		echo "📋 Table: $$table"; \
		echo "────────────────────────────────────────"; \
		psql "$(GOOSE_DBSTRING)" -c "SELECT column_name, data_type, is_nullable, column_default FROM information_schema.columns WHERE table_name='$$table' ORDER BY ordinal_position;" --quiet; \
	done
