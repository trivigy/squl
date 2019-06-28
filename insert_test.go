package squl

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/suite"
)

type InsertSuite struct {
	suite.Suite
}

func (r *InsertSuite) TestInsertSuiteDump() {
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
							&Var{250},
							&Var{"Anderson"},
							&Var{"Jane"},
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
							&Const{"John"},
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
							&Var{250},
							&Var{"Anderson"},
							&Var{"Jane"},
							&Default{},
						},
						&List{
							&Var{251},
							&Var{"Smith"},
							&Var{"John"},
							&Var{"US"},
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
							&Const{"Joe"},
							&Const{"Cool"},
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
							&Const{"Joe"},
							&Const{"Cool"},
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
						RHS:  &Const{4000},
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
			assert.Fail(r.T(), failMsg, err)
		}
		assert.EqualValues(r.T(), testCase.args, counter.args())
		assert.Equal(r.T(), testCase.output, actual, failMsg)
	}
}

func TestInsertSuite(t *testing.T) {
	suite.Run(t, new(InsertSuite))
}
