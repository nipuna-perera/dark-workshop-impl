package workshop_test

import (
	"testing"

	"github.com/nipuna-perera/dark-workshop-impl"
)

func TestVersion(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		want string
	}{
		{name: "initial release", want: "0.1.0"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := workshop.Version(); got != tt.want {
				t.Fatalf("Version() = %q, want %q", got, tt.want)
			}
		})
	}
}
