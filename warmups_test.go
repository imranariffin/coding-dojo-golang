package main;

import "testing"

func TestEmptyArray(t *testing.T) {
	input := []int{};
	want := -1;

	got := max(input);

	if got != want {
		t.Errorf("wanted %d but got %d", want, got);
	}
}

func TestReverseSorted(t *testing.T) {
	input := []int{
		3,
		2,
		1,
	};
	want := 3;

	got := max(input);

	if got != want {
		t.Errorf("wanted %d but got %d", want, got);
	}
}

func TestSorted(t *testing.T) {
	input := []int{
		4,
		5,
		6,
	};
	want := 6;

	got := max(input);

	if got != want {
		t.Errorf("wanted %d but got %d", want, got);
	}
}

func TestArbitraryOrder(t *testing.T) {
	input := []int{
		6,
		1,
		5,
		3,
		4,
		7,
		2,
		2,
	};
	want := 7;

	got := max(input);

	if got != want {
		t.Errorf("wanted %d but got %d", want, got);
	}
}

