# TODO App

## 概要

このプロジェクトは、Next.jsをフロントエンド、Golangをバックエンドに使用したTODOアプリです。データベースにはPostgreSQLを使用しています。

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

2. フロントエンドサーバーを起動します。

    ```bash
    cd ../frontend
    npm run dev
    ```

## APIエンドポイント

- `POST /todos` - TODOを作成
- `GET /todos` - TODOの一覧を取得
- `PUT /todos/{id}` - TODOを更新
- `DELETE /todos/{id}` - TODOを削除

## データベース

このプロジェクトではPostgreSQLを使用しています。データベースの設定は、`.env`ファイルで行います。接続設定を変更するには、`.env`ファイルの内容を修正してください。
