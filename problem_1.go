package main

func Solution1(target int) int {
	sum := 0
	for i := 0; i < target; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	return sum
}

func Solution2(target int) int {
	sdb := func(n int) int {
		p := (target - 1) / n
		return n * (p * (p + 1)) / 2
	}
	return sdb(3) + sdb(5) - sdb(15)
}
