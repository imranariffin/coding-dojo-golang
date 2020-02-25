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

func TestAverageEmptyArray(t *testing.T) {
	input := []int{};
	want := float64(0);

	got := average(input);

	if got != want {
		t.Errorf("wanted %f but got %f", want, got);
	}
}

func TestAverageFewPositiveEvenNumbers(t *testing.T) {
	input := []int{2,4};
	want := float64(3);

	got := average(input);

	if got != want {
		t.Errorf("wanted %f but got %f", want, got);
	}
}

func TestAverageFewPositiveOddNumbers(t *testing.T) {
	input := []int{3,2};
	want := 2.5;

	got := average(input);

	if got != want {
		t.Errorf("wanted %f but got %f", want, got);
	}
}

func TestAverageManyPositiveNegativeNumbers(t *testing.T) {
	input := []int{
		-6,
		1,
		5,
		3,
		-4,
		7,
		2,
		2,
	};
	want := float64(1.25);

	got := average(input);

	if got != want {
		t.Errorf("wanted %f but got %f", want, got);
	}
}

