package main

import (
  "fmt"
  "os"
  "reflect"
)

const (
  Version = "Alpha v0.1"
)

var (
  Mode = ""
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
  
  fmt.Println("\nThis is the compiled version of simple.\n")
  
  fmt.Println("Usage: csimple [file] | [-v | --version] | [-h | --help]\n")
  
  fmt.Println("Run 'simple [file] to run code from file;'")
  fmt.Println("Run 'simple -v' or 'simple --version' to show the version number;")
  fmt.Println("Run 'simple -h' or 'simple --help' to show this help message.")
  
  fmt.Println("\nhttps://github.com/JuniorBecari10/CSimple")
  fmt.Println("\nhttps://github.com/JuniorBecari10/Simple - The original project")
}

func Compile(content string) {
  
}
