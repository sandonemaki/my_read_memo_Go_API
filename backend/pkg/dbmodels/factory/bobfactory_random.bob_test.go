// Code generated by BobGen psql v0.38.0. DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package factory

import (
	"testing"

	"github.com/stephenafamo/bob"
)

// Set the testDB to enable tests that use the database
var testDB bob.Transactor

func TestRandom_int32(t *testing.T) {
	t.Parallel()

	val1 := random_int32(nil)
	val2 := random_int32(nil)

	if val1 == val2 {
		t.Fatalf("random_int32() returned the same value twice: %v", val1)
	}
}

func TestRandom_int64(t *testing.T) {
	t.Parallel()

	val1 := random_int64(nil)
	val2 := random_int64(nil)

	if val1 == val2 {
		t.Fatalf("random_int64() returned the same value twice: %v", val1)
	}
}

func TestRandom_string(t *testing.T) {
	t.Parallel()

	val1 := random_string(nil)
	val2 := random_string(nil)

	if val1 == val2 {
		t.Fatalf("random_string() returned the same value twice: %v", val1)
	}
}

func TestRandom_time_Time(t *testing.T) {
	t.Parallel()

	val1 := random_time_Time(nil)
	val2 := random_time_Time(nil)

	if val1.Equal(val2) {
		t.Fatalf("random_time_Time() returned the same value twice: %v", val1)
	}
}
