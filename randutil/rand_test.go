package randutil_test

import (
	"fmt"
	"testing"

	"github.com/f1zm0/utl/randutil"
	"github.com/stretchr/testify/assert"
)

const strPoolSize = 255

func TestLengthGetRandomString(t *testing.T) {
	type args struct {
		slen      int
		generated string
	}
	var testCases []args
	for i := 1; i < strPoolSize; i += 5 {
		s := randutil.GetRandomStringN(i)
		testCases = append(testCases, args{
			slen:      i,
			generated: s,
		})
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Length_%d", tc.slen), func(t *testing.T) {
			assert.Equal(t, tc.slen, len(tc.generated))
		})
	}
}

func TestUniqueGetRandomString(t *testing.T) {
	var (
		testCases = make(map[string]int)
		fixedSlen = 16
		strToGen  = strPoolSize
	)

	for i := 0; i < strToGen; i++ {
		s := randutil.GetRandomStringN(fixedSlen)
		testCases[s] = 1
	}

	t.Run("Random string uniqueness", func(t *testing.T) {
		// the testCases map contains <strToGen unique keys
		assert.Equal(t, strToGen, len(testCases))
	})
}
