package pagefilter

import (
	"errors"
	"net/http"
	"net/url"
	"testing"

	"github.com/jacobbrewer1/pagefilter/common"
	"github.com/stretchr/testify/require"
)

func TestGetLimit(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name      string
		query     url.Values
		wantLimit int
		wantErr   bool
	}{
		{"default limit", url.Values{}, queryLimitDefault, false},
		{"valid limit", url.Values{QueryLimit: {"10"}}, 10, false},
		{"invalid limit", url.Values{QueryLimit: {"invalid"}}, 0, true},
		{"limit exceeds max", url.Values{QueryLimit: {"100000"}}, queryLimitMax, false},
		{"zero limit", url.Values{QueryLimit: {"0"}}, 1, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			gotLimit, err := getLimit(tt.query)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
			require.Equal(t, tt.wantLimit, gotLimit)
		})
	}
}

func TestGetPaginatorDetails(t *testing.T) {
	limit := 10
	lastVal := "lastVal"
	lastID := "lastID"
	offset := 0
	sortBy := "name"
	sortDir := common.SortDirection("asc")

	details := GetPaginatorDetails(&limit, &lastVal, &lastID, &offset, &sortBy, &sortDir)

	require.Equal(t, 10, details.Limit)
	require.Equal(t, "lastVal", details.LastVal)
	require.Equal(t, "lastID", details.LastID)
	require.Equal(t, 0, details.Offset)
	require.Equal(t, "name", details.SortBy)
	require.Equal(t, "asc", details.SortDir)
}

func TestGetPaginatorDetails_NoOffset(t *testing.T) {
	limit := 10
	lastVal := "lastVal"
	lastID := "lastID"
	sortBy := "name"
	sortDir := common.SortDirection("asc")

	details := GetPaginatorDetails(&limit, &lastVal, &lastID, nil, &sortBy, &sortDir)

	require.Equal(t, 10, details.Limit)
	require.Equal(t, "lastVal", details.LastVal)
	require.Equal(t, "lastID", details.LastID)
	require.Equal(t, 0, details.Offset)
	require.Equal(t, "name", details.SortBy)
	require.Equal(t, "asc", details.SortDir)
}

func TestRemoveLimit(t *testing.T) {
	details := &PaginatorDetails{Limit: 10}
	details.RemoveLimit()
	require.Equal(t, -1, details.Limit)
}

func TestDetailsFromRequest(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		query   url.Values
		want    *PaginatorDetails
		wantErr error
	}{
		{
			name: "valid request",
			query: url.Values{
				QueryLimit:   {"10"},
				QueryLastVal: {"lastVal"},
				QueryLastID:  {"lastID"},
				QueryOffset:  {"5"},
				QuerySortBy:  {"name"},
				QuerySortDir: {"asc"},
			},
			want: &PaginatorDetails{
				Limit:   10,
				LastVal: "lastVal",
				LastID:  "lastID",
				Offset:  5,
				SortBy:  "name",
				SortDir: "asc",
			},
			wantErr: nil,
		},
		{
			name: "invalid limit",
			query: url.Values{
				QueryLimit: {"invalid"},
			},
			want:    nil,
			wantErr: errors.New("invalid limit: strconv.Atoi: parsing \"invalid\""),
		},
		{
			name: "invalid offset",
			query: url.Values{
				QueryOffset: {"invalid"},
			},
			want:    nil,
			wantErr: errors.New("invalid offset: strconv.Atoi: parsing \"invalid\""),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			req := &http.Request{URL: &url.URL{RawQuery: tt.query.Encode()}}
			got, err := DetailsFromRequest(req)
			if tt.wantErr != nil {
				require.ErrorContains(t, err, tt.wantErr.Error())
			} else {
				require.NoError(t, err)
			}
			require.Equal(t, tt.want, got)
		})
	}
}
