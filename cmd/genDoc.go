/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
	"os"
)

// genDocCmd represents the genDoc command
var genDocCmd = &cobra.Command{
	Use:   "genDoc",
	Short: "ドキュメントを生成します",
	Long:  `APIドキュメントを出力します`,
	Run: func(cmd *cobra.Command, args []string) {
		// ドキュメントを保存するディレクトリ
		outputDir := "./docs"

		// ディレクトリが存在しない場合は作成
		if _, err := os.Stat(outputDir); os.IsNotExist(err) {
			err := os.Mkdir(outputDir, os.ModePerm)
			if err != nil {
				panic("ディレクトリの作成に失敗: " + err.Error())
			}
		}
		err := doc.GenMarkdownTree(rootCmd, outputDir)
		if err != nil {
			fmt.Printf("ドキュメントの生成に失敗しました%s", err)
			os.Exit(1)
		}
		fmt.Println("./docsにドキュメントを生成しました")
	},
}

func init() {
	rootCmd.AddCommand(genDocCmd)
}
