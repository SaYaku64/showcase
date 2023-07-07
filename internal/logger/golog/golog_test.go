package golog

import "testing"

func Test_newSpace(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want string
	}{
		{
			name: "1 symb",
			str:  "0",
			want: "0\n",
		},
		{
			name: "exists",
			str:  "My string\n",
			want: "My string\n",
		},
		{
			name: "absent",
			str:  "My second string",
			want: "My second string\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newSpace(tt.str); got != tt.want {
				t.Errorf("newSpace() = %v, want %v", got, tt.want)
			}
		})
	}
}
