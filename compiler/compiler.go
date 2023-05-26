package compiler

import (
  "csimple/ast"
)

type Compiler struct {
  statements []ast.Statement
}

func New(stats []ast.Statement) *Compiler {
  return &Compiler { stats }
}
