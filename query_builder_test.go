package pagefilter

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRetrieve(t *testing.T) {
	mockDB := NewMockDB(t)
	mockFilter := NewMockFilter(t)

	defer func() {
		mockDB.AssertExpectations(t)
		mockFilter.AssertExpectations(t)
	}()

	mockFilter.On("Join").Return("JOIN table 1 ON table 1.id = table 2.id AND table 1.id = ?", []any{"1"})
	mockFilter.On("Where").Return("t.age > ?", []any{18})

	mockDB.On("Select", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything,
		mock.Anything, mock.Anything, mock.Anything).Return(nil)

	p := &Paginator{
		db: mockDB,
		details: &PaginatorDetails{
			SortBy:         "id",
			sortOperator:   ">",
			sortComparator: "ASC",
			LastID:         "1",
			Limit:          10,
		},
		filter: mockFilter,
		table:  "test_table",
	}

	// Test cases
	tests := []struct {
		name    string
		pivot   string
		dest    any
		wantErr bool
	}{
		{"nil destination", "pivot", nil, true},
		{"non-pointer destination", "pivot", []struct{}{}, true},
		{"non-slice pointer destination", "pivot", &struct{}{}, true},
		{"non-struct slice element", "pivot", &[]int{}, true},
		{"valid destination", "pivot", &[]struct{ ID int }{}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := p.Retrieve(tt.pivot, tt.dest)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestCounts(t *testing.T) {
	mockDB := NewMockDB(t)
	mockFilter := NewMockFilter(t)

	defer func() {
		mockDB.AssertExpectations(t)
		mockFilter.AssertExpectations(t)
	}()

	mockFilter.On("Join").Return("JOIN table 1 ON table 1.id = table 2.id AND table 1.id = ?", []any{"1"})
	mockFilter.On("Where").Return("t.age > ?", []any{18})

	mockDB.On("Get", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything,
		mock.Anything, mock.Anything).Return(func(dest any, sql string, args ...any) error {
		return nil
	})

	p := &Paginator{
		db: mockDB,
		details: &PaginatorDetails{
			SortBy:         "id",
			sortOperator:   ">",
			sortComparator: "ASC",
			LastID:         "1",
			Limit:          10,
		},
		filter: mockFilter,
		table:  "test_table",
	}

	var count int64 = 0
	err := p.Counts(&count)
	assert.NoError(t, err)
}

func TestTrimWherePrefix(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"with AND prefix", "AND condition", "condition"},
		{"with OR prefix", "OR condition", "condition"},
		{"without prefix", "condition", "condition"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := trimWherePrefix(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}
