package main

func main() {
	var grid [][]int
	var i, x, y = 1, 0, 0

	for true {
		grid[y][x] = i
		x++
		i++
		grid[y][x] = i
		x++
		i++
	}
	/*
	   17  16  15  14  13
	   18   5   4   3  12
	   19   6   1   2  11
	   20   7   8   9  10
	   21  22  23---> ...

	   56	55	54	53	52	51	50	49
	   31	17  16  15  14  13	30	48
	   32	18   5   4   3  12	29	47
	   33	19   6   1   2  11	28	46
	   34	20   7   8   9  10	27	45
	   35	21  22  23	24	25	26	44
	   36	37	38	39	40	41	42	43

	   Data from square 1 is carried 0 steps, since it's at the access port.
	   Data from square 12 is carried 3 steps, such as: down, left, left.
	   Data from square 23 is carried only 2 steps: up twice.
	   Data from square 1024 must be carried 31 steps.

	   //Input: 277678
	*/
	//TODO
}
