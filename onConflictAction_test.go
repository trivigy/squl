package squl

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewOnConflictAction(t *testing.T) {
	testCases := []struct {
		shouldFail  bool
		enumTypeStr string
		enumType    OnConflictAction
	}{
		{true, unknownStr, OnConflictAction(Unknown)},
		{false, "nothing", OnConflictNothing},
		{false, "update", OnConflictUpdate},
	}

	for i, testCase := range testCases {
		actual, err := NewOnConflictAction(testCase.enumTypeStr)
		failMsg := fmt.Sprintf("testCase: %d %v", i, testCase)
		if (err != nil) != testCase.shouldFail {
			assert.Fail(t, failMsg, err)
		}
		assert.Equal(t, testCase.enumType, actual, failMsg)
	}
}

func TestOnConflictAction_String(t *testing.T) {
	testCases := []struct {
		enumType    OnConflictAction
		enumTypeStr string
	}{
		{OnConflictAction(Unknown), unknownStr},
		{OnConflictNothing, "nothing"},
		{OnConflictUpdate, "update"},
	}

	for i, testCase := range testCases {
		failMsg := fmt.Sprintf("testCase: %d %v", i, testCase)
		assert.Equal(t, testCase.enumTypeStr, testCase.enumType.String(), failMsg)
	}
}

func TestOnConflictAction_UnmarshalJSON(t *testing.T) {
	testCases := []struct {
		shouldFail      bool
		enumTypeJSONStr string
		enumType        OnConflictAction
	}{
		{true, unknownJSONStr, OnConflictAction(Unknown)},
		{false, `"nothing"`, OnConflictNothing},
		{false, `"update"`, OnConflictUpdate},
	}

	for i, testCase := range testCases {
		actual := OnConflictAction(Unknown)
		err := actual.UnmarshalJSON([]byte(testCase.enumTypeJSONStr))
		failMsg := fmt.Sprintf("testCase: %d %v", i, testCase)
		if (err != nil) != testCase.shouldFail {
			assert.Fail(t, failMsg, err)
		}
		assert.Equal(t, testCase.enumType, actual, failMsg)
	}
}

func TestOnConflictAction_MarshalJSON(t *testing.T) {
	testCases := []struct {
		enumType        OnConflictAction
		enumTypeJSONStr string
	}{
		{OnConflictAction(Unknown), unknownJSONStr},
		{OnConflictNothing, `"nothing"`},
		{OnConflictUpdate, `"update"`},
	}

	for i, testCase := range testCases {
		failMsg := fmt.Sprintf("testCase: %d %v", i, testCase)
		actual, err := testCase.enumType.MarshalJSON()
		assert.Nil(t, err, failMsg)
		assert.Equal(t, testCase.enumTypeJSONStr, string(actual), failMsg)
	}
}
