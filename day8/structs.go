package day8

type Coords struct {
	X int
	Y int
	Z int
}

type Edge struct {
	Weight int64
	U, V   int
}

type DSU struct {
	Parent []int
	Size   []int
}

func NewDSU(n int) *DSU {
	parent := make([]int, n)
	size := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		size[i] = 1
	}
	return &DSU{Parent: parent, Size: size}
}

func (d *DSU) Find(i int) int {
	if d.Parent[i] == i {
		return i
	}
	d.Parent[i] = d.Find(d.Parent[i])
	return d.Parent[i]
}

func (d *DSU) Union(i, j int) bool {
	rootI := d.Find(i)
	rootJ := d.Find(j)

	if rootI != rootJ {
		if d.Size[rootI] < d.Size[rootJ] {
			rootI, rootJ = rootJ, rootI
		}
		d.Parent[rootJ] = rootI
		d.Size[rootI] += d.Size[rootJ]
		return true
	}
	return false
}
