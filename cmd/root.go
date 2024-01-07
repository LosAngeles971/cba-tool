package cmd

import (
	"fmt"
	"os"

	"github.com/LosAngeles971/cba-tool/business/ui"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

const (
	description = `cba`
)

var debug bool
var format string
var filename string

var rootCmd = &cobra.Command{
	Use:   "cba-tool",
	Short: "Cost-Benefit Analysis tool",
	Long:  description,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if debug {
			log.SetLevel(log.DebugLevel)
			log.Print("enabled debug")
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		app := ui.Build()
		if err := app.Run(); err != nil {
			panic(err)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	//rootCmd.AddCommand(PortfolioCmd, tickerCmd)
	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "enable/disable debug")
	rootCmd.PersistentFlags().StringVar(&filename, "filename", "", "cba filename")
	rootCmd.PersistentFlags().StringVar(&format, "format", "stdout", "output format")
}
