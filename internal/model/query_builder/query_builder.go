package querybuilder

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

var ErrInvalidWhereCondition = errors.New("invalid where condition")

type Where struct {
	Left      *Where
	Right     *Where
	Operation operable
	Value     any
}

type operable interface {
	operation() string
}

type Or struct {
}

func (o Or) operation() string {
	return "OR"
}

type And struct {
}

func (a And) operation() string {
	return "AND"
}

type Equal struct {
}

func (e Equal) operation() string {
	return "="
}

type DatabaseTable interface {
	GetTableName() string
}

func BuildUpdateQuery(data DatabaseTable, fieldMask []string, where Where) (string, []any, error) {
	fieldMaskMap := map[string]bool{}
	for _, field := range fieldMask {
		fieldMaskMap[strings.ToLower(field)] = true
	}
	query := fmt.Sprintf("UPDATE %s SET", data.GetTableName())
	queryValues := []any{}

	s := reflect.ValueOf(data).Elem()
	typeOfT := s.Type()
	noOfUpdateFields := 0
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		if _, ok := fieldMaskMap[strings.ToLower(typeOfT.Field(i).Name)]; ok {
			if noOfUpdateFields != 0 {
				query += ","
			}
			query += fmt.Sprintf(" %s = ?", typeOfT.Field(i).Tag.Get("db"))
			queryValues = append(queryValues, f.Interface())
			noOfUpdateFields++
		}
	}
	whereCondition, whereQueryValues, err := BuildWhereCondition(where)
	if err != nil {
		return "", nil, err
	}
	query += fmt.Sprintf(" where %s", whereCondition)
	queryValues = append(queryValues, whereQueryValues...)
	return query, queryValues, nil
}

func BuildWhereCondition(where Where) (string, []any, error) {
	// Refactor where condition to return the query and query values seperately
	if where.Value == nil && where.Operation == nil {
		return "", nil, ErrInvalidWhereCondition
	}
	if where.Value != nil {
		return fmt.Sprintf("%v", where.Value), nil, nil
	}
	var left, right string
	var queryValue []any
	if where.Left != nil {
		res, values, err := BuildWhereCondition(*where.Left)
		if res == "" {
			return "", nil, ErrInvalidWhereCondition
		}
		if err != nil {
			return "", nil, nil
		}
		queryValue = append(queryValue, values...)
		left = fmt.Sprintf("%v", res)
	}
	if where.Right != nil {
		res, values, err := BuildWhereCondition(*where.Right)
		if res == "" {
			return "", nil, ErrInvalidWhereCondition
		}
		if err != nil {
			return "", nil, nil
		}
		if isLeaf(where.Right) {
			queryValue = append(queryValue, res)
			right = "?"
		} else {
			queryValue = append(queryValue, values...)
			right = fmt.Sprintf("%v", res)
		}
	}
	return fmt.Sprintf("(%v %s %v)", left, where.Operation.operation(), right), queryValue, nil
}

func isLeaf(where *Where) bool {
	return where.Value != nil
}
