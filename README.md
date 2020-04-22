# SQuL
[![CircleCI](https://img.shields.io/circleci/project/github/trivigy/squl/master.svg?label=master&logo=circleci)](https://circleci.com/gh/trivigy/workflows/squl)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE.md)
[![go.dev](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white)](https://pkg.go.dev/github.com/trivigy/squl)
[![SemVer](https://img.shields.io/github/tag/trivigy/squl.svg?style=flat&color=e36397&label=release)](https://github.com/trivigy/squl/releases/latest)
[![CodeCov](https://codecov.io/gh/trivigy/squl/branch/master/graph/badge.svg)](https://codecov.io/gh/trivigy/squl)

Unlike all the other sql query builder tools that I have encountered out there 
in the golang ecosystem, this tool does not introduce a new custom designed 
fluent set of methods. Instead the goal is to provide golang structs which can 
be then built into custom queries. It's almost like writing the queries by hand 
except that you construct them using golang structs.

## Installation
To get the latest version of the library simply run:
```bash
go get -u github.com/trivigy/squl
```

## Usage
For a whole lot more examples please make sure to explore the unittest files.

### `INSERT`
> `INSERT INTO contacts (last_name,first_name) SELECT last_name,first_name FROM customers WHERE customer_id > 4000`
```
package main

import (
    "github.com/trivigy/squl"
)

func main() {
	query, args, err := squl.Build(&squl.Insert{
		Relation: &squl.RangeVar{
			Name: "contacts",
		},
		Columns: &squl.List{
			&squl.ResTarget{
				Value: &squl.ColumnRef{Fields: "last_name"},
			},
			&squl.ResTarget{
				Value: &squl.ColumnRef{Fields: "first_name"},
			},
		},
		Select: &squl.Select{
			Targets: &squl.List{
				&squl.ResTarget{
					Value: &squl.ColumnRef{Fields: "last_name"},
				},
				&squl.ResTarget{
					Value: &squl.ColumnRef{Fields: "first_name"},
				},
			},
			From: &squl.RangeVar{
				Name: "customers",
			},
			Where: &squl.Expr{
				Type: squl.ExprTypeOp,
				Name: ">",
				LHS:  &squl.ColumnRef{Fields: "customer_id"},
				RHS:  &squl.Const{Value: 4000},
			},
		},
	})
}
```

### `SELECT`
> `SELECT a.id,first_name,last_name FROM customer AS a,laptops AS l INNER JOIN payment ON payment.id = a.id LEFT JOIN people ON people.id = a.id RIGHT JOIN cars ON cars.id = a.id`
```
package main

import (
	"github.com/trivigy/squl"
)

func main() {
	query, args, err := squl.Build(&squl.Select{
        Targets: &squl.List{
            &squl.ResTarget{
                Value: &squl.ColumnRef{Fields: []string{"a", "id"}},
            },
            &squl.ResTarget{
                Value: &squl.ColumnRef{Fields: "first_name"},
            },
            &squl.ResTarget{
                Value: &squl.ColumnRef{Fields: "last_name"},
            },
        },
        From: &squl.List{
            &squl.RangeVar{
                Name:  "customer",
                Alias: "a",
            },
            &squl.JoinExpr{
                Type: squl.JoinTypeRight,
                LHS: &squl.JoinExpr{
                    Type: squl.JoinTypeLeft,
                    LHS: &squl.JoinExpr{
                        Type: squl.JoinTypeInner,
                        LHS: &squl.RangeVar{
                            Name:  "laptops",
                            Alias: "l",
                        },
                        RHS: &squl.RangeVar{Name: "payment"},
                        Qualifiers: &squl.Expr{
                            Type: squl.ExprTypeOp,
                            Name: "=",
                            LHS:  &squl.ColumnRef{Fields: []string{"payment", "id"}},
                            RHS:  &squl.ColumnRef{Fields: []string{"a", "id"}},
                        },
                    },
                    RHS: &squl.RangeVar{Name: "people"},
                    Qualifiers: &squl.Expr{
                        Type: squl.ExprTypeOp,
                        Name: "=",
                        LHS:  &squl.ColumnRef{Fields: []string{"people", "id"}},
                        RHS:  &squl.ColumnRef{Fields: []string{"a", "id"}},
                    },
                },
                RHS: &squl.RangeVar{Name: "cars"},
                Qualifiers: &squl.Expr{
                    Type: squl.ExprTypeOp,
                    Name: "=",
                    LHS:  &squl.ColumnRef{Fields: []string{"cars", "id"}},
                    RHS:  &squl.ColumnRef{Fields: []string{"a", "id"}},
                },
            },
        },
    })
}
```

### `UPDATE`
> `UPDATE stock SET retail = stock_backup.retail FROM stock_backup WHERE stock.isbn = stock_backup.isbn`
```
package main

import (
    "github.com/trivigy/squl"
)

func main() {
	query, args, err := squl.Build(&squl.Update{
        Relation: &squl.RangeVar{
            Name: "stock",
        },
        Targets: &squl.ResTarget{
            Name:  &squl.ColumnRef{Fields: "retail"},
            Value: &squl.ColumnRef{Fields: []string{"stock_backup", "retail"}},
        },
        From: &squl.RangeVar{
            Name: "stock_backup",
        },
        Where: &squl.Expr{
            Type: squl.ExprTypeOp,
            Name: "=",
            LHS:  &squl.ColumnRef{Fields: []string{"stock", "isbn"}},
            RHS:  &squl.ColumnRef{Fields: []string{"stock_backup", "isbn"}},
        },
    })
}
```

### `DELETE`
> `DELETE FROM products WHERE obsoletion_date = 'today' RETURNING *`
```
package main

import (
    "github.com/trivigy/squl"
)

func main() {
	query, args, err := squl.Build(&squl.Delete{
        Relation: &squl.RangeVar{
            Name: "products",
        },
        Where: &squl.Expr{
            Type: squl.ExprTypeOp,
            Name: "=",
            LHS:  &squl.ColumnRef{Fields: "obsoletion_date"},
            RHS:  &squl.Const{Value: "today"},
        },
        Returning: &squl.ResTarget{
            Value: &squl.ColumnRef{Fields: "*"},
        },
    })
}
```
