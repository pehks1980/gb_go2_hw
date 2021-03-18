package main_test

import (
	"fmt"
	"github.com/pehks1980/gb_go2_hw/hw8/fscan"
	Logger "github.com/pehks1980/gb_go2_hw/hw8/logger"
	"os"
	"reflect"
	"strings"
	"testing"
)

// ErrorContains checks if the error message in out contains the text in
// want.
//
// This is safe when out is nil. Use an empty string for want if you want to
// test that err is nil.
func ErrorContains(out error, want string) bool {
	if out == nil {
		return want == ""
	}
	if want == "" {
		return false
	}
	return strings.Contains(out.Error(), want)
}
///////////////////////////// test 1 table test IOReadDir
func TestIOReadDir(t *testing.T) {

	tests := []struct {
		name string
		path string
		want    string
		wantErr string
	}{
		{"test1 - not found",
			"/homeless",
			"no such file or directory",
			"", // No error expected
		},
		{"test2 - read ok",
			"/home",
			"",
			"out of coconuts",
		},
	}

	fileSet := fscan.NewRWSet()
	err := Logger.InitLoggers("log_test.txt")
	if err != nil {
		fmt.Printf("error opening log. exiting.")
		return
	}

	deepScan := false
	for _, tc := range tests {
		_, err := fscan.IOReadDir(tc.path,fileSet,&deepScan)
		if !ErrorContains(err, tc.want) {
			t.Errorf("%s ===== unexpected error: %v", tc.name, err)
		}
	}
	_ = os.Remove("log_test.txt")
}
///////////////////////////// test 2 table test GetHash
func TestGetHash(t *testing.T) {

	tests := []struct {
		name string
		size int64
		fh string
		want string
	}{
		{"test1 - hash ok!",
			0,
			"000fff",
			"3857029705",
		},
		{"test2 - hash error!",
			100,
			"",
			"1766481806",
		},
	}

	for _, tc := range tests {
		result, _ := fscan.GetHash(tc.size,tc.name,tc.fh)
		if !reflect.DeepEqual(tc.want, result) {
			t.Fatalf("%s: expected: %v, got: %v", tc.name, tc.want, result)
		}
	}
}

///////////////////////////// test 3 table test fileSet adding editing having
func TestHashTable(t *testing.T) {

	fileSet := fscan.NewRWSet()
	// add one member

	elemMM := &fscan.FileElem{
		FullPath: "fullFilePath",
		Filesize: 123,
		FileHash: "hash",
		DubPaths: nil,
	}
	fileSet.Add("first", *elemMM)
	//check if we have it inside map
	want := true
	if got := fileSet.Has("first"); got != want {
		t.Errorf("fileSet Has = %t, want %t", got, want)
	}
	want = true
	if got := fileSet.Edit("first", "file"); got != want {
		t.Errorf("fileSet Has = %t, want %t", got, want)
	}
	want = true
	if got := fileSet.Edit("first", "file"); got != want {
		t.Errorf("fileSet Has = %t, want %t", got, want)
	}
	want = false
	if got := fileSet.Edit("second", "file"); got != want {
		t.Errorf("fileSet Has = %t, want %t", got, want)
	}
}

/////////////////// GetFileMd5Hash - not existing file
func TestGetFileMd5Hash(t *testing.T) {

	_, err := fscan.GetFileMd5Hash("aaa")
	if !ErrorContains(err, "no such file or directory") {
		t.Errorf("%s =====> unexpected error: %v", "expected: no such file or directory", err)
	}

}

/*

if !ErrorContains(err, "unexpected banana") {
    t.Errorf("unexpected error: %v", err)
}

if !ErrorContains(err, "") {
    t.Errorf("unexpected error: %v", err)
}

 */