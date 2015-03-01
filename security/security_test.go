package security
import "testing"

func TestCompare_pass(t *testing.T) {
    if !compare_pass("3627909a29c31381a071ec27f7c9ca97726182aed29a7ddd2e54353322cfb30abb9e3a6df2ac2c20fe23436311d678564d0c8d305930575f60e2d3d048184d79","12345") {
        t.Error("compare_pass function is broken")
    }
}