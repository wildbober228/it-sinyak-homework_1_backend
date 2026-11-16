package main

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
)

const maxLines = 120

var (
	nonTaskFile = regexp.MustCompile(`^(handler\w*|\w*store|types|constants|utils|errors)\.go$`)
	storeFile   = regexp.MustCompile(`^\w*store\.go$`)
	handlerFile = regexp.MustCompile(`^handler\w*\.go$`)
)

func TestEveryFile(t *testing.T) {
	module := modulePath(t)
	features := featureDirs(t)

	for _, path := range goFiles(t) {
		file := parse(t, path)
		source, err := os.ReadFile(path)
		if err != nil {
			t.Fatal(err)
		}

		if lines := strings.Count(string(source), "\n"); lines > maxLines {
			t.Errorf("%s: %d строк, лимит %d — разнеси ответственности", path, lines, maxLines)
		}

		if len(file.Comments) > 0 {
			t.Errorf("%s: комментарии не пишем — переименуй или разбей код", path)
		}

		for _, imported := range importPaths(file) {
			feature := strings.TrimPrefix(imported, module+"/")
			if filepath.Dir(path) != "." && features[feature] {
				t.Errorf("%s: фичу %s подключает только main.go", path, feature)
			}
		}

		if !strings.HasSuffix(path, "_test.go") {
			checkDeclarations(t, path, file)
		}
	}
}

func TestFeatures(t *testing.T) {
	module := modulePath(t)

	for dir := range featureDirs(t) {
		files := sourceFiles(t, dir)
		storeVars := map[string]bool{}

		for _, path := range files {
			if storeFile.MatchString(filepath.Base(path)) {
				for _, name := range topLevelVars(parse(t, path)) {
					storeVars[name] = true
				}
			}
		}

		for _, path := range files {
			name := filepath.Base(path)

			switch {
			case handlerFile.MatchString(name):
				checkHandler(t, path, storeVars)
			case !nonTaskFile.MatchString(name):
				checkTask(t, path, module)
			}
		}
	}
}
