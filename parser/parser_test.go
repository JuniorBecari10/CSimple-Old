package parser

import (
  "testing"
  "reflect"
  
  "csimple/token"
  "csimple/lexer"
  "csimple/ast"
)

func TestParser(t *testing.T) {
  input := `jump :hello`
  
  tokens := lexer.Lex(input)
  checkLexerErrors(t, tokens)
  
  stats := Parse(tokens)
  
  expect := []ast.Statement {
    ast.JumpStatement {
      Token: token.Token { token.JumpKw, "jump", 0 },
      Label: ":hello",
    },
    ast.EndStatement {},
  }
  
  for i, s := range stats {
    if !reflect.DeepEqual(s, expect[i]) {
      t.Fatalf("not equal. len %d.\nexpected %+v\ngot      %+v", len(stats), expect[i], s)
    }
  }
}

func checkLexerErrors(t *testing.T, tokens []token.Token) {
  for _, tk := range tokens {
    if tk.Type == token.Error {
      t.Fatalf("lexer error: %s", tk.Content)
    }
  }
}