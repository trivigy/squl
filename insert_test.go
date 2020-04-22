package squl

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsert_dump(t *testing.T) {
	testCases := []struct {
		shouldFail bool
		output     string
		args       []interface{}
		expected   *Insert
	}{
		{
			false,
			"INSERT INTO contacts (contact_id,last_name,first_name,country) VALUES ($1,$2,$3,DEFAULT)",
			[]interface{}{250, "Anderson", "Jane"},
			&Insert{
				Relation: &RangeVar{
					Name: "contacts",
				},
				Columns: &List{
					&ResTarget{
						Value: &ColumnRef{Fields: "contact_id"},
					},
					&ResTarget{
						Value: &ColumnRef{Fields: "last_name"},
					},
					&ResTarget{
						Value: &ColumnRef{Fields: "first_name"},
					},
					&ResTarget{
						Value: &ColumnRef{Fields: "country"},
					},
				},
				Select: &Select{
					Values: []Node{
						&List{
							&Var{Value: 250},
							&Var{Value: "Anderson"},
							&Var{Value: "Jane"},
							&Default{},
						},
					},
				},
			},
		},
		{
			false,
			"INSERT INTO contacts (contact_id,first_name) VALUES (NULL,'John')",
			[]interface{}{},
			&Insert{
				Relation: &RangeVar{
					Name: "contacts",
				},
				Columns: &List{
					&ResTarget{
						Value: &ColumnRef{Fields: "contact_id"},
					},
					&ResTarget{
						Value: &ColumnRef{Fields: "first_name"},
					},
				},
				Select: &Select{
					Values: []Node{
						&List{
							&Null{},
							&Const{Value: "John"},
						},
					},
				},
			},
		},
		{
			false,
			"INSERT INTO contacts (contact_id,last_name,first_name,country) VALUES ($1,$2,$3,DEFAULT),($4,$5,$6,$7)",
			[]interface{}{250, "Anderson", "Jane", 251, "Smith", "John", "US"},
			&Insert{
				Relation: &RangeVar{
					Name: "contacts",
				},
				Columns: &List{
					&ResTarget{
						Value: &ColumnRef{Fields: "contact_id"},
					},
					&ResTarget{
						Value: &ColumnRef{Fields: "last_name"},
					},
					&ResTarget{
						Value: &ColumnRef{Fields: "first_name"},
					},
					&ResTarget{
						Value: &ColumnRef{Fields: "country"},
					},
				},
				Select: &Select{
					Values: []Node{
						&List{
							&Var{Value: 250},
							&Var{Value: "Anderson"},
							&Var{Value: "Jane"},
							&Default{},
						},
						&List{
							&Var{Value: 251},
							&Var{Value: "Smith"},
							&Var{Value: "John"},
							&Var{Value: "US"},
						},
					},
				},
			},
		},
		{
			false,
			"INSERT INTO contacts DEFAULT VALUES",
			[]interface{}{},
			&Insert{
				Relation: &RangeVar{
					Name: "contacts",
				},
			},
		},
		{
			false,
			"INSERT INTO users (firstname,lastname) VALUES ('Joe','Cool') RETURNING id,firstname",
			[]interface{}{},
			&Insert{
				Relation: &RangeVar{
					Name: "users",
				},
				Columns: &List{
					&ResTarget{
						Value: &ColumnRef{Fields: "firstname"},
					},
					&ResTarget{
						Value: &ColumnRef{Fields: "lastname"},
					},
				},
				Select: &Select{
					Values: []Node{
						&List{
							&Const{Value: "Joe"},
							&Const{Value: "Cool"},
						},
					},
				},
				Returning: &List{
					&ResTarget{
						Value: &ColumnRef{Fields: "id"},
					},
					&ResTarget{
						Value: &ColumnRef{Fields: "firstname"},
					},
				},
			},
		},
		{
			false,
			"INSERT INTO users (firstname,lastname) VALUES ('Joe','Cool') RETURNING *",
			[]interface{}{},
			&Insert{
				Relation: &RangeVar{
					Name: "users",
				},
				Columns: &List{
					&ResTarget{
						Value: &ColumnRef{Fields: "firstname"},
					},
					&ResTarget{
						Value: &ColumnRef{Fields: "lastname"},
					},
				},
				Select: &Select{
					Values: []Node{
						&List{
							&Const{Value: "Joe"},
							&Const{Value: "Cool"},
						},
					},
				},
				Returning: &ResTarget{
					Value: &ColumnRef{Fields: "*"},
				},
			},
		},
		{
			false,
			"INSERT INTO contacts (last_name,first_name) SELECT last_name,first_name FROM customers WHERE customer_id > 4000",
			[]interface{}{},
			&Insert{
				Relation: &RangeVar{
					Name: "contacts",
				},
				Columns: &List{
					&ResTarget{
						Value: &ColumnRef{Fields: "last_name"},
					},
					&ResTarget{
						Value: &ColumnRef{Fields: "first_name"},
					},
				},
				Select: &Select{
					Targets: &List{
						&ResTarget{
							Value: &ColumnRef{Fields: "last_name"},
						},
						&ResTarget{
							Value: &ColumnRef{Fields: "first_name"},
						},
					},
					From: &RangeVar{
						Name: "customers",
					},
					Where: &Expr{
						Type: ExprTypeOp,
						Name: ">",
						LHS:  &ColumnRef{Fields: "customer_id"},
						RHS:  &Const{Value: 4000},
					},
				},
			},
		},
		{
			false,
			"INSERT INTO customers (name,email) VALUES ('Microsoft','hotline@microsoft.com') ON CONFLICT (name) DO NOTHING",
			[]interface{}{},
			&Insert{
				Relation: &RangeVar{
					Name: "customers",
				},
				Columns: &List{
					&ResTarget{
						Value: &ColumnRef{Fields: "name"},
					},
					&ResTarget{
						Value: &ColumnRef{Fields: "email"},
					},
				},
				Select: &Select{
					Values: []Node{
						&List{
							&Const{Value: "Microsoft"},
							&Const{Value: "hotline@microsoft.com"},
						},
					},
				},
				OnConflict: &OnConflict{
					Infer: &Infer{
						IndexElems: &List{
							&IndexElem{Name: "name"},
						},
					},
					Action: OnConflictNothing,
				},
			},
		},
		{
			false,
			`INSERT INTO users (id,level) VALUES (1,0) ON CONFLICT (id) DO UPDATE SET level = users.level + 1`,
			[]interface{}{},
			&Insert{
				Relation: &RangeVar{
					Name: "users",
				},
				Columns: &List{
					&ResTarget{
						Value: &ColumnRef{Fields: "id"},
					},
					&ResTarget{
						Value: &ColumnRef{Fields: "level"},
					},
				},
				Select: &Select{
					Values: []Node{
						&List{
							&Const{Value: 1},
							&Const{Value: 0},
						},
					},
				},
				OnConflict: &OnConflict{
					Action: OnConflictUpdate,
					Infer: &Infer{
						IndexElems: &List{
							&IndexElem{Name: "id"},
						},
					},
					TargetList: &List{
						&ResTarget{
							Name: "level",
							Value: &Expr{
								Type: ExprTypeOp,
								Name: "+",
								LHS:  &ColumnRef{Fields: []string{"users", "level"}},
								RHS:  &Const{Value: 1},
							},
						},
					},
				},
			},
		},
		{
			false,
			"INSERT INTO customers (name,email) VALUES ('Microsoft','hotline@microsoft.com') ON CONFLICT (id,name) DO NOTHING",
			[]interface{}{},
			&Insert{
				Relation: &RangeVar{
					Name: "customers",
				},
				Columns: &List{
					&ResTarget{
						Value: &ColumnRef{Fields: "name"},
					},
					&ResTarget{
						Value: &ColumnRef{Fields: "email"},
					},
				},
				Select: &Select{
					Values: []Node{
						&List{
							&Const{Value: "Microsoft"},
							&Const{Value: "hotline@microsoft.com"},
						},
					},
				},
				OnConflict: &OnConflict{
					Infer: &Infer{
						IndexElems: &List{
							&IndexElem{Name: "id"},
							&IndexElem{Name: "name"},
						},
					},
					Action: OnConflictNothing,
				},
			},
		},
		{
			false,
			`INSERT INTO sal_emp VALUES ('Bill',ARRAY[10000,10000,10000,10000],ARRAY[ARRAY['meeting','lunch'],ARRAY['training','presentation']])`,
			[]interface{}{},
			&Insert{
				Relation: &RangeVar{
					Name: "sal_emp",
				},
				Select: &Select{
					Values: []Node{
						&List{
							&Const{Value: "Bill"},
							&Const{Value: &ArrayExpr{
								&Const{Value: 10000},
								&Const{Value: 10000},
								&Const{Value: 10000},
								&Const{Value: 10000},
							}},
							&Const{Value: &ArrayExpr{
								&Const{Value: &ArrayExpr{
									&Const{Value: "meeting"},
									&Const{Value: "lunch"},
								}},
								&Const{Value: &ArrayExpr{
									&Const{Value: "training"},
									&Const{Value: "presentation"},
								}},
							}},
						},
					},
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
