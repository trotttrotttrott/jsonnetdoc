package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-clix/cli"
)

func main() {
	rootCmd := &cli.Command{
		Use:   "jsonnetdoc <input-file|dir> <output-dir>",
		Short: "Documentation parser for Jsdoc style comments in Jsonnet",
		Args:  cli.ArgsExact(2),
		Run:   rootCmd,
	}
	if err := rootCmd.Execute(); err != nil {
		if err := rootCmd.Execute(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}

func rootCmd(cmd *cli.Command, args []string) error {
	inputPath := args[0]
	files, err := getJsonnetFiles(inputPath)
	if err != nil {
		return err
	}
	fmt.Println(files)
	return nil
}

func getJsonnetFiles(path string) ([]string, error) {
	var files []string
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() &&
			strings.HasSuffix(info.Name(), ".jsonnet") ||
			strings.HasSuffix(info.Name(), ".libsonnet") {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}
