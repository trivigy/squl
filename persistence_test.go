package squl

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type PersistenceSuite struct {
	suite.Suite
}

func (r *PersistenceSuite) TestUserType_NewUserType() {
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
			assert.Fail(r.T(), failMsg, err)
		}
		assert.Equal(r.T(), testCase.enumType, actual, failMsg)
	}
}

func (r *PersistenceSuite) TestUserType_String() {
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
		assert.Equal(r.T(), testCase.enumTypeStr, testCase.enumType.String(), failMsg)
	}
}

func (r *PersistenceSuite) TestUserType_UnmarshalJSON() {
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
			assert.Fail(r.T(), failMsg, err)
		}
		assert.Equal(r.T(), testCase.enumType, actual, failMsg)
	}
}

func (r *PersistenceSuite) TestUserType_MarshalJSON() {
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
		assert.Nil(r.T(), err, failMsg)
		assert.Equal(r.T(), testCase.enumTypeJSONStr, string(actual), failMsg)
	}
}

func TestPersistenceSuite(t *testing.T) {
	suite.Run(t, new(PersistenceSuite))
}
