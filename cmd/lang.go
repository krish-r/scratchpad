package cmd

import (
	"strings"
)

type languageInfo map[string]command

func (c *languageInfo) keys() SupportedLanguages {
	keys := make(SupportedLanguages, len(*c))
	i := 0
	for k := range *c {
		keys[i] = k
		i += 1
	}
	return keys
}

var mappings = languageInfo{
	"BASH": {
		requiresDir: true,
		layers: []layer{
			{executable: "touch", args: []string{"scratch.sh"}},
		},
	},
	"GO": {
		requiresDir: true,
		layers: []layer{
			{executable: "go", args: []string{"mod", "init", "{dirName}"}},
			{executable: "touch", args: []string{"main.go"}},
		},
	},
	"JAVA": {
		requiresDir: false,
		layers: []layer{
			{executable: "mvn", args: []string{"archetype:generate", "-DgroupId=com.example", "-DartifactId={dirName}", "-DarchetypeArtifactId=maven-archetype-quickstart", "-DarchetypeVersion=1.4", "-DinteractiveMode=false"}},
		},
	},
	"JAVASCRIPT": {
		requiresDir: true,
		layers: []layer{
			{executable: "npm", args: []string{"init", "-y"}},
		},
	},
	"PYTHON": {
		requiresDir: true,
		layers: []layer{
			{executable: "touch", args: []string{"main.py"}},
		},
	},
	"RUST": {
		requiresDir: false,
		layers: []layer{
			{executable: "cargo", args: []string{"new", "{dirName}"}},
		},
	},
	"LUA": {
		requiresDir: true,
		layers: []layer{
			{executable: "touch", args: []string{"init.lua"}},
		},
	},
}

type SupportedLanguages []string

var supportedLanguages SupportedLanguages = mappings.keys()

func (l *SupportedLanguages) contains(lang string) bool {
	lang = strings.ToUpper(lang)
	for _, supportedLanguage := range *l {
		if lang == supportedLanguage {
			return true
		}
	}
	return false
}

func (l *SupportedLanguages) String() string {
	return strings.Join(*l, ", ")
}
