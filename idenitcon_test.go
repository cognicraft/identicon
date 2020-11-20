package identicon

import (
	"fmt"
	"os"
	"testing"
)

func TestGenerate(t *testing.T) {
	i := Generate([]byte("bart"))
	f, err := os.Create(i.Name + ".png")
	if err != nil {
		fmt.Printf("error: failed creating file for output png")
		return
	}
	defer f.Close()

	if err := i.WriteImage(f); err != nil {
		fmt.Printf("error failed writing image to file")
		return
	}
}
