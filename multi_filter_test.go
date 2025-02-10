package pagefilter

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type multiFilterSuite struct {
	suite.Suite

	mf *MultiFilter
}

func TestMultiFilterSuite(t *testing.T) {
	suite.Run(t, new(multiFilterSuite))
}

func (s *multiFilterSuite) SetupTest() {
	s.mf = NewMultiFilter()
}

func (s *multiFilterSuite) TearDownTest() {
	s.mf = nil
}

func (s *multiFilterSuite) TestNewMultiFilter() {
	mf := NewMultiFilter()
	s.NotNil(mf)
	s.NotNil(mf.joinArgs)
	s.NotNil(mf.whereArgs)
	s.NotNil(mf.groupCols)
}

func (s *multiFilterSuite) TestAdd_Where() {
	mf := s.mf
	s.NotNil(mf)

	w := NewMockWherer(s.T())
	w.On("Where").Return("col = ?", []any{"val"})
	mf.Add(w)

	s.Equal("AND col = ?\n", mf.whereSQL.String())
	s.Equal([]any{"val"}, mf.whereArgs)

	w.AssertExpectations(s.T())
}

func (s *multiFilterSuite) TestAdd_WhereTyper() {
	mf := s.mf
	s.NotNil(mf)

	wt := NewMockWhereTyper(s.T())
	wt.On("Where").Return("col = ?", []any{"val"})
	wt.On("WhereType").Return(WhereTypeOr)
	mf.Add(wt)

	s.Equal("OR col = ?\n", mf.whereSQL.String())
	s.Equal([]any{"val"}, mf.whereArgs)

	wt.AssertExpectations(s.T())
}

func (s *multiFilterSuite) TestAdd_Joiner() {
	mf := s.mf
	s.NotNil(mf)

	j := NewMockJoiner(s.T())
	j.On("Join").Return("JOIN table ON table.id = other_table.id", []any{})
	mf.Add(j)

	s.Equal("JOIN table ON table.id = other_table.id\n", mf.joinSQL.String())

	j.AssertExpectations(s.T())
}

func (s *multiFilterSuite) TestAdd_Grouper() {
	mf := s.mf
	s.NotNil(mf)

	g := NewMockGrouper(s.T())
	g.On("Group").Return([]string{"col1", "col2"})
	mf.Add(g)

	s.Equal([]string{"col1", "col2"}, mf.groupCols)

	g.AssertExpectations(s.T())
}

func (s *multiFilterSuite) TestJoin() {
	mf := s.mf
	s.NotNil(mf)

	mf.joinSQL.WriteString("JOIN table ON table.id = other_table.id\n")
	mf.joinArgs = []any{}

	sql, args := mf.Join()
	s.Equal("JOIN table ON table.id = other_table.id", sql)
	s.Empty(args)
}

func (s *multiFilterSuite) TestWhere() {
	mf := s.mf
	s.NotNil(mf)

	mf.whereSQL.WriteString("AND col = ?\n")
	mf.whereArgs = []any{"val"}

	sql, args := mf.Where()
	s.Equal("AND col = ?", sql)
	s.Equal([]any{"val"}, args)
}

func (s *multiFilterSuite) TestAdd_Invalid() {
	mf := s.mf
	s.NotNil(mf)

	mf.Add("invalid")

	s.Empty(mf.joinSQL.String())
	s.Empty(mf.whereSQL.String())
	s.Empty(mf.groupCols)
}

func (s *multiFilterSuite) TestAdd_InvalidWhereTyper() {
	mf := s.mf
	s.NotNil(mf)

	mw := NewMockWhereTyper(s.T())
	mw.On("Where").Return("col = ?", []any{"val"})
	mw.On("WhereType").Return(WhereType("invalid"))

	mf.Add(mw)

	s.Empty(mf.joinSQL.String())
	s.Equal("AND col = ?\n", mf.whereSQL.String())
	s.Equal([]any{"val"}, mf.whereArgs)
}

func (s *multiFilterSuite) TestAdd_InvalidGrouper() {
	mf := s.mf
	s.NotNil(mf)

	mf.Add("invalid")

	s.Empty(mf.joinSQL.String())
	s.Empty(mf.whereSQL.String())
	s.Empty(mf.groupCols)
}

func (s *multiFilterSuite) TestAdd_InvalidJoiner() {
	mf := s.mf
	s.NotNil(mf)

	mf.Add("invalid")

	s.Empty(mf.joinSQL.String())
	s.Empty(mf.whereSQL.String())
	s.Empty(mf.groupCols)
}
