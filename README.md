# TODO App

## 概要

このプロジェクトは、Next.jsをフロントエンド、Golangをバックエンドに使用したTODOアプリです。データベースにはPostgreSQLを使用しています。

## 必要なソフトウェア

- [Node.js](https://nodejs.org/)（フロントエンド）
- [Go](https://golang.org/)（バックエンド）
- [PostgreSQL](https://www.postgresql.org/)（データベース）

## インストール手順

### PostgreSQLのインストール

1. **PostgreSQLのダウンロード**

   PostgreSQLの公式ウェブサイトからインストーラーをダウンロードします: [PostgreSQL Downloads](https://www.postgresql.org/download/)

2. **インストーラーの実行**

   ダウンロードしたインストーラーを実行し、指示に従ってインストールを進めます。インストール中に、以下の設定を行います:
   - **パスワードの設定**: デフォルトのスーパーユーザー（postgres）用のパスワードを設定します。
   - **ポート番号**: デフォルトでポート5432が使用されます。特別な理由がない限り、そのままで問題ありません。

3. **PostgreSQLの確認**

   インストールが完了したら、PostgreSQLサーバーが実行中であることを確認します。

   - **Windows**:
     ```powershell
     Get-Service -Name postgresql*
     ```

   - **Unix系システム（Linux、macOS）**:
     ```bash
     sudo service postgresql status
     ```

データベース名やパスワードは、config/.env ファイルで設定した値を使用します。

## ディレクトリ構成

- `backend/` - バックエンドのコード
- `frontend/` - フロントエンドのコード

## セットアップ

### 環境設定

1. プロジェクトのルートに`.env`ファイルを作成し、以下の設定を追加します。

    ```env
    DB_USER=postgres
    DB_PASSWORD=yourpassword
    DB_ADDR=localhost:5432
    DB_NAME=yourdatabase
    PORT=8080
    ```

2. 必要な依存関係をインストールします。

    ```bash
    # フロントエンドの依存関係をインストール
    cd frontend
    npm install

    # バックエンドの依存関係をインストール
    cd ../backend
    go mod tidy
    ```

### 実行

1. バックエンドサーバーを起動します。

    ```bash
    cd backend
    go run cmd/server/main.go
    ```
バックエンドは http://localhost:8080 でアクセスできます。
2. フロントエンドサーバーを起動します。

    ```bash
    cd ../frontend
    npm run dev
    ```
これで、フロントエンドは http://localhost:3000 でアクセスできます。

## APIエンドポイント

- `POST /todos` - TODOを作成
- `GET /todos` - TODOの一覧を取得
- `PUT /todos/{id}` - TODOを更新
- `DELETE /todos/{id}` - TODOを削除

## データベース

このプロジェクトではPostgreSQLを使用しています。データベースの設定は、`.env`ファイルで行います。接続設定を変更するには、`.env`ファイルの内容を修正してください。
