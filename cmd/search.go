package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/PicoSushi/irsty/pkg/irasutoya"
	"github.com/spf13/cobra"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search from irasutoya",
	Long: `Search from irasutoya.

This command returns result as JSON.`,
	Run: search,
}

func init() {
	rootCmd.AddCommand(searchCmd)
}

func search(cmd *cobra.Command, args []string) {
	query := args[0]
	results, err := irasutoya.Search(query)
	if err != nil {
		panic(err)
	}
	i, err := json.Marshal(results)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(i))
}
