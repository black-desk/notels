package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"strings"

	"github.com/black-desk/lib/go/logger"
)

var log = logger.Get("jsonrpc2gen")

func main() {
	sourceDir, outputDir := getSourceAndOutputDir()

	log.Infow("start",
		"source dir", sourceDir,
		"output dir", outputDir)

	packagePath := os.Args[1]

	fset, lsp := parsePackage(sourceDir)

	for fileName := range lsp.Files {
		processFile(fset, packagePath, lsp, lsp.Files[fileName])
	}

	log.Info("done")

}

func getSourceAndOutputDir() (string, string) {
	sourceDir, err := os.Getwd()
	if err != nil {
		log.Fatalw("failed to getwd",
			"error", err)
	}

	err = os.Mkdir("rpc", 0755)
	if err != nil && (!os.IsExist(err)) {
		log.Fatalw("failed to mkdir",
			"error", err)
	}

	err = os.Chdir("rpc")
	if err != nil {
		log.Fatalw("failed to chdir",
			"error", err)
	}

	outputDir, err := os.Getwd()
	if err != nil {
		log.Fatalw("failed to getwd",
			"error", err)
	}

	return sourceDir, outputDir
}

func parsePackage(sourceDir string) (*token.FileSet, *ast.Package) {
	fset := token.NewFileSet()
	packages, err := parser.ParseDir(
		fset,
		sourceDir,
		nil,
		parser.ParseComments|parser.AllErrors,
	)
	if err != nil {
		log.Fatalw("failed to parse source dir", err)
	}

	for name := range packages {
		if !strings.Contains(name, "test") {
			if name != "lsp" {
				log.Fatalw("unexpected package name",
					"name", name)
			}
			return fset, packages[name]
		}
	}

	log.Fatalw("unexpected situation")
	panic("")
}

func processFile(
	fset *token.FileSet,
	packagePath string,
	pack *ast.Package,
	file *ast.File,
) {
	adaptorNames := []string{}
	proxyNames := []string{}
	printAST := false

	for i := range file.Comments {

		for j := range file.Comments[i].List {
			line := file.Comments[i].List[j].Text
			if !strings.HasPrefix(
				line,
				"// jsonrpc2gen: ",
			) {
				continue
			}

			log.Debugw("comment mark found", "comment", line)

			components := strings.Split(
				line,
				" ",
			)

			if len(components) != 4 {
				log.Fatalw(
					"unexpected jsonrpc2gen comment",
					"line", line,
				)
			}

			if components[3] == "adaptor" {
				adaptorNames = append(
					adaptorNames,
					components[2],
				)
				printAST = true
			} else if components[3] == "proxy" {
				proxyNames = append(
					proxyNames,
					components[2],
				)
				printAST = true
			}
		}
	}

	if printAST {
		ast.Print(fset, file)
	}

	for _, name := range adaptorNames {
		adaptorToGen := file.Scope.Objects[name]
		if adaptorToGen.Kind != ast.Typ {
			log.Fatalw(
				"we can only gen adaptor for interface",
				"kind", adaptorToGen.Kind)
		}

		ts, ok := adaptorToGen.Decl.(*ast.TypeSpec)
		if !ok {
			log.Fatalw(
				"we can only gen adaptor for interface",
				"decl", adaptorToGen.Decl)

		}

		it, ok := ts.Type.(*ast.InterfaceType)
		if !ok {
			log.Fatalw(
				"we can only gen adaptor for interface",
				"type", ts.Type)
		}

		generateAdaptor(packagePath, pack, name, it, fset)
	}

	for _, name := range proxyNames {
		proxyToGen := file.Scope.Objects[name]
		if proxyToGen.Kind != ast.Typ {
			log.Fatalw(
				"we can only gen proxy for interface",
				"kind", proxyToGen.Kind)
		}

		ts, ok := proxyToGen.Decl.(*ast.TypeSpec)
		if !ok {
			log.Fatalw(
				"we can only gen proxy for interface",
				"decl", proxyToGen.Decl)

		}

		it, ok := ts.Type.(*ast.InterfaceType)
		if !ok {
			log.Fatalw(
				"we can only gen proxy for interface",
				"type", ts.Type)
		}

		generateProxy(packagePath, pack, name, it, fset)
	}

}
