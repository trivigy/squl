package squl

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewOrderByDirection(t *testing.T) {
	testCases := []struct {
		shouldFail  bool
		enumTypeStr string
		enumType    OrderByDirection
	}{
		{true, unknownStr, OrderByDirection(Unknown)},
		{false, "asc", OrderByDirectionAsc},
		{false, "desc", OrderByDirectionDesc},
		{false, "using", OrderByDirectionUsing},
	}

	for i, testCase := range testCases {
		actual, err := NewOrderByDirection(testCase.enumTypeStr)
		failMsg := fmt.Sprintf("testCase: %d %v", i, testCase)
		if (err != nil) != testCase.shouldFail {
			assert.Fail(t, failMsg, err)
		}
		assert.Equal(t, testCase.enumType, actual, failMsg)
	}
}

func TestOrderByDirection_String(t *testing.T) {
	testCases := []struct {
		enumType    OrderByDirection
		enumTypeStr string
	}{
		{OrderByDirection(Unknown), unknownStr},
		{OrderByDirectionAsc, "asc"},
		{OrderByDirectionDesc, "desc"},
		{OrderByDirectionUsing, "using"},
	}

	for i, testCase := range testCases {
		failMsg := fmt.Sprintf("testCase: %d %v", i, testCase)
		assert.Equal(t, testCase.enumTypeStr, testCase.enumType.String(), failMsg)
	}
}

func TestOrderByDirection_UnmarshalJSON(t *testing.T) {
	testCases := []struct {
		shouldFail      bool
		enumTypeJSONStr string
		enumType        OrderByDirection
	}{
		{true, unknownJSONStr, OrderByDirection(Unknown)},
		{false, `"asc"`, OrderByDirectionAsc},
		{false, `"desc"`, OrderByDirectionDesc},
		{false, `"using"`, OrderByDirectionUsing},
	}

	for i, testCase := range testCases {
		actual := OrderByDirection(Unknown)
		err := actual.UnmarshalJSON([]byte(testCase.enumTypeJSONStr))
		failMsg := fmt.Sprintf("testCase: %d %v", i, testCase)
		if (err != nil) != testCase.shouldFail {
			assert.Fail(t, failMsg, err)
		}
		assert.Equal(t, testCase.enumType, actual, failMsg)
	}
}

func TestOrderByDirection_MarshalJSON(t *testing.T) {
	testCases := []struct {
		enumType        OrderByDirection
		enumTypeJSONStr string
	}{
		{OrderByDirection(Unknown), unknownJSONStr},
		{OrderByDirectionAsc, `"asc"`},
		{OrderByDirectionDesc, `"desc"`},
		{OrderByDirectionUsing, `"using"`},
	}

	for i, testCase := range testCases {
		failMsg := fmt.Sprintf("testCase: %d %v", i, testCase)
		actual, err := testCase.enumType.MarshalJSON()
		assert.Nil(t, err, failMsg)
		assert.Equal(t, testCase.enumTypeJSONStr, string(actual), failMsg)
	}
}
