# Todo CRUD API

GinフレームワークとGORMを使用したシンプルなTodo管理REST APIです。

## 🚀 機能

- ✅ Todo作成 (CREATE)
- ✅ Todo一覧取得 (READ)
- ✅ Todo更新 (UPDATE)
- ✅ Todo削除 (DELETE)
- ✅ CORS対応 (フロントエンド連携)
- ✅ SQLiteデータベース
- ✅ 自動マイグレーション

## 🛠 技術スタック

- **Language**: Go
- **Framework**: Gin
- **ORM**: GORM
- **Database**: SQLite
- **CORS**: gin-contrib/cors

## 📋 前提条件

- Go 1.19以上
- Git

## ⚡ クイックスタート

### 1. リポジトリのクローン

```bash
git clone https://github.com/yourusername/todo-crud-api.git
cd todo-crud-api
```

### 2. 依存関係のインストール

```bash
go mod init todo-crud-api
go get github.com/gin-gonic/gin
go get github.com/gin-contrib/cors
go get gorm.io/gorm
go get gorm.io/driver/sqlite
```

### 3. アプリケーションの起動

```bash
go run main.go
```

サーバーは `http://localhost:8080` で起動します。

## 📊 API エンドポイント

### Base URL
```
http://localhost:8080
```

### エンドポイント一覧

| Method | Endpoint | Description | Request Body |
|--------|----------|-------------|--------------|
| POST | `/todos` | 新しいTodoを作成 | `{"title": "string", "completed": boolean}` |
| GET | `/todos` | 全てのTodoを取得 | - |
| PUT | `/todos/:id` | 指定IDのTodoを更新 | `{"title": "string", "completed": boolean}` |
| DELETE | `/todos/:id` | 指定IDのTodoを削除 | - |

## 🔍 データモデル

```json
{
  "id": 1,
  "title": "サンプルタスク",
  "completed": false
}
```

## 📝 API使用例

### Todo作成

```bash
curl -X POST http://localhost:8080/todos \
  -H "Content-Type: application/json" \
  -d '{"title":"買い物リストを作成","completed":false}'
```

**レスポンス:**
```json
{
  "id": 1,
  "title": "買い物リストを作成",
  "completed": false
}
```

### Todo一覧取得

```bash
curl http://localhost:8080/todos
```

**レスポンス:**
```json
[
  {
    "id": 1,
    "title": "買い物リストを作成",
    "completed": false
  },
  {
    "id": 2,
    "title": "プレゼンテーション準備",
    "completed": true
  }
]
```

### Todo更新

```bash
curl -X PUT http://localhost:8080/todos/1 \
  -H "Content-Type: application/json" \
  -d '{"title":"買い物リストを作成","completed":true}'
```

**レスポンス:**
```json
{
  "id": 1,
  "title": "買い物リストを作成",
  "completed": true
}
```

### Todo削除

```bash
curl -X DELETE http://localhost:8080/todos/1
```

**レスポンス:**
```json
{
  "message": "Todo deleted successfully"
}
```

## 🔧 設定

### CORS設定

現在の設定では `http://localhost:3000` からのアクセスを許可しています。
フロントエンドのURL変更時は、`main.go`の以下の部分を修正してください：

```go
AllowOrigins: []string{"http://localhost:3000"},
```

### データベース

- SQLiteファイル: `todos.db`
- 自動マイグレーション対応
- 初回起動時に自動でテーブル作成

## 🧪 テスト

### 手動テスト

```bash
# サーバー起動
go run main.go

# 別のターミナルでテスト実行
curl -X POST http://localhost:8080/todos \
  -H "Content-Type: application/json" \
  -d '{"title":"テストタスク","completed":false}'

curl http://localhost:8080/todos
```

## 📁 プロジェクト構造

```
.
├── main.go          # メインアプリケーションファイル
├── todos.db         # SQLiteデータベースファイル（自動生成）
├── go.mod           # Go モジュールファイル
├── go.sum           # 依存関係チェックサムファイル
└── README.md        # このファイル
```

## 🚀 デプロイ

### Docker対応（オプション）

```dockerfile
FROM golang:1.19-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o main .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["./main"]
```

## 🤝 コントリビューション

1. このリポジトリをフォーク
2. フィーチャーブランチを作成 (`git checkout -b feature/amazing-feature`)
3. 変更をコミット (`git commit -m 'Add some amazing feature'`)
4. ブランチにプッシュ (`git push origin feature/amazing-feature`)
5. プルリクエストを作成

## 📄 ライセンス

このプロジェクトはMITライセンスの下で公開されています。

## ✨ 今後の改善予定

- [ ] ユーザー認証機能
- [ ] バリデーション強化
- [ ] ログ機能
- [ ] テストコード追加
- [ ] Docker対応
- [ ] 設定ファイル外部化
- [ ] ページネーション機能
