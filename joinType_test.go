package squl

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type JoinTypeSuite struct {
	suite.Suite
}

func (r *JoinTypeSuite) TestJoinType_NewJoinType() {
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
			assert.Fail(r.T(), failMsg, err)
		}
		assert.Equal(r.T(), testCase.enumType, actual, failMsg)
	}
}

func (r *JoinTypeSuite) TestJoinType_String() {
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
		assert.Equal(r.T(), testCase.enumTypeStr, testCase.enumType.String(), failMsg)
	}
}

func (r *JoinTypeSuite) TestJoinType_UnmarshalJSON() {
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
			assert.Fail(r.T(), failMsg, err)
		}
		assert.Equal(r.T(), testCase.enumType, actual, failMsg)
	}
}

func (r *JoinTypeSuite) TestJoinType_MarshalJSON() {
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
		assert.Nil(r.T(), err, failMsg)
		assert.Equal(r.T(), testCase.enumTypeJSONStr, string(actual), failMsg)
	}
}

func TestJoinTypeSuite(t *testing.T) {
	suite.Run(t, new(JoinTypeSuite))
}
