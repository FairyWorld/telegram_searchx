package run

import (
	"context"
	"github.com/fatih/color"
	"github.com/iyear/searchx/app/usr/run"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
)

var Cmd = &cobra.Command{
	Use:   "run",
	Short: "run the (user)bot",
	Run: func(cmd *cobra.Command, args []string) {
		ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
		defer cancel()

		cfg, err := cmd.Flags().GetString("config")
		if err != nil {
			color.Red("value of config flag not found")
			return
		}
		if err := run.Run(ctx, cfg); err != nil {
			color.Red("run error: %v", err)
		}
	},
}
