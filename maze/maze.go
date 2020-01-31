package main

import (
	"os"
	"fmt"
)

func readMaze(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	var row, col int
	fmt.Fscanf(file, "%d %d", &row, &col)

	maze := make([][]int, row)
	for i := range maze {
		maze[i] = make([]int, col)
		for j := range maze[i] {
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}
	return maze
}

type point struct {
	i, j int
}

var dirs = [4]point{
	{-1, 0}, {0, -1}, {1, 0}, {0, 1},
}
//是否越界们,是否符合规则
func (p point) at(grid [][]int) (int, bool) {
	if p.i < 0 || p.i >= len(grid) {
		return 0, false
	}
	if p.j < 0 || p.j >= len(grid[p.i]) {
		return 0, false

	}
	return grid[p.i][p.j], true
}

//下一个点
func (p point) add(r point) point {

	return point{p.i + r.i, p.j + r.j}
}

func walk(maze [][]int, start, end point) [][]int {
	//走过的点不为0
	step := make([][]int, len(maze))
	for i := range step {
		step[i] = make([]int, len(maze[i]))
	}

	Q := []point{start}
	for len(Q) > 0 {
		cur := Q[0]
		Q = Q[1:]
		//发现结束，退出， 也可以在next中查找终点退出
		if cur == end {
			break
		}

		for _, dir := range dirs {
			next := cur.add(dir)

			//	迷宫的下一个节点是0 才可以探索
			// and step at next  is 0 下一个没走过
			// and next ！= start 不是起点
			val, ok := next.at(maze)
			if !ok || val == 1 {
				continue
			}
			val, ok = next.at(step)
			if !ok || val != 0 {
				continue
			}
			if next == start {
				continue
			}
			//下一步的值为上一步的值+1
			curstep, _ := cur.at(step)
			step[next.i][next.j] = curstep + 1

			Q = append(Q, next)
		}
	}
	return step
}

func main() {
	maze := readMaze("maze/maze.in")
	for _, row := range maze {
		for _, col := range row {
			fmt.Printf("%d ", col)
		}
		fmt.Println()
	}

	p1 := point{1, 0}
	p2 := point{1, 0}
	fmt.Println(p1 == p2)

	step := walk(maze, point{0, 0}, point{len(maze) - 1, len(maze[0]) - 1})

	for _, row := range step {
		for _, val := range row {
			fmt.Printf("%3d", val)
		}
		fmt.Println()
	}


}
