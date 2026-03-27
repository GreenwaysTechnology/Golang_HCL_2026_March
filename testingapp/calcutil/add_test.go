package calcutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// single possiblity
//func TestAdd(t *testing.T) {
//	result := Add(1, 4)
//	if result != 3 {
//		//
//		t.Errorf("Add(1,2) = %d; want %d", result, 3)
//	}
//}

//func setup() (int, int) {
//	return 2, 3
//}
//func tearDown() {
//	//clean up activity
//	fmt.Println("tearDown()")
//}

// multiple possiblity
//func TestAdd_TableDriven(t *testing.T) {
//	tests := []struct {
//		name     string //label
//		a, b     int    //input
//		expected int    //output
//	}{
//		{"add Result should be 5", 2, 3, 5},
//		{"zero", 0, 0, 0},
//		{"negative", -1, -10, -11},
//	}
//	for _, test := range tests {
//		t.Run(test.name, func(t *testing.T) {
//			if result := Add(test.a, test.b); result != test.expected {
//				t.Errorf("add(%d, %d) = %d, want %d", test.a, test.b, result, test.expected)
//			}
//		})
//	}
//}

//func TestAdd_WithSetup(t *testing.T) {
//	data1, data2 := setup()
//	defer tearDown()
//	if result := Add(data1, data2); result != 5 {
//		t.Fail()
//	}
//}

func TestAdd_WithTestify(t *testing.T) {
	result := Add(1, 2)
	assert.Equal(t, 3, result)
}
