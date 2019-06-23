package squl

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"strings"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func init() {
	gob.Register(MockUserType(Unknown))
}

const (
	Unknown    = iota
	unknownStr = "unknown"
)

type MockUserType int

const (
	MockUserTypeIndividual MockUserType = iota + 1
	MockUserTypeOrganization
)

const (
	mockUserTypeIndividualStr   = "individual"
	mockUserTypeOrganizationStr = "organization"
)

var toStringUserType = map[MockUserType]string{
	MockUserType(Unknown):    unknownStr,
	MockUserTypeIndividual:   mockUserTypeIndividualStr,
	MockUserTypeOrganization: mockUserTypeOrganizationStr,
}

func NewUserType(raw string) (MockUserType, error) {
	switch raw {
	case mockUserTypeIndividualStr:
		return MockUserTypeIndividual, nil
	case mockUserTypeOrganizationStr:
		return MockUserTypeOrganization, nil
	default:
		return MockUserType(Unknown), errors.Errorf("unknown type %q", raw)
	}
}

func (r MockUserType) String() string {
	return toStringUserType[r]
}

func (r *MockUserType) UnmarshalJSON(rbytes []byte) error {
	var s string
	if err := json.Unmarshal(rbytes, &s); err != nil {
		return err
	}
	raw := strings.ToLower(s)
	switch raw {
	case mockUserTypeIndividualStr:
		*r = MockUserTypeIndividual
	case mockUserTypeOrganizationStr:
		*r = MockUserTypeOrganization
	default:
		*r = Unknown
		return errors.Errorf("unknown type %q", raw)
	}
	return nil
}

func (r MockUserType) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	if _, err := buffer.WriteString(toStringUserType[r]); err != nil {
		return nil, errors.WithStack(err)
	}
	if _, err := buffer.WriteString(`"`); err != nil {
		return nil, errors.WithStack(err)
	}
	return buffer.Bytes(), nil
}

type MockUser struct {
	ID       int64
	Type     MockUserType
	Name     string
	Password []byte
	Salt     []byte
	Email    string
	FullName string
}

type SqulSuite struct {
	suite.Suite
}

func (r *SqulSuite) TestBuild() {
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
	assert.Nil(r.T(), err)
	assert.Equal(r.T(), "\n\t\tINSERT INTO users (\n\t\t\tid,\n\t\t\ttype,\n\t\t\tname,\n\t\t\tpassword,\n\t\t\tsalt,\n\t\t\temail,\n\t\t\tfull_name\n\t\t)\n\t\tVALUES (\n\t\t\t$1,\n\t\t\t$2,\n\t\t\t$3,\n\t\t\t$4,\n\t\t\t$5,\n\t\t\t$6,\n\t\t\t$7\n\t\t)", query)
	assert.EqualValues(r.T(), []interface{}{int64(11), "individual", "unittest", []byte("password"), []byte("salt"), "unittest@unittest.com", "unittest unittest"}, args)
}

func (r *SqulSuite) TestBuild_RepeatParameters() {
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
		MockUser{
			ID:   int64(11),
			Name: "unittest",
		},
	)
	assert.Nil(r.T(), err)
	assert.Equal(r.T(), "\n\t\tINSERT INTO users (\n\t\t\tid,\n\t\t\tname,\n\t\t\tpasswd,\n\t\t\temail,\n\t\t\tfull_name,\n\t\t\tsalt\n\t\t)\n\t\tVALUES (\n\t\t\t$1,\n\t\t\t$2,\n\t\t\t$1,\n\t\t\t$2,\n\t\t\t$1,\n\t\t\t$2,\n\t\t)", query)
	assert.EqualValues(r.T(), []interface{}{int64(11), "unittest"}, args)
}

func TestSqulSuite(t *testing.T) {
	suite.Run(t, new(SqulSuite))
}
