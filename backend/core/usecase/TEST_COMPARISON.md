# UseCase層テスト手法の詳細比較ガイド

## 目次
1. [sqlmock版の詳細解説](#1-sqlmock版の詳細解説)
2. [インターフェースモック版の詳細解説](#2-インターフェースモック版の詳細解説)
3. [動作原理の違い](#3-動作原理の違い)
4. [実装パターンの比較](#4-実装パターンの比較)
5. [使い分けガイド](#5-使い分けガイド)

---

## 1. sqlmock版の詳細解説

### 1.1 sqlmockとは何か？

**sqlmock**は、Goのdatabase/sqlパッケージのモックライブラリです。実際のデータベースに接続せずに、SQLクエリの実行をシミュレートします。

```go
// sqlmockの基本的な仕組み
db, mock, err := sqlmock.New()  // モックDBを作成
mock.ExpectExec("INSERT INTO...").WithArgs(...).WillReturnResult(...)  // 期待するSQLを定義
// アプリケーションコードがSQLを実行
mock.ExpectationsWereMet()  // 期待通りのSQLが実行されたか検証
```

### 1.2 sqlmockの動作原理

```
[テストコード]
    ↓ (1) モックDB作成
[sqlmock]
    ↓ (2) 期待するSQL登録
[期待値リスト]
    INSERT INTO "publishers"...
    SELECT * FROM "publishers"...
    
[アプリケーション実行]
    ↓ (3) SQL実行
[sqlmock]
    ↓ (4) 実際のSQLと期待値を比較
[検証結果]
```

### 1.3 sqlmockテストの詳細実装解説

```go
func TestMockCreatePublisher(t *testing.T) {
    // ===== Step 1: テストデータの準備 =====
    const (
        TestName = "テスト出版社"  // 固定値（動的値は使わない）
        TestID   = int64(1)       // DBが返すIDを固定
    )

    vectors := map[string]struct {
        // ... テストケースの定義 ...
        prepare func(mock sqlmock.Sqlmock)  // ← ここがsqlmockの核心
    }{
        "OK": {
            prepare: func(mock sqlmock.Sqlmock) {
                // ===== Step 2: 期待するSQLを正確に定義 =====
                
                // Bob ORMが生成する実際のSQL形式を記述
                // 注意点：
                // - エスケープが必要（\( \) など）
                // - カラム名の順序も正確に
                // - DEFAULTキーワードも含める
                insertQuery := `INSERT INTO "publishers" AS "publishers"\("name", "created_at", "updated_at"\) VALUES \(\$1, DEFAULT, DEFAULT\) RETURNING`
                
                // ===== Step 3: SQLの引数と戻り値を定義 =====
                mock.ExpectExec(insertQuery).
                    WithArgs(TestName).              // $1 = "テスト出版社"
                    WillReturnResult(
                        sqlmock.NewResult(TestID, 1) // LastInsertId=1, RowsAffected=1
                    )
            },
        },
    }

    // ===== Step 4: モックDBのセットアップ =====
    db, mock, err := sqlmock.New()
    
    // ===== Step 5: 期待値の設定 =====
    v.prepare(mock)  // 上で定義したSQLの期待値を登録
    
    // ===== Step 6: アプリケーションコードの実行 =====
    // ここでBob ORMがSQLを生成・実行
    // sqlmockは実際のSQLと期待値を比較
    
    // ===== Step 7: 検証 =====
    if err := mock.ExpectationsWereMet(); err != nil {
        // 期待したSQLと実際のSQLが異なる場合エラー
        t.Errorf("unfulfilled expectations: %v", err)
    }
}
```

### 1.4 sqlmockの落とし穴と注意点

#### ① SQL形式の完全一致が必要
```go
// ❌ 失敗例：スペースが異なる
mock.ExpectExec("INSERT INTO publishers")  // スペース1つ
// 実際：      "INSERT INTO  publishers"  // スペース2つ

// ❌ 失敗例：エスケープ忘れ
mock.ExpectExec("INSERT INTO publishers (name)")  // カッコのエスケープなし
// 実際：      "INSERT INTO publishers \(name\)"  // Bob ORMはエスケープする
```

#### ② 引数の型と順序
```go
// ❌ 失敗例：型が異なる
mock.ExpectExec(query).WithArgs("1")  // 文字列
// 実際のコード: WithArgs(1)          // 数値

// ✅ 正しい例：型を正確に
mock.ExpectExec(query).WithArgs(int64(1))
```

#### ③ Bob ORM特有のSQL形式
```go
// Bob ORMは独特なSQL形式を生成
// AS句、DEFAULT値、RETURNINGなど
`INSERT INTO "publishers" AS "publishers"\("name", "created_at", "updated_at"\) VALUES \(\$1, DEFAULT, DEFAULT\) RETURNING`

// 通常のSQLとは異なるので注意
```

---

## 2. インターフェースモック版の詳細解説

### 2.1 インターフェースモックとは何か？

**インターフェースモック**は、Goのインターフェースを利用して、実装を差し替え可能にする手法です。SQLを一切書かず、メソッドの振る舞いのみを定義します。

```go
// インターフェースモックの基本的な仕組み
type MockRepository struct {
    CreateFunc func(ctx context.Context, model *Model) error
}

// テスト時に振る舞いを定義
mock := &MockRepository{
    CreateFunc: func(ctx context.Context, model *Model) error {
        // テスト用の振る舞いを定義
        model.ID = 1  // 固定値を設定
        return nil     // 成功を返す
    },
}
```

### 2.2 インターフェースモックの動作原理

```
[インターフェース定義]
    type Publisher interface {
        Create(ctx, model) error
    }
    
[本番実装]                    [モック実装]
    SQLを実行                   振る舞いのみ定義
    DBに接続                    メモリ上で動作
    
[UseCase層]
    interfaceに依存（実装に依存しない）
    テスト時はモック、本番は実装を注入
```

### 2.3 インターフェースモックテストの詳細実装解説

```go
// ===== Step 1: モック構造体の定義 =====
type mockPublisherRepository struct {
    // 各メソッドの振る舞いを関数として保持
    CreateFunc func(ctx context.Context, publisher *model.Publisher) error
    // なぜ関数フィールド？
    // → テストケースごとに異なる振る舞いを定義できるため
}

// ===== Step 2: インターフェースの実装 =====
func (m *mockPublisherRepository) Create(ctx context.Context, publisher *model.Publisher) error {
    if m.CreateFunc != nil {
        return m.CreateFunc(ctx, publisher)  // 定義された振る舞いを実行
    }
    return nil  // デフォルトの振る舞い
}

func TestMockCreatePublisher_WithInterfaceMock(t *testing.T) {
    vectors := map[string]struct {
        // ...
        setupMock func() (*mockPublisherQuery, *mockPublisherRepository)
    }{
        "OK": {
            setupMock: func() (*mockPublisherQuery, *mockPublisherRepository) {
                // ===== Step 3: テスト用の振る舞いを定義 =====
                mockRepo := &mockPublisherRepository{
                    CreateFunc: func(ctx context.Context, publisher *model.Publisher) error {
                        // ===== ビジネスロジックの検証 =====
                        
                        // 入力値の検証（SQLではなくビジネスルール）
                        if publisher.Name != TestName {
                            t.Errorf("unexpected name: got %s, want %s", 
                                publisher.Name, TestName)
                        }
                        
                        // DBの振る舞いをシミュレート
                        // （SQLは書かない、振る舞いのみ）
                        publisher.ID = TestID  // DBの自動採番を模倣
                        
                        return nil  // 成功
                    },
                }
                return mockQuery, mockRepo
            },
        },
        "RepositoryError": {
            setupMock: func() (*mockPublisherQuery, *mockPublisherRepository) {
                mockRepo := &mockPublisherRepository{
                    CreateFunc: func(ctx context.Context, publisher *model.Publisher) error {
                        // エラーケースの振る舞い
                        return errors.New("repository error")
                    },
                }
                return mockQuery, mockRepo
            },
        },
    }

    // ===== Step 4: モックを使ったテスト実行 =====
    mockQuery, mockRepo := v.setupMock()
    usecase := NewPublisher(mockQuery, mockRepo)  // モックを注入
    
    // ===== Step 5: ビジネスロジックの実行 =====
    actual, err := usecase.Create(context.Background(), v.params)
    
    // ===== Step 6: 結果の検証（SQLではなく出力値） =====
    if diff := cmp.Diff(v.expected, actual); diff != "" {
        t.Errorf("output mismatch (-want +got):\n%s", diff)
    }
}
```

### 2.4 インターフェースモックの利点

#### ① シンプルで理解しやすい
```go
// SQLを知らなくても理解できる
CreateFunc: func(ctx context.Context, publisher *model.Publisher) error {
    publisher.ID = 1  // 明確：IDを1に設定
    return nil        // 明確：成功を返す
}

// vs sqlmock（SQLの知識が必要）
mock.ExpectExec(`INSERT INTO "publishers"...複雑なSQL...`).
    WithArgs(...).WillReturnResult(...)
```

#### ② テストケースごとの振る舞い制御
```go
// 成功ケース
CreateFunc: func(...) error {
    return nil
}

// バリデーションエラー
CreateFunc: func(...) error {
    return ErrValidation
}

// DB接続エラー
CreateFunc: func(...) error {
    return ErrDBConnection
}
```

#### ③ ビジネスロジックに集中
```go
// ビジネスルールのテスト
CreateFunc: func(ctx context.Context, publisher *model.Publisher) error {
    // ビジネスルール：名前は必須
    if publisher.Name == "" {
        return ErrNameRequired
    }
    
    // ビジネスルール：名前の長さ制限
    if len(publisher.Name) > 100 {
        return ErrNameTooLong
    }
    
    publisher.ID = 1
    return nil
}
```

---

## 3. 動作原理の違い

### 3.1 sqlmock版の動作フロー

```
[テスト開始]
    ↓
[sqlmock.New()] → モックDB作成
    ↓
[mock.ExpectExec("INSERT...")] → 期待するSQL登録
    ↓
[usecase.Create()] 実行
    ↓
[Bob ORM] → SQL生成
    ↓
[db.Exec()] → sqlmockがSQL受信
    ↓
[sqlmock] → 期待値と比較
    ↓
    一致 → 定義済みの結果を返す
    不一致 → エラー
    ↓
[mock.ExpectationsWereMet()] → 全期待値の確認
```

### 3.2 インターフェースモック版の動作フロー

```
[テスト開始]
    ↓
[モック構造体作成] → 振る舞いを定義
    ↓
[usecase.Create()] 実行
    ↓
[repository.Create()] 呼び出し
    ↓
[モックのCreateFunc] 実行
    ↓
    定義された振る舞いを実行
    （SQL実行なし、DB接続なし）
    ↓
[結果を返す]
```

### 3.3 決定的な違い

| 観点 | sqlmock | インターフェースモック |
|------|---------|----------------------|
| **テスト対象** | SQL + ビジネスロジック | ビジネスロジックのみ |
| **依存関係** | Bob ORM、SQL形式に依存 | インターフェースのみ依存 |
| **実行速度** | 遅い（SQL解析あり） | 速い（メモリ上のみ） |
| **保守性** | 低い（SQL変更で修正必要） | 高い（実装変更の影響なし） |
| **学習コスト** | 高い（SQL知識必要） | 低い（Go知識のみ） |
| **エラー検出** | SQL構文エラーも検出 | ビジネスロジックエラーのみ |

---

## 4. 実装パターンの比較

### 4.1 Create操作のテスト

#### sqlmock版
```go
// SQL文字列を正確に記述する必要
prepare: func(mock sqlmock.Sqlmock) {
    // 1. INSERT文の期待値
    insertQuery := `INSERT INTO "publishers"...`  // 長く複雑
    mock.ExpectExec(insertQuery).
        WithArgs("テスト出版社").
        WillReturnResult(sqlmock.NewResult(1, 1))
    
    // 2. SELECT文の期待値（Bob ORMがRETURNINGを使わない場合）
    selectQuery := `SELECT * FROM "publishers" WHERE id = $1`
    rows := sqlmock.NewRows([]string{"id", "name"}).
        AddRow(1, "テスト出版社")
    mock.ExpectQuery(selectQuery).
        WithArgs(1).
        WillReturnRows(rows)
}
```

#### インターフェースモック版
```go
// 振る舞いのみを定義
CreateFunc: func(ctx context.Context, publisher *model.Publisher) error {
    // シンプルで明確
    publisher.ID = 1
    return nil
}
```

### 4.2 エラーハンドリングのテスト

#### sqlmock版
```go
// DB接続エラーをシミュレート
prepare: func(mock sqlmock.Sqlmock) {
    mock.ExpectExec(insertQuery).
        WithArgs(TestName).
        WillReturnError(sql.ErrConnDone)  // DB特有のエラー
}
```

#### インターフェースモック版
```go
// ビジネスエラーを返す
CreateFunc: func(ctx context.Context, publisher *model.Publisher) error {
    return ErrDuplicateName  // ビジネスルールのエラー
}
```

### 4.3 バリデーションのテスト

#### sqlmock版
```go
// バリデーションエラーの場合、SQLは実行されない
prepare: func(mock sqlmock.Sqlmock) {
    // 何も期待値を設定しない
    // → SQLが実行されるとエラーになる
}
```

#### インターフェースモック版
```go
// メソッドが呼ばれないことを検証
CreateFunc: func(ctx context.Context, publisher *model.Publisher) error {
    t.Error("CreateFunc should not be called")
    return nil
}
```

---

## 5. 使い分けガイド

### 5.1 テストピラミッドでの位置づけ

```
        △
       / \     E2Eテスト（実DB）
      /   \    統合テスト（sqlmock可）
     /     \   
    /       \  単体テスト
   /_________\ （インターフェースモック推奨）
   
   多 ← 数 → 少
   速 ← 実行速度 → 遅
   低 ← コスト → 高
```

### 5.2 選択フローチャート

```
Q: 何をテストしたい？
├─ ビジネスロジック → インターフェースモック
├─ SQL文の正確性 → sqlmock
├─ DB接続・トランザクション → sqlmock or 実DB
└─ 全体の動作 → 実DB（統合テスト）
```

### 5.3 具体的な使い分け例

#### UseCase層のテスト → **インターフェースモック**
```go
// 理由：ビジネスロジックに集中すべき
// - 入力値のバリデーション
// - ビジネスルールの適用
// - エラーハンドリング
// - 出力値の変換
```

#### Infra層のテスト → **sqlmock**
```go
// 理由：SQL実装の正確性を検証
// - Bob ORMの使い方
// - SQL文の構築
// - トランザクション処理
// - DB特有のエラー処理
```

#### 統合テスト → **実DB or TestContainers**
```go
// 理由：実際の動作を確認
// - パフォーマンス
// - 同時実行制御
// - 外部キー制約
// - インデックスの効果
```

### 5.4 移行戦略

現在sqlmockを使っている場合の段階的移行：

```
Phase 1: 理解フェーズ（1-2週間）
├─ sqlmockで基本を理解
└─ Bob ORMの動作を学習

Phase 2: 並行フェーズ（2-4週間）
├─ 新規テストはインターフェースモック
└─ 重要な既存テストにモック版追加

Phase 3: 移行フェーズ（1-2ヶ月）
├─ UseCase層を順次モック化
└─ Infra層テストを新規作成

Phase 4: 完成フェーズ
├─ UseCase層：インターフェースモック
├─ Infra層：sqlmock
└─ 統合：実DB
```

---

## まとめ

### sqlmockを選ぶべき時
- SQL文の正確性を検証したい
- Bob ORMの使い方を学習中
- Infra層の実装をテスト
- DB固有の機能をテスト

### インターフェースモックを選ぶべき時
- ビジネスロジックをテスト
- 高速なテストが必要
- 保守性を重視
- チーム開発で認識を統一

### 重要な原則
1. **単一責任の原則**: 各層は自分の責任のみテスト
2. **依存性逆転の原則**: 上位層は下位層の実装に依存しない
3. **テストの独立性**: テスト間で状態を共有しない
4. **明確性**: 何をテストしているか一目で分かる

最終的に、両手法を理解し、適切に使い分けることが重要です。