package main

import (
	"errors"
	"fmt"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/TheRiverBank/cli_test/internal/github/api"
	"github.com/TheRiverBank/cli_test/internal/github/format"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get <detail>",
	Short: "Display one or many repositories details",
	Example: heredoc.Doc(`
		Get stars count for a given repository
		$ name get stars -r golang/go
	`),
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("Requires a detail argument")
		}

		possibleDetails := []string{"stars"}
		validResource := false

		for _, resource := range possibleDetails {
			if args[0] == resource {
				validResource = true
				break
			}
		}

		if validResource == false {
			return errors.New(fmt.Sprintf(`Detail "%s" is invalid`, args[0]))
		}

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		if !format.IsRepositoryValid("golang/go") {
			return errors.New("Repository flag is required")
		}

		detail := args[0]

		switch detail {
		case "stars":
			{
				starsCount, err := api.GetStarsCount("golang/go")
				if err != nil {
					return errors.New("Could not get stars count for this repository: " + err.Error())
				}

				fmt.Println("golang/go" + " has " + fmt.Sprint(*starsCount) + " stars")
			}
		}

		return nil
	},
}
