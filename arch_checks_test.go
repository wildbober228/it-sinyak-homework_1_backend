package main

import (
	"go/ast"
	"go/token"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
)

func checkDeclarations(t *testing.T, path string, file *ast.File) {
	name := filepath.Base(path)

	for _, decl := range file.Decls {
		gen, ok := decl.(*ast.GenDecl)

		switch {
		case !ok:
		case gen.Tok == token.CONST && name != "constants.go":
			t.Errorf("%s: константы лежат в constants.go", path)
		case gen.Tok == token.TYPE && name != "types.go" && !storeFile.MatchString(name):
			t.Errorf("%s: типы лежат в types.go, тип стора — в файле стора", path)
		}
	}
}

func checkHandler(t *testing.T, path string, storeVars map[string]bool) {
	ast.Inspect(parse(t, path), func(node ast.Node) bool {
		if ident, ok := node.(*ast.Ident); ok && storeVars[ident.Name] {
			t.Errorf("%s: хендлер не трогает стор %s — вызови таску", path, ident.Name)
		}

		return true
	})
}

func checkTask(t *testing.T, path, module string) {
	file := parse(t, path)

	for _, imported := range importPaths(file) {
		if imported == "net/http" || imported == module+"/web" {
			t.Errorf("%s: таска не зависит от HTTP, убери импорт %s", path, imported)
		}
	}

	funcs := topLevelFuncs(file)
	want := strings.ReplaceAll(strings.TrimSuffix(filepath.Base(path), ".go"), "_", "")

	if len(funcs) != 1 || !strings.EqualFold(funcs[0], want) {
		t.Errorf("%s: в файле таски одна функция с именем файла, найдено %v", path, funcs)
	}
}

func importPaths(file *ast.File) []string {
	paths := make([]string, 0, len(file.Imports))

	for _, spec := range file.Imports {
		path, _ := strconv.Unquote(spec.Path.Value)
		paths = append(paths, path)
	}

	return paths
}

func topLevelFuncs(file *ast.File) []string {
	var names []string

	for _, decl := range file.Decls {
		if fn, ok := decl.(*ast.FuncDecl); ok && fn.Recv == nil {
			names = append(names, fn.Name.Name)
		}
	}

	return names
}

func topLevelVars(file *ast.File) []string {
	var names []string

	for _, decl := range file.Decls {
		if gen, ok := decl.(*ast.GenDecl); ok && gen.Tok == token.VAR {
			for _, spec := range gen.Specs {
				for _, name := range spec.(*ast.ValueSpec).Names {
					names = append(names, name.Name)
				}
			}
		}
	}

	return names
}
