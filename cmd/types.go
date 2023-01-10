package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// A command is made of multiple layers
type command struct {
	requiresDir bool
	layers      []layer
}

// Each layer has an executable and optional args
type layer struct {
	executable string
	args       []string
}

func parseFlags(lang, dir string) (string, string) {
	lang = parseLanguage(lang)
	dir = parseDirectory(dir, lang)
	return lang, dir
}

func parseLanguage(l string) string {
	language := strings.TrimSpace(l)
	if language == "" {
		fmt.Println("Language cannot be empty", l)
		os.Exit(1)
	}
	if !supportedLanguages.contains(language) {
		fmt.Printf("Received invalid value for language %q\n", l)
		fmt.Printf("Accepted values are: %q\n", sortedCopy(supportedLanguages))
		os.Exit(1)
	}
	return language
}

func parseDirectory(n string, lang string) string {
	dirName := strings.ReplaceAll(strings.TrimSpace(n), " ", "_")
	if dirName == default_dir_name {
		dirName = fmt.Sprintf("%s_%s", dirName, cases.Title(language.English).String(lang))
	}
	return dirName
}

func createScratchpad(language string, dirName string) {
	language = strings.ToUpper(language)
	commands := mappings[language]
	if commands.requiresDir {
		mkDirs(dirName)
		check(os.Chdir(dirName))
	}
	for _, command := range commands.layers {
		substituteDirNameInPlace(command.args, dirName)
		cmd := exec.Command(command.executable, command.args...)
		fmt.Printf("Executing command: %s\n", cmd)
		err := cmd.Run()
		check(err)
	}
}

func substituteDirNameInPlace(args []string, dirName string) {
	length := len(args)
	for i := 0; i < length; i++ {
		if strings.Contains(args[i], "{dirName}") {
			args[i] = strings.ReplaceAll(args[i], "{dirName}", dirName)
		}
	}
}
