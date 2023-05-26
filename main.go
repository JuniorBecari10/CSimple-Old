package main

import (
  "fmt"
  "os"
  "strings"
  "strconv"
  "reflect"
  
  "csimple/lexer"
  "csimple/parser"
)

const (
  Version = "Alpha v0.1"
)

func main() {
  if len(os.Args) == 2 {
    if os.Args[1] == "-v" || os.Args[1] == "--version" {
      fmt.Println("CSimple - " + Version)
      fmt.Println("Made by JuniorBecari10")
      return
    }
    
    if os.Args[1] == "help" || os.Args[1] == "-h" || os.Args[1] == "--help" {
      help()
      return
    }
    
    content, err := os.ReadFile(os.Args[1])
    
    if err != nil {
      fmt.Println("File '" + os.Args[1] + "' doesn't exist.")
      fmt.Println("Verify if you typed the name correctly.")
      
      return
    }
    
    Compile(string(content))
    return
  }
  
  help()
}

func help() {
  fmt.Println("CSimple - " + Version)
  
  fmt.Println("\nThe compiled version of Simple.\n")
  
  fmt.Println("Usage: csimple [file] | [-v | --version] | [-h | --help]\n")
  
  fmt.Println("Run 'csimple [file]' to run code from file;")
  fmt.Println("Run 'csimple -v' or 'csimple --version' to show the version number;")
  fmt.Println("Run 'csimple -h' or 'csimple --help' to show this help message.")
  
  fmt.Println("\nhttps://github.com/JuniorBecari10/CSimple")
  fmt.Println("https://github.com/JuniorBecari10/Simple - The original project")
  
  fmt.Println("\nMade by JuniorBecari10.")
}

func Compile(code string) {
  lines := strings.Split(strings.TrimSpace(code), "\n")

  for i, _ := range lines {
    tokens := lexer.Lex(lines[i])
    errs := lexer.CheckErrors(tokens)
    
    if len(errs) > 0 {
      for _, e := range errs {
        ShowError("Error in lexer: " + e, "Consider removing them.", lines[i], i)
      }
      
      return
    }

    stats := parser.Parse(tokens)
    errs = parser.CheckErrors(stats)
    
    if len(errs) > 0 {
      for _, e := range errs {
        ShowError("Error in parser: " + e, "", lines[i], i)
      }
      
      return
    }
    
    for _, s := range stats {
      fmt.Printf("%s | %+v\n", reflect.TypeOf(s), s)
    }

    if len(stats) == 0 {
      continue
    }
  }
}

func ShowError(msg, hint, lineCode string, line int) {
  fmt.Println("\n-------------\n")
  
  fmt.Println("ERROR: On line " + strconv.Itoa(line + 1) + ".")
  fmt.Println("\n" + msg)
  
  fmt.Println()
  
  if line > 0 {
    fmt.Printf("%d |\n", line)
  }
  
  fmt.Printf("%d | %s\n", line + 1, lineCode)
  fmt.Printf("%d |\n\n", line + 2)
  
  if hint != "" {
    fmt.Println(hint)
    fmt.Println()
  }
}