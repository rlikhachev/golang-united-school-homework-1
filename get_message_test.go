package solution

import (
	"testing"

	"github.com/kyokomi/emoji/v2"
)

func TestGetMessage(t *testing.T) {
	message := GetMessage() 
	if message != emoji.Sprint("Hello, :world_map:!") {
		t.Error("Incorrect message!")
	}
}