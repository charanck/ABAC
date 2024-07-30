package querybuilder

import (
	"testing"

	"github.com/charanck/ABAC/internal/model"
)

func TestBuildWhereCondition(t *testing.T) {
	tt := []struct {
		name           string
		where          Where
		wantQuery      string
		wantQueryValue []any
		wantErr        error
	}{
		{
			name: "should generate where condition",
			where: Where{
				Left: &Where{
					Value: "name",
				},
				Right: &Where{
					Value: "charan",
				},
				Operation: Equal{},
			},
			wantQuery: "(name = ?)",
		},
		{
			name: "should prepare where condition for nested condition",
			where: Where{
				Left: &Where{
					Left: &Where{
						Value: "name",
					},
					Right: &Where{
						Value: "Charan",
					},
					Operation: Equal{},
				},
				Right: &Where{
					Left: &Where{
						Value: "age",
					},
					Right: &Where{
						Value: 23,
					},
					Operation: Equal{},
				},
				Operation: Or{},
			},
			wantQuery: "((name = ?) OR (age = ?))",
		},
		{
			name: "should fail when value is empty",
			where: Where{
				Left: &Where{
					Value: "",
				},
				Right: &Where{
					Value: "charan",
				},
				Operation: Equal{},
			},
			wantErr: ErrInvalidWhereCondition,
		},
	}

	for _, test := range tt {
		got, gotQueryValue, gotErr := BuildWhereCondition(test.where)
		if test.wantErr != gotErr || test.wantQuery != got {
			t.Fatalf("failed %s, WantQuery: %s, wantQueryValue:%s, GotQuery: %s, GotQueryValue: %s", test.name, test.wantQuery, test.wantQuery, got, gotQueryValue)
		}
	}
}

func TestBuildUpdateQuery(t *testing.T) {
	type buildUpdateQueryParameters struct {
		data      DatabaseTable
		fieldMask []string
		where     Where
	}
	tt := []struct {
		name            string
		parameters      buildUpdateQueryParameters
		wantQuery       string
		wantQueryValues []any
		wantErr         error
	}{
		{
			name: "should build update query successfully",
			parameters: buildUpdateQueryParameters{
				data:      &model.Resource{},
				fieldMask: []string{"ownerId", "policyId"},
				where: Where{
					Left: &Where{
						Value: "id",
					},
					Right: &Where{
						Value: "abc",
					},
					Operation: Equal{},
				},
			},
			wantQueryValues: []any{
				"",
				"",
			},
			wantQuery: "UPDATE resource SET owner_id = ?, policy_id = ? where (id = ?)",
		},
	}

	for _, test := range tt {
		gotQuery, gotQueryValue, gotErr := BuildUpdateQuery(test.parameters.data, test.parameters.fieldMask, test.parameters.where)
		if test.wantErr != gotErr || test.wantQuery != gotQuery {
			t.Fatalf("failed %s, wantQuery: %s, wantQueryValue:%v, wantErr:%s gotQuery: %s, gotQueryValue:%v, gotErr:%s", test.name, test.wantQuery, test.wantQueryValues, test.wantErr, gotQuery, gotQueryValue, gotErr)
		}
	}

}
