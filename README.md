# WebSocketテストツール

## このツールについて

WebSocketを用いた通信を実装するときの検証用サーバーとして使用できます

## 準備

Windowsの場合は、特に準備は必要なく、実行ファイルをPowerShellやコマンドプロンプトから実行することができます。

### Mac

Macの場合は、実行ファイルに実行権限を付与する必要があります。
ターミナルを開いて、バイナリを配置したディレクトリに移動して以下のコマンドを実行してください。

```sh
chmod +x ws
```

## 使用方法

詳しいコマンドなどは、`-hフラグ`をつけてサブコマンドを実行して確認してください。

面倒くさい方は、以下のリンクに`-h`フラグをつけたコマンドの説明があります。

[コマンド詳細](https://github.com/Daiki-Iijima/go-ws-dev-tool/docs/ws.md)

## 使用例

### WebSocketサーバーを起動する

`7777`ポートを指定してWebSocketサーバーを起動する(指定しない場合のデフォルトポートは、`8080`です)

```sh
./ws server -p 7777
```

### WebSocketクライアントを起動する

`IP`、`ポート`、`URL`を指定してWebSocketクライアントで接続します。

`ws://192.168.2.12:8080/aaa`に接続します。

```sh
./ws client -i 192.168.2.12 -p 8080 -u aaa
```

接続後は、メッセージを入力して`Enter`を押すと、サーバーにメッセージが送信されます。
