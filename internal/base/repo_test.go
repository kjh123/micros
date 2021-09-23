package base

import (
	"context"
	"testing"
)

func TestRepo(t *testing.T) {
	r := NewRepo("http://112.126.91.233:8090/jiahui/365micros-layout.git", "")
	if err := r.Clone(context.Background()); err != nil {
		t.Fatal(err)
	}
	if err := r.CopyTo(context.Background(), "/tmp/test_repo", "112.126.91.233:8090/jiahui/365micros-layout", nil); err != nil {
		t.Fatal(err)
	}
}
