/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "バージョン出力",
	Long:  `このアプリのバージョンが出力されます`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("v0.0.3")
	},
}

// インポート時に呼び出される
// UnityのStartみたいな役割の規定関数
func init() {
	rootCmd.AddCommand(versionCmd)
}
