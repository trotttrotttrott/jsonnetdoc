package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/go-clix/cli"
)

type jsonnetFunction struct {
	description string
	params      map[string]string
	retrn       string
}

type jsonnetFile struct {
	name      string
	functions []jsonnetFunction
}

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
	var apiDocs []jsonnetFile
	for _, f := range files {
		jf, err := parseJsonnetFile(f)
		if err != nil {
			return err
		}
		apiDocs = append(apiDocs, jf)
	}
	return nil
}

func getJsonnetFiles(p string) ([]string, error) {
	var files []string
	err := filepath.Walk(p, func(p string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() &&
			strings.HasSuffix(info.Name(), ".jsonnet") ||
			strings.HasSuffix(info.Name(), ".libsonnet") {
			files = append(files, p)
		}
		return nil
	})
	return files, err
}

func parseJsonnetFile(p string) (jf jsonnetFile, err error) {
	_, f := path.Split(p)
	name := strings.TrimSuffix(f, path.Ext(f))
	jf.name = name

	return
}
