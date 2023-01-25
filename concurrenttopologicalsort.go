import (
	"fmt"
	"sync"
)

// FindOrder finds the topological order of the graph represented by numCourses and prerequisites
func FindOrder(numCourses int, prerequisites [][]int) []int {
	graph := make(map[int][]int)
	indegree := make(map[int]int)
	var order []int
	var wg sync.WaitGroup

	// Create graph and indegree map
	for _, edge := range prerequisites {
		from, to := edge[1], edge[0]
		graph[from] = append(graph[from], to)
		indegree[to]++
	}

	// Create channels for each node
	channels := make(map[int]chan bool)
	for i := 0; i < numCourses; i++ {
		channels[i] = make(chan bool)
	}

	// Create goroutines for each node
	for i := 0; i < numCourses; i++ {
		wg.Add(1)
		go func(node int) {
			defer wg.Done()
			if indegree[node] == 0 {
				channels[node] <- true
			}
			<-channels[node]
			order = append(order, node)
			for _, neighbor := range graph[node] {
				indegree[neighbor]--
				if indegree[neighbor] == 0 {
					channels[neighbor] <- true
				}
			}
		}(i)
	}
	wg.Wait()
	return order
}

func main() {
	numCourses := 4
	prerequisites := [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}
	order := FindOrder(numCourses, prerequisites)
	fmt.Println(order)
}

