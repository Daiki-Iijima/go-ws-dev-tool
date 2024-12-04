/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/spf13/cobra"
	"net/url"
)

var wsClientIp string
var wsClientPort string
var wsPath string

// clientCmd represents the client command
var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "WebSocketクライアントを起動します",
	Long: `WebSocketクライアントを起動します。
正常に起動させるためには、オプションを設定してください`,
	Run: func(cmd *cobra.Command, args []string) {
		u := url.URL{Scheme: "ws", Host: wsClientIp + ":" + wsClientPort, Path: wsPath}
		fmt.Printf("接続開始 : %s\n", u.String())
		c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
		if err != nil {
			fmt.Println("接続に失敗しました:", err)
			return
		}

		defer c.Close()

		fmt.Println("送信したい文章を入力してEnterを押してください")

		for {
			var str string
			_, err := fmt.Scan(&str)
			if err != nil {
				return
			}
			if str == "" {
				continue
			}

			err = c.WriteMessage(websocket.TextMessage, []byte(str))
			if err != nil {
				fmt.Println("書き込みエラー:", err)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(clientCmd)

	clientCmd.Flags().StringVarP(&wsClientIp, "ip", "i", "127.0.0.1", "接続先のIPを指定します")
	clientCmd.Flags().StringVarP(&wsClientPort, "port", "p", "8079", "接続先のPortを指定します")
	clientCmd.Flags().StringVarP(&wsPath, "url", "u", "/ws", "接続先のパス/URLを指定します")
}
