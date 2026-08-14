package main

import (
	"strings"
	"testing"
)

// TestAddNodeRejectsExcessiveVirtualNodes verifies that AddNode enforces
// a reasonable upper bound on virtualNodes to prevent memory exhaustion.
func TestAddNodeRejectsExcessiveVirtualNodes(t *testing.T) {
	s := NewService()

	// A very large virtualNodes value should be rejected
	_, err := s.AddNode("bignode", 100000)
	if err == nil {
		t.Fatal("AddNode should reject virtualNodes > 10000 to prevent memory exhaustion")
	}
	if !strings.Contains(err.Error(), "exceed") && !strings.Contains(err.Error(), "too large") {
		t.Fatalf("error should mention limit, got: %v", err)
	}

	// Values at or below the limit should still work
	_, err = s.AddNode("oknode", 10000)
	if err != nil {
		t.Fatalf("AddNode with 10000 vnodes should succeed: %v", err)
	}
}
