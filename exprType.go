package squl

import (
	"bytes"
	"encoding/json"

	fmt "golang.org/x/xerrors"

	"github.com/trivigy/squl/internal/global"
)

// ExprType defines possible types of expressions.
type ExprType int

const (
	// ExprTypeOp describes normal operator
	ExprTypeOp ExprType = iota + 1

	// ExprTypeOpAny describes scalar op ANY (array)
	ExprTypeOpAny

	// ExprTypeOpAll describes scalar op ALL (array)
	ExprTypeOpAll

	// ExprTypeDistinct describes IS DISTINCT FROM - name must be "="
	ExprTypeDistinct

	// ExprTypeNotDistinct describes IS NOT DISTINCT FROM - name must be "="
	ExprTypeNotDistinct

	// ExprTypeNullIf describes NULLIF - name must be "="
	ExprTypeNullIf

	// ExprTypeOf describes IS [NOT] OF - name must be "=" or "<>"
	ExprTypeOf

	// ExprTypeIn describes [NOT] IN - name must be "=" or "<>"
	ExprTypeIn

	// ExprTypeLike describes [NOT] LIKE - name must be "~~" or "!~~"
	ExprTypeLike

	// ExprTypeILike describes [NOT] ILIKE - name must be "~~*" or "!~~*"
	ExprTypeILike

	// ExprTypeSimilar describes [NOT] SIMILAR - name must be "~" or "!~"
	ExprTypeSimilar

	// ExprTypeBetween indicates that the name must be "BETWEEN".
	ExprTypeBetween

	// ExprTypeNotBetween indicates that the name must be "NOT BETWEEN".
	ExprTypeNotBetween

	// ExprTypeBetweenSym indicates that the name must be "BETWEEN SYMMETRIC".
	ExprTypeBetweenSym

	// ExprTypeNotBetweenSym indicates that the name must be "NOT BETWEEN SYMMETRIC".
	ExprTypeNotBetweenSym

	// ExprTypeParen indicates nameless dummy node for parentheses.
	ExprTypeParen
)

const (
	exprTypeOpStr            = "op"
	exprTypeOpAnyStr         = "opAny"
	exprTypeOpAllStr         = "opAll"
	exprTypeDistinctStr      = "distinct"
	exprTypeNotDistinctStr   = "notDistinct"
	exprTypeNullIfStr        = "nullIf"
	exprTypeOfStr            = "of"
	exprTypeInStr            = "in"
	exprTypeLikeStr          = "like"
	exprTypeILikeStr         = "iLike"
	exprTypeSimilarStr       = "similar"
	exprTypeBetweenStr       = "between"
	exprTypeNotBetweenStr    = "notBetween"
	exprTypeBetweenSymStr    = "betweenSym"
	exprTypeNotBetweenSymStr = "notBetweenSym"
	exprTypeParenStr         = "paren"
)

var toStringExprType = map[ExprType]string{
	ExprType(Unknown):     unknownStr,
	ExprTypeOp:            exprTypeOpStr,
	ExprTypeOpAny:         exprTypeOpAnyStr,
	ExprTypeOpAll:         exprTypeOpAllStr,
	ExprTypeDistinct:      exprTypeDistinctStr,
	ExprTypeNotDistinct:   exprTypeNotDistinctStr,
	ExprTypeNullIf:        exprTypeNullIfStr,
	ExprTypeOf:            exprTypeOfStr,
	ExprTypeIn:            exprTypeInStr,
	ExprTypeLike:          exprTypeLikeStr,
	ExprTypeILike:         exprTypeILikeStr,
	ExprTypeSimilar:       exprTypeSimilarStr,
	ExprTypeBetween:       exprTypeBetweenStr,
	ExprTypeNotBetween:    exprTypeNotBetweenStr,
	ExprTypeBetweenSym:    exprTypeBetweenSymStr,
	ExprTypeNotBetweenSym: exprTypeNotBetweenSymStr,
	ExprTypeParen:         exprTypeParenStr,
}

// NewExprType creates a new instance of the enum from raw string.
func NewExprType(raw string) (ExprType, error) {
	switch raw {
	case exprTypeOpStr:
		return ExprTypeOp, nil
	case exprTypeOpAnyStr:
		return ExprTypeOpAny, nil
	case exprTypeOpAllStr:
		return ExprTypeOpAll, nil
	case exprTypeDistinctStr:
		return ExprTypeDistinct, nil
	case exprTypeNotDistinctStr:
		return ExprTypeNotDistinct, nil
	case exprTypeNullIfStr:
		return ExprTypeNullIf, nil
	case exprTypeOfStr:
		return ExprTypeOf, nil
	case exprTypeInStr:
		return ExprTypeIn, nil
	case exprTypeLikeStr:
		return ExprTypeLike, nil
	case exprTypeILikeStr:
		return ExprTypeILike, nil
	case exprTypeSimilarStr:
		return ExprTypeSimilar, nil
	case exprTypeBetweenStr:
		return ExprTypeBetween, nil
	case exprTypeNotBetweenStr:
		return ExprTypeNotBetween, nil
	case exprTypeBetweenSymStr:
		return ExprTypeBetweenSym, nil
	case exprTypeNotBetweenSymStr:
		return ExprTypeNotBetweenSym, nil
	case exprTypeParenStr:
		return ExprTypeParen, nil
	default:
		return ExprType(Unknown), fmt.Errorf(global.ErrFmt, pkg.Name(), fmt.Errorf("unknown type %q", raw))
	}
}

// String returns the string representation of the enum type
func (r ExprType) String() string {
	return toStringExprType[r]
}

// UnmarshalJSON unmarshals a quoted json string to enum type.
func (r *ExprType) UnmarshalJSON(rbytes []byte) error {
	var s string
	if err := json.Unmarshal(rbytes, &s); err != nil {
		return err
	}
	switch s {
	case exprTypeOpStr:
		*r = ExprTypeOp
	case exprTypeOpAnyStr:
		*r = ExprTypeOpAny
	case exprTypeOpAllStr:
		*r = ExprTypeOpAll
	case exprTypeDistinctStr:
		*r = ExprTypeDistinct
	case exprTypeNotDistinctStr:
		*r = ExprTypeNotDistinct
	case exprTypeNullIfStr:
		*r = ExprTypeNullIf
	case exprTypeOfStr:
		*r = ExprTypeOf
	case exprTypeInStr:
		*r = ExprTypeIn
	case exprTypeLikeStr:
		*r = ExprTypeLike
	case exprTypeILikeStr:
		*r = ExprTypeILike
	case exprTypeSimilarStr:
		*r = ExprTypeSimilar
	case exprTypeBetweenStr:
		*r = ExprTypeBetween
	case exprTypeNotBetweenStr:
		*r = ExprTypeNotBetween
	case exprTypeBetweenSymStr:
		*r = ExprTypeBetweenSym
	case exprTypeNotBetweenSymStr:
		*r = ExprTypeNotBetweenSym
	case exprTypeParenStr:
		*r = ExprTypeParen
	default:
		*r = Unknown
		return fmt.Errorf(global.ErrFmt, pkg.Name(), fmt.Errorf("unknown type %q", s))
	}
	return nil
}

// MarshalJSON marshals the enum as a quoted json string.
func (r ExprType) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	if _, err := buffer.WriteString(toStringExprType[r]); err != nil {
		return nil, fmt.Errorf(global.ErrFmt, pkg.Name(), err)
	}
	if _, err := buffer.WriteString(`"`); err != nil {
		return nil, fmt.Errorf(global.ErrFmt, pkg.Name(), err)
	}
	return buffer.Bytes(), nil
}
