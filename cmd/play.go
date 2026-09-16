/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os/exec"

	utils "github.com/anjanragh/bento/internal"
	"github.com/spf13/cobra"
)

// playCmd represents the play command
var playCmd = &cobra.Command{
	Use:   "play",
	Short: "Lets you play a song.",
	Long:  `bento play <path_to_audio_file>`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if !utils.FileExists(args[0]) {
			fmt.Println("Oh no, the file doesn't exist")
			return
		}
		fmt.Println("play called on file ", args[0])

		if _, err := exec.LookPath("mpv"); err != nil {
			fmt.Println("mpv not installed! Please install it using `brew install mpv`")
			return
		}
		fmt.Println("mpv exists!")
	},
}

func init() {
	rootCmd.AddCommand(playCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// playCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// playCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
