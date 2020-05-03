package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/go-clix/cli"
)

type jsonnetFunction struct {
	Description string            `json:"description"`
	Params      map[string]string `json:"params"`
	Return      string            `json:"return"`
}

type jsonnetFile struct {
	Name      string            `json:"name"`
	Functions []jsonnetFunction `json:"functions"`
}

func main() {
	rootCmd := &cli.Command{
		Use:   "jsonnetdoc <input-file|dir>",
		Short: "Documentation parser for Jsdoc style comments in Jsonnet",
		Args:  cli.ArgsExact(1),
		Run:   rootCmd,
	}
	rootCmd.Flags().Bool("markdown", false, "output markdown instead of JSON")
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
	markdown, err := strconv.ParseBool(cmd.Flags().Lookup("markdown").Value.String())
	if err != nil {
		return err
	}
	if markdown {
		md, err := generateMarkdown(apiDocs)
		if err != nil {
			return err
		}
		fmt.Println(md)
	} else {
		j, err := json.Marshal(apiDocs)
		if err != nil {
			return err
		}
		fmt.Println(string(j))
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
	jf.Name = name
	content, err := ioutil.ReadFile(p)
	if err != nil {
		return
	}
	r := regexp.MustCompile(`/\*\*(.|[\n])*\*/`)
	docs := r.FindAll(content, -1)
	for _, doc := range docs {
		var desc [][]byte
		descRegexp := regexp.MustCompile(`(\* [^@].+|\s\*$)`)
		params := map[string]string{}
		paramRegexp := regexp.MustCompile(`\* @param.+`)
		var retrn []byte
		retrnRegexp := regexp.MustCompile(`\* @return.+`)
		for _, l := range bytes.Split(doc, []byte("\n")) {
			switch {
			case descRegexp.Match(l):
				desc = append(desc, bytes.TrimLeft(l, "* "))
			case paramRegexp.Match(l):
				param := bytes.SplitN(bytes.TrimLeft(l, "* @param"), []byte(" "), 2)
				if len(param) > 1 {
					params[string(param[0])] = string(param[1])
				} else if len(param) == 1 {
					params[string(param[0])] = ""
				}
			case retrnRegexp.Match(l):
				retrn = bytes.TrimLeft(l, "* @return")
			}
		}
		jf.Functions = append(
			jf.Functions,
			jsonnetFunction{
				Description: string(bytes.Join(desc, []byte("\n"))),
				Params:      params,
				Return:      string(retrn),
			},
		)
	}
	return
}

func generateMarkdown(apiDocs []jsonnetFile) (string, error) {
	md := []string{"# API Docs"}
	for _, jfile := range apiDocs {
		md = append(md, fmt.Sprintf("## %s", jfile.Name))
		for _, jfunc := range jfile.Functions {
			md = append(md, jfunc.Description)
			md = append(md, "### Params")
			for pn, pd := range jfunc.Params {
				md = append(md, fmt.Sprintf("#### %s", pn))
				md = append(md, pd)
			}
			md = append(md, "### Return")
			md = append(md, jfunc.Return)
		}
	}
	return strings.Join(md, "\n"), nil
}
