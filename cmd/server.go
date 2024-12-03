/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/spf13/cobra"
	"net/http"
	"strconv"
	"time"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

type Message struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	//	初期のリクエストをWebSocketにアップグレード
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer ws.Close()

	for {
		//	ブラウザからのメッセージを読み込む
		_, msg, err := ws.ReadMessage()
		if err != nil {
			fmt.Println("読み込みに失敗しました:", err)
			break
		}

		// 現在の日時を取得
		currentTime := time.Now()
		formattedDateTime := currentTime.Format("01-02 15:04:05")
		fmt.Printf("%s : [受信] : %s\n", formattedDateTime, msg)

		if err := ws.WriteMessage(websocket.TextMessage, msg); err != nil {
			fmt.Println("書き込みに失敗しました:", err)
			break
		}
	}
}

var port int
var useBroadcast bool

var isConnected bool

// wssCmd represents the wss command
var wssCmd = &cobra.Command{
	Use:   "server",
	Short: "WebSocketサーバーを起動",
	Long: `WebSocketサーバーを現在のIPアドレスを使用して起動します。
終了したい場合は、Ctrl + cを実行してください`,
	Run: func(cmd *cobra.Command, args []string) {

		http.HandleFunc("/ws", handleConnections)

		ip, ipErr := LocalIP()
		if ipErr != nil {
			fmt.Println("IPアドレスを取得できませんでした")
			return
		}

		if useBroadcast {
			fmt.Println("ブロードキャスト開始")
			go Broadcast()
		}

		fmt.Printf("WebSocketサーバー起動 [ws://%s:%s/ws]\n", ip, strconv.Itoa(port))

		err := http.ListenAndServe(":"+strconv.Itoa(port), nil)
		if err != nil {
			fmt.Println("サーバー起動エラー:", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(wssCmd)
	wssCmd.Flags().IntVarP(&port, "port", "p", 8080, "ポート番号を指定する")
	wssCmd.Flags().BoolVarP(&useBroadcast, "broadcast", "b", false, "ブロードキャストを実行する")
}
