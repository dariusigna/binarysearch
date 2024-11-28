package index

import "testing"

func TestFinder_FindIndex(t *testing.T) {
	tests := []struct {
		name   string
		values []int
		value  int
		want   int
		found  bool
	}{
		{
			name:   "empty list",
			values: []int{},
			value:  3,
			want:   -1,
			found:  false,
		},
		{
			name:   "value found",
			values: []int{1, 2, 3, 4, 5},
			value:  3,
			want:   2,
			found:  true,
		},
		{
			name:   "value found with conformation - index below",
			values: []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100},
			value:  22,
			want:   1,
			found:  true,
		},
		{
			name:   "value found with conformation - index above",
			values: []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100},
			value:  28,
			want:   2,
			found:  true,
		},
		{
			name:   "value not found and not in conformation - greater",
			values: []int{1, 2, 3, 4, 5, 8, 9, 10},
			value:  12,
			want:   -1,
			found:  false,
		},
		{
			name:   "value not found and not in conformation - lesser",
			values: []int{4, 5, 8, 9, 10},
			value:  2,
			want:   -1,
			found:  false,
		},
		{
			name:   "value not found with conformation",
			values: []int{10, 20, 30, 40, 50},
			value:  25,
			want:   -1,
			found:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := NewFinder(tt.values, 10)
			got, found := f.FindIndex(tt.value)
			if got != tt.want {
				t.Errorf("Finder.FindIndex() got = %v, want %v", got, tt.want)
			}
			if found != tt.found {
				t.Errorf("Finder.FindIndex() found = %v, want %v", found, tt.found)
			}
		})
	}
}
