/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var importFilename string

// importCmd represents the import command
var importCmd = &cobra.Command{
	Use:   "import",
	Short: "Lets you import a song from youtube",
	Long:  `bento import <url>`,
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if _, err := exec.LookPath("yt-dlp"); err != nil {
			return fmt.Errorf("yt-dlp not installed! Please install it using `brew install yt-dlp`")
		}

		ytArgs := []string{
			"--ignore-config",
			"--no-playlist",
			"--format", "bestaudio",
			"--extract-audio",
			"--audio-format", "best",
			"--embed-metadata",
			"--embed-thumbnail",
			"--print", "after_move:filepath",
		}

		outputTemplate := "%(title)s [%(id)s].%(ext)s"

		if importFilename != "" {
			outputTemplate = importFilename + ".%(ext)s"
		}

		ytArgs = append(ytArgs, "--output", outputTemplate)

		ytArgs = append(ytArgs, "--", args[0])
		downloader := exec.Command("yt-dlp", ytArgs...)

		downloader.Stdout = os.Stdout
		downloader.Stderr = os.Stderr

		if err := downloader.Run(); err != nil {
			return fmt.Errorf("import failed: %w", err)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(importCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// importCmd.PersistentFlags().String("foo", "", "A help for foo")
	importCmd.Flags().StringVar(&importFilename, "filename", "", "Output filename without an extension")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// importCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
