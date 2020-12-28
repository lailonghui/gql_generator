package table

import (
	"fmt"
	"gqlgen-generator/config"
	"gqlgen-generator/db"
	"testing"
)

func init() {
	config.Setup("../config.yaml")
	db.Setup()
}

func Test_isTableExist(t *testing.T) {
	type args struct {
		table string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "enterprise",
			args: args{table: "enterprise"},
			want: true,
		}, {
			name: "enter",
			args: args{table: "etnet"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isTableExist(tt.args.table); got != tt.want {
				t.Errorf("isTableExist() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getTableDesc(t *testing.T) {
	type args struct {
		table string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "true",
			args: args{
				table: "enterprise",
			},
			want: "企业",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getTableDesc(tt.args.table); got != tt.want {
				t.Errorf("getTableDesc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_resolverTableInfos(t *testing.T) {
	tables := []string{"enterprise", "department", "vehicle_location_last"}
	gotTableInfos, err := ResolverTableInfos(tables)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(gotTableInfos)
}

func Test_cleanColumnType(t *testing.T) {
	c := cleanColumnType("timestamptz(6)")
	fmt.Println(c)
	c = cleanColumnType("timestamptz")
	fmt.Println(c)
}
