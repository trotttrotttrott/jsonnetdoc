package main

import (
	"testing"
)

func TestGetJsonnetFiles(t *testing.T) {
	tests := map[string]struct {
		path      string
		expectLen int
		expectErr bool
	}{
		"dir": {
			"testdata", 2, false,
		},
		"jsonnet-file": {
			"testdata/foo.jsonnet", 1, false,
		},
		"libsonnet-file": {
			"testdata/bar.libsonnet", 1, false,
		},
		"notjsonnet-file": {
			"testdata/baz.notjsonnet", 0, false,
		},
		"does-not-exist": {
			"testdata/does-not-exist", 0, true,
		},
	}
	for testName, test := range tests {
		t.Logf("Running test case, %q...", testName)
		files, err := getJsonnetFiles(test.path)
		if err != nil && !test.expectErr {
			t.Errorf("Unexpected error getting Jsonnet files: %s", err)
		}
		if len(files) != test.expectLen && !test.expectErr {
			t.Errorf("Expected %d file(s), got %d", test.expectLen, len(files))
		}
	}
}

func TestParseJsonnetFile(t *testing.T) {
	tests := map[string]struct {
		path   string
		expect jsonnetFile
	}{
		"foo": {
			"testdata/foo.jsonnet",
			jsonnetFile{
				Name: "foo",
				Functions: []jsonnetFunction{
					jsonnetFunction{
						Description: "A jsonnet file called \"foo\"\n",
						Params: map[string]string{
							"foo": "a param called \"foo\"",
							"moo": "a param called \"moo\"",
							"poo": "a param called \"poo\"",
							"roo": "a param called \"roo\"",
							"aoo": "a param called \"aoo\"",
						},
						Return: "a new \"foo\"",
					},
				},
			},
		},
		"bar": {
			"testdata/bar.libsonnet",
			jsonnetFile{
				Name: "bar",
				Functions: []jsonnetFunction{
					jsonnetFunction{
						Description: "A libsonnet file called \"bar\"\nIts got a multi-line description!\n\nMulti-paragraph as well!\n",
						Params: map[string]string{
							"bar":            "a param called \"bar\"",
							"barbar":         "a param called \"barbar\"",
							"no_description": "",
						},
						Return: "a new \"bar\"",
					},
				},
			},
		},
	}
	for testName, test := range tests {
		t.Logf("Running test case, %q...", testName)
		jf, err := parseJsonnetFile(test.path)
		if err != nil {
			t.Errorf("Unexpected error getting Jsonnet files: %s", err)
		}
		if jf.Name != test.expect.Name {
			t.Errorf("Expected jsonnetFile name %q, got %q", test.expect.Name, jf.Name)
		}
		if len(jf.Functions) != len(test.expect.Functions) {
			t.Errorf("Expected %d function(s), got %d", len(test.expect.Functions), len(jf.Functions))
		}
		for i, fn := range jf.Functions {
			if fn.Description != test.expect.Functions[i].Description {
				t.Errorf("Expected description %q, got %q", test.expect.Functions[i].Description, fn.Description)
			}
			if len(fn.Params) != len(test.expect.Functions[i].Params) {
				t.Errorf("Expected %d param(s), got %d", len(test.expect.Functions[i].Params), len(fn.Params))
			}
			for param, desc := range fn.Params {
				if desc != test.expect.Functions[i].Params[param] {
					t.Errorf("Expected param description %q for %q, got %q", test.expect.Functions[i].Params[param], param, desc)
				}
			}
			if fn.Return != test.expect.Functions[i].Return {
				t.Errorf("Expected return %q, got %q", test.expect.Functions[i].Return, fn.Return)
			}
		}
	}
}

func TestGenerateMarkdown(t *testing.T) {

	t.Skip()

	tests := map[string]struct {
		path   string
		expect string
	}{
		"foo": {
			path: "testdata/foo.jsonnet",
			expect: `# API Docs
## foo
A jsonnet file called "foo"

* **foo**: a param called "foo"
* **moo**: a param called "moo"
* **poo**: a param called "poo"
* **roo**: a param called "roo"
* **aoo**: a param called "aoo"

_returns_ a new "foo"`,
		},
	}
	for testName, test := range tests {
		t.Logf("Running test case, %q...", testName)
		var apiDocs []jsonnetFile
		jf, _ := parseJsonnetFile(test.path)
		apiDocs = append(apiDocs, jf)
		md, err := generateMarkdown(apiDocs)
		if err != nil {
			t.Errorf("Unexpected error parsing markdown: %s", err)
		}
		if md != test.expect {
			t.Errorf("Expected:\n%s\n\ngot:\n\n%s", test.expect, md)
		}
	}
}
