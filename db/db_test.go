package db
import (
    "testing"
)

func TestDbInit(t *testing.T) {
    err  := Init()
    if err != nil {
        t.Fatalf("db not init success")
    }
}