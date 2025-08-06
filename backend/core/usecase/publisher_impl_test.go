package usecase

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/core/domain/model"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/core/usecase/input"
	"github.com/sandonemaki/my_read_memo_Go_API/backend/core/usecase/output"
)

// ================================================================================
// sqlmock版テスト - SQL文を直接記述してテストする手法
// 
// 【sqlmockとは】
// - database/sqlパッケージのモックライブラリ
// - 実際のDBに接続せずSQLクエリの実行をシミュレート
// - 期待するSQLと実際のSQLを比較して検証
//
// 【メリット】
// - 実際のSQL文が正しく生成されることを確認できる
// - Bob ORMの挙動を詳細に検証できる
// - DB層まで含めた統合的なテストが可能
//
// 【デメリット】
// - UseCase層がSQL実装の詳細を知る必要がある（抽象化の漏れ）
// - SQL形式が変わるとテストも修正が必要（保守性が低い）
// - テストが複雑（SQLのエスケープや形式を正確に記述）
//
// 【いつ使うべきか】
// - Infra層のテスト
// - Bob ORMの使い方を学習したい場合
// - SQLの正確性を検証したい場合
// ================================================================================

func TestMockCreatePublisher(t *testing.T) {
	// ===== Step 1: テストデータの準備 =====
	// 固定値を使用することで、テストの再現性を保証
	// 動的値（time.Now()やuuid.New()など）は使用禁止
	const (
		TestName = "テスト出版社"
		TestID   = int64(1)
	)
	
	// 固定時刻（必要な場合に使用）
	fixedTime := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)

	// ===== Step 2: テストケースの定義 =====
	vectors := map[string]struct {
		params   input.CreatePublisher      // 入力パラメータ
		expected *output.CreatePublisher     // 期待する出力
		wantErr  error                       // 期待するエラー
		options  cmp.Options                 // 比較オプション（動的フィールドの無視など）
		prepare  func(mock sqlmock.Sqlmock)  // SQLの期待値を設定する関数（sqlmockの核心）
	}{
		// ===== 正常系テストケース =====
		"OK": {
			params: input.CreatePublisher{
				Name: TestName,
			},
			expected: &output.CreatePublisher{
				Publisher: &model.Publisher{
					ID:   TestID,
					Name: TestName,
				},
			},
			wantErr: nil,
			options: cmp.Options{
				// 時刻フィールドは無視（DB側で自動設定されるため）
				// sqlmockでは時刻の制御が難しいため、比較から除外
				cmpopts.IgnoreFields(output.CreatePublisher{}, "Publisher.CreatedAt", "Publisher.UpdatedAt"),
			},
			prepare: func(mock sqlmock.Sqlmock) {
				// ===== sqlmockの期待値設定 =====
				// Bob ORMが生成する実際のSQL形式を正確に記述する必要がある
				// 
				// 【重要な注意点】
				// 1. エスケープが必要: \( \) \$ など
				// 2. AS句: Bob ORMは"publishers" AS "publishers"という形式を使う
				// 3. DEFAULT値: created_at, updated_atはDEFAULTキーワードを使用
				// 4. RETURNING句: PostgreSQLの機能（IDを返す）
				insertQuery := `INSERT INTO "publishers" AS "publishers"\("name", "created_at", "updated_at"\) VALUES \(\$1, DEFAULT, DEFAULT\) RETURNING`
				
				// ExpectExec: INSERT/UPDATE/DELETE文の期待値を設定
				mock.ExpectExec(insertQuery).
					WithArgs(TestName).                          // $1 = "テスト出版社"
					WillReturnResult(sqlmock.NewResult(TestID, 1)) // LastInsertId=1, RowsAffected=1
			},
		},
		// ===== バリデーションエラーのテストケース =====
		"EmptyName": {
			params: input.CreatePublisher{
				Name: "", // 空の名前（validation error）
			},
			expected: nil,
			wantErr:  errors.New("validation error"), // バリデーションエラーを期待
			options:  cmp.Options{},
			prepare: func(mock sqlmock.Sqlmock) {
				// バリデーションエラーの場合、SQLは実行されない
				// そのため、何も期待値を設定しない
				// もしSQLが実行されると、期待値が設定されていないためエラーになる
				// これにより、バリデーションが正しく機能していることを確認できる
			},
		},
		// ===== データベースエラーのテストケース =====
		"DatabaseError": {
			params: input.CreatePublisher{
				Name: TestName,
			},
			expected: nil,
			wantErr:  errors.New("database error"),
			options:  cmp.Options{},
			prepare: func(mock sqlmock.Sqlmock) {
				// DB層でのエラーをシミュレート
				// WillReturnErrorを使用してエラーを返す
				insertQuery := `INSERT INTO "publishers" AS "publishers"\("name", "created_at", "updated_at"\) VALUES \(\$1, DEFAULT, DEFAULT\) RETURNING`
				mock.ExpectExec(insertQuery).
					WithArgs(TestName).
					WillReturnError(errors.New("database error")) // DB接続エラーなどをシミュレート
			},
		},
	}

	// ===== Step 3: 各テストケースの実行 =====
	for k, v := range vectors {
		t.Run(k, func(t *testing.T) {
			// ===== Step 4: sqlmockのセットアップ =====
			// sqlmock.New()でモックDBを作成
			// これは実際のDBに接続しない仮想的なDB接続
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("sqlmock error: %v", err)
			}
			defer db.Close()

			// ===== Step 5: SQL期待値の設定 =====
			// prepareで定義したSQL期待値をmockに登録
			// この順序でSQLが実行されることを期待
			v.prepare(mock)

			// ===== Step 6: アプリケーションコードの実行 =====
			// TODO: 実際のテスト実行部分はPublisherDIが実装されてから追加
			// usecase := NewPublisherDI(db)  // Wire DIでusecaseを生成
			// actual, err := usecase.Create(context.Background(), v.params)
			//
			// ここでBob ORMがSQLを生成し、dbに対して実行
			// sqlmockは実際のSQLと期待値を比較
			
			// 現在はコンパイルエラーを防ぐためコメントアウト
			t.Skip("TODO: PublisherDI実装後に有効化")

			// ===== Step 7: 期待値の検証 =====
			// ExpectationsWereMet()で全ての期待値が満たされたか確認
			// - 期待したSQLが実行されたか
			// - 余分なSQLが実行されていないか
			// - 順序は正しいか
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("unfulfilled expectations: %v", err)
			}
		})
	}
}