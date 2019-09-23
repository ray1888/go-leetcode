package graph

import "testing"

func TestCourseSchedule2DFS(t *testing.T) {
	var tests = []struct {
		in       CourseInput
		expected [][]int
	}{
		// Add Testcase Here
		//{CourseInput{2,[][]int{[]int{1,0}, []int{0,1}}},
		//	[][]int{}},
		{CourseInput{2, [][]int{[]int{1, 0}}},
			[][]int{[]int{0, 1}}},
		{CourseInput{4, [][]int{
			[]int{1, 0}, []int{2, 0}, []int{3, 1}, []int{3, 2},
		}},
			[][]int{[]int{0, 1, 2, 3}, []int{0, 2, 1, 3}}},
		{CourseInput{2, [][]int{}},
			[][]int{[]int{1, 0}}},
		{CourseInput{1, [][]int{}},
			[][]int{[]int{0}}},
	}
	for _, tt := range tests {
		actual := findOrder(tt.in.numCourse, tt.in.prerequisites)
		res := false
		for _, i := range tt.expected {
			breakFlag := false
			for v := 0; v < len(i); v++ {
				if actual[v] != i[v] {
					breakFlag = true
					break
				}
			}
			if !breakFlag {
				res = true
			}
		}
		if !res {
			t.Fatalf("get two sum failed, actual is %v, expected is %v",
				actual, tt.expected)
		}
	}
}

//func TestCourseSchedule2TOPO(t *testing.T) {
//	var tests = []struct {
//		in       CourseInput
//		expected [][]int
//	}{
//		// Add Testcase Here
//		{CourseInput{2,[][]int{[]int{1,0}, []int{0,1}}},
//			[][]int{}},
//		{CourseInput{2,[][]int{[]int{1,0}}},
//			[][]int{[]int{0,1}}},
//		{CourseInput{2,[][]int{[]int{1,0}}},
//			[][]int{[]int{0,1}}},
//		{CourseInput{4,[][]int{
//			[]int{1,0}, []int{2,0}, []int{3,1}, []int{3,2},
//		}},
//			[][]int{[]int{0,1,2,3}, []int{0,2,1,3}}},
//	}
//	for _, tt := range tests {
//		actual := findOrderTOPOSort(tt.in.numCourse, tt.in.prerequisites)
//		res := false
//		for _,i := range tt.expected{
//			breakFlag := false
//			for v:=0;v<len(i);v++{
//				if actual[v] != i[v]{
//					breakFlag = true
//					break
//				}
//			}
//			if !breakFlag{
//				res = true
//			}
//		}
//		if !res{
//			t.Fatalf("get two sum failed, actual is %t, expected is %t",
//				actual, tt.expected)
//		}
//	}
//
//}
