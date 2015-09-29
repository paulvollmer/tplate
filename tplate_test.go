package main

import "testing"

func Test_Tplate_Process(t *testing.T) {
	source, err := Process("./fixture/hello.tplate", []string{})
	if err != nil {
		t.Error(err)
	}
	if string(source) != "hello world\n" {
		t.Error("./fixture/hello.tplate result source not equal")
	}
}

func Test_Tplate_Process_Error(t *testing.T) {
	_, err := Process("./file/not/found", []string{})
	if err == nil {
		t.Error("Error check failed")
	}
}

func Test_Tplate_Process_Vars(t *testing.T) {
	var paramtests = []struct {
		filepath string
		args     []string
		result   string
	}{
		{"./fixture/foo/var1.tplate", []string{"Foo=123"}, "hello world\n123\n"},
		{"./fixture/foo/var2.tplate", []string{"Foo=123", "Bar=456"}, "hello world\n123456\n"},
	}
	for _, tt := range paramtests {
		source, err := Process(tt.filepath, tt.args)
		if err != nil {
			t.Error(err)
		}
		if string(source) != tt.result {
			t.Error(tt.filepath + " result source not equal")
		}
	}
}
