/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"

	utils "github.com/anjanragh/bento/internal"
	"github.com/spf13/cobra"
)

const socketPath = "/tmp/bento.sock"

// playCmd represents the play command
var playCmd = &cobra.Command{
	Use:   "play",
	Short: "Lets you play a song.",
	Long:  `bento play <path_to_audio_file>`,
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if !utils.FileExists(args[0]) {
			return fmt.Errorf("file does not exist")
		}
		fmt.Println("play called on file ", args[0])

		if _, err := exec.LookPath("mpv"); err != nil {
			return fmt.Errorf("mpv not installed! Please install it using `brew install mpv`")
		}
		fmt.Println("mpv exists!")

		if utils.SocketActive(socketPath) {
			return fmt.Errorf("bento is already playing something!")
		}

		if utils.FileExists(socketPath) {
			err := os.Remove(socketPath)
			if err != nil {
				return fmt.Errorf("file could not be removed: %w", err)
			}
		}

		player := exec.Command("mpv", "--no-video", "--input-terminal=no", "--input-ipc-server="+socketPath, "--", args[0])

		player.Stdout = os.Stdout
		player.Stderr = os.Stderr

		if err := player.Start(); err != nil {
			return fmt.Errorf("failed to start mpv: %w", err)
		}

		fmt.Println("Playing : ", args[0])

		if err := player.Wait(); err != nil {
			return fmt.Errorf("mpv stopped with err: %w ", err)
		}
		return nil
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
