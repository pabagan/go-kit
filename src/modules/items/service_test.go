package items

import "testing"

func Test_itemService_Create(t *testing.T) {
	tests := []struct {
		name    string
		args    Item
		want    Item
		wantErr bool
	}{
		{"Upper a", Item{Name: "a"}, Item{Name: "a"}, false},
		{"Error empty", Item{Name: ""}, Item{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			is := itemService{}
			got, err := is.Save(tt.args)
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
