#!/bin/bash

# Goose + PostgreSQL用のschema.sql自動生成スクリプト
# Rails schema.rbと同様の機能を提供

DB_URL="postgres://yondeco:yondeco@localhost:5432/yondeco?sslmode=disable"
OUTPUT_FILE="schema.sql"

# 現在の日時を取得
TIMESTAMP=$(date '+%Y-%m-%d %H:%M:%S')

# 最新のマイグレーションバージョンを取得
LATEST_VERSION=$(goose -dir migration postgres "$DB_URL" status | grep "Applied At" -A 100 | tail -1 | awk '{print $NF}' | sed 's/--//' | awk '{print $1}')

# ヘッダーコメントを生成
cat > "$OUTPUT_FILE" << EOF
-- This file is auto-generated from the current state of the database. Instead
-- of editing this file, please use the goose migrations feature to 
-- incrementally modify your database, and then regenerate this schema definition.
--
-- This file represents the schema when all migrations have been applied.
-- When creating a new database, you can load this schema directly instead of
-- running all migrations from scratch.
--
-- It's strongly recommended that you check this file into your version control system.
--
-- Generated at: $TIMESTAMP
-- Latest migration: $LATEST_VERSION

-- Extensions (if any)
EOF

# ENUM型を追加
echo "" >> "$OUTPUT_FILE"
echo "-- Custom types" >> "$OUTPUT_FILE"
psql "$DB_URL" -t -c "SELECT 'CREATE TYPE ' || typname || ' AS ENUM (' || string_agg('''' || enumlabel || '''', ', ') || ');' FROM pg_enum JOIN pg_type ON pg_enum.enumtypid = pg_type.oid GROUP BY typname;" | grep -v "^$" >> "$OUTPUT_FILE"

# テーブル構造をダンプ（外部キー制約は除く）
echo "" >> "$OUTPUT_FILE"
echo "-- Tables" >> "$OUTPUT_FILE"
pg_dump --schema-only --no-owner --no-privileges --no-comments \
  "$DB_URL" \
  | grep -v "CREATE TYPE" \
  | grep -v "ALTER TYPE" \
  | grep -v "goose_db_version" \
  | sed '/^$/N;/^\n$/d' \
  >> "$OUTPUT_FILE"

echo ""
echo "✅ Schema generated successfully: $OUTPUT_FILE"
echo "📄 Latest migration: $LATEST_VERSION"
echo "🕐 Generated at: $TIMESTAMP"