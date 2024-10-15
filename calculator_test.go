package calculator

import "testing"

// create a test function to test the add function
func TestAdd(t *testing.T) {
	// create a test table
	tests := []struct {
		x        int
		y        int
		expected int
	}{
		{1, 1, 2},
		{1, 2, 3},
		{2, 2, 4},
		{2, 3, 5},
	}

	// loop through the test table
	for _, test := range tests {
		// run the test
		result := add(test.x, test.y)
		if result != test.expected {
			t.Errorf("Expected %d but got %d", test.expected, result)
		}
	}
}

func TestMultiply(t *testing.T) {
    // create a test table
    tests := []struct {
        x        int
        y        int
        expected int
    }{
        {1, 1, 1},
        {2, 2, 4},
        {3, 3, 9},
        {4, 5, 20},
    }

    // loop through the test table
    for _, test := range tests {
        // run the test
        result := multiply(test.x, test.y)
        if result != test.expected {
            t.Errorf("Expected %d but got %d", test.expected, result)
        }
    }
}
