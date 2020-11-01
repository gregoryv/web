package files

import "testing"

func TestMustLoad_panics(t *testing.T) {
	defer ExpectPanic(t)
	MustLoad("no-such-file")
}

func TestMustLoad_existing_file(t *testing.T) {
	got := MustLoad("package.go")
	if len(got) < 10 {
		t.Error("Did not load package.go correctly")
	}
}

func TestMustLoadLines_panics(t *testing.T) {
	defer ExpectPanic(t)
	MustLoadLines("no-such-file", 0, -1)
}

func TestMustLoadLines(t *testing.T) {
	got := MustLoadLines("package.go", 1, 1)
	if got[:2] != "/*" {
		t.Error(got)
	}
	MustLoadLines("package.go", 1, -1)
	MustLoadLines("package.go", 3, 10)
}

func ExpectPanic(t *testing.T) {
	t.Helper()
	e := recover()
	if e == nil {
		t.Fatal("expected panic")
	}
}
