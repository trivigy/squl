package squl

import (
	"encoding/gob"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

const (
	Unknown    = iota
	unknownStr = "unknown"
)

type UserType int

const (
	UserTypeIndividual UserType = iota + 1
	UserTypeOrganization
)

const (
	userTypeIndividualStr   = "individual"
	userTypeOrganizationStr = "organization"
)

var toStringUserType = map[UserType]string{
	UserType(Unknown):    unknownStr,
	UserTypeIndividual:   userTypeIndividualStr,
	UserTypeOrganization: userTypeOrganizationStr,
}

func NewUserType(raw string) (UserType, error) {
	switch raw {
	case userTypeIndividualStr:
		return UserTypeIndividual, nil
	case userTypeOrganizationStr:
		return UserTypeOrganization, nil
	default:
		return UserType(Unknown), errors.Errorf("unknown type %q", raw)
	}
}

func (r UserType) String() string {
	return toStringUserType[r]
}

type User struct {
	ID       int64
	Type     UserType
	Name     string
	Password []byte
	Salt     []byte
	Email    string
	FullName string
}

func (r *UserType) Scan(src interface{}) error {
	var err error
	if *r, err = NewUserType(string(src.([]byte))); err != nil {
		return err
	}
	return nil
}

type SQuLSuite struct {
	suite.Suite
}

func (r *SQuLSuite) SetupSuite() {
	gob.Register(UserType(Unknown))
}

func (r *SQuLSuite) TestBuild() {
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
		User{
			ID:       int64(11),
			Type:     UserTypeIndividual,
			Name:     "unittest",
			Password: []byte("password"),
			Salt:     []byte("salt"),
			Email:    "unittest@unittest.com",
			FullName: "unittest unittest",
		},
	)
	assert.Nil(r.T(), err)
	assert.Equal(r.T(), "\n\t\tINSERT INTO users (\n\t\t\tid,\n\t\t\ttype,\n\t\t\tname,\n\t\t\tpassword,\n\t\t\tsalt,\n\t\t\temail,\n\t\t\tfull_name\n\t\t)\n\t\tVALUES (\n\t\t\t$1,\n\t\t\t$2,\n\t\t\t$3,\n\t\t\t$4,\n\t\t\t$5,\n\t\t\t$6,\n\t\t\t$7\n\t\t)", query)
	assert.EqualValues(r.T(), []interface{}{int64(11), "individual", "unittest", []byte("password"), []byte("salt"), "unittest@unittest.com", "unittest unittest"}, args)
}

func (r *SQuLSuite) TestBuild_RepeatParameters() {
	squl := Builder{}
	query, args, err := squl.Build(`
		INSERT INTO users (
			id,
			name,
			passwd,
			email,
			full_name,
			salt
		)
		VALUES (
			{{if .ID}}{{$.Params.Mark "id" .ID}}{{else}}DEFAULT{{end}},
			{{if .Name}}{{$.Params.Mark "name" .Name}}{{else}}DEFAULT{{end}},
			{{if .ID}}{{$.Params.Mark "id" .ID}}{{else}}DEFAULT{{end}},
			{{if .Name}}{{$.Params.Mark "name" .Name}}{{else}}DEFAULT{{end}},
			{{if .ID}}{{$.Params.Mark "id" .ID}}{{else}}DEFAULT{{end}},
			{{if .Name}}{{$.Params.Mark "name" .Name}}{{else}}DEFAULT{{end}},
		)`,
		User{
			ID:   int64(11),
			Name: "unittest",
		},
	)
	assert.Nil(r.T(), err)
	assert.Equal(r.T(), "\n\t\tINSERT INTO users (\n\t\t\tid,\n\t\t\tname,\n\t\t\tpasswd,\n\t\t\temail,\n\t\t\tfull_name,\n\t\t\tsalt\n\t\t)\n\t\tVALUES (\n\t\t\t$1,\n\t\t\t$2,\n\t\t\t$1,\n\t\t\t$2,\n\t\t\t$1,\n\t\t\t$2,\n\t\t)", query)
	assert.EqualValues(r.T(), []interface{}{int64(11), "unittest"}, args)
}

func TestSqulSuite(t *testing.T) {
	suite.Run(t, new(SQuLSuite))
}
