package testcase

// refer: https://www.youtube.com/watch?v=yszygk1cpEc
import (
	"os"
	"testing"
)

func testChdir(t *testing.T, dir string) func() {
	old, err := os.Getwd()
	if err != nil {
		t.Fatalf("err :%s", err)
	}

	if err := os.Chdir(dir); err != nil {
		t.Fatalf("err :%s", err)
	}

	return func() { os.Chdir(old) }
}

func TestChdir(t *testing.T) {
	defer testChdir(t, "/tmp")
	os.Chdir("/tmp")
	pwd, _ := os.Getwd()
	t.Logf("pwd = %s", pwd)
}
