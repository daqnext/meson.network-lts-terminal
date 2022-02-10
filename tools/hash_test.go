package tools

import (
	"log"
	"testing"
)

func Test_hash(t *testing.T) {
	log.Println(FilePathToHash("bindname", "sdf/sdf/sdfs/a"))
}
