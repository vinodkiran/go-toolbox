/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package info

import (
	"fmt"
	"github.com/ricochet2200/go-disk-usage/du"
	"github.com/spf13/cobra"
)

var KB = uint64(1024)

// diskUsageCmd represents the diskUsage command
var diskUsageCmd = &cobra.Command{
	Use:   "disk-usage",
	Short: "Prints the disk usage statistics of the current directory",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		usage := du.NewDiskUsage("/Users/vinod/code")
		fmt.Println("Disk Usage: ")
		fmt.Println("  Free:", usage.Free()/(KB*KB), "GB")
		fmt.Println("  Available:", usage.Available()/(KB*KB), "GB")
		fmt.Println("  Size:", usage.Size()/(KB*KB), "GB")
		fmt.Println("  Used:", usage.Used()/(KB*KB), "GB")
		fmt.Println("  Usage:", usage.Usage()*100, "%")
	},
}

func init() {
	InfoCmd.AddCommand(diskUsageCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// diskUsageCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// diskUsageCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
