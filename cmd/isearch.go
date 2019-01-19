package cmd

import (
	"fmt"

	"github.com/PicoSushi/irsty/pkg/irasutoya"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// isearchCmd represents the isearch command
var isearchCmd = &cobra.Command{
	Use:   "isearch",
	Short: "Interactive search from irasutoya",
	Long: `Interactive search from irasutoya.

This command is interactive version of search.`,
	Run: isearch,
}

func init() {
	rootCmd.AddCommand(isearchCmd)
}

func isearch(cmd *cobra.Command, args []string) {
	query := args[0]
	results, err := irasutoya.Search(query)
	if err != nil {
		panic(err)
	}

	titleResult := make(map[string]irasutoya.SearchResult)
	pItems := []string{}
	for _, result := range results {
		titleResult[result.Description] = result
		pItems = append(pItems, result.Description)

	}

	prompt := promptui.Select{
		Label: "Select Entry",
		Items: pItems,
	}

	_, choosed, err := prompt.Run()

	if err != nil {
		return
	}

	entry, err := irasutoya.NewEntry(titleResult[choosed].EntryURL)
	if entry.IsSpecial {
		panic(fmt.Errorf("Sorry, it's special page and there're no irasutoes."))
	}
	irasutoes := entry.Irasutoes
	if len(irasutoes) < 1 {
		panic(fmt.Errorf("Couldn't fetch irasutoes for %s", entry.URL))
	}

	imageURL := ""
	if len(irasutoes) == 1 {
		imageURL = irasutoes[0].ImageURL
	} else {
		titleURL := make(map[string]irasutoya.Irasuto)
		pItems = []string{}
		for _, irasuto := range irasutoes {
			titleURL[irasuto.Title] = irasuto
			pItems = append(pItems, irasuto.Title)
		}

		prompt = promptui.Select{
			Label: "Select Irasuto",
			Items: pItems,
		}
		_, choosed, err = prompt.Run()

		if err != nil {
			return
		}

		imageURL = titleURL[choosed].ImageURL
	}
	fmt.Println(imageURL)
}
