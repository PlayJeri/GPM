package cmd

import (
	"fmt"

	storage "github.com/playjeri/gpm/pkg"
	"github.com/spf13/cobra"
)


var addCmd = &cobra.Command{
	Use: "add <key> <value>",
	Short: "Add a new package",
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		key := args[0]
		value := args[1]

		packages, err := storage.ReadPackages()
		if err != nil {
			fmt.Println("Error reading packages:", err)
			return
		}

		packages[key] = value

		err = storage.WritePackages(packages)
		if err != nil {
			fmt.Println("Error writing packages:", err)
			return
		}

		fmt.Printf("Package %s added with URL %s\n", key, value)
	},
}
