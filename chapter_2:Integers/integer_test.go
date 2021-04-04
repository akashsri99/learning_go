package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	t.Run("this is test the addition opf 2 numbers", func(t *testing.T) {
		sum := Add(2, 2)
		expected := 4
		fmt.Println(sum)
		if sum != expected {
			t.Errorf("expected %d got %d", expected, sum)
		}
		// 5
		// expected := 4
	})

}
