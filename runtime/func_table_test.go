package runtime

import "testing"

func TestFuncTable(t *testing.T) {
	f := NewFuncTable(`
push 1
push 2
$sum

FUNC @sum:
    arg a, b

    push a
    push b
    add
    ret ~
ENDFUNC
	`)
	err := f.Initialize()
	if err != nil {
		t.Fatal(err)
	}

	fn, ok := f.SearchFunc("sum")
	if !ok {
		t.Fatal("should be ok")
	}

	cmd, _ := fn.Next()
	if cmd.Operate != "arg" {
		t.Fatalf("should be `arg` cmd, but `%v`", cmd)
	}

	if len(cmd.Values) != 2 {
		t.Fatalf("should be 2 value, but `%v`", cmd)
	}
}
