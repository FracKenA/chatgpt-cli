package main

import (
	"os"
)

const (
	defaultConfigFilename      = "chatgpt"
	envPrefix                  = "CHATGPT"
	replaceHyphenWithCamelCase = false
)

// TODO add comments to functions. This was a quick will it work.

func main() {
	cmd := rootCommands()
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
