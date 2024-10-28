package cmd

import (
	"fmt"
	"os"
	"os/exec"

	storage "github.com/playjeri/gpm/pkg"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use: "install <key>",
	Short: "Install a package",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		key := args[0]

		packages, err := storage.ReadPackages()
		if err != nil {
			fmt.Println("Error reading packages:", err)
			return
		}

		value, exists := packages[key]
		if !exists {
			fmt.Printf("Package %s not found\n", key)
			return
		}

		execCmd := exec.Command("go", "get", value)
		fmt.Printf("Running the command: go get %s\n", value)
		execCmd.Stdout = os.Stdout
		execCmd.Stderr = os.Stderr
		err = execCmd.Run()
		if err != nil {
			fmt.Println("Error installing package:", err)
			return
		}

		fmt.Printf("Package %s installed from URL %s\n", key, value)
	},
}
