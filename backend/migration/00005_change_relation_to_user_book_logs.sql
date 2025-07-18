-- +goose Up
-- +goose StatementBegin
-- ENUM型を作成
CREATE TYPE reading_type AS ENUM ('smooth_randoku', 'slowly_seidoku', 'smooth_tudoku');

-- まず外部キー制約を削除
ALTER TABLE randoku_images
DROP CONSTRAINT randoku_images_master_book_id_fkey;

ALTER TABLE randoku_memos
DROP CONSTRAINT randoku_memos_master_book_id_fkey;

ALTER TABLE seidoku_memos
DROP CONSTRAINT seidoku_memos_master_book_id_fkey;

-- カラムの名前を変更し、新しい外部キー制約を追加
ALTER TABLE randoku_images
RENAME COLUMN master_book_id TO user_book_logs_id;

ALTER TABLE randoku_memos
RENAME COLUMN master_book_id TO user_book_logs_id;

ALTER TABLE seidoku_memos
RENAME COLUMN master_book_id TO user_book_logs_id;

-- 新しい外部キー制約を追加
ALTER TABLE randoku_images
ADD CONSTRAINT randoku_images_user_book_logs_id_fkey
FOREIGN KEY (user_book_logs_id) REFERENCES user_book_logs(id);

ALTER TABLE randoku_memos
ADD CONSTRAINT randoku_memos_user_book_logs_id_fkey
FOREIGN KEY (user_book_logs_id) REFERENCES user_book_logs(id);

ALTER TABLE seidoku_memos
ADD CONSTRAINT seidoku_memos_user_book_logs_id_fkey
FOREIGN KEY (user_book_logs_id) REFERENCES user_book_logs(id);

-- user_book_logsのstatusカラムをINTEGERからreading_typeに変更
-- Step 1: DEFAULT値を削除
ALTER TABLE user_book_logs
ALTER COLUMN status DROP DEFAULT;

-- Step 2: 型変更（既存データの変換）
ALTER TABLE user_book_logs
ALTER COLUMN status TYPE reading_type USING
  CASE status
    WHEN 0 THEN 'smooth_randoku'::reading_type
    WHEN 1 THEN 'slowly_seidoku'::reading_type
    WHEN 2 THEN 'smooth_tudoku'::reading_type
  END;

-- Step 3: 新しいDEFAULT値を設定
ALTER TABLE user_book_logs
ALTER COLUMN status SET DEFAULT 'smooth_randoku'::reading_type;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- 外部キー制約を削除
ALTER TABLE randoku_images
DROP CONSTRAINT randoku_images_user_book_logs_id_fkey;

ALTER TABLE randoku_memos
DROP CONSTRAINT randoku_memos_user_book_logs_id_fkey;

ALTER TABLE seidoku_memos
DROP CONSTRAINT seidoku_memos_user_book_logs_id_fkey;

-- カラムの名前を元に戻し、元の外部キー制約を追加
ALTER TABLE randoku_images
RENAME COLUMN user_book_logs_id TO master_book_id;

ALTER TABLE randoku_memos
RENAME COLUMN user_book_logs_id TO master_book_id;

ALTER TABLE seidoku_memos
RENAME COLUMN user_book_logs_id TO master_book_id;

-- 元の外部キー制約を追加
ALTER TABLE randoku_images
ADD CONSTRAINT randoku_images_master_book_id_fkey
FOREIGN KEY (master_book_id) REFERENCES master_books(id);

ALTER TABLE randoku_memos
ADD CONSTRAINT randoku_memos_master_book_id_fkey
FOREIGN KEY (master_book_id) REFERENCES master_books(id);

ALTER TABLE seidoku_memos
ADD CONSTRAINT seidoku_memos_master_book_id_fkey
FOREIGN KEY (master_book_id) REFERENCES master_books(id);

-- user_book_logsのstatusカラムをreading_typeからINTEGERに戻す
-- Step 1: DEFAULT値を削除
ALTER TABLE user_book_logs
ALTER COLUMN status DROP DEFAULT;

-- Step 2: 型をINTEGERに戻す
ALTER TABLE user_book_logs 
ALTER COLUMN status TYPE INTEGER USING 
  CASE status::text
    WHEN 'smooth_randoku' THEN 0
    WHEN 'slowly_seidoku' THEN 1
    WHEN 'smooth_tudoku' THEN 2
    ELSE 0
  END;

-- Step 3: 元のDEFAULT値を設定
ALTER TABLE user_book_logs
ALTER COLUMN status SET DEFAULT 0;

-- ENUM型を削除
DROP TYPE IF EXISTS reading_type;
-- +goose StatementEnd
