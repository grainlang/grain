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
		standard_library.CreateAddAst(),
		standard_library.CreateSubtractAst(),
		standard_library.CreateMultiplyAst(),
		standard_library.CreateDivideAst(),
		standard_library.CreateIntEqualsAst(),
		hello.CreateGetCharPutCharAst(),
		hello.CreatePutCharConstGAst(),
		hello.CreatePutCharConst97Ast(),
		hello.CreateGetCharToUpperPutCharAst(),
		hello.CreateAddTwoCharactersAst(),
		hello.CreateVoodooCalculationsAst(),
		hello.CreateGetCharEqualsAAst(),
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
	allUsedFunctions := make([]ast.Function, 0)
	allUsedFunctions = fillAllUsedFunctions(allUsedFunctions, matchedFunctions[0], allFunctions)
	modules := make(map[string]llvm.Module)
	modules["main"] = llvm_module.CreateMainModuleWithCallToFunction(matchedFunctions[0])
	for _, usedFunc := range allUsedFunctions {
		module := llvm_module.CreateLlvmModuleFromFunction(usedFunc, allFunctions)
		module.Dump()
		modules[usedFunc.Id] = module
	}
	machine, err := initMachine()
	if err != nil {
		os.Exit(-1)
	}
	clangParams := make([]string, len(modules) + 2)
	for id, module := range modules {
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

func fillAllUsedFunctions(allUsedFunctions []ast.Function, current ast.Function, allFunctions []ast.Function) []ast.Function {
	if !isIn(current, allUsedFunctions) {
		allUsedFunctions = append(allUsedFunctions, current)
		for _, part := range current.Body {
			if use, ok := part.(ast.FunctionUse); ok {
				fn := llvm_module.FindUsedFunction(use, allFunctions)
				childrenFunctions := fillAllUsedFunctions(allUsedFunctions, fn, allFunctions)
				for _, child := range childrenFunctions {
					if !isIn(child, allUsedFunctions) {
						allUsedFunctions = append(allUsedFunctions, child)
					}
				}
			}
		}
	}
	return allUsedFunctions
}

func isIn(function ast.Function, functions []ast.Function) bool {
	for _, f := range functions {
		if function.Id == f.Id {
			return true
		}
	}
	return false
}