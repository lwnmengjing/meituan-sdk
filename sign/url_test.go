package sign

import "testing"

func TestValues_Encode(t *testing.T) {
	tests := []struct {
		name string
		v    Values
		want string
	}{
		// TODO: Add test cases.
		{
			"test01",
			Values{map[string][]string{
				"k": {"v"},
				"a": {"a"},
				"b": {"b"},
				"键": {"值"},
			}},
			"a=a&b=b&k=v&键=值",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.Encode(); got != tt.want {
				t.Errorf("Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}
