/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"net"
	"strconv"
	"time"
)

var udpPort int
var interval int

func Broadcast() {

	broadcastAddr := "255.255.255.255:" + strconv.Itoa(udpPort)

	ip, ipErr := LocalIP()
	if ipErr != nil {
		fmt.Println("IPアドレスの取得に失敗しました")
		return
	}

	fmt.Printf("ブロードキャストアドレス:%s\n", broadcastAddr)

	//	フォーマットを最適化
	message := fmt.Sprintf("%s", ip)

	udpAddr, err := net.ResolveUDPAddr("udp", broadcastAddr)
	if err != nil {
		fmt.Printf("UDPアドレスの作成に失敗しました%v\n", err)
		return
	}

	//	UDP接続を作成
	conn, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		fmt.Printf("UDPコネクションの作成に失敗しました%v\n", err)
		return
	}

	defer conn.Close()

	//fmt.Println("ブロードキャスト:", message)

	// 無限ループで送信
	for {
		// メッセージをブロードキャスト
		_, err = conn.Write([]byte(message))
		if err != nil {
			fmt.Printf("Error sending broadcast message: %v\n", err)
			return
		}

		// 指定間隔待機
		time.Sleep(time.Duration(interval) * time.Second)
	}
}

var udpBroadcastCmd = &cobra.Command{
	Use:   "broadcast",
	Short: "udpプロトコルで指定されたポートに向けてブロードキャストを行います",
	Long:  `udpプロトコルで指定されたポートに向けてブロードキャストを行います`,
	Run: func(cmd *cobra.Command, args []string) {
		Broadcast()
	},
}

func init() {
	rootCmd.AddCommand(udpBroadcastCmd)

	udpBroadcastCmd.Flags().IntVarP(&udpPort, "port", "p", 12345, "ブロードキャスト時に使用するポート番号を指定できます")
	udpBroadcastCmd.Flags().IntVarP(&interval, "interval", "i", 2, "ブロードキャストする間隔を設定できます(秒)")
}
