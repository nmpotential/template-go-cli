package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"reflect"
	"strings"
	"testing"
)

var binaryName = "tempconverterTestBinary"

func TestMain(m *testing.M) {
	build := exec.Command("go", "build", "-o", binaryName)
	err := build.Run()
	if err != nil {
		fmt.Printf("could not make binary %v", err)
		os.Exit(1)
	}
	exitCode := m.Run()

	cleanUp := exec.Command("rm", "-f", binaryName)
	cleanUperr := cleanUp.Run()
	if cleanUperr != nil {
		fmt.Println("could not clean up", err)
	}

	os.Exit(exitCode)
}

func TestCheckForArgumentsM1(t *testing.T) {
	t.Run("invalid args", func(t *testing.T) {
		dir, err := os.Getwd()
		if err != nil {
			t.Fatal(err)
		}
		// Runs the program with not enough arguments.
		cmd := exec.Command(path.Join(dir, binaryName), []string{}...)
		output, err := cmd.CombinedOutput()
		if err == nil || !strings.Contains(string(output), errInvalidArguments.Error()) {
			t.Fatal("Did not validate command line arguments properly")
		}

		// Runs the program with more than enough
		cmd2 := exec.Command(path.Join(dir, binaryName), []string{"one", "two"}...)
		output2, err2 := cmd2.CombinedOutput()
		if err2 == nil || !strings.Contains(string(output2), errInvalidArguments.Error()) {
			t.Fatal("Did not validate command line arguments properly")
		}
	})
}

func TestAssignsToOriginUnitM1(t *testing.T) {
	isAssigningToOriginUnit := false
	isAssigningToOriginUnitFromCorrectFunction := false

	src, err := ioutil.ReadFile("main.go")
	if err != nil {
		log.Fatal(err)
	}

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", src, 0)
	if err != nil {
		panic(err)
	}

	for _, decl := range f.Decls {
		if reflect.TypeOf(decl).String() == "*ast.FuncDecl" {
			funcDecl := decl.(*ast.FuncDecl)
			// Only interested in main() function
			if funcDecl.Name.Name != "main" {
				continue
			}

			for _, decl2 := range funcDecl.Body.List {
				if reflect.TypeOf(decl2).String() == "*ast.AssignStmt" {
					assignToOriginUnit := decl2.(*ast.AssignStmt)
					for _, ouVariable := range assignToOriginUnit.Lhs {
						ident := ouVariable.(*ast.Ident)
						if ident.Obj.Name == "originUnit" {
							isAssigningToOriginUnit = true
						}
					}
					for _, ouValue := range assignToOriginUnit.Rhs {
						if reflect.TypeOf(ouValue).String() == "*ast.CallExpr" {
							v1 := ouValue.(*ast.CallExpr)
							v2 := v1.Fun.(*ast.SelectorExpr)
							v3 := v2.X.(*ast.Ident)
							if v3.Name == "strings" && v2.Sel.Name == "ToUpper" {
								isAssigningToOriginUnitFromCorrectFunction = true
							}
						}
					}
				}
			}

			if !isAssigningToOriginUnit || !isAssigningToOriginUnitFromCorrectFunction {
				t.Error("Did not assign proper value to originUnit")
			}
		}
	}
}

func TestReadsCurrentTemperatureM2(t *testing.T) {
	t.Run("reads current temperature", func(t *testing.T) {
		dir, err := os.Getwd()
		if err != nil {
			t.Fatal(err)
		}

		cmd := exec.Command(path.Join(dir, binaryName), []string{"C"}...)

		stdin, e := cmd.StdinPipe()
		if e != nil {
			panic(e)
		}
		stderr, e := cmd.StderrPipe()
		if e != nil {
			panic(e)
		}

		if e := cmd.Start(); e != nil {
			panic(e)
		}
		_, e = stdin.Write([]byte("FFFF\n"))
		if e != nil {
			panic(e)
		}
		_, e = stdin.Write([]byte("n\n"))
		if e != nil {
			panic(e)
		}
		stdin.Close()
		errPrint, _ := ioutil.ReadAll(stderr)
		if !strings.Contains(string(errPrint), errReadingInput.Error()) {
			t.Fatal("Did not print errReadingInput error")
		}
	})
}

func TestCheckConversionM2(t *testing.T) {
	t.Run("converts temperature", func(t *testing.T) {
		dir, err := os.Getwd()
		if err != nil {
			t.Fatal(err)
		}

		cmd := exec.Command(path.Join(dir, binaryName), []string{"F"}...)

		stdin, e := cmd.StdinPipe()
		if e != nil {
			panic(e)
		}
		outPipe, e := cmd.StdoutPipe()
		if e != nil {
			panic(e)
		}

		if e := cmd.Start(); e != nil {
			panic(e)
		}
		_, e = stdin.Write([]byte("32\n"))
		if e != nil {
			panic(e)
		}
		_, e = stdin.Write([]byte("n\n"))
		if e != nil {
			panic(e)
		}
		stdin.Close()
		outPrint, _ := ioutil.ReadAll(outPipe)
		if !strings.Contains(string(outPrint), "32 F = 0 C") {
			t.Fatal("Did not properly convert temperature")
		}
	})

}

func TestPromptAgainM2(t *testing.T) {
	t.Run("convert again", func(t *testing.T) {
		dir, err := os.Getwd()
		if err != nil {
			t.Fatal(err)
		}

		cmd := exec.Command(path.Join(dir, binaryName), []string{"F"}...)

		stdin, e := cmd.StdinPipe()
		if e != nil {
			panic(e)
		}
		outPipe, e := cmd.StdoutPipe()
		if e != nil {
			panic(e)
		}

		if e := cmd.Start(); e != nil {
			panic(e)
		}
		_, e = stdin.Write([]byte("32\n"))
		if e != nil {
			panic(e)
		}
		_, e = stdin.Write([]byte(" Y\n"))
		if e != nil {
			panic(e)
		}
		_, e = stdin.Write([]byte("40\n"))
		if e != nil {
			panic(e)
		}
		stdin.Close()
		outPrint, _ := ioutil.ReadAll(outPipe)
		if !strings.Contains(string(outPrint), "32 F = 0 C") ||
			!strings.Contains(string(outPrint), "40 F = 4 C") {
			t.Fatal("Did not prompt user for another run")
		}

	})
}

func TestParsePromptToUpperM2(t *testing.T) {
	t.Run("prompt to upper", func(t *testing.T) {
		dir, err := os.Getwd()
		if err != nil {
			t.Fatal(err)
		}

		cmd := exec.Command(path.Join(dir, binaryName), []string{"F"}...)

		stdin, e := cmd.StdinPipe()
		if e != nil {
			panic(e)
		}
		outPipe, e := cmd.StdoutPipe()
		if e != nil {
			panic(e)
		}

		if e := cmd.Start(); e != nil {
			panic(e)
		}
		_, e = stdin.Write([]byte("32\n"))
		if e != nil {
			panic(e)
		}
		_, e = stdin.Write([]byte("y\n"))
		if e != nil {
			panic(e)
		}
		_, e = stdin.Write([]byte("212\n"))
		if e != nil {
			panic(e)
		}
		stdin.Close()
		outPrint, _ := ioutil.ReadAll(outPipe)
		if !strings.Contains(string(outPrint), "32 F = 0 C") ||
			!strings.Contains(string(outPrint), "212 F = 100 C") {
			t.Fatal("Did not properly parse user input")
		}
	})
}
