//
// Simple testing of our embedded resource.
//
//
package main

import (
	"io/ioutil"
	"testing"
)

//
// Test that we have one embedded resource.
//
func TestResourceCount(t *testing.T) {
	out := getResources()
	if len(out) != 1 {
		t.Errorf("We expected one resource but found %d.", len(out))
	}
}

//
// Test that our resource is identical to our static-file.
//
func TestResourceMatches(t *testing.T) {

	//
	// Get the data from our embedded copy
	//
	data, err := getResource("data/static.tmpl")
	if err != nil {
		t.Errorf("Loading our resource failed.")
	}

	//
	// Get the data from our master-copy.
	//
	master, err := ioutil.ReadFile("data/static.tmpl")
	if err != nil {
		t.Errorf("Loading our master-resource failed.")
	}

	//
	// Test the lengths match
	//
	if len(master) != len(data) {
		t.Errorf("Embedded and real resources have different sizes.")
	}

	//
	// Now test the length is the same as generated in the file.
	//
	if len(master) != getResources()[0].Length {
		t.Errorf("Data length didn't match the generated size")
	}

	//
	// Test the data-matches
	//
	if string(master) != string(data) {
		t.Errorf("Embedded and real resources have different content.")
	}
}