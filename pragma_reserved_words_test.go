package lexpragma

import (
	"github.com/bhbosman/yaccpragma"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHandler_Hex(t *testing.T) {
	t.Run("Test Identifier", func(t *testing.T) {
		handler, e := NewPragmaLexFromData("(test stream)", "a", nil)
		assert.NoError(t, e)
		lexem := handler.ReadLexem()
		assert.Equal(t, "a", lexem.StringValue)
		assert.Equal(t, yaccpragma.Identifier, lexem.TypeKind)
	})

	t.Run("Test #pragma", func(t *testing.T) {
		handler, e := NewPragmaLexFromData("(test stream)", "#pragma", nil)
		assert.NoError(t, e)
		lexem := handler.ReadLexem()
		assert.Equal(t, yaccpragma.RwPragma, lexem.TypeKind)
	})

	t.Run("Test prefix", func(t *testing.T) {
		handler, e := NewPragmaLexFromData("(test stream)", "prefix", nil)
		assert.NoError(t, e)
		lexem := handler.ReadLexem()

		assert.Equal(t, yaccpragma.RwPrefix, lexem.TypeKind)
	})

	t.Run("Test ID", func(t *testing.T) {
		handler, e := NewPragmaLexFromData("(test stream)", "ID", nil)
		assert.NoError(t, e)
		lexem := handler.ReadLexem()

		assert.Equal(t, yaccpragma.RwId, lexem.TypeKind)
	})

	t.Run("Test string", func(t *testing.T) {
		handler, e := NewPragmaLexFromData("(test stream)", "\"abc\"", nil)
		assert.NoError(t, e)
		lexem := handler.ReadLexem()

		assert.Equal(t, yaccpragma.StringLiteral, lexem.TypeKind)
	})

	t.Run("Test Whitespace", func(t *testing.T) {
		handler, e := NewPragmaLexFromData("(test stream)", "\t\t\n ", nil)
		assert.NoError(t, e)
		lexem := handler.ReadLexem()

		assert.Equal(t, yaccpragma.WhiteSpace, lexem.TypeKind)
	})

	t.Run("Test Whitespace", func(t *testing.T) {
		handler, e := NewPragmaLexFromData("(test stream)", "version", nil)
		assert.NoError(t, e)
		lexem := handler.ReadLexem()

		assert.Equal(t, yaccpragma.RwVersion, lexem.TypeKind)
	})

}
