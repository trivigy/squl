package squl

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewJoinType(t *testing.T) {
	testCases := []struct {
		shouldFail  bool
		enumTypeStr string
		enumType    JoinType
	}{
		{true, unknownStr, JoinType(Unknown)},
		{false, "default", JoinTypeDefault},
		{false, "inner", JoinTypeInner},
		{false, "left", JoinTypeLeft},
		{false, "outerLeft", JoinTypeOuterLeft},
		{false, "right", JoinTypeRight},
		{false, "outerRight", JoinTypeOuterRight},
		{false, "full", JoinTypeFull},
		{false, "outerFull", JoinTypeOuterFull},
		{false, "cross", JoinTypeCross},
	}

	for i, testCase := range testCases {
		actual, err := NewJoinType(testCase.enumTypeStr)
		failMsg := fmt.Sprintf("testCase: %d %v", i, testCase)
		if (err != nil) != testCase.shouldFail {
			assert.Fail(t, failMsg, err)
		}
		assert.Equal(t, testCase.enumType, actual, failMsg)
	}
}

func TestJoinType_String(t *testing.T) {
	testCases := []struct {
		enumType    JoinType
		enumTypeStr string
	}{
		{JoinType(Unknown), unknownStr},
		{JoinTypeDefault, "default"},
		{JoinTypeInner, "inner"},
		{JoinTypeLeft, "left"},
		{JoinTypeOuterLeft, "outerLeft"},
		{JoinTypeRight, "right"},
		{JoinTypeOuterRight, "outerRight"},
		{JoinTypeFull, "full"},
		{JoinTypeOuterFull, "outerFull"},
		{JoinTypeCross, "cross"},
	}

	for i, testCase := range testCases {
		failMsg := fmt.Sprintf("testCase: %d %v", i, testCase)
		assert.Equal(t, testCase.enumTypeStr, testCase.enumType.String(), failMsg)
	}
}

func TestJoinType_UnmarshalJSON(t *testing.T) {
	testCases := []struct {
		shouldFail      bool
		enumTypeJSONStr string
		enumType        JoinType
	}{
		{true, unknownJSONStr, JoinType(Unknown)},
		{false, `"default"`, JoinTypeDefault},
		{false, `"inner"`, JoinTypeInner},
		{false, `"left"`, JoinTypeLeft},
		{false, `"outerLeft"`, JoinTypeOuterLeft},
		{false, `"right"`, JoinTypeRight},
		{false, `"outerRight"`, JoinTypeOuterRight},
		{false, `"full"`, JoinTypeFull},
		{false, `"outerFull"`, JoinTypeOuterFull},
		{false, `"cross"`, JoinTypeCross},
	}

	for i, testCase := range testCases {
		actual := JoinType(Unknown)
		err := actual.UnmarshalJSON([]byte(testCase.enumTypeJSONStr))
		failMsg := fmt.Sprintf("testCase: %d %v", i, testCase)
		if (err != nil) != testCase.shouldFail {
			assert.Fail(t, failMsg, err)
		}
		assert.Equal(t, testCase.enumType, actual, failMsg)
	}
}

func TestJoinType_MarshalJSON(t *testing.T) {
	testCases := []struct {
		enumType        JoinType
		enumTypeJSONStr string
	}{
		{JoinType(Unknown), unknownJSONStr},
		{JoinTypeDefault, `"default"`},
		{JoinTypeInner, `"inner"`},
		{JoinTypeLeft, `"left"`},
		{JoinTypeOuterLeft, `"outerLeft"`},
		{JoinTypeRight, `"right"`},
		{JoinTypeOuterRight, `"outerRight"`},
		{JoinTypeFull, `"full"`},
		{JoinTypeOuterFull, `"outerFull"`},
		{JoinTypeCross, `"cross"`},
	}

	for i, testCase := range testCases {
		failMsg := fmt.Sprintf("testCase: %d %v", i, testCase)
		actual, err := testCase.enumType.MarshalJSON()
		assert.Nil(t, err, failMsg)
		assert.Equal(t, testCase.enumTypeJSONStr, string(actual), failMsg)
	}
}
