package pagefilter

import (
	"net/url"
	"testing"

	"github.com/jacobbrewer1/pagefilter/common"
	"github.com/stretchr/testify/assert"
)

func TestGetLimit(t *testing.T) {
	tests := []struct {
		name      string
		query     url.Values
		wantLimit int
		wantErr   bool
	}{
		{"default limit", url.Values{}, defaultPageLimit, false},
		{"valid limit", url.Values{QueryLimit: {"10"}}, 10, false},
		{"invalid limit", url.Values{QueryLimit: {"invalid"}}, -1, true},
		{"limit exceeds max", url.Values{QueryLimit: {"100000"}}, maxLimit, false},
		{"zero limit", url.Values{QueryLimit: {"0"}}, defaultPageLimit, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotLimit, err := getLimit(tt.query)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.wantLimit, gotLimit)
		})
	}
}

func TestGetPaginatorDetails(t *testing.T) {
	limit := "10"
	lastVal := "lastVal"
	lastID := "lastID"
	sortBy := "name"
	sortDir := common.SortDirection("asc")

	details := GetPaginatorDetails(&limit, &lastVal, &lastID, &sortBy, &sortDir)

	assert.Equal(t, 10, details.Limit)
	assert.Equal(t, "lastVal", details.LastVal)
	assert.Equal(t, "lastID", details.LastID)
	assert.Equal(t, "name", details.SortBy)
	assert.Equal(t, "asc", details.SortDir)
}

func TestRemoveLimit(t *testing.T) {
	details := &PaginatorDetails{Limit: 10}
	details.RemoveLimit()
	assert.Equal(t, -1, details.Limit)
}
