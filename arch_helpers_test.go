package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func goFiles(t *testing.T) []string {
	t.Helper()

	var files []string

	err := filepath.WalkDir(".", func(path string, entry fs.DirEntry, err error) error {
		switch {
		case err != nil:
			return err
		case entry.IsDir() && path != "." && strings.HasPrefix(entry.Name(), "."):
			return filepath.SkipDir
		case !entry.IsDir() && strings.HasSuffix(path, ".go"):
			files = append(files, path)
		}

		return nil
	})
	if err != nil {
		t.Fatal(err)
	}

	return files
}

func sourceFiles(t *testing.T, dir string) []string {
	var files []string

	for _, path := range goFiles(t) {
		if filepath.Dir(path) == dir && !strings.HasSuffix(path, "_test.go") {
			files = append(files, path)
		}
	}

	return files
}

func featureDirs(t *testing.T) map[string]bool {
	dirs := map[string]bool{}

	for _, path := range goFiles(t) {
		if filepath.Base(path) == "handler.go" {
			dirs[filepath.Dir(path)] = true
		}
	}

	return dirs
}

func modulePath(t *testing.T) string {
	source, err := os.ReadFile("go.mod")
	if err != nil {
		t.Fatal(err)
	}

	firstLine, _, _ := strings.Cut(string(source), "\n")

	return strings.TrimPrefix(firstLine, "module ")
}

func parse(t *testing.T, path string) *ast.File {
	t.Helper()

	file, err := parser.ParseFile(token.NewFileSet(), path, nil, parser.ParseComments)
	if err != nil {
		t.Fatal(err)
	}

	return file
}
