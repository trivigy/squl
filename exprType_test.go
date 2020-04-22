package squl

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewExprType(t *testing.T) {
	testCases := []struct {
		shouldFail  bool
		enumTypeStr string
		enumType    ExprType
	}{
		{true, unknownStr, ExprType(Unknown)},
		{false, "op", ExprTypeOp},
		{false, "opAny", ExprTypeOpAny},
		{false, "opAll", ExprTypeOpAll},
		{false, "distinct", ExprTypeDistinct},
		{false, "notDistinct", ExprTypeNotDistinct},
		{false, "nullIf", ExprTypeNullIf},
		{false, "of", ExprTypeOf},
		{false, "in", ExprTypeIn},
		{false, "like", ExprTypeLike},
		{false, "iLike", ExprTypeILike},
		{false, "similar", ExprTypeSimilar},
		{false, "between", ExprTypeBetween},
		{false, "notBetween", ExprTypeNotBetween},
		{false, "betweenSym", ExprTypeBetweenSym},
		{false, "notBetweenSym", ExprTypeNotBetweenSym},
		{false, "paren", ExprTypeParen},
	}

	for i, testCase := range testCases {
		actual, err := NewExprType(testCase.enumTypeStr)
		failMsg := fmt.Sprintf("testCase: %d %v", i, testCase)
		if (err != nil) != testCase.shouldFail {
			assert.Fail(t, failMsg, err)
		}
		assert.Equal(t, testCase.enumType, actual, failMsg)
	}
}

func TestExprType_String(t *testing.T) {
	testCases := []struct {
		enumType    ExprType
		enumTypeStr string
	}{
		{ExprType(Unknown), unknownStr},
		{ExprTypeOp, "op"},
		{ExprTypeOpAny, "opAny"},
		{ExprTypeOpAll, "opAll"},
		{ExprTypeDistinct, "distinct"},
		{ExprTypeNotDistinct, "notDistinct"},
		{ExprTypeNullIf, "nullIf"},
		{ExprTypeOf, "of"},
		{ExprTypeIn, "in"},
		{ExprTypeLike, "like"},
		{ExprTypeILike, "iLike"},
		{ExprTypeSimilar, "similar"},
		{ExprTypeBetween, "between"},
		{ExprTypeNotBetween, "notBetween"},
		{ExprTypeBetweenSym, "betweenSym"},
		{ExprTypeNotBetweenSym, "notBetweenSym"},
		{ExprTypeParen, "paren"},
	}

	for i, testCase := range testCases {
		failMsg := fmt.Sprintf("testCase: %d %v", i, testCase)
		assert.Equal(t, testCase.enumTypeStr, testCase.enumType.String(), failMsg)
	}
}

func TestExprType_UnmarshalJSON(t *testing.T) {
	testCases := []struct {
		shouldFail      bool
		enumTypeJSONStr string
		enumType        ExprType
	}{
		{true, unknownJSONStr, ExprType(Unknown)},
		{false, `"op"`, ExprTypeOp},
		{false, `"opAny"`, ExprTypeOpAny},
		{false, `"opAll"`, ExprTypeOpAll},
		{false, `"distinct"`, ExprTypeDistinct},
		{false, `"notDistinct"`, ExprTypeNotDistinct},
		{false, `"nullIf"`, ExprTypeNullIf},
		{false, `"of"`, ExprTypeOf},
		{false, `"in"`, ExprTypeIn},
		{false, `"like"`, ExprTypeLike},
		{false, `"iLike"`, ExprTypeILike},
		{false, `"similar"`, ExprTypeSimilar},
		{false, `"between"`, ExprTypeBetween},
		{false, `"notBetween"`, ExprTypeNotBetween},
		{false, `"betweenSym"`, ExprTypeBetweenSym},
		{false, `"notBetweenSym"`, ExprTypeNotBetweenSym},
		{false, `"paren"`, ExprTypeParen},
	}

	for i, testCase := range testCases {
		actual := ExprType(Unknown)
		err := actual.UnmarshalJSON([]byte(testCase.enumTypeJSONStr))
		failMsg := fmt.Sprintf("testCase: %d %v", i, testCase)
		if (err != nil) != testCase.shouldFail {
			assert.Fail(t, failMsg, err)
		}
		assert.Equal(t, testCase.enumType, actual, failMsg)
	}
}

func TestExprType_MarshalJSON(t *testing.T) {
	testCases := []struct {
		enumType        ExprType
		enumTypeJSONStr string
	}{
		{ExprType(Unknown), unknownJSONStr},
		{ExprTypeOp, `"op"`},
		{ExprTypeOpAny, `"opAny"`},
		{ExprTypeOpAll, `"opAll"`},
		{ExprTypeDistinct, `"distinct"`},
		{ExprTypeNotDistinct, `"notDistinct"`},
		{ExprTypeNullIf, `"nullIf"`},
		{ExprTypeOf, `"of"`},
		{ExprTypeIn, `"in"`},
		{ExprTypeLike, `"like"`},
		{ExprTypeILike, `"iLike"`},
		{ExprTypeSimilar, `"similar"`},
		{ExprTypeBetween, `"between"`},
		{ExprTypeNotBetween, `"notBetween"`},
		{ExprTypeBetweenSym, `"betweenSym"`},
		{ExprTypeNotBetweenSym, `"notBetweenSym"`},
		{ExprTypeParen, `"paren"`},
	}

	for i, testCase := range testCases {
		failMsg := fmt.Sprintf("testCase: %d %v", i, testCase)
		actual, err := testCase.enumType.MarshalJSON()
		assert.Nil(t, err, failMsg)
		assert.Equal(t, testCase.enumTypeJSONStr, string(actual), failMsg)
	}
}
