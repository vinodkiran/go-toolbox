/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package net

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/vinodkiran/go-toolbox/internal/net"
)

var (
	host string
)

var portScanCmd = &cobra.Command{
	Use:   "port-scan",
	Short: "Scan for open ports.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		net.ScanPorts(host)
	},
}

func init() {
	host = "localhost"
	portScanCmd.Flags().StringVarP(&host, "url", "u", "", "the host (url) to scan")
	if err := portScanCmd.MarkFlagRequired("url"); err != nil {
		fmt.Println(err)
	}

	NetCmd.AddCommand(portScanCmd)
}
