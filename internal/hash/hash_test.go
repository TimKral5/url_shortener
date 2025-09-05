package hash_test

import (
	"testing"

	"github.com/timkral5/url_shortener/internal/hash"
)

func TestGenerateSHA256Hex(t *testing.T) {
	t.Parallel()

	ctrlText := "Hello World"
	ctrlHash := "A591A6D40BF420404A011733CFB7B190D62C65BF0BCDA32B57B277D9AD9F146E"

	hash := hash.GenerateSHA256Hex(ctrlText)
	if hash != ctrlHash {
		t.Error("Hash does not match the expected value.")
		t.Error("Expected", ctrlHash, "got", hash)

		return
	}
}
