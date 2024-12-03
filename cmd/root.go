/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd サブコマンドなしで呼び出された場合の基本コマンドを表します
var rootCmd = &cobra.Command{
	Use:   "ws",
	Short: "動作チェック用Webサーバーを作成します",
	Long:  `WebSocketを用いた通信を実装するときの検証用サーバーとして使用できます`,
	// ベアアプリケーションの場合は、次の行のコメントを解除します。
	// Run: func(cmd *cobra.Command, args []string) { },
	//	ベアアプリケーションはベアメタルサーバーで動かすアプリケーションで、OSを使わずに動かすアプリケーション
}

// Execute すべての子コマンドをルート コマンドに追加し、フラグを適切に設定します。
// これは main.main() によって呼び出されます。これは rootCmd に対して 1 回だけ実行する必要があります。
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}
