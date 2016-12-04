package gorgloo

import (
	"fmt"
	"time"
)

// Priority is a level of importance
type Priority int

// Enums of priority
const (
	ILLEGAL_PRIORITY Priority = iota
	A
	B
	C
)

// Node is an unit of item in org mode
type Node struct {
	State    string
	Priority Priority
	Headline string
	Tag      string
	Sheduled time.Time
	Deadline time.Time
	Closed   time.Time
	Body     string
	Child    []*Node
}

// ArchivedNode is an unit of archived item in org mode
type ArchivedNode struct {
	Node
	Time     time.Time
	File     string
	Category string
	ToDo     string
}

func String2Priority(str string) (Priority, error) {
	switch str {
	case "[#A]":
		return A, nil
	case "[#B]":
		return B, nil
	case "[#C]":
		return C, nil
	default:
		return ILLEGAL_PRIORITY, fmt.Errorf("Invalid priority: %s", str)
	}
}
