# SQuL
[![CircleCI branch](https://img.shields.io/circleci/project/github/trivigy/squl/master.svg?label=master&logo=circleci)](https://circleci.com/gh/trivigy/workflows/squl)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE.md)
[![](https://godoc.org/github.com/trivigy/squl?status.svg&style=flat)](http://godoc.org/github.com/trivigy/squl)
[![GitHub tag (latest SemVer)](https://img.shields.io/github/tag/trivigy/squl.svg?style=flat&color=e36397&label=release)](https://github.com/trivigy/squl/releases/latest)

Unlike all the other sql query builder tools that I have encountered out there 
in the golang ecosystem, this tool does not introduce a new custom designed 
fluent set of operations. Instead the goal in the sign is to completely give the 
developer the power to write their own queries using plain old sql syntax but 
with support for special in template operations that help generate the queries 
easily for a standard `sql.DB` object.

## Installation
To get the latest version of the library simply run:
```bash
go get -u github.com/trivigy/squl
```

## Usage
The way this template processing works is quiet simple. When the template 
executes, the user can rely on the existence of a special global variable 
called `$` which gives access to special helper tools for generating easy sql 
queries. The reference to all those special objects can be found 
[HERE](http://godoc.org/github.com/trivigy/squl/#Session).


### Example
```
package main

import (
	"os"

	"github.com/trivigy/squl"
)

func main() {
squl := Builder{}
	query, args, err := squl.Build(`
		INSERT INTO users (
			id,
			type,
			name,
			password,
			salt,
			email,
			full_name
		)
		VALUES (
			{{if .ID}}{{$.Params.Mark "id" .ID}}{{else}}DEFAULT{{end}},
			{{if .Type}}{{$.Params.Mark "type" .Type.String}}{{else}}DEFAULT{{end}},
			{{if .Name}}{{$.Params.Mark "name" .Name}}{{else}}DEFAULT{{end}},
			{{if .Password}}{{$.Params.Mark "password" .Password}}{{else}}DEFAULT{{end}},
			{{if .Salt}}{{$.Params.Mark "salt" .Salt}}{{else}}DEFAULT{{end}},
			{{if .Email}}{{$.Params.Mark "email" .Email}}{{else}}DEFAULT{{end}},
			{{if .FullName}}{{$.Params.Mark "full_name" .FullName}}{{else}}DEFAULT{{end}}
		)`,
		MockUser{
			ID:       int64(11),
			Type:     MockUserTypeIndividual,
			Name:     "unittest",
			Password: []byte("password"),
			Salt:     []byte("salt"),
			Email:    "unittest@unittest.com",
			FullName: "unittest unittest",
		},
	)
}

```

This will produce the following printout. What you see is that in this case is 
that `$.Params.Mark` will come to your help to automatically mark the parameters 
and assign ordinal parameter marking to them which play well with 
`sql.DB.Exec()`. 
```
INSERT INTO users (
    id,
    type,
    name,
    password,
    salt,
    email,
    full_name
)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7
)
[11 individual unittest [112 97 115 115 119 111 114 100] [115 97 108 116] unittest@unittest.com unittest unittest]
```

> The ordinal marking uses a key to tag the specific value. If later you choose 
to call {{$.Params.Mark "param_name" .ParamValue}} for a particular parameter, 
the query will reuse an ordinal marking that it had previously assigned.
