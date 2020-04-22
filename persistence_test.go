package squl

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPersistence(t *testing.T) {
	testCases := []struct {
		shouldFail  bool
		enumTypeStr string
		enumType    Persistence
	}{
		{true, unknownStr, Persistence(Unknown)},
		{false, "permanent", PersistencePermanent},
		{false, "unlogged", PersistenceUnlogged},
		{false, "temporary", PersistenceTemporary},
	}

	for i, testCase := range testCases {
		actual, err := NewPersistence(testCase.enumTypeStr)
		failMsg := fmt.Sprintf("testCase: %d %v", i, testCase)
		if (err != nil) != testCase.shouldFail {
			assert.Fail(t, failMsg, err)
		}
		assert.Equal(t, testCase.enumType, actual, failMsg)
	}
}

func TestPersistence_String(t *testing.T) {
	testCases := []struct {
		enumType    Persistence
		enumTypeStr string
	}{
		{Persistence(Unknown), unknownStr},
		{PersistencePermanent, "permanent"},
		{PersistenceUnlogged, "unlogged"},
		{PersistenceTemporary, "temporary"},
	}

	for i, testCase := range testCases {
		failMsg := fmt.Sprintf("testCase: %d %v", i, testCase)
		assert.Equal(t, testCase.enumTypeStr, testCase.enumType.String(), failMsg)
	}
}

func TestPersistence_UnmarshalJSON(t *testing.T) {
	testCases := []struct {
		shouldFail      bool
		enumTypeJSONStr string
		enumType        Persistence
	}{
		{true, unknownJSONStr, Persistence(Unknown)},
		{false, `"permanent"`, PersistencePermanent},
		{false, `"unlogged"`, PersistenceUnlogged},
		{false, `"temporary"`, PersistenceTemporary},
	}

	for i, testCase := range testCases {
		actual := Persistence(Unknown)
		err := actual.UnmarshalJSON([]byte(testCase.enumTypeJSONStr))
		failMsg := fmt.Sprintf("testCase: %d %v", i, testCase)
		if (err != nil) != testCase.shouldFail {
			assert.Fail(t, failMsg, err)
		}
		assert.Equal(t, testCase.enumType, actual, failMsg)
	}
}

func TestPersistence_MarshalJSON(t *testing.T) {
	testCases := []struct {
		enumType        Persistence
		enumTypeJSONStr string
	}{
		{Persistence(Unknown), unknownJSONStr},
		{PersistencePermanent, `"permanent"`},
		{PersistenceUnlogged, `"unlogged"`},
		{PersistenceTemporary, `"temporary"`},
	}

	for i, testCase := range testCases {
		failMsg := fmt.Sprintf("testCase: %d %v", i, testCase)
		actual, err := testCase.enumType.MarshalJSON()
		assert.Nil(t, err, failMsg)
		assert.Equal(t, testCase.enumTypeJSONStr, string(actual), failMsg)
	}
}
