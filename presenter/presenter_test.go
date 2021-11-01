package presenter

import (
	"reflect"
	"testing"
)

type FakeStructureCode struct {
	Code string
}

func TestGivenFakeStructureCodeWhenHeaderThenGetCodeString(t *testing.T) {
	fake := FakeStructureCode{}
	expected := []string{"CODE"}
	current := Header(fake)
	if !reflect.DeepEqual(current, expected) {
		t.Errorf("got %q instead of %q", current, expected)
	}
}

type FakeStructureName struct {
	Name string
}

func TestGivenFakeStructureNameWhenHeaderThenGetNameString(t *testing.T) {
	fake := FakeStructureName{Name: "name"}
	expected := []string{"NAME"}
	current := Header(fake)
	if !reflect.DeepEqual(current, expected) {
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
	expected := []string{"CODE", "NAME", "DESCRIPTION"}
	current := Header(fake)
	if !reflect.DeepEqual(current, expected) {
		t.Errorf("got %q instead of %q", current, expected)
	}
}
