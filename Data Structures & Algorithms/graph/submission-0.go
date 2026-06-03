import "slices"

type Graph struct {
	// it will be A -> {B, C, D}; B -> {C}; D -> {}
	adj map[int][]int
}

func NewGraph() *Graph {
	adj := make(map[int][]int)
	return &Graph{adj}
}

func (g *Graph) AddEdge(src, dst int) {
	// initialize if vertex don't exist
	if _, ok := g.adj[dst]; !ok {
		var list []int
		g.adj[dst] = list
	}
	if list, ok := g.adj[src]; !ok {
		g.adj[src] = []int{dst}
	} else {
		// append if dst not in src's
		if !slices.Contains(list, dst) {
			g.adj[src] = append(list, dst)
		}
	}
}

func (g *Graph) RemoveEdge(src, dst int) bool {
	// this early exit heavily assume that graph is valid with addEdge above
	if _, ok := g.adj[dst]; !ok { return false }

	if list, ok := g.adj[src]; ok {
		i := slices.Index(list, dst)
		if i != -1 {
			g.adj[src] = slices.Delete(list, i, i+1)
			return true
		}
	}
	return false

	// this does not remove the vertex, only edges
}

func (g *Graph) HasPath(src, dst int) bool {
	// "assume both src and dst exist" from the problem statement
	// this is hasPath, not hasEdge. careful with cycle. hmm

	// early check
	if slices.Contains(g.adj[src], dst) { return true }

	// we use BFS here
	visited := make(map[int]bool)
	visited[src] = true
	queue := slices.Clone(g.adj[src])
	for len(queue) > 0 {
		visit := queue[0]
		queue = queue[1:]
		if visited[visit] { continue }
		visited[visit] = true
		
		list := g.adj[visit]
		if slices.Contains(list, dst) { return true }
		queue = append(queue, list...)
	}
	return false
}
