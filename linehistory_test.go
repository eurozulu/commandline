package commandline

import "testing"

func TestLinehistory_Add(t *testing.T) {
	lh := &linehistory{}
	lh.Add("one")
	if len(*lh) != 1 {
		t.Fatalf("unexpected length of history after Add.  Expected %d, found %d", 1, len(*lh))
	}
	if (*lh)[0] != "one" {
		t.Fatalf("unexpected value of history after Add.  Expected %q, found %q", "one", (*lh)[0])
	}
	lh.Add("two")
	if len(*lh) != 2 {
		t.Fatalf("unexpected length of history after Add.  Expected %d, found %d", 2, len(*lh))
	}
	if (*lh)[0] != "one" {
		t.Fatalf("unexpected value at start of history after Add.  Expected %q, found %q", "one", (*lh)[0])
	}
	if (*lh)[1] != "two" {
		t.Fatalf("unexpected value at end of history after Add.  Expected %q, found %q", "two", (*lh)[1])
	}
	// Readd existing value
	lh.Add("two")
	if len(*lh) != 2 {
		t.Fatalf("unexpected length of history after Add.  Expected %d, found %d", 2, len(*lh))
	}
	if (*lh)[0] != "one" {
		t.Fatalf("unexpected value at start of history after Add.  Expected %q, found %q", "one", (*lh)[0])
	}
	if (*lh)[1] != "two" {
		t.Fatalf("unexpected value at end of history after Add.  Expected %q, found %q", "two", (*lh)[1])
	}
	// Readd existing value
	lh.Add("one")
	if len(*lh) != 2 {
		t.Fatalf("unexpected length of history after Add.  Expected %d, found %d", 2, len(*lh))
	}
	if (*lh)[0] != "two" {
		t.Fatalf("unexpected value at start of history after Add.  Expected %q, found %q", "two", (*lh)[0])
	}
	if (*lh)[1] != "one" {
		t.Fatalf("unexpected value at end of history after Add.  Expected %q, found %q", "one", (*lh)[1])
	}
}

func TestLinehistory_Previous(t *testing.T) {
	lh := &linehistory{}
	p := lh.Previous("")
	if p != "" {
		t.Fatalf("unexpected next value, expected empty string, found %q", p)
	}
	lh = testHistory()
	p = lh.Previous("")
	if p != "three" {
		t.Fatalf("unexpected previous value, expected %q, found %q", "three", p)
	}
	lh = testHistory()
	p = lh.Previous("one")
	if p != "" {
		t.Fatalf("unexpected previous value, expected empty string, found %q", p)
	}
	lh = testHistory()
	p = lh.Previous("two")
	if p != "one" {
		t.Fatalf("unexpected previous value, expected %q, found %q", "one", p)
	}
	lh = testHistory()
	p = lh.Previous("three")
	if p != "two" {
		t.Fatalf("unexpected previous value, expected %q, found %q", "two", p)
	}

	p = lh.Previous("one")
	if p != "" {
		t.Fatalf("unexpected previous value, expected empty string, found %q", p)
	}
	p = lh.Previous("one")
	p = lh.Previous("three")
	if p != "two" {
		t.Fatalf("unexpected previous value, expected %q, found %q", "three", p)
	}
}

func TestLinehistory_Next(t *testing.T) {
	lh := &linehistory{}
	n := lh.Next("")
	if n != "" {
		t.Fatalf("unexpected next value, expected empty string, found %q", n)
	}
	lh = testHistory()
	n = lh.Next("one")
	if n != "two" {
		t.Fatalf("unexpected next value, expected %q, found %q", "two", n)
	}

	lh = testHistory()
	n = lh.Next("two")
	if n != "three" {
		t.Fatalf("unexpected next value, expected %q, found %q", "three", n)
	}
	lh = testHistory()
	n = lh.Next("three")
	if n != "" {
		t.Fatalf("unexpected next value, expected empty string, found %q", n)
	}

	n = lh.Next("")
	if n != "one" {
		t.Fatalf("unexpected previous value, expected %q, found %q", "one", n)
	}
	n = lh.Previous("one")
	n = lh.Previous("three")
	if n != "two" {
		t.Fatalf("unexpected previous value, expected %q, found %q", "three", n)
	}
}

func testHistory() *linehistory {
	lh := &linehistory{}
	lh.Add("one")
	lh.Add("two")
	lh.Add("three")
	return lh
}
