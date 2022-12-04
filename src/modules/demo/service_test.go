package demo

import "testing"

func Test_stringService_Count(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Count 1", args{s: "I"}, 1},
		{"Count 2", args{s: "am"}, 2},
		{"Count 4", args{s: "here"}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			st := stringService{}
			if got := st.Count(tt.args.s); got != tt.want {
				t.Errorf("Count() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_stringService_Uppercase(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"Upper a", args{s: "a"}, "A", false},
		{"Upper aba", args{s: "aba"}, "ABA", false},
		{"Error empty", args{s: ""}, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			st := stringService{}
			got, err := st.Uppercase(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("Uppercase() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Uppercase() got = %v, want %v", got, tt.want)
			}
		})
	}
}
