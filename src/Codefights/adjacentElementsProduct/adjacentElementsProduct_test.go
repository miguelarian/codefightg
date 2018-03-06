package AdjacentElementsProduct

import "testing"

func Test_1(t *testing.T) {
	expected := 1
	output := adjacentElementsProduct([]int{1, 2})
	if expected != output {
		t.Errorf("Failed Test_1 \n Expected: %d \n Output: %d ", expected, output)
	}
}
