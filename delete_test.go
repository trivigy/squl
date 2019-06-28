package squl

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type DeleteSuite struct {
	suite.Suite
}

func (r *DeleteSuite) TestDeleteSuiteDump() {
	testCases := []struct {
		shouldFail bool
		output     string
		args       []interface{}
		expected   *Delete
	}{
		{
			false,
			"DELETE FROM products WHERE obsoletion_date = 'today' RETURNING *",
			[]interface{}{},
			&Delete{
				Relation: &RangeVar{
					Name: "products",
				},
				Where: &Expr{
					Type: ExprTypeOp,
					Name: "=",
					LHS:  &ColumnRef{Fields: "obsoletion_date"},
					RHS:  &Const{Value: "today"},
				},
				Returning: &ResTarget{
					Value: &ColumnRef{Fields: "*"},
				},
			},
		},
		{
			false,
			"DELETE FROM link USING link_tmp WHERE link.id = link_tmp.id",
			[]interface{}{},
			&Delete{
				Relation: &RangeVar{
					Name: "link",
				},
				Using: &RangeVar{
					Name: "link_tmp",
				},
				Where: &Expr{
					Type: ExprTypeOp,
					Name: "=",
					LHS:  &ColumnRef{Fields: []string{"link", "id"}},
					RHS:  &ColumnRef{Fields: []string{"link_tmp", "id"}},
				},
			},
		},
	}

	for i, testCase := range testCases {
		counter := &ordinalMarker{}
		actual, err := testCase.expected.dump(counter)
		failMsg := fmt.Sprintf("testCase: %d %v", i, testCase)
		if (err != nil) != testCase.shouldFail {
			assert.Fail(r.T(), failMsg, err)
		}
		assert.EqualValues(r.T(), testCase.args, counter.args())
		assert.Equal(r.T(), testCase.output, actual, failMsg)
	}
}

func TestDeleteSuite(t *testing.T) {
	suite.Run(t, new(DeleteSuite))
}
