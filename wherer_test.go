package pagefilter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Mock implementation of Wherer
type mockWherer struct{}

func (m mockWherer) Where() (string, []any) {
	return "id = ?", []any{1}
}

// Mock implementation of WhereTyper
type mockWhereTyper struct {
	typeValue WhereType
}

func (m mockWhereTyper) Where() (string, []any) {
	return "name = ?", []any{"test"}
}

func (m mockWhereTyper) WhereType() WhereType {
	return m.typeValue
}

func TestWhereType_IsValid(t *testing.T) {
	assert.True(t, WhereTypeAnd.IsValid())
	assert.True(t, WhereTypeOr.IsValid())
	assert.False(t, WhereType("INVALID").IsValid())
}

func TestMockWherer(t *testing.T) {
	w := mockWherer{}
	query, args := w.Where()
	assert.Equal(t, "id = ?", query)
	assert.Equal(t, []any{1}, args)
}

func TestMockWhereTyper(t *testing.T) {
	w := mockWhereTyper{typeValue: WhereTypeAnd}
	query, args := w.Where()
	assert.Equal(t, "name = ?", query)
	assert.Equal(t, []any{"test"}, args)
	assert.Equal(t, WhereTypeAnd, w.WhereType())
}
