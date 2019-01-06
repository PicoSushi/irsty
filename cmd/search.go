package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/PicoSushi/irsty/pkg/irasutoya"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: search,
}

func init() {
	rootCmd.AddCommand(searchCmd)
	// searchCmd.PersistentFlags().Bool("all", false, "If specifyied, search returns all pages.")
}

func search(cmd *cobra.Command, args []string) {
	fmt.Println("search called")
}
