package reader

import (
	"fmt"
	"testing"
)

func TestReader(t *testing.T) {
	Reader("./testdata")
	fmt.Println(configMap)
}
