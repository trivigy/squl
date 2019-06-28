package squl

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type BoolExprTypeSuite struct {
	suite.Suite
}

func (r *BoolExprTypeSuite) TestBoolExprType_NewBoolExprType() {
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
			assert.Fail(r.T(), failMsg, err)
		}
		assert.Equal(r.T(), testCase.enumType, actual, failMsg)
	}
}

func (r *BoolExprTypeSuite) TestBoolExprType_String() {
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
		assert.Equal(r.T(), testCase.enumTypeStr, testCase.enumType.String(), failMsg)
	}
}

func (r *BoolExprTypeSuite) TestBoolExprType_UnmarshalJSON() {
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
			assert.Fail(r.T(), failMsg, err)
		}
		assert.Equal(r.T(), testCase.enumType, actual, failMsg)
	}
}

func (r *BoolExprTypeSuite) TestBoolExprType_MarshalJSON() {
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
		assert.Nil(r.T(), err, failMsg)
		assert.Equal(r.T(), testCase.enumTypeJSONStr, string(actual), failMsg)
	}
}

func TestBoolExprTypeSuite(t *testing.T) {
	suite.Run(t, new(BoolExprTypeSuite))
}
