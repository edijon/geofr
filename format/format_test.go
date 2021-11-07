package format

import (
	"reflect"
	"testing"
)

type fakeStructureCode struct {
	Code string
}

func TestGivenFakeStructureCodeWhenHeaderThenGetCodeString(t *testing.T) {
	fake := fakeStructureCode{}
	expected := []string{"CODE"}
	current := Header(fake)
	if !reflect.DeepEqual(current, expected) {
		t.Errorf("got %q instead of %q", current, expected)
	}
}

type fakeStructureName struct {
	Name string
}

func TestGivenFakeStructureNameWhenHeaderThenGetNameString(t *testing.T) {
	fake := fakeStructureName{Name: "name"}
	expected := []string{"NAME"}
	current := Header(fake)
	if !reflect.DeepEqual(current, expected) {
		t.Errorf("got %q instead of %q", current, expected)
	}
}

type fakeStructureComposed struct {
	Code        string
	Name        string
	Description string
}

func TestGivenFakeStructureComposedWhenHeaderThenGetCodeNameDescriptionStrings(t *testing.T) {
	fake := fakeStructureComposed{}
	expected := []string{"CODE", "NAME", "DESCRIPTION"}
	current := Header(fake)
	if !reflect.DeepEqual(current, expected) {
		t.Errorf("got %q instead of %q", current, expected)
	}
}

func TestGivenFakeStructureCodeWhenRowThenGetCodeValue(t *testing.T) {
	fake := fakeStructureCode{Code: "test"}
	expected := []string{"test"}
	current := Row(fake)
	if !reflect.DeepEqual(current, expected) {
		t.Errorf("got %q instead of %q", current, expected)
	}
}

type fakeStructureSlice struct {
	Code []string
}

func TestGivenFakeStructureWithSliceWhenRowThenGetCodeValue(t *testing.T) {
	fake := fakeStructureSlice{Code: []string{"test", "test"}}
	expected := []string{"test, test"}
	current := Row(fake)
	if !reflect.DeepEqual(current, expected) {
		t.Errorf("got %q instead of %q", current, expected)
	}
}

type fakeStructureInt struct {
	Age int
}

func TestGivenFakeStructureIntWhenRowThenGetCodeValue(t *testing.T) {
	fake := fakeStructureInt{Age: 30}
	expected := []string{"30"}
	current := Row(fake)
	if !reflect.DeepEqual(current, expected) {
		t.Errorf("got %q instead of %q", current, expected)
	}
}

func TestGivenFakeStructureComposedWhenRowThenGetCodeNameDescriptionValues(t *testing.T) {
	fake := fakeStructureComposed{Code: "test1", Name: "test2", Description: "test3"}
	expected := []string{"test1", "test2", "test3"}
	current := Row(fake)
	if !reflect.DeepEqual(current, expected) {
		t.Errorf("got %q instead of %q", current, expected)
	}
}
