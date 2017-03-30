package main

func main() {

}

func Collatz(start uint) (steps []uint) {
	steps = []uint{}
	for start > 1 {
		steps = append(steps, start)
		if start % 2 == 1 {
			start = start*3 + 1
		} else {
			start = start/2
		}
	}
	steps = append(steps, start)
	return
}
