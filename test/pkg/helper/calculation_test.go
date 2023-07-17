package helper_test

import (
	"fmt"
	"log"
	"sesi-11/pkg/helper"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSum(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	expected := 15

	got := helper.Sum(nums...)

	// if expected != got {
	// 	t.Error("expected", expected, "got", got)
	// 	// t.Error()
	// }

	// assert.Equal(t, expected, got)
	require.Equal(t, expected, got)

	log.Println("done")
}

func TestSum_WithTableTest(t *testing.T) {
	type testData struct {
		expected   int
		got        int
		errMessage string
	}

	var tests = []testData{
		{
			expected:   15,
			got:        helper.Sum(1, 2, 3, 4, 5),
			errMessage: "not equal",
		},
		{
			expected:   10,
			got:        helper.Sum(1, 2, 3, 4),
			errMessage: "not equal",
		},
		{
			expected:   0,
			got:        helper.Sum(),
			errMessage: "not equal",
		},
	}

	for i, te := range tests {
		t.Run(fmt.Sprintf("running %v", i+1), func(t *testing.T) {
			require.Equal(t, te.expected, te.got, te.errMessage)
		})
	}
}
