package stabledifussion

import (
	"context"
	"testing"
)

func TestSynxis(t *testing.T) {
	stableDiffusion := StableDiffusion{}
	stableDiffusion.Process(context.Background(), "")
}
