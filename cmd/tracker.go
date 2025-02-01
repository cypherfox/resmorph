/*
Copyright © 2025 Lutz Behnke <lutz.behnke@gmx.de>
This file is part of {{ .appName }}
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// trackerCmd represents the tracker command
var trackerCmd = &cobra.Command{
	Use:   "tracker",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("tracker called")
	},
}

func init() {
	rootCmd.AddCommand(trackerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// trackerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// trackerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
