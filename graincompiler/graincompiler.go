package main

import (
	"llvm.org/llvm/bindings/go/llvm"
	"fmt"
	"os"
	"io/ioutil"
	"os/exec"
	"github.com/grainlang/grain/ast"
	"github.com/grainlang/grain/standard_library"
	"github.com/grainlang/grain/hello"
	"github.com/grainlang/grain/llvm_module"
	"flag"
	"strings"
)

func main() {
	functionFlag := flag.String("f", "", "part of function id, name or description")
	outputFlag := flag.String("o", "program", "name of executable to be created")
	flag.Parse()
	functionId := *functionFlag
	outputName := *outputFlag
	if functionId == "" {
		fmt.Println("Choose function to compile. -h for help.")
		os.Exit(-1)
	}
	allFunctions := []ast.Function{
		standard_library.CreateGetCharacterAst(),
		standard_library.CreatePutCharacterAst(),
		standard_library.CreateToUppercaseAst(),
		hello.CreateGetCharPutCharAst(),
		hello.CreateGetCharToUpperPutCharAst(),
	}
	matchedFunctions := make([]ast.Function, 0)
	for _, function := range allFunctions {
		if strings.Contains(function.Id, functionId) || strings.Contains(function.Name, functionId) || strings.Contains(function.Description, functionId) {
			matchedFunctions = append(matchedFunctions, function)
		}
	}
	if len(matchedFunctions) == 0 {
		fmt.Println("No function matched name \"" + functionId + "\"")
		os.Exit(-1)
	} else if len(matchedFunctions) > 1 {
		list := ""
		for _, matchedFunc := range matchedFunctions {
			list += matchedFunc.Id + " " + matchedFunc.Name + " " + matchedFunc.Description + "\n"
		}
		fmt.Print("Too many function matched name \"" + functionId + "\"\n" + list)
		os.Exit(-1)
	}
	modules := make(map[string]llvm.Module)
	modules["main"] = llvm_module.CreateMainModuleWithCallToFunction(matchedFunctions[0])
	processedFunctions := make([]ast.Function, 1)
	for len(matchedFunctions) != 0 {
		current := matchedFunctions[0]
		matchedFunctions = matchedFunctions[1:]
		for _, bodyPart := range current.Body {
			if binding, ok := bodyPart.(ast.Binding); ok {
				f1, f2 := llvm_module.FindUsedFunctions(binding, allFunctions)
				if !isIn(f1, processedFunctions) && !isIn(f1, matchedFunctions) {
					matchedFunctions = append(matchedFunctions, f1)
				}
				if !isIn(f2, processedFunctions) && !isIn(f2, matchedFunctions) {
					matchedFunctions = append(matchedFunctions, f2)
				}
			}
		}
		modules[current.Id] = llvm_module.CreateLlvmModuleFromFunction(current, allFunctions)
		processedFunctions = append(processedFunctions, current)
	}
	machine, err := initMachine()
	if err != nil {
		os.Exit(-1)
	}
	clangParams := make([]string, len(modules) + 2)
	for id, module := range modules {
		module.Dump()
		buffer, err := machine.EmitToMemoryBuffer(module, llvm.ObjectFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Cannot emit object file to memory buffer:")
			fmt.Fprintln(os.Stderr, err)
			os.Exit(-1)
		}
		fileName := id + ".o"
		err = ioutil.WriteFile(fileName, buffer.Bytes(), 0644)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Cannot save file:")
			fmt.Fprintln(os.Stderr, err)
			os.Exit(-1)
		}
		clangParams = append(clangParams, fileName)
	}
	clangParams = append(clangParams, "-o")
	clangParams = append(clangParams, outputName)
	cmd := exec.Command("clang", clangParams...)
	err = cmd.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Cannot run clang:")
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	}
	for id := range modules {
		os.Remove(id + ".o")
	}
}

func isIn(function ast.Function, functions []ast.Function) bool {
	for _, f := range functions {
		if function.Id == f.Id {
			return true
		}
	}
	return false
}