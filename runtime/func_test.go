package runtime

import "testing"

func TestFunc(t *testing.T) {
	f := NewFunc(`
push 1   ; 1 argu
push 2
push 3
mul
add
pop x ; pop command
		_NPRINT:;label 
print "hello world" ; hello
		var a,b,c; vaiables
`)
	err := f.Initialize()
	if err != nil {
		t.Fatal(err)
	}

	cmd, _ := f.Next()
	if cmd.Operate != "push" {
		t.Fatalf("should be `push` cmd, but `%v`", cmd)
	}

	if len(cmd.Values) != 1 {
		t.Fatalf("should be 1 arg, but `%v`", cmd)
	}

	cmd, _ = f.Next()
	if cmd.Values[0].(int) != 2 {
		t.Fatalf("should be 2 value, but `%v`", cmd)
	}

	f.Goto("_NPRINT")
	cmd, _ = f.Next()
	if cmd.Operate != "print" {
		t.Fatalf("should be `print` cmd, but `%v`", cmd)
	}

	if len(cmd.Values) != 1 {
		t.Fatalf("should be 1 arg, but `%v`", cmd)
	}

	if cmd.Values[0].(string) != "hello world" {
		t.Fatalf("should be `hello world` value, but `%v`", cmd)
	}

	cmd, _ = f.Next()
	if len(cmd.Values) != 3 {
		t.Fatalf("should be 3 arg, but `%v`", cmd)
	}

	if cmd.Values[1].(string) != "b" {
		t.Fatalf("should be `b` value, but `%v`", cmd)
	}

}
