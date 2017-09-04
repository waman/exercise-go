package main

import "fmt"

// 適当に addEdge と hasEdge を使うコード
func main(){
	addEdge("a", "b")
	addEdge("a", "c")
	addEdge("b", "c")
	addEdge("b", "z")
	fmt.Println(graph)

	fmt.Printf("a -> b : %v\n", hasEdge("a", "b"))
	fmt.Printf("a -> z : %v\n", hasEdge("a", "z"))
}

var graph = make(map[string]map[string]bool)

func addEdge(from, to string){
	edges := graph[from]
	if edges == nil {
		edges = make(map[string]bool)
		graph[from] = edges
	}
	edges[to] = true
}

func hasEdge(from, to string) bool {
	return graph[from][to]
}
