package lexpragmaids

import (
	"github.com/bhbosman/yaccpragmaids"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPragmaIdsDeclaredTokens(t *testing.T) {

	t.Run("", func(t *testing.T) {
		handler, e := NewPragmaLexIdsFromData("(test stream)", "1234")
		assert.NoError(t, e)

		lexem, err := handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccpragmaids.Value, lexem.TypeKind)
		assert.Equal(t, "1234", lexem.Value)
	})

	t.Run("Test Identifier", func(t *testing.T) {
		handler, e := NewPragmaLexIdsFromData("(test stream)", "xxx")
		assert.NoError(t, e)

		lexem, err := handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, "xxx", lexem.Value)
		assert.Equal(t, yaccpragmaids.Value, lexem.TypeKind)
	})
	t.Run("", func(t *testing.T) {
		handler, e := NewPragmaLexIdsFromData("(test stream)", "IDL")
		assert.NoError(t, e)

		lexem, err := handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccpragmaids.RWIdl, lexem.TypeKind)
	})

	t.Run("", func(t *testing.T) {
		handler, e := NewPragmaLexIdsFromData("(test stream)", "DCE")
		assert.NoError(t, e)

		lexem, err := handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccpragmaids.RWDce, lexem.TypeKind)
	})

	t.Run("", func(t *testing.T) {
		handler, e := NewPragmaLexIdsFromData("(test stream)", "RMI")
		assert.NoError(t, e)

		lexem, err := handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccpragmaids.RWRmi, lexem.TypeKind)
	})

	t.Run("", func(t *testing.T) {
		handler, e := NewPragmaLexIdsFromData("(test stream)", "LOCAL")
		assert.NoError(t, e)

		lexem, err := handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccpragmaids.RWLocal, lexem.TypeKind)
	})

	t.Run("", func(t *testing.T) {
		handler, e := NewPragmaLexIdsFromData("(test stream)", "0")
		assert.NoError(t, e)

		lexem, err := handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccpragmaids.Value, lexem.TypeKind)
	})

	t.Run("", func(t *testing.T) {
		handler, e := NewPragmaLexIdsFromData("(test stream)", "foo/bar;")
		assert.NoError(t, e)

		lexem, err := handler.ReadLexem()
		assert.NoError(t, err)
		assert.Equal(t, yaccpragmaids.Value, lexem.TypeKind)
	})
}
