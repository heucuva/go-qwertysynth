package command

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "qwertysynth",
	Short: "Qwertysynth is a silly little computer keyboard synth",
	Long:  `Qwertysynth is a silly little computer keyboard synth written entirely in Go`,
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Usage()
	},
}

func Execute() {
	cmd, args, err := rootCmd.Find(os.Args[1:])
	if err == nil && cmd == rootCmd {
		// assume play command if command argument not provided
		os.Args = append([]string{os.Args[0], "play"}, args...)
		cmd = playCmd
	}

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if err := cmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
