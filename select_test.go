package squl

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSelect_dump(t *testing.T) {
	testCases := []struct {
		shouldFail bool
		output     string
		args       []interface{}
		expected   *Select
	}{
		{
			false,
			"SELECT * FROM weather",
			[]interface{}{},
			&Select{
				Targets: &ResTarget{
					Value: &ColumnRef{
						Fields: "*",
					},
				},
				From: &RangeVar{
					Name: "weather",
				},
			},
		},
		{
			false,
			"SELECT city,(temp_hi + temp_lo) / 2 AS temp_avg,date FROM weather",
			[]interface{}{},
			&Select{
				Targets: &List{
					&ResTarget{
						Value: &ColumnRef{
							Fields: "city",
						},
					},
					&ResTarget{
						Value: &Expr{
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
						Alias: "temp_avg",
					},
					&ResTarget{
						Value: &ColumnRef{
							Fields: "date",
						},
					},
				},
				From: &RangeVar{
					Name: "weather",
				},
			},
		},
		{
			false,
			"SELECT * FROM weather WHERE city = 'San Francisco'",
			[]interface{}{},
			&Select{
				Targets: &ResTarget{
					Value: &ColumnRef{
						Fields: "*",
					},
				},
				From: &RangeVar{
					Name: "weather",
				},
				Where: &Expr{
					Type: ExprTypeOp,
					Name: "=",
					LHS:  &ColumnRef{Fields: "city"},
					RHS:  &Const{Value: "San Francisco"},
				},
			},
		},
		{
			false,
			"SELECT * FROM weather WHERE city = 'San Francisco' AND prcp > 0",
			[]interface{}{},
			&Select{
				Targets: &ResTarget{
					Value: &ColumnRef{
						Fields: "*",
					},
				},
				From: &RangeVar{
					Name: "weather",
				},
				Where: &BoolExpr{
					Type: BoolExprTypeAnd,
					Args: []Node{
						&Expr{
							Type: ExprTypeOp,
							Name: "=",
							LHS:  &ColumnRef{Fields: "city"},
							RHS:  &Const{Value: "San Francisco"},
						},
						&Expr{
							Type: ExprTypeOp,
							Name: ">",
							LHS:  &ColumnRef{Fields: "prcp"},
							RHS:  &Const{Value: float64(0)},
						},
					},
				},
			},
		},
		{
			false,
			"SELECT * FROM weather ORDER BY city,temp_lo",
			[]interface{}{},
			&Select{
				Targets: &ResTarget{
					Value: &ColumnRef{
						Fields: "*",
					},
				},
				From: &RangeVar{
					Name: "weather",
				},
				OrderBy: &List{
					&OrderBy{
						Value: &ColumnRef{Fields: "city"},
					},
					&OrderBy{
						Value: &ColumnRef{Fields: "temp_lo"},
					},
				},
			},
		},
		{
			false,
			"SELECT DISTINCT city FROM weather",
			[]interface{}{},
			&Select{
				Distinct: &List{},
				Targets: &ResTarget{
					Value: &ColumnRef{
						Fields: "city",
					},
				},
				From: &RangeVar{
					Name: "weather",
				},
			},
		},
		{
			false,
			"SELECT DISTINCT ON (bcolor) bcolor,fcolor FROM t1 ORDER BY bcolor,fcolor",
			[]interface{}{},
			&Select{
				Distinct: &List{
					&ColumnRef{Fields: "bcolor"},
				},
				Targets: &List{
					&ResTarget{
						Value: &ColumnRef{
							Fields: "bcolor",
						},
					},
					&ResTarget{
						Value: &ColumnRef{
							Fields: "fcolor",
						},
					},
				},
				From: &RangeVar{
					Name: "t1",
				},
				OrderBy: &List{
					&OrderBy{
						Value: &ColumnRef{Fields: "bcolor"},
					},
					&OrderBy{
						Value: &ColumnRef{Fields: "fcolor"},
					},
				},
			},
		},
		{
			false,
			"SELECT isbn,retail,cost FROM stock ORDER BY isbn ASC,cost DESC,retail USING > LIMIT 3",
			[]interface{}{},
			&Select{
				Targets: &List{
					&ResTarget{
						Value: &ColumnRef{
							Fields: "isbn",
						},
					},
					&ResTarget{
						Value: &ColumnRef{
							Fields: "retail",
						},
					},
					&ResTarget{
						Value: &ColumnRef{
							Fields: "cost",
						},
					},
				},
				From: &RangeVar{
					Name: "stock",
				},
				OrderBy: &List{
					&OrderBy{
						Value:     &ColumnRef{Fields: "isbn"},
						Direction: OrderByDirectionAsc,
					},
					&OrderBy{
						Value:     &ColumnRef{Fields: "cost"},
						Direction: OrderByDirectionDesc,
					},
					&OrderBy{
						Value:     &ColumnRef{Fields: "retail"},
						Direction: OrderByDirectionUsing,
						UsingOp:   ">",
					},
				},
				Limit: &Const{3},
			},
		},
		{
			false,
			"SELECT * FROM stock OFFSET 33 + 1 LIMIT 3",
			[]interface{}{},
			&Select{
				Targets: &List{
					&ResTarget{
						Value: &ColumnRef{
							Fields: "*",
						},
					},
				},
				From: &RangeVar{
					Name: "stock",
				},
				Offset: &Expr{
					Type: ExprTypeOp,
					Name: "+",
					LHS:  &Const{33},
					RHS:  &Const{1},
				},
				Limit: &Const{3},
			},
		},
		{
			false,
			"SELECT a.id,first_name,last_name FROM customer AS a,laptops AS l INNER JOIN payment ON payment.id = a.id LEFT JOIN people ON people.id = a.id RIGHT JOIN cars ON cars.id = a.id",
			[]interface{}{},
			&Select{
				Targets: &List{
					&ResTarget{
						Value: &ColumnRef{Fields: []string{"a", "id"}},
					},
					&ResTarget{
						Value: &ColumnRef{Fields: "first_name"},
					},
					&ResTarget{
						Value: &ColumnRef{Fields: "last_name"},
					},
				},
				From: &List{
					&RangeVar{
						Name:  "customer",
						Alias: "a",
					},
					&JoinExpr{
						Type: JoinTypeRight,
						LHS: &JoinExpr{
							Type: JoinTypeLeft,
							LHS: &JoinExpr{
								Type: JoinTypeInner,
								LHS: &RangeVar{
									Name:  "laptops",
									Alias: "l",
								},
								RHS: &RangeVar{Name: "payment"},
								Qualifiers: &Expr{
									Type: ExprTypeOp,
									Name: "=",
									LHS:  &ColumnRef{Fields: []string{"payment", "id"}},
									RHS:  &ColumnRef{Fields: []string{"a", "id"}},
								},
							},
							RHS: &RangeVar{Name: "people"},
							Qualifiers: &Expr{
								Type: ExprTypeOp,
								Name: "=",
								LHS:  &ColumnRef{Fields: []string{"people", "id"}},
								RHS:  &ColumnRef{Fields: []string{"a", "id"}},
							},
						},
						RHS: &RangeVar{Name: "cars"},
						Qualifiers: &Expr{
							Type: ExprTypeOp,
							Name: "=",
							LHS:  &ColumnRef{Fields: []string{"cars", "id"}},
							RHS:  &ColumnRef{Fields: []string{"a", "id"}},
						},
					},
				},
			},
		},
		// {
		// 	false,
		// 	"SELECT name, substr(address, 1, 40) || '...' AS short_address FROM publishers WHERE id = 113",
		// 	[]interface{}{},
		// 	&Select{},
		// },
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
