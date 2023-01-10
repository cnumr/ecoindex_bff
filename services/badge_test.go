package services

import "testing"

func TestGetColor(t *testing.T) {
	type args struct {
		grade string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "A",
			args: args{grade: "A"},
			want: "#349A47",
		},
		{
			name: "B",
			args: args{grade: "B"},
			want: "#51B84B",
		},
		{
			name: "C",
			args: args{grade: "C"},
			want: "#CADB2A",
		},
		{
			name: "D",
			args: args{grade: "D"},
			want: "#F6EB15",
		},
		{
			name: "E",
			args: args{grade: "E"},
			want: "#FECD06",
		},
		{
			name: "F",
			args: args{grade: "F"},
			want: "#F99839",
		},
		{
			name: "G",
			args: args{grade: "G"},
			want: "#ED2124",
		},
		{
			name: "Unkonwn",
			args: args{grade: ""},
			want: "lightgrey",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetColor(tt.args.grade); got != tt.want {
				t.Errorf("GetColor() = %v, want %v", got, tt.want)
			}
		})
	}
}
