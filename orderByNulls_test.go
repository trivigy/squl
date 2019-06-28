package squl

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type OrderByNullsSuite struct {
	suite.Suite
}

func (r *OrderByNullsSuite) TestOrderByNulls_NewOrderByNulls() {
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
			assert.Fail(r.T(), failMsg, err)
		}
		assert.Equal(r.T(), testCase.enumType, actual, failMsg)
	}
}

func (r *OrderByNullsSuite) TestOrderByNulls_String() {
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
		assert.Equal(r.T(), testCase.enumTypeStr, testCase.enumType.String(), failMsg)
	}
}

func (r *OrderByNullsSuite) TestOrderByNulls_UnmarshalJSON() {
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
			assert.Fail(r.T(), failMsg, err)
		}
		assert.Equal(r.T(), testCase.enumType, actual, failMsg)
	}
}

func (r *OrderByNullsSuite) TestUserType_MarshalJSON() {
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
		assert.Nil(r.T(), err, failMsg)
		assert.Equal(r.T(), testCase.enumTypeJSONStr, string(actual), failMsg)
	}
}

func TestOrderByNullsSuite(t *testing.T) {
	suite.Run(t, new(OrderByNullsSuite))
}
