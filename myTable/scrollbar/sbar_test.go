package scrollbar

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/test"
	"testing"
)

type testInput struct {
	startPos float32
	dragDist float32
	expected uint32
}

func TestSBar(t *testing.T) {
	testList := []testInput{
		testInput{
			startPos: 60,
			dragDist: 10,
			expected: 40,
		},
		testInput{
			startPos: 60,
			dragDist: -60,
			expected: 0,
		},
		testInput{
			startPos: 150,
			dragDist: 50,
			expected: 100,
		},
		testInput{
			startPos: 195,
			dragDist: -100,
			expected: 56,
		},
		testInput{
			startPos: 5,
			dragDist: 100,
			expected: 62,
		},
	}
	var sbar *ScrollBar
	var input testInput
	var caseID int
	f := func(off uint32) {
		t.Logf("got offset %d", off)
		if input.expected != off {
			t.Fatalf("got offset %d, expect %d", off, input.expected)
		}
		caseID++
	}
	sbar = NewSBar(f, false)
	w := test.NewWindow(sbar)
	w.Resize(fyne.NewSize(100, 200))
	var i int
	for i, input = range testList {
		t.Logf("testing case %+v", input)
		test.MoveMouse(w.Canvas(), fyne.NewPos(90, input.startPos))
		test.Drag(w.Canvas(), fyne.NewPos(90, input.startPos), 0, input.dragDist)
		if caseID != i+1 {
			t.Fatalf("case %d didn't get triggered", i)
		}

	}
}
