package main;

import "testing"

func TestOneRowAllSimilar(t *testing.T) {
	input := []string{
		"alice@mail.com,alic.e@mail.com,a.lice@mail.com",
	}
	want := 1;

	got := getSimilarEmailsCount(input);

	if got != want {
		t.Errorf("wanted %d but got %d", want, got);
	}
}

func TestTwoRowOneNotSimilar(t *testing.T) {
	input := []string{
		"alice@mail.com,alic.e@mail.com,a.lice@mail.com",
		"alice@mail.com,alic.e@mail.com,a.lice.dev@mail.com",
	}
	want := 1;

	got := getSimilarEmailsCount(input);

	if got != want {
		t.Errorf("wanted %d but got %d", want, got);
	}
}

func TestTwoRowOneBothSimilar(t *testing.T) {
	input := []string{
		"alice@mail.com,alic.e@mail.com,a.lice@mail.com",
		"bob@mail.com,bob+dev@mail.com,b.ob+dev@mail.com",
	}
	want := 2;

	got := getSimilarEmailsCount(input);

	if got != want {
		t.Errorf("wanted %d but got %d", want, got);
	}
}

func TestTwoColOneRejected(t *testing.T) {
	input := []string{
		"alice@mail.com,alic.e@mail.com",
		"bob@mail.com,b.o.b@mail.com",
		"foo@bar.com,foo@bar.xyz",
		"john@mail.com,j.o.h.n+123@mail.com",
	}
	want := 3;

	got := getSimilarEmailsCount(input);

	if got != want {
		t.Errorf("wanted %d but got %d", want, got);
	}
}
