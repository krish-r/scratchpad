/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const (
	default_dir_name = "Scratch"
)

var (
	programmingLanguage string
	dirName             string
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a scratchpad to play/prototype with your favorite programming language",
	Long: `Create a scratchpad to play/prototype with your favorite programming language.

For example:
scratchpad create --lang java --name "Scratch_Java"
# (or)
scratchpad create -l go -n "Go_Playground"`,
	Run: func(cmd *cobra.Command, args []string) {
		lang, dir := parseFlags(programmingLanguage, dirName)
		createScratchpad(lang, dir)
		fmt.Printf("Created scratchpad for %q under %q directory\n", lang, dir)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	createCmd.Flags().StringVarP(&programmingLanguage, "lang", "l", "", "Language")
	createCmd.Flags().StringVarP(&dirName, "name", "n", default_dir_name, "Optional directory name")
	check(createCmd.MarkFlagRequired("lang"))
}
