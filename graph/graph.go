package graph

// 207. Course Schedule
//
// There are a total of numCourses courses you have to take, labeled from 0 to
// numCourses - 1. You are given an array prerequisites where prerequisites[i]
// = [ai, bi] indicates that you must take course bi first if you want to take
// course ai.
//
// For example, the pair [0, 1], indicates that to take course 0 you have to
// first take course 1.
//
// Return true if you can finish all courses. Otherwise, return false.
//
// Example 1:
//
// Input: numCourses = 2, prerequisites = [[1,0]] Output: true Explanation:
// There are a total of 2 courses to take. To take course 1 you should have
// finished course 0. So it is possible.
//
// Example 2:
//
// Input: numCourses = 2, prerequisites = [[1,0],[0,1]] Output: false
// Explanation: There are a total of 2 courses to take. To take course 1 you
// should have finished course 0, and to take course 0 you should also have
// finished course 1. So it is impossible.
//
// Constraints:
//
//  1. 1 <= numCourses <= 2000
//  2. 0 <= prerequisites.length <= 5000
//  3. prerequisites[i].length == 2
//  4. 0 <= ai, bi < numCourses
//  5. All the pairs prerequisites[i] are unique.
func canFinish(numCourses int, prerequisites [][]int) bool {
	visited := make([]bool, numCourses)
	inStack := make([]bool, numCourses)
	courseMap := make(map[int][]int)
	for _, preprerequisite := range prerequisites {
		a, b := preprerequisite[0], preprerequisite[1]
		courseMap[a] = append(courseMap[a], b)
	}

	var checkCycle func(int) bool
	checkCycle = func(course int) bool {
		if inStack[course] {
			return true
		}

		if visited[course] {
			return false
		}

		visited[course] = true
		inStack[course] = true

		for _, p := range courseMap[course] {
			if checkCycle(p) {
				return true
			}
		}

		inStack[course] = false
		return false
	}

	for course := 0; course < numCourses; course++ {
		if checkCycle(course) {
			return false
		}
	}

	return true
}
