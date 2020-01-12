package a_star_test

import (
	"math"
	"testing"
	"github.com/a-star"
	"fmt"
)

type mynode struct {
	x int
	y int

	value int

	sub_node_list []a_star.Node
}

func (n *mynode) Location() int {
	return n.x * 1000 + n.y
}

func (n *mynode) SurNodeList() []a_star.Node {
	return n.sub_node_list
}

func (n *mynode) Distance(dst a_star.Node) float64 {
	x := dst.Location() / 1000
	y := dst.Location() % 1000

	return math.Sqrt(float64((n.x - x) * (n.x - x) + (n.y - y) * (n.y - y)))
}

func (n *mynode) Value() float64 {
	return float64(n.value)
}

var graph [10][10]int = [10][10]int{
	{ 0, 0, 0, 1, 0, 0, 0, 0, 0, 0},
	{ 0, 1, 0, 0, 0, 1, 0, 1, 0, 1},
	{ 0, 0, 1, 1, 1, 1, 0, 0, 0, 0},
	{ 4, 0, 7, 0, 0, 0, 0, 0, 1, 0},
	{ 4, 1, 6, 1, 1, 0, 1, 0, 1, 0},
	{ 4, 0, 2, 0, 0, 0, 1, 0, 0, 0},
	{ 1, 4, 1, 0, 1, 0, 0, 1, 1, 1},
	{ 0, 0, 0, 0, 0, 9, 1, 0, 0, 0},
	{ 0, 1, 1, 9, 1, 9, 1, 0, 1, 1},
	{ 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
}

func buildNodeList() [10][10]*mynode {

	var node_list [10][10]*mynode

	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			n := &mynode{}
			n.x = i
			n.y = j

			if graph[i][j] == 0 {
				n.value = 1
			} else if graph[i][j] ==1 {
				n.value = -1
			} else {
				n.value = graph[i][j]
			}

			node_list[i][j] = n
		}
	}

	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if (node_list[i][j].value == -1) {
				continue
			}

			if i-1 >= 0 && node_list[i-1][j].value != -1 {
				node_list[i][j].sub_node_list = append(node_list[i][j].sub_node_list, node_list[i-1][j])
			}

			if i+1 < 10 && node_list[i+1][j].value != -1 {
				node_list[i][j].sub_node_list = append(node_list[i][j].sub_node_list, node_list[i+1][j])
			}

			if j-1 >= 0 && node_list[i][j-1].value != -1 {
				node_list[i][j].sub_node_list = append(node_list[i][j].sub_node_list, node_list[i][j-1])
			}

			if j+1 < 10 && node_list[i][j+1].value != -1 {
				node_list[i][j].sub_node_list = append(node_list[i][j].sub_node_list, node_list[i][j+1])
			}
		}
	}

	return node_list
}

func TestAStar_BFS(t *testing.T) {

	node_list := buildNodeList()

	var astar a_star.AStar

	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			fmt.Print(graph[i][j])
		}
		fmt.Print("\n");
	}

	road := astar.BFS(node_list[0][0], node_list[9][9])

	fmt.Println("find road")
	for _, v := range road {
		x := v.Location() / 1000
		y := v.Location() % 1000

		graph[x][y] = 8
	}

	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			fmt.Print(graph[i][j])
		}
		fmt.Print("\n");
	}
}