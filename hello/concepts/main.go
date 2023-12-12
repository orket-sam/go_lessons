package main

func isPrime(x int) bool {
	var isPrime bool
	if x > 2 {
		for i := 2; i < x; i++ {
			if x%i == 0 {
				isPrime = false
				break
			}
			isPrime = true
		}
	}
	return isPrime
}

func main() {
	println(isPrime(3))
}
