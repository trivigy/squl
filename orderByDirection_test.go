package squl

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type OrderByDirectionSuite struct {
	suite.Suite
}

func (r *OrderByDirectionSuite) TestOrderByDirection_NewOrderByDirection() {
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
			assert.Fail(r.T(), failMsg, err)
		}
		assert.Equal(r.T(), testCase.enumType, actual, failMsg)
	}
}

func (r *OrderByDirectionSuite) TestOrderByDirection_String() {
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
		assert.Equal(r.T(), testCase.enumTypeStr, testCase.enumType.String(), failMsg)
	}
}

func (r *OrderByDirectionSuite) TestOrderByDirection_UnmarshalJSON() {
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
			assert.Fail(r.T(), failMsg, err)
		}
		assert.Equal(r.T(), testCase.enumType, actual, failMsg)
	}
}

func (r *OrderByDirectionSuite) TestUserType_MarshalJSON() {
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
		assert.Nil(r.T(), err, failMsg)
		assert.Equal(r.T(), testCase.enumTypeJSONStr, string(actual), failMsg)
	}
}

func TestOrderByDirectionSuite(t *testing.T) {
	suite.Run(t, new(OrderByDirectionSuite))
}
