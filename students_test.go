package coverage

import (
	"fmt"
	"os"
	"reflect"
	"testing"
	"time"
)

// DO NOT EDIT THIS FUNCTION
func init() {
	content, err := os.ReadFile("students_test.go")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("autocode/students_test", content, 0644)
	if err != nil {
		panic(err)
	}
}

// WRITE YOUR CODE BELOW



var tPoeple = People{
	{
		firstName: "Alexa",
		lastName: "Dean",
		birthDay: time.Date(1992, 3, 8, 0, 0, 0, 0, time.UTC),
	},
	{
		firstName: "Thomas",
		lastName: "Anderson",
		birthDay: time.Date(1975, 9, 13, 0, 0, 0, 0, time.UTC),
	},
	{
		firstName: "Ronald",
		lastName: "Gregory",
		birthDay: time.Date(1998, 6, 10, 0, 0, 0, 0, time.UTC),
	},
	{
		firstName: "Crystal",
		lastName: "Hansen",
		birthDay: time.Date(1949, 1, 5, 0, 0, 0, 0, time.UTC),
	},
	{
		firstName: "Russel",
		lastName: "Rhodes",
		birthDay: time.Date(1998, 6, 10, 0, 0, 0, 0, time.UTC),
	},
	{
		firstName: "Russel",
		lastName: "Jackson",
		birthDay: time.Date(1998, 6, 10, 0, 0, 0, 0, time.UTC),
	},
}

func TestLen(t *testing.T){
	tData:=[]struct{
		input People
		expected int
	}{
		{
			input: tPoeple,
			expected: 6,
		},
		{
			input: tPoeple[2:],
			expected: 4,
		},
		{
			input: tPoeple[1:4],
			expected: 3,
		},
		{
			input: People{},
			expected: 0,
		},
	}

	for _, v:=range tData{
		result:=v.input.Len()
		if result != v.expected {
			t.Errorf("expected %d, but got %d", v.expected, result)
		}
	}
}

func TestLess(t *testing.T){
	tData:=[]struct{
		i int
		j int
		expected bool
	}{
		{
			i: 0,
			j: 1,
			expected: true,
		},
		{
			i: 4,
			j: 2,
			expected: false,
		},
		{
			i: 3,
			j: 4,
			expected: false,
		},{
			i: 4,
			j: 5,
			expected: false,
		},
	}

	for _, v:=range tData{
		result:= tPoeple.Less(v.i, v.j)
		if result != v.expected{
			t.Errorf("expected %t, but got %t", v.expected, result)
		}
	}
}

func TestSwap(t *testing.T){
	tData:=[]struct{
		i int
		j int
	}{
		{
			i: 0,
			j: 3,
		},
		{
			i: 1,
			j: 5,
		},
		{
			i: 2,
			j: 4,
		},
	}

	for _, v:=range tData{
		v1:=tPoeple[v.i]
		v2:=tPoeple[v.j]

		tPoeple.Swap(v.i, v.j)
		if v1 != tPoeple[v.j] && v2 != tPoeple[v.i]{
			t.Error("error, did not swap")
		}
	}
}

/////////////

var tMatrixList = []Matrix{
	{
		rows: 2,
		cols: 2,
		data: []int{
			1, 2, 3, 4,
		},
	},
	{
		rows: 1,
		cols: 3,
		data: []int{
			1, 2, 3,
		},
	},
	{
		rows: 2,
		cols: 1,
		data: []int{
			1, 2,
		},
	},
}


func TestNew(t *testing.T){
	tData:=[]struct{
		str string
		mtx Matrix
		err error
	}{
		{
			str: "0 1\n2 3",
			mtx: tMatrixList[0],
			err: nil,
		},
		{
			str: "100",
			mtx: tMatrixList[1],
			err: nil,
		},
		{
			str: "1\n2 3",
			mtx: tMatrixList[2],
			err: fmt.Errorf("Rows need to be the same length"),
		},
		{
			str: "1\na",
			err: fmt.Errorf("parsing error"),
		},
	}

	for _, v:=range tData{
		result, err := New(v.str)
		if result != &v.mtx && err!=v.err {
			t.Errorf("Rows need to be the same length")
		}
	}
	
}

func TestRows(t *testing.T){
	tData:=[]struct{
		input Matrix
		expected [][]int
	}{
		{
			input: tMatrixList[0],
			expected: [][]int{
				{1, 2}, {3, 4},
			},
		},
		{
			input: tMatrixList[1],
			expected: [][]int{
				{1, 2, 3},
			},
		},
		{
			input: tMatrixList[2],
			expected: [][]int{
				{1}, {2},
			},
		},
	}

	for _, v:=range tData{
		result := v.input.Rows()
		if !reflect.DeepEqual(result, v.expected){
			t.Errorf("rows method: not equal matrix")
		}
	}
}

func TestCols(t *testing.T){
	tData:=[]struct{
		input Matrix
		expected [][]int
	}{
		{
			input: tMatrixList[0],
			expected: [][]int{
				{1, 3}, {2, 4},
			},
		},{
			input: tMatrixList[1],
			expected: [][]int{
				{1}, {2}, {3},
			},
		},
		{
			input: tMatrixList[2],
			expected: [][]int{
				{1, 2},
			},
		},
	}

	for _, v:= range tData{
		result:=v.input.Cols()
		if !reflect.DeepEqual(result, v.expected){
			t.Errorf("cols method: not equal matrix")
		}
	}
}

func TestSet(t *testing.T){
	tData:=[]struct{
		input Matrix
		args struct{
			rows int
			cols int
			values int
		}
		expected bool
	}{
		{
			input: tMatrixList[0],
			args: struct{rows int; cols int; values int}{
				2, 3, 5,
			},
			expected: false,
		},
		{
			input: tMatrixList[0],
			args: struct{rows int; cols int; values int}{
				1, 1, 5,
			},
			expected: true,
		},
		{
			input: tMatrixList[1],
			args: struct{rows int; cols int; values int}{
				1, 0, 5,
			},
			expected: true,
		},
	}

	for _, v:=range tData{
		result:=v.input.Set(v.args.cols, v.args.cols, v.args.values)
		if result != v.expected{
			t.Errorf("expected %t, but got %t", v.expected, result)
		}
	}
}