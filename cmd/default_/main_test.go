package default_

import (
	"log"
	"testing"
	"time"
)

func TestName(t *testing.T) {
	log.Println("run test")
	log.Println(time.Now().UTC().String())
}
