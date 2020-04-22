package squl

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBoolExprType(t *testing.T) {
	testCases := []struct {
		shouldFail  bool
		enumTypeStr string
		enumType    BoolExprType
	}{
		{true, unknownStr, BoolExprType(Unknown)},
		{false, "and", BoolExprTypeAnd},
		{false, "or", BoolExprTypeOr},
	}

	for i, testCase := range testCases {
		actual, err := NewBoolExprType(testCase.enumTypeStr)
		failMsg := fmt.Sprintf("testCase: %d %v", i, testCase)
		if (err != nil) != testCase.shouldFail {
			assert.Fail(t, failMsg, err)
		}
		assert.Equal(t, testCase.enumType, actual, failMsg)
	}
}

func TestBoolExprType_String(t *testing.T) {
	testCases := []struct {
		enumType    BoolExprType
		enumTypeStr string
	}{
		{BoolExprType(Unknown), unknownStr},
		{BoolExprTypeAnd, "and"},
		{BoolExprTypeOr, "or"},
	}

	for i, testCase := range testCases {
		failMsg := fmt.Sprintf("testCase: %d %v", i, testCase)
		assert.Equal(t, testCase.enumTypeStr, testCase.enumType.String(), failMsg)
	}
}

func TestBoolExprType_UnmarshalJSON(t *testing.T) {
	testCases := []struct {
		shouldFail      bool
		enumTypeJSONStr string
		enumType        BoolExprType
	}{
		{true, unknownJSONStr, BoolExprType(Unknown)},
		{false, `"and"`, BoolExprTypeAnd},
		{false, `"or"`, BoolExprTypeOr},
	}

	for i, testCase := range testCases {
		actual := BoolExprType(Unknown)
		err := actual.UnmarshalJSON([]byte(testCase.enumTypeJSONStr))
		failMsg := fmt.Sprintf("testCase: %d %v", i, testCase)
		if (err != nil) != testCase.shouldFail {
			assert.Fail(t, failMsg, err)
		}
		assert.Equal(t, testCase.enumType, actual, failMsg)
	}
}

func TestBoolExprType_MarshalJSON(t *testing.T) {
	testCases := []struct {
		enumType        BoolExprType
		enumTypeJSONStr string
	}{
		{BoolExprType(Unknown), unknownJSONStr},
		{BoolExprTypeAnd, `"and"`},
		{BoolExprTypeOr, `"or"`},
	}

	for i, testCase := range testCases {
		failMsg := fmt.Sprintf("testCase: %d %v", i, testCase)
		actual, err := testCase.enumType.MarshalJSON()
		assert.Nil(t, err, failMsg)
		assert.Equal(t, testCase.enumTypeJSONStr, string(actual), failMsg)
	}
}
