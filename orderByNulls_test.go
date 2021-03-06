package squl

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewOrderByNulls(t *testing.T) {
	testCases := []struct {
		shouldFail  bool
		enumTypeStr string
		enumType    OrderByNulls
	}{
		{true, unknownStr, OrderByNulls(Unknown)},
		{false, "first", OrderByNullsFirst},
		{false, "last", OrderByNullsLast},
	}

	for i, testCase := range testCases {
		actual, err := NewOrderByNulls(testCase.enumTypeStr)
		failMsg := fmt.Sprintf("testCase: %d %v", i, testCase)
		if (err != nil) != testCase.shouldFail {
			assert.Fail(t, failMsg, err)
		}
		assert.Equal(t, testCase.enumType, actual, failMsg)
	}
}

func TestOrderByNulls_String(t *testing.T) {
	testCases := []struct {
		enumType    OrderByNulls
		enumTypeStr string
	}{
		{OrderByNulls(Unknown), unknownStr},
		{OrderByNullsFirst, "first"},
		{OrderByNullsLast, "last"},
	}

	for i, testCase := range testCases {
		failMsg := fmt.Sprintf("testCase: %d %v", i, testCase)
		assert.Equal(t, testCase.enumTypeStr, testCase.enumType.String(), failMsg)
	}
}

func TestOrderByNulls_UnmarshalJSON(t *testing.T) {
	testCases := []struct {
		shouldFail      bool
		enumTypeJSONStr string
		enumType        OrderByNulls
	}{
		{true, unknownJSONStr, OrderByNulls(Unknown)},
		{false, `"first"`, OrderByNullsFirst},
		{false, `"last"`, OrderByNullsLast},
	}

	for i, testCase := range testCases {
		actual := OrderByNulls(Unknown)
		err := actual.UnmarshalJSON([]byte(testCase.enumTypeJSONStr))
		failMsg := fmt.Sprintf("testCase: %d %v", i, testCase)
		if (err != nil) != testCase.shouldFail {
			assert.Fail(t, failMsg, err)
		}
		assert.Equal(t, testCase.enumType, actual, failMsg)
	}
}

func TestOrderByNulls_MarshalJSON(t *testing.T) {
	testCases := []struct {
		enumType        OrderByNulls
		enumTypeJSONStr string
	}{
		{OrderByNulls(Unknown), unknownJSONStr},
		{OrderByNullsFirst, `"first"`},
		{OrderByNullsLast, `"last"`},
	}

	for i, testCase := range testCases {
		failMsg := fmt.Sprintf("testCase: %d %v", i, testCase)
		actual, err := testCase.enumType.MarshalJSON()
		assert.Nil(t, err, failMsg)
		assert.Equal(t, testCase.enumTypeJSONStr, string(actual), failMsg)
	}
}
