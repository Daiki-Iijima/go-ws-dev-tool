/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"net"

	"github.com/spf13/cobra"
)

func LocalIP() (net.IP, error) {
	ift, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	for _, ifi := range ift {
		addrs, err := ifi.Addrs()
		if err != nil {
			return nil, err
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if isPrivateIP(ip) {
				return ip, nil
			}
		}
	}

	return nil, errors.New("no IP")
}

func isPrivateIP(ip net.IP) bool {
	var prvMasks []*net.IPNet

	for _, cidr := range []string{
		"10.0.0.0/8",
		"172.16.0.0/12",
		"192.168.0.0/16",
	} {
		_, mask, _ := net.ParseCIDR(cidr)
		prvMasks = append(prvMasks, mask)
	}

	for _, mask := range prvMasks {
		if mask.Contains(ip) {
			return true
		}
	}
	return false
}

// ipCmd represents the ip command
var ipCmd = &cobra.Command{
	Use:   "ip",
	Short: "IPアドレスを出力",
	Long:  `IPアドレスを出力`,
	Run: func(cmd *cobra.Command, args []string) {
		ip, err := LocalIP()
		if err != nil {
			fmt.Println("IPの取得に失敗しました")
			return
		}
		fmt.Printf("IP : %s\n", ip)
	},
}

func init() {
	rootCmd.AddCommand(ipCmd)
}
