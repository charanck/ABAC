package querybuilder

import "testing"

func TestBuildWhereCondition(t *testing.T) {
	tt := []struct {
		name    string
		where   Where
		want    string
		wantErr error
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
			want: "(name = charan)",
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
			want: "((name = Charan) OR (age = 23))",
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
			wantErr: ErrInvalidValue,
		},
	}

	for _, test := range tt {
		got, gotErr := BuildWhereCondition(test.where)
		if test.wantErr != nil && test.wantErr != gotErr {
			t.Fatalf("failed %s, Want: %s, Got: %s", test.name, test.want, got)
		} else if test.want != got {
			t.Fatalf("failed %s, Want: %s, Got: %s", test.name, test.want, got)
		}
	}
}
