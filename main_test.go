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
				name: "foo",
			},
		},
		"bar": {
			"testdata/bar.libsonnet",
			jsonnetFile{
				name: "bar",
			},
		},
	}
	for testName, test := range tests {
		t.Logf("Running test case, %q...", testName)
		jf, err := parseJsonnetFile(test.path)
		if err != nil {
			t.Errorf("Unexpected error getting Jsonnet files: %s", err)
		}
		if jf.name != test.expect.name {
			t.Errorf("Expected jsonnetFile name %q, got %q", test.expect.name, jf.name)
		}
		if len(jf.functions) != len(test.expect.functions) {
			t.Errorf("Expected %d file(s), got %d", len(test.expect.functions), len(jf.functions))
		}
	}
}
