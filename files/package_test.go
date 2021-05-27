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

// ----------------------------------------

func TestMustLoadFunc_panics(t *testing.T) {
	defer ExpectPanic(t)
	MustLoadFunc("no-such-file", "x")
}

func TestMustLoadFunc(t *testing.T) {
	got := MustLoadFunc("package.go", "MustLoadFunc")
	if got == "" {
		t.Error("empty func")
	}
	// check start and end
	if got[0:4] != "func" || got[len(got)-1] != '}' {
		t.Error(got)
	}
}

func TestLoadFunc_fails(t *testing.T) {
	got, err := LoadFunc("package.go", "no")
	if err == nil {
		t.Errorf("found something, should fail: %s", got)
	}
}

// ----------------------------------------

func ExpectPanic(t *testing.T) {
	t.Helper()
	e := recover()
	if e == nil {
		t.Fatal("expected panic")
	}
}
