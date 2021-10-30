package presenter

import "testing"

type FakeStructureCode struct {
	Code string
}

func TestGivenFakeStructureCodeWhenHeaderThenGetCodeString(t *testing.T) {
	fake := FakeStructureCode{}
	expected := "CODE"
	current := Header(fake)
	if current != expected {
		t.Errorf("got %q instead of %q", current, expected)
	}
}

type FakeStructureName struct {
	Name string
}

func TestGivenFakeStructureNameWhenHeaderThenGetNameString(t *testing.T) {
	fake := FakeStructureName{Name: "name"}
	expected := "NAME"
	current := Header(fake)
	if current != expected {
		t.Errorf("got %q instead of %q", current, expected)
	}
}

type FakeStructureComposed struct {
	Code        string
	Name        string
	Description string
}

func TestGivenFakeStructureComposedWhenHeaderThenGetCodeNameDescriptionStrings(t *testing.T) {
	fake := FakeStructureComposed{}
	expected := "CODE; NAME; DESCRIPTION"
	current := Header(fake)
	if current != expected {
		t.Errorf("got %q instead of %q", current, expected)
	}
}
