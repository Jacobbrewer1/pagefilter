package pagefilter

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

// Mock implementation of Wherer
type mockWherer struct{}

func (m mockWherer) Where() (sqlStr string, args []any) {
	return "id = ?", []any{1}
}

// Mock implementation of WhereTyper
type mockWhereTyper struct {
	typeValue WhereType
}

func (m mockWhereTyper) Where() (sqlStr string, args []any) {
	return "name = ?", []any{"test"}
}

func (m mockWhereTyper) WhereType() WhereType {
	return m.typeValue
}

type whereTypeSuite struct {
	suite.Suite
}

func TestWhereTypeSuite(t *testing.T) {
	suite.Run(t, new(whereTypeSuite))
}

func (w *whereTypeSuite) TestWhereType_IsValid() {
	w.True(WhereTypeAnd.IsValid())
	w.True(WhereTypeOr.IsValid())
	w.False(WhereType("INVALID").IsValid())
}

func (w *whereTypeSuite) TestMockWherer() {
	mw := new(mockWherer)
	query, args := mw.Where()
	w.Equal("id = ?", query)
	w.Equal([]any{1}, args)
}

func (w *whereTypeSuite) TestMockWhereTyper() {
	mw := &mockWhereTyper{typeValue: WhereTypeAnd}
	query, args := mw.Where()
	w.Equal("name = ?", query)
	w.Equal([]any{"test"}, args)
	w.Equal(WhereTypeAnd, mw.WhereType())
}
