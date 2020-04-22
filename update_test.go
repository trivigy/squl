package squl

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpdate_dump(t *testing.T) {
	testCases := []struct {
		shouldFail bool
		output     string
		args       []interface{}
		expected   *Update
	}{
		{
			false,
			"UPDATE products SET price = price * 1.1 WHERE price <= 99.99 RETURNING name,price AS new_price",
			[]interface{}{},
			&Update{
				Relation: &RangeVar{
					Name: "products",
				},
				Targets: &ResTarget{
					Name: "price",
					Value: &Expr{
						Type: ExprTypeOp,
						Name: "*",
						LHS:  &ColumnRef{Fields: "price"},
						RHS:  &Const{1.1},
					},
				},
				Where: &Expr{
					Type: ExprTypeOp,
					Name: "<=",
					LHS:  &ColumnRef{Fields: "price"},
					RHS:  &Const{99.99},
				},
				Returning: &List{
					&ResTarget{
						Value: &ColumnRef{Fields: "name"},
					},
					&ResTarget{
						Value: &ColumnRef{Fields: "price"},
						Alias: "new_price",
					},
				},
			},
		},
		{
			false,
			"UPDATE stock SET retail = cost * ((retail / cost) + 0.1::numeric)",
			[]interface{}{},
			&Update{
				Relation: &RangeVar{
					Name: "stock",
				},
				Targets: &ResTarget{
					Name: "retail",
					Value: &Expr{
						Type: ExprTypeOp,
						Name: "*",
						LHS:  &ColumnRef{Fields: "cost"},
						RHS: &Expr{
							Type: ExprTypeOp,
							Wrap: true,
							Name: "+",
							LHS: &Expr{
								Type: ExprTypeOp,
								Wrap: true,
								Name: "/",
								LHS:  &ColumnRef{Fields: "retail"},
								RHS:  &ColumnRef{Fields: "cost"},
							},
							RHS: &TypeCast{
								Arg:  &Const{0.1},
								Type: "numeric",
							},
						},
					},
				},
			},
		},
		{
			false,
			"UPDATE stock SET retail = stock_backup.retail FROM stock_backup WHERE stock.isbn = stock_backup.isbn",
			[]interface{}{},
			&Update{
				Relation: &RangeVar{
					Name: "stock",
				},
				Targets: &ResTarget{
					Name:  "retail",
					Value: &ColumnRef{Fields: []string{"stock_backup", "retail"}},
				},
				From: &RangeVar{
					Name: "stock_backup",
				},
				Where: &Expr{
					Type: ExprTypeOp,
					Name: "=",
					LHS:  &ColumnRef{Fields: []string{"stock", "isbn"}},
					RHS:  &ColumnRef{Fields: []string{"stock_backup", "isbn"}},
				},
			},
		},
	}

	for i, testCase := range testCases {
		counter := &ordinalMarker{}
		actual, err := testCase.expected.dump(counter)
		failMsg := fmt.Sprintf("testCase: %d %v", i, testCase)
		if (err != nil) != testCase.shouldFail {
			assert.Fail(t, failMsg, err)
		}
		assert.EqualValues(t, testCase.args, counter.args())
		assert.Equal(t, testCase.output, actual, failMsg)
	}
}
