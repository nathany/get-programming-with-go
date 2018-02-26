package main

import "testing"

var grid = [rows][columns]int8{
	{5, 3, 0, 0, 7, 0, 0, 0, 0},
	{6, 0, 0, 1, 9, 5, 0, 0, 0},
	{0, 9, 8, 0, 0, 0, 0, 6, 0},
	{8, 0, 0, 0, 6, 0, 0, 0, 3},
	{4, 0, 0, 8, 0, 3, 0, 0, 1},
	{7, 0, 0, 0, 2, 0, 0, 0, 6},
	{0, 6, 0, 0, 0, 0, 2, 8, 0},
	{0, 0, 0, 4, 1, 9, 0, 0, 5},
	{0, 0, 0, 0, 8, 0, 0, 7, 9},
}

func TestOk(t *testing.T) {
	s := NewSudoku(grid)
	tests := []struct {
		row, column int
		digit       int8
	}{
		{1, 1, 4},
		{0, 3, 6},
		{1, 6, 4},
	}

	for _, tt := range tests {
		err := s.Set(tt.row, tt.column, tt.digit)
		if err != nil {
			t.Errorf("Expected no error, got %v for row: %d, column: %d, digit: %d.", err, tt.row, tt.column, tt.digit)
		}
		err = s.Clear(tt.row, tt.column)
		if err != nil {
			t.Fatalf("Unexpected error clearing row: %d, column: %d", tt.row, tt.column)
		}
	}
}

func TestClear(t *testing.T) {
	s := NewSudoku(grid)
	err := s.Clear(0, 0)
	if err != ErrFixedDigit {
		t.Errorf("Clear expected ErrFixedDigit, got %v", err)
	}
	err = s.Clear(1, 1)
	if err != nil {
		t.Errorf("Clear expected no error, got %v", err)
	}
}

func TestErrors(t *testing.T) {
	s := NewSudoku(grid)
	tests := []struct {
		row, column int
		digit       int8
		err         error
	}{
		{-1, 0, 8, ErrBounds},
		{0, -1, 8, ErrBounds},
		{9, 0, 8, ErrBounds},
		{0, 9, 8, ErrBounds},
		{0, 0, -1, ErrDigit},
		{0, 0, 10, ErrDigit},
		{0, 2, 5, ErrInRow},
		{2, 0, 5, ErrInColumn},
		{1, 1, 8, ErrInRegion},
		{2, 3, 7, ErrInRegion},
		{0, 0, 1, ErrFixedDigit},
	}

	for _, tt := range tests {
		err := s.Set(tt.row, tt.column, tt.digit)
		if err != tt.err {
			t.Errorf("Expected error %q, got %v for row: %d, column: %d, digit: %d.", tt.err, err, tt.row, tt.column, tt.digit)
		}
	}
}
