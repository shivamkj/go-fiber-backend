package db

import (
	"errors"
	"strconv"
	"strings"

	"github.com/ganigeorgiev/fexpr"
	"github.com/qnify/api-server/utils/fiber"
)

// Parameter used by Postgres driver, also need parameter number
// after this, ex. - $1, $2, which is added while building query
var PgParam = []byte("$")

// Parameter used by MySQL driver
var MsParam = []byte("?")

type queryBuilder struct {
	strb      strings.Builder
	params    []any
	filter    fiber.QueryFilters
	parameter []byte
}

func NewQuery(filter *fiber.QueryFilters, param []byte) *queryBuilder {
	builder := queryBuilder{
		filter:    *filter,
		parameter: param,
	}
	return &builder
}

func (qb *queryBuilder) Query() string {
	qb.addFilters()
	qb.addPagination()
	return qb.strb.String()
}

func (qb *queryBuilder) Params() []any {
	return qb.params
}

var errInValidFilter = errors.New("invalid filter passed in query params")

func GetQuery(filter fiber.QueryFilters) (string, error) {
	qb := queryBuilder{
		filter: filter,
		params: make([]any, 0, len(filter.Conditions)+2),
	}

	if err := qb.addFilters(); err != nil {
		return "", err
	}
	qb.addPagination()

	return qb.strb.String(), nil
}

func (q *queryBuilder) addPagination() {

	q.strb.WriteString(" id > ")
	q.addParams(q.filter.Page)
	q.strb.WriteString(" ORDER BY id ASC")

	q.strb.WriteString(" LIMIT ")
	q.addParams(q.filter.Limit)
}

func (q *queryBuilder) addFilters() error {
	q.strb.WriteString(" WHERE ")

	for index, exprGrp := range q.filter.Conditions {
		switch item := exprGrp.Item.(type) {

		case fexpr.Expr:
			q.addJoinExpr(exprGrp.Join, index)
			if err := q.addCondition(item); err != nil {
				return err
			}

		case []fexpr.ExprGroup:
			q.addJoinExpr(exprGrp.Join, index)
			q.strb.WriteString("( ")

			for innerIndx, exp := range item {
				if nestedExprG, ok := exp.Item.(fexpr.Expr); ok {
					q.addJoinExpr(exp.Join, innerIndx)
					if err := q.addCondition(nestedExprG); err != nil {
						return err
					}
				} else {
					return errors.New("only one level of nesting supported in filters")
				}
			}

			q.strb.WriteString(" )")
		}
	}

	if len(q.filter.Conditions) != 0 {
		q.strb.WriteString(" AND")
	}

	return nil
}

func (q *queryBuilder) addJoinExpr(opr fexpr.JoinOp, index int) {
	if index != 0 {
		q.strb.WriteString(" ")
		q.strb.WriteString(string(opr))
		q.strb.WriteString(" ")
	}
}

func (q *queryBuilder) addCondition(exp fexpr.Expr) error {
	q.strb.WriteString(exp.Left.Literal)
	q.strb.WriteString(" ")
	q.strb.WriteString(string(exp.Op))
	q.strb.WriteString(" ")

	if exp.Right.Type == fexpr.TokenIdentifier {
		if exp.Right.Literal == "true" {
			q.addParams(true)
		} else if exp.Right.Literal == "false" {
			q.addParams(false)
		} else {
			return errInValidFilter
		}
	} else if exp.Right.Type == fexpr.TokenNumber {
		num, err := strconv.Atoi(exp.Right.Literal)
		if err != nil {
			return errInValidFilter
		}
		q.addParams(num)
	} else if exp.Right.Type == fexpr.TokenText {
		q.addParams(exp.Right.Literal)
	} else {
		return errInValidFilter
	}

	return nil
}

func (q *queryBuilder) addParams(param any) {
	q.params = append(q.params, param)
	q.strb.Write(q.parameter)
	// also add parameter number for Postgres
	if q.parameter[0] == PgParam[0] {
		q.strb.WriteString(strconv.Itoa(len(q.params)))
	}
}
