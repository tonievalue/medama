package util_test

import (
	"fmt"
	"testing"

	"github.com/medama-io/medama/util"
	"github.com/stretchr/testify/assert"
)

func TestGeneratesRandomString(t *testing.T) {
	stringLengths := []int{3, 99, 150, 30}

	for _, sl := range stringLengths {
		t.Run(fmt.Sprintf("String length %d", sl), func(t *testing.T) {
			s := util.GenerateRandomString(sl)
			assert.Len(t, s, sl)
		})
	}
}
