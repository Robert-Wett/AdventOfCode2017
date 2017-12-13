package main

import (
	"fmt"
	"math"
)

/*
def spiral(X, Y):
    x = y = 0
    dx = 0
    dy = -1
    for i in range(max(X, Y)**2):
        if (-X/2 < x <= X/2) and (-Y/2 < y <= Y/2):
            print (x, y)
            # DO STUFF...
        if x == y or (x < 0 and x == -y) or (x > 0 and x == 1-y):
            dx, dy = -dy, dx
		x, y = x+dx, y+dy
*/

func spiral(X, Y int) {
	hi := make([][]int, X, Y)
	var x, y = 0, 0
	var dx, dy = 0, -1
	max := math.Exp2(math.Max(float64(X), float64(Y)))
	for i := 0; i < int(max); i++ {
		if (-X/2 < x && x <= X/2) && (-Y/2 < y && y <= Y/2) {
			hi[y][x] = i
		}
		if x == y || (x < 0 && x == -y) || (x > 0 && x == 1-y) {
			dx, dy = -dy, dx
		}
		x, y = x+dx, y+dx
	}
	fmt.Println(hi)
}

func main() {
	spiral(20, 20)
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
