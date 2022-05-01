package lsystemstr

import (
	"testing"
)

func TestIncGeneration(t *testing.T) {
	system := New("F")
	rule := NewRule("F", "FF[FFA]")
	system.AddRules(rule)

	got := system.Sentence()
	want := "F"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}

	system.Increment()
	{
		got := system.Sentence()
		want := "FF[FFA]"

		if got != want {
			t.Errorf("got %q, wanted %q", got, want)
		}
	}
	system.Increment()
	{
		got := system.Sentence()
		want := "FF[FFA]FF[FFA][FF[FFA]FF[FFA]A]"

		if got != want {
			t.Errorf("got %q, wanted %q", got, want)
		}
	}
}

func TestIterate(t *testing.T) {
	system := New("F")
	rule := NewRule("F", "FF[FFA]")
	system.AddRules(rule)

	system.Iterate(2)

	got := system.Sentence()
	want := "FF[FFA]FF[FFA][FF[FFA]FF[FFA]A]"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestGetSuccessor(t *testing.T) {
	system := New("F")
	system.AddRules(NewRule("F", "FFAFF"))
	got := system.GetSuccessor("FF")
	want := "FFAFFFFAFF"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestMultiRule(t *testing.T) {
	system := New("F")
	system.AddRules(NewRule("F", "FF[A]"), NewRule("A", "FB"))

	system.Iterate(3)

	got := system.Sentence()
	want := "FF[A]FF[A][FB]FF[A]FF[A][FB][FF[A]B]"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestReplaceRules(t *testing.T) {
	system := New("F")
	system.AddRules(NewRule("F", "FF[A]"), NewRule("A", "FB"))

	system.Iterate(3)

	got := system.Sentence()
	want := "FF[A]FF[A][FB]FF[A]FF[A][FB][FF[A]B]"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestAddRulesStr(t *testing.T) {
	system := New("F")
	system.AddRulesStr("F", "FF[A]", "A", "FB")

	system.Iterate(3)

	got := system.Sentence()
	want := "FF[A]FF[A][FB]FF[A]FF[A][FB][FF[A]B]"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
