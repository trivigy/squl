package squl

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ResTargetSuite struct {
	suite.Suite
}

func (r *ResTargetSuite) TestResTargetDump() {
	testCases := []struct {
		shouldFail bool
		output     string
		args       []interface{}
		expected   *ResTarget
	}{
		{
			false,
			"one.two AS u",
			[]interface{}{},
			&ResTarget{
				Value: &ColumnRef{
					Fields: []string{"one", "two"},
				},
				Alias: "u",
			},
		},
		{
			false,
			"one.two[$1].four AS c",
			[]interface{}{3},
			&ResTarget{
				Value: &ColumnRef{
					Fields: []interface{}{"one", "two", Var{3}, "four"},
				},
				Alias: "c",
			},
		},
	}

	for i, testCase := range testCases {
		counter := &ordinalMarker{}
		actual, err := testCase.expected.dump(counter)
		failMsg := fmt.Sprintf("testCase: %d %v", i, testCase)
		if (err != nil) != testCase.shouldFail {
			assert.Fail(r.T(), failMsg, err)
		}
		assert.EqualValues(r.T(), testCase.args, counter.args())
		assert.Equal(r.T(), testCase.output, actual, failMsg)
	}
}

func TestResTargetSuite(t *testing.T) {
	suite.Run(t, new(ResTargetSuite))
}
