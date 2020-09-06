package squl

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestExpr_dump(t *testing.T) {
	testCases := []struct {
		shouldFail bool
		output     string
		args       []interface{}
		expr       *Expr
	}{
		{
			false,
			"col1 = col2",
			[]interface{}{},
			&Expr{
				Type: ExprTypeOp,
				Name: "=",
				LHS:  &ColumnRef{Fields: "col1"},
				RHS:  &ColumnRef{Fields: "col2"},
			},
		},
		{
			false,
			"4 OPERATOR(pg_catalog.*) 4",
			[]interface{}{},
			&Expr{
				Type: ExprTypeOp,
				Name: []string{"pg_catalog", "*"},
				LHS:  &Const{Value: 4},
				RHS:  &Const{Value: 4},
			},
		},
		{
			false,
			"(temp_hi + temp_lo) / 2",
			[]interface{}{},
			&Expr{
				Type: ExprTypeOp,
				Name: "/",
				LHS: &Expr{
					Type: ExprTypeOp,
					Wrap: true,
					Name: "+",
					LHS: &ColumnRef{
						Fields: "temp_hi",
					},
					RHS: &ColumnRef{
						Fields: "temp_lo",
					},
				},
				RHS: &Const{Value: 2},
			},
		},
		{
			false,
			"col1 = 'b720bca3-19f0-4ff4-97d2-53d865e74ab7'",
			[]interface{}{},
			&Expr{
				Type: ExprTypeOp,
				Name: "=",
				LHS:  &ColumnRef{Fields: "col1"},
				RHS: func() *Const {
					guid, _ := uuid.Parse("b720bca3-19f0-4ff4-97d2-53d865e74ab7")
					return &Const{Value: guid}
				}(),
			},
		},
		{
			false,
			"unittest IS NOT NULL",
			[]interface{}{},
			&Expr{
				Type: ExprTypeOf,
				Name: "<>",
				LHS:  &ColumnRef{Fields: "unittest"},
				RHS:  &Null{},
			},
		},
		// {
		// 	false,
		// 	"max(11, $1) = 20",
		// 	[]interface{}{20},
		// 	&Expr{
		// 		Type: ExprTypeOp,
		// 		Alias: []string{"="},
		// 		LHS: FuncCall{
		// 			Alias: "max",
		// 			Args: []interface{}{11, Var{Value: 20}},
		// 		},
		// 		RHS: 20,
		// 	},
		// },
		// {
		// 	false,
		// 	"'100'::integer != 100",
		// 	[]interface{}{20},
		// 	&Expr{
		// 		Type: ExprTypeOp,
		// 		Alias: []string{"<>"},
		// 		LHS: TypeCast{
		// 			Arg:  "100",
		// 			Type: "integer",
		// 		},
		// 		RHS: 100,
		// 	},
		// },
		// {false, "col1", &ColumnRef{Fields: "col1"}},
		// {false, "col1", &ColumnRef{Fields: []string{"col1"}}},
		// {false, "col1.col2", &ColumnRef{Fields: []string{"col1", "col2"}}},
		// {true, "", &ColumnRef{Fields: 11}},
	}

	for i, testCase := range testCases {
		counter := &ordinalMarker{}
		actual, err := testCase.expr.dump(counter)
		failMsg := fmt.Sprintf("testCase: %d %v", i, testCase)
		if (err != nil) != testCase.shouldFail {
			assert.Fail(t, failMsg, err)
		}
		assert.EqualValues(t, testCase.args, counter.args())
		assert.Equal(t, testCase.output, actual, failMsg)
	}

	// query, args, err := (&Expr{
	// 	Type: ExprTypeOp,
	// 	Alias: []string{"="},
	// 	LHS:  ColumnRef{},
	// 	RHS:  ColumnRef{},
	// }).dump()
}
