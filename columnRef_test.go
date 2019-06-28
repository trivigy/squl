package squl

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ColumnRefSuite struct {
	suite.Suite
}

func (r *ColumnRefSuite) TestColumnRefDump() {
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
			assert.Fail(r.T(), failMsg, err)
		}
		assert.EqualValues(r.T(), testCase.args, counter.args())
		assert.Equal(r.T(), testCase.output, actual, failMsg)
	}
}

func TestColumnRefSuite(t *testing.T) {
	suite.Run(t, new(ColumnRefSuite))
}
