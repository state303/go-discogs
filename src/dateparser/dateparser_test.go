package dateparser

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseYMD(t *testing.T) {
	var of = func(n int) *int16 {
		if n <= 0 {
			return nil
		}
		i := int16(n)
		return &i
	}
	var tests = []struct {
		name  string
		ymd   string
		wantY *int16
		wantM *int16
		wantD *int16
	}{
		{
			name:  "4 digit year",
			ymd:   "1991",
			wantY: of(1991),
			wantM: nil,
			wantD: nil,
		},
		{
			name:  "2 digit year",
			ymd:   "03",
			wantY: of(2003),
			wantM: nil,
			wantD: nil,
		},
		{
			name:  "4 digit year and 2 digit month",
			ymd:   "1973-01",
			wantY: of(1973),
			wantM: of(1),
			wantD: nil,
		},
		{
			name:  "4 digit year and 1 digit month",
			ymd:   "2013-1",
			wantY: of(2013),
			wantM: of(1),
			wantD: nil,
		},
		{
			name:  "4 digit year and letters",
			ymd:   "2013-xx",
			wantY: of(2013),
			wantM: nil,
			wantD: nil,
		},
		{
			name:  "2 digit year and letters",
			ymd:   "98-x",
			wantY: of(1998),
			wantM: nil,
			wantD: nil,
		},
		{
			name:  "6 digit ymd",
			ymd:   "100322",
			wantY: of(2010),
			wantM: of(3),
			wantD: of(22),
		},
		{
			name:  "6 digit ym",
			ymd:   "201312",
			wantY: of(2013),
			wantM: of(12),
			wantD: nil,
		},
		{
			name:  "4 digit y 1 letter",
			ymd:   "1990x",
			wantY: of(1990),
			wantM: nil,
			wantD: nil,
		},
		{
			name:  "4 digit y 2 letter",
			ymd:   "1990xx",
			wantY: of(1990),
			wantM: nil,
			wantD: nil,
		},
		{
			name:  "4 digit year",
			ymd:   "2012",
			wantY: of(2012),
			wantM: nil,
			wantD: nil,
		},
		{
			name:  "7digit ymd",
			ymd:   "2012-2-12",
			wantY: of(2012),
			wantM: of(2),
			wantD: of(12),
		},
		{
			name:  "invalid day",
			ymd:   "2012-2-30",
			wantY: nil,
			wantM: nil,
			wantD: nil,
		},
		{
			name:  "leap year",
			ymd:   "2012-2-29",
			wantY: of(2012),
			wantM: of(2),
			wantD: of(29),
		},
		{
			name:  "zero day",
			ymd:   "2012-03-00",
			wantY: of(2012),
			wantM: of(3),
			wantD: nil,
		},
		{
			name:  "year month 6 digit",
			ymd:   "201203",
			wantY: of(2012),
			wantM: of(3),
			wantD: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotY, gotM, gotD := ParseYMD(tt.ymd)
			assert.Equalf(t, tt.wantY, gotY, "ParseYMD(%v)", tt.ymd)
			assert.Equalf(t, tt.wantM, gotM, "ParseYMD(%v)", tt.ymd)
			assert.Equalf(t, tt.wantD, gotD, "ParseYMD(%v)", tt.ymd)
		})
	}
}
