package coord


import (
    "testing"
    "fmt"
)

func TestGetCoord(t *testing.T) {
	var d,e = GetDirection(UP)
	fmt.Println(d)
	if e != nil{
		t.Errorf("Error")
	}
	if d.di != -1 || d.dj != 0  {
		t.Errorf("Wrong direction: expected (%d,%d), found (%d,%d)", -1, 0, d.di, d.dj)
	}
}
