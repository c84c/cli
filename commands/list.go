package commands

import (
	"fmt"
	"os"

	"github.com/project-flogo/cli/api"
	"github.com/project-flogo/cli/common"
	"github.com/spf13/cobra"
)

var json bool
var all bool

var listCmd = &cobra.Command{
	Use:   "list [flags]",
	Short: "list used flogo contributions",
	Long:  "list used flogo contributions",
	Run: func(cmd *cobra.Command, args []string) {

		err := api.ListPackages(common.CurrentProject(), json, all)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	listCmd.Flags().BoolVarP(&json, "json", "j", true, "print in json format")
	listCmd.Flags().BoolVarP(&all, "all", "a", false, "print all imports")
	rootCmd.AddCommand(listCmd)
}
