# WebSocketテストツール（ws-dev-tool）

> WebSocket通信の開発・検証を手軽に行えるCLIツール

## 概要

WebSocketを用いた通信を実装するときの検証用サーバー・クライアントとして使用できるCLIツールです。
サーバー起動、クライアント接続、UDPブロードキャストなどの機能を1つのバイナリで提供します。
Linux / macOS / Windows 向けのビルド済みバイナリが提供されています。

## 技術スタック

- 言語: Go 1.23.1
- CLIフレームワーク: Cobra
- WebSocketライブラリ: gorilla/websocket
- ビルド・リリース: GoReleaser（linux/darwin/windows × amd64/arm64）

## 機能

- WebSocketサーバーの起動（ポート指定可能、デフォルト: 8080）
- WebSocketクライアントの起動（IP・ポート・パス指定可能）
- UDPブロードキャスト（自分のIPアドレスをLANへ定期配信）
- シェル補完スクリプトの生成（bash / zsh / fish / PowerShell）
- コマンドドキュメントの自動生成

## 使い方 / 動かし方

### インストール

[Releases](https://github.com/Daiki-Iijima/go-ws-dev-tool/releases) からお使いのOS・アーキテクチャに合ったバイナリをダウンロードしてください。

Macの場合は実行権限の付与が必要です。

```sh
chmod +x ws-dev-tool
```

### WebSocketサーバーを起動する

```sh
# デフォルトポート(8080)で起動
./ws-dev-tool server

# ポートを指定して起動
./ws-dev-tool server -p 7777

# UDPブロードキャストと同時に起動
./ws-dev-tool server -p 7777 --broadcast
```

### WebSocketクライアントを起動する

```sh
# ws://192.168.2.12:8080/aaa に接続
./ws-dev-tool client -i 192.168.2.12 -p 8080 -u aaa
```

接続後はメッセージを入力して `Enter` を押すと、サーバーにメッセージが送信されます。

### UDPブロードキャスト

```sh
# 自分のIPをLANへブロードキャスト（デフォルトポート: 12345、間隔: 2秒）
./ws-dev-tool broadcast -p 12345 -i 2
```

### ヘルプ

```sh
./ws-dev-tool -h
./ws-dev-tool server -h
./ws-dev-tool client -h
```

詳細は [コマンドドキュメント](https://github.com/Daiki-Iijima/go-ws-dev-tool/blob/main/docs/ws.md) を参照してください。

## 状態

完成・安定稼働中。WebSocket開発時の検証ツールとして実用に足る状態です。GoReleaserによるマルチプラットフォームリリースが整備されています。
