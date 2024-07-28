package querybuilder

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

var ErrInvalidValue = errors.New("invalid value")

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
	getTableName() string
}

func BuildUpdateQuery(data DatabaseTable, fieldMask []string, where Where) (string, []any, error) {
	fieldMaskMap := map[string]bool{}
	for _, field := range fieldMask {
		fieldMaskMap[strings.ToLower(field)] = true
	}
	// Implement check to remove where condition fields in the fieldMaskMap to prevent them from updating
	query := fmt.Sprintf("UPDATE %s SET ", data.getTableName())
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
	return query, queryValues, nil
}

func BuildWhereCondition(where Where) (string, error) {
	if where.Value != nil {
		return fmt.Sprintf("%v", where.Value), nil
	}
	var left, right string
	if where.Left != nil {
		res, err := BuildWhereCondition(*where.Left)
		if res == "" {
			return "", ErrInvalidValue
		}
		if err != nil {
			return "", nil
		}
		left = fmt.Sprintf("%v", res)
	}
	if where.Right != nil {
		res, err := BuildWhereCondition(*where.Right)
		if res == "" {
			return "", ErrInvalidValue
		}
		if err != nil {
			return "", nil
		}
		right = fmt.Sprintf("%v", res)
	}
	return fmt.Sprintf("(%v %s %v)", left, where.Operation.operation(), right), nil
}
