package config

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// ShowCmd represents the config:show command
var ShowCmd = &cobra.Command{
	Use:   "config:show",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("--- Local Flags ---")
		cmd.Flags().VisitAll(func(flag *pflag.Flag) {
			fmt.Printf("%s: %v\n", flag.Name, flag.Value)
		})

		fmt.Println("--- Inherited Flags ---")
		cmd.InheritedFlags().VisitAll(func(flag *pflag.Flag) {
			fmt.Printf("%s: %v\n", flag.Name, flag.Value)
		})
		return nil
	},
}
