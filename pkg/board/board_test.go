package board


import (
    "testing"
    "fmt"
    "slices"
    "15puzzle/pkg/coord"
)

func TestPopulateData(t *testing.T) {
	var width = 4
	data := populateData(width)
	fmt.Println(data)
	if len(data) != width {
		t.Errorf("Wrong length of data: expected %d, found %d", width, len(data))
	}
	if len(data)*len(data[0]) != width*width {
		t.Errorf("Wrong length of data: expected %d, found %d", width*width, 
		len(data)*len(data[0]))
	}
}

func TestMove(t *testing.T) {
	var width = 4
	b := NewBoard(width)
	fmt.Println(b)

	var d1, _ = coord.GetDirection(coord.DOWN)
	var d2, _ = coord.GetDirection(coord.UP)
	var d3, _ = coord.GetDirection(coord.RIGHT)
	var d4, _ = coord.GetDirection(coord.LEFT)

	var e1 = b.Move(d1)
	fmt.Println(b)
	var e2 = b.Move(d2)
	fmt.Println(b)
	var e3 = b.Move(d3)
	fmt.Println(b)
	var e4 = b.Move(d4)
	fmt.Println(b)

	if e1 != nil {
		t.Errorf("Move DOWN should be possible")
	}
	if e2 != nil {
		t.Errorf("Move UP should be possible")
	}
	if e3 != nil {
		t.Errorf("Move RIGHT should be possible")
	}
	if e4 != nil {
		t.Errorf("Move LEFT should be possible")
	}
}

func TestShuffle(t *testing.T) {
	var width = 4
	b := NewBoard(width)
	b2 := b
	fmt.Println(b)
	b.Shuffle()
	fmt.Println(b)
	var equal = true
	for i,r := range b.data {
		for j, _ := range r {
			equal = b.data[i][j] == b2.data[i][j] 
		}
	}
	if !equal {
		t.Errorf("No effect found: before %v, after %v", b, b2)
	}
}

func TestLinearizeTable(t *testing.T) {
	var width = 4
	b := NewBoard(width)
	exp := make([]int, width*width)
	for i := 1; i<len(exp); i++ {
		exp[i-1] = i
	}
	var found = linearizeTable(b.data, width)
	fmt.Println(exp)
	fmt.Println(found)
	if !slices.Equal(found, exp) {
		t.Errorf("Found unexpected value: expected %v, found %v", exp, found)
	}
}

func TestShuffleArray(t *testing.T) {
	var data = []int{1,2,3,4,5,6,7,8,9,10}
	var data2 = make([]int, 10)
	copy(data2, data)
	shuffleArray(&data2)
	fmt.Println(data)
	fmt.Println(data2)
	if slices.Equal(data, data2) {
		t.Errorf("No effect found: before %v, after %v", data, data2)
	}
}

func TestNewBoard(t *testing.T) {
	var width = 4
	b := NewBoard(width)
	fmt.Println(b)
	var data = b.data
	if b.width != width {
		t.Errorf("Wrong attribute value: expected %d, found %d", width, b.width)
	}
	if len(data) != width {
		t.Errorf("Wrong length of data: expected %d, found %d", width, len(data))
	}
	if len(data)*len(data[0]) != width*width {
		t.Errorf("Wrong length of data: expected %d, found %d", width*width, 
		len(data)*len(data[0]))
	}
}
