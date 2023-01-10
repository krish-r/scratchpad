package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "scratchpad",
	Short: "Create a scratchpad to play/prototype with your favorite programming language",
	Long: `Create a scratchpad to play/prototype with your favorite programming language.
(Supported Languages: Bash, Go, Java, Javascript, Lua, Python, Rust)
Note: This assumes that the necessary tooling is already installed for the programming language.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
