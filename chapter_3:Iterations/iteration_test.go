package iteration

import "testing"

func TestIteration(t *testing.T) {
	t.Run("this is the basic iteration test", func(t *testing.T) {
		repeated := Repeat("a")
		expected := "aaaaa"

		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})
}
