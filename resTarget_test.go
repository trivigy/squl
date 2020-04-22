package squl

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResTarget_dump(t *testing.T) {
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
			assert.Fail(t, failMsg, err)
		}
		assert.EqualValues(t, testCase.args, counter.args())
		assert.Equal(t, testCase.output, actual, failMsg)
	}
}
