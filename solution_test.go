package solution

import (
	"fmt"
	"testing"

	"github.com/kyokomi/emoji/v2"
)

func TestGetMessage(t *testing.T) {
	message := GetMessage()
	fmt.Println(message) 
	if message != emoji.Sprint("Hello, :world_map:!") {
		t.Error("Incorrect message!")
	}
	
}
