package squl

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestColumnRef_dump(t *testing.T) {
	testCases := []struct {
		shouldFail bool
		output     string
		args       []interface{}
		columnRef  *ColumnRef
	}{
		{
			false,
			"*",
			[]interface{}{},
			&ColumnRef{Fields: "*"},
		},
		{
			false,
			"col1",
			[]interface{}{},
			&ColumnRef{Fields: "col1"},
		},
		{
			false,
			"col1",
			[]interface{}{},
			&ColumnRef{Fields: []string{"col1"}},
		},
		{
			false,
			"col1.col2",
			[]interface{}{},
			&ColumnRef{Fields: []string{"col1", "col2"}},
		},
		{
			false,
			"col1.col2",
			[]interface{}{},
			&ColumnRef{Fields: []string{"col1", "col2"}},
		},
		{
			false,
			"col1.col2[1]",
			[]interface{}{},
			&ColumnRef{Fields: []interface{}{"col1", "col2", 1}},
		},
		{
			false,
			"col1[1][3].col2",
			[]interface{}{},
			&ColumnRef{Fields: []interface{}{"col1", 1, 3, "col2"}},
		},
		{
			false,
			"$1.col2[$2]",
			[]interface{}{"col1", 1},
			&ColumnRef{Fields: []interface{}{
				Var{"col1"},
				"col2",
				Var{1},
			}},
		},
		{
			true,
			"",
			[]interface{}{},
			&ColumnRef{Fields: 11},
		},
	}

	for i, testCase := range testCases {
		counter := &ordinalMarker{}
		actual, err := testCase.columnRef.dump(counter)
		failMsg := fmt.Sprintf("testCase: %d %v", i, testCase)
		if (err != nil) != testCase.shouldFail {
			assert.Fail(t, failMsg, err)
		}
		assert.EqualValues(t, testCase.args, counter.args())
		assert.Equal(t, testCase.output, actual, failMsg)
	}
}
