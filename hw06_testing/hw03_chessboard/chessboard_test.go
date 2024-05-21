package hw03

import (
	"bufio"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type MockReader struct {
	data   string
	offset int
}

func NewMockReader(data string) *MockReader {
	return &MockReader{data: data}
}

func (m *MockReader) Read(p []byte) (n int, err error) {
	if m.offset >= len(m.data) {
		return 0, io.EOF
	}
	n = copy(p, m.data[m.offset:])
	m.offset += n
	return n, nil
}

func (m *MockReader) ReadString(delim byte) (string, error) {
	i := strings.IndexByte(m.data[m.offset:], delim)
	if i < 0 {
		return "", io.EOF
	}
	result := m.data[m.offset : m.offset+i+1]
	m.offset += i + 1
	return result, nil
}

func TestGetChessCellValue(t *testing.T) {
	type want struct {
		cell ChessCell
		err  bool
	}

	testCases := []struct {
		name string
		cell int
		want want
	}{
		{
			name: "should return chess cell when value is valid #1",
			cell: 0,
			want: want{
				cell: white,
				err:  false,
			},
		},
		{
			name: "should return chess cell when value is valid #2",
			cell: 1,
			want: want{
				cell: black,
				err:  false,
			},
		},
		{
			name: "should return error when value is invalid",
			cell: -1,
			want: want{
				cell: invalid,
				err:  true,
			},
		},
	}

	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			got, err := getChessCellValue(tC.cell)

			if tC.want.err {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}

			assert.Equal(t, tC.want.cell, got)
		})
	}
}

func TestGenerateChessBoard(t *testing.T) {
	type want struct {
		cells [][]ChessCell
		err   bool
	}

	testCases := []struct {
		name   string
		height int
		width  int
		want   want
	}{
		{
			name:   "should return chess board when height and width are valid #1",
			height: 5,
			width:  2,
			want: want{
				cells: [][]ChessCell{
					{white, black},
					{black, white},
					{white, black},
					{black, white},
					{white, black},
				},
				err: false,
			},
		},
		{
			name:   "should return chess board when height and width are valid #2",
			height: 2,
			width:  5,
			want: want{
				cells: [][]ChessCell{
					{white, black, white, black, white},
					{black, white, black, white, black},
				},
				err: false,
			},
		},
		{
			name:   "should return chess board when height and width are valid #3",
			height: 5,
			width:  5,
			want: want{
				cells: [][]ChessCell{
					{white, black, white, black, white},
					{black, white, black, white, black},
					{white, black, white, black, white},
					{black, white, black, white, black},
					{white, black, white, black, white},
				},
				err: false,
			},
		},
		{
			name:   "should return error when height is invalid",
			height: -1,
			width:  5,
			want: want{
				err: true,
			},
		},
		{
			name:   "should return error when width is invalid",
			height: 5,
			width:  -1,
			want: want{
				err: true,
			},
		},
	}

	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			got, err := generateChessBoard(tC.height, tC.width)
			if tC.want.err {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tC.want.cells, got)
			}
		})
	}
}

func TestChessboardToString(t *testing.T) {
	testCases := []struct {
		name       string
		chessBoard [][]ChessCell
		want       string
	}{
		{
			name:       "should return empty string when chessboard is empty",
			chessBoard: [][]ChessCell{},
			want:       "",
		},
		{
			name: "should return string with formatted chessboard when chessboard is not empty #1",
			chessBoard: [][]ChessCell{
				{white, black},
				{black, white},
				{white, black},
				{black, white},
				{white, black},
			},
			want: " #\n# \n #\n# \n #\n",
		},
		{
			name: "should return string with formatted chessboard when chessboard is not empty #2",
			chessBoard: [][]ChessCell{
				{white, black, white, black, white},
				{black, white, black, white, black},
			},
			want: " # # \n# # #\n",
		},
		{
			name: "should return string with formatted chessboard when chessboard is not empty #3",
			chessBoard: [][]ChessCell{
				{white, black, white, black, white},
				{black, white, black, white, black},
				{white, black, white, black, white},
				{black, white, black, white, black},
				{white, black, white, black, white},
			},
			want: " # # \n# # #\n # # \n# # #\n # # \n",
		},
	}

	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			got := chessboardToString(tC.chessBoard)
			assert.Equal(t, tC.want, got)
		})
	}
}

func TestPromptChessboardParameters(t *testing.T) {
	type want struct {
		height int
		width  int
		err    bool
	}

	testCases := []struct {
		name string
		data string
		want want
	}{
		{
			name: "should return width and height when input is valid #1",
			data: "1\n5\n",
			want: want{
				height: 1,
				width:  5,
				err:    false,
			},
		},
		{
			name: "should return width and height when input is valid #2",
			data: "5\n1\n",
			want: want{
				height: 5,
				width:  1,
				err:    false,
			},
		},
		{
			name: "should return width and height when input is valid #3",
			data: "5\n5\n",
			want: want{
				height: 5,
				width:  5,
				err:    false,
			},
		},
		{
			name: "should return error when input is empty",
			data: "",
			want: want{
				err: true,
			},
		},
		{
			name: "should return error when height input is invalid",
			data: "s\n5\n",
			want: want{
				err: true,
			},
		},
		{
			name: "should return error when width input is invalid",
			data: "5\ns\n",
			want: want{
				err: true,
			},
		},
	}

	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			reader := NewMockReader(tC.data)
			bufReader := bufio.NewReader(reader)
			gotHeight, gotWidth, err := promptChessboardParameters(bufReader)
			if tC.want.err {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, gotHeight, tC.want.height)
				assert.Equal(t, gotWidth, tC.want.width)
			}
		})
	}
}
