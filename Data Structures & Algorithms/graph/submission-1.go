import "slices"

// in this 2nd solution, we simplify and optimize a bit by
// in 1st sol, i explicitly track vertices. here, vertices are
//   implicitly defined by being referenced by the edges

type Graph struct {
	// it will be A -> {B, C, D}; B -> {C}
	adj map[int][]int
}

func NewGraph() *Graph {
	adj := make(map[int][]int)
	return &Graph{adj}
}

func (g *Graph) AddEdge(src, dst int) {
	// nil slice is valid for Contains and append
	if !slices.Contains(g.adj[src], dst) {
		g.adj[src] = append(g.adj[src], dst)
	}
}

func (g *Graph) RemoveEdge(src, dst int) bool {
	list, ok := g.adj[src]
	if !ok { return false }
	i := slices.Index(list, dst)
	if i == -1 { return false }
	g.adj[src] = slices.Delete(list, i, i+1)
	return true
}

func (g *Graph) HasPath(src, dst int) bool {
	visited := make(map[int]bool) // avoid cycle
	visited[src] = true
	queue := slices.Clone(g.adj[src]) // use BFS here
	head := 0 // use index instead of q[0] then q[1:]
	for head < len(queue) {
		visit := queue[head]
		head++
		if visited[visit] { continue }
		visited[visit] = true
		if visit == dst { return true }
		queue = append(queue, g.adj[visit]...)
	}
	return false
}
