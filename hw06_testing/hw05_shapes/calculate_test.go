package hw05

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCalculate(t *testing.T) {
	type want struct {
		square float64
		err    bool
	}

	testCases := []struct {
		name string
		obj  any
		want want
	}{
		{
			name: "should correctly calculate the square when argument is a shape #1",
			obj:  NewCircle(5),
			want: want{
				square: 78.53981633974483,
				err:    false,
			},
		},
		{
			name: "should correctly calculate the square when argument is a shape #2",
			obj:  NewRectangle(10, 5),
			want: want{
				square: 50,
				err:    false,
			},
		},
		{
			name: "should correctly calculate the square when argument is a shape #3",
			obj:  NewTriangle(8, 6),
			want: want{
				square: 24,
				err:    false,
			},
		},
		{
			name: "should return error when argument is not a shape",
			obj:  NewApple(),
			want: want{
				err: true,
			},
		},
	}

	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			square, err := CalculateArea(tC.obj)
			if tC.want.err {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tC.want.square, square)
			}
		})
	}
}
