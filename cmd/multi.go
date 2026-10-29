/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"helthy/app"
	"sync"

	"github.com/spf13/cobra"
)

// multiCmd represents the multi command
var multiCmd = &cobra.Command{
	Use:   "multi",
	Short: "accepts multiple inputs from command line",
	Long:  `multi cmdline version`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args)
		i := 0
		var wg sync.WaitGroup
		for i < len(args) {
			wg.Add(1)
			go app.Basic(args[i], &wg)
			i++
		}
		wg.Wait()
	},
}

func init() {
	rootCmd.AddCommand(multiCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// multiCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// multiCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
