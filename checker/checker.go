package checker

import (
  "csimple/ast"
)

type Checker struct {
  statement ast.Statement
}

func New(stat ast.Statement) *Checker {
  return &Checker { stat }
}

func (this *Checker) Check() {
  
}

func (this *Checker) CheckZeroDivision() error {
  switch this.statement.(type) {
    case ast.EndStatement:
    case ast.LabelStatement:
    case ast.RetStatement:
      return nil
    
    
  }
}