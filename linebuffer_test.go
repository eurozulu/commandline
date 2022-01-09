package commandline

import (
	"bytes"
	"strings"
	"testing"
)

func TestLineBuffer_Insert(t *testing.T) {
	testValue := "testvalue"
	testOut := bytes.NewBuffer(nil)
	lb := lineBuffer{Out: testOut}
	lb.SetValue([]byte(testValue))

	lb.Insert([]byte("abc"))
	if testValue+"abc" != lb.String() {
		t.Fatalf("unexpected value after insert, expected %q, found %q", testValue+"abc", lb.String())
	}

	lb.CursorHome()
	lb.CursorRight()
	lb.Insert([]byte("cba"))
	if "tcba"+testValue[1:]+"abc" != lb.String() {
		t.Fatalf("unexpected value after insert, expected %q, found %q", "acba"+testValue[1:], lb.String())
	}
}

func TestLineBuffer_Delete(t *testing.T) {
	testValue := "testvalue"
	testOut := bytes.NewBuffer(nil)
	lb := lineBuffer{Out: testOut}
	lb.SetValue([]byte(testValue))
	lb.Delete()
	if lb.String() != testValue[:len(testValue)-1] {
		t.Fatalf("unexpected value after delete, expected %q, found %q", testValue[:len(testValue)-1], lb.String())
	}
	lb.CursorHome()
	lb.CursorRight()
	lb.Delete()
	if lb.String() != testValue[1:len(testValue)-1] {
		t.Fatalf("unexpected value after delete, expected %q, found %q", testValue[1:len(testValue)-1], lb.String())
	}

}

func TestLineBuffer_SetValue(t *testing.T) {
	testValue := "testvalue"
	testOut := bytes.NewBuffer(nil)
	lb := lineBuffer{Out: testOut}
	lb.SetValue([]byte(testValue))

	if lb.String() != testValue {
		t.Fatalf("expected value %s, found %s", testValue, lb.String())
	}
	if testOut.String() != testValue {
		t.Fatalf("expected output %q, found %q", testValue, testOut.String())
	}
}

func TestLineBuffer_Reset(t *testing.T) {
	testValue := "testvalue"
	testOut := bytes.NewBuffer(nil)
	lb := lineBuffer{Out: testOut}
	lb.SetValue([]byte(testValue))
	if lb.String() != testValue {
		t.Fatalf("expected value %s, found %s", testValue, lb.String())
	}
	lb.CursorHome()
	if lb.CursorPosition() != 0 {

	}
}

func TestLineBuffer_CursorHome(t *testing.T) {
	testValue := "testvalue"
	testOut := bytes.NewBuffer(nil)
	lb := lineBuffer{Out: testOut}
	lb.SetValue([]byte(testValue))
	lb.CursorHome()
	if lb.CursorPosition() != 0 {
		t.Fatalf("Unexpected cursor position after home. Expected 0, found %d", lb.CursorPosition())
	}
	if len(testOut.Bytes()) < len(testValue)*2 {
		t.Fatalf("Unexpected output length after home. Expected %d, found %d", len(testValue)*2, len(testOut.Bytes()))
	}
	out := string(testOut.Bytes()[testOut.Len()-len(testValue)*2:])
	bs := strings.Repeat(BACKSPACE.String(), len(testValue))
	if out != testValue+bs {
		t.Fatalf("Unexpected output end after home. Expected %d backspaces, found %v", len(testValue), testOut.String()[len(testValue):])
	}
}

func TestLineBuffer_CursorEnd(t *testing.T) {
	testValue := "testvalue"
	testOut := bytes.NewBuffer(nil)
	lb := lineBuffer{Out: testOut}
	lb.SetValue([]byte(testValue))
	lb.CursorHome()
	if lb.CursorPosition() != 0 {
		t.Fatalf("Unexpected cursor position after home. Expected 0, found %d", lb.CursorPosition())
	}
	lb.CursorEnd()
	if lb.CursorPosition() != len(testValue) {
		t.Fatalf("Unexpected cursor position after End. Expected 0, found %d", lb.CursorPosition())
	}

	if len(testOut.Bytes()) < len(testValue) {
		t.Fatalf("Unexpected output length after home. Expected %d, found %d", len(testValue)*2, len(testOut.Bytes()))
	}
	out := string(testOut.Bytes()[testOut.Len()-len(testValue):])
	if out != testValue {
		t.Fatalf("Unexpected output end after home. Expected %d backspaces, found %v", len(testValue), testOut.String()[len(testValue):])
	}
}
