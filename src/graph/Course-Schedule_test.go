package graph

import "testing"

type CourseInput struct {
	numCourse     int
	prerequisites [][]int
}

func TestCourseScheduleDFS(t *testing.T) {
	var tests = []struct {
		in       CourseInput
		expected bool
	}{
		// Add Testcase Here
		{CourseInput{2, [][]int{[]int{1, 0}, []int{0, 1}}},
			false},
		{CourseInput{2, [][]int{[]int{0, 1}}},
			true},
	}
	for _, tt := range tests {
		actual := canFinishDFS(tt.in.numCourse, tt.in.prerequisites)
		if actual != tt.expected {
			t.Fatalf("get two sum failed, actual is %t, expected is %t",
				actual, tt.expected)
		}
	}
}

func TestCourseScheduleTOPO(t *testing.T) {
	var tests = []struct {
		in       CourseInput
		expected bool
	}{
		// Add Testcase Here
		{CourseInput{2, [][]int{[]int{1, 0}, []int{0, 1}}},
			false},
		{CourseInput{2, [][]int{[]int{0, 1}}},
			true},
	}
	for _, tt := range tests {
		actual := canFinishTOPOSort(tt.in.numCourse, tt.in.prerequisites)
		if actual != tt.expected {
			t.Fatalf("get two sum failed, actual is %t, expected is %t",
				actual, tt.expected)
		}
	}
}
