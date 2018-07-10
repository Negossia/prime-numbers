package main

import (
	"fmt"
)

func main() {
	fmt.Println(sieve(100))
	fmt.Println(GetPrimes1(100))
	fmt.Println(GetPrimes2(100))
}

// Решето Эратосфена
func sieve(max int) (primes []int) {
	primesBool := make([]bool, max-1)
	for i := 2; i*i <= max; i++ {
		if !primesBool[i-2] {
			for j := i * i; j <= max; j += i {
				primesBool[j-2] = true
			}
		}
	}
	for k, v := range primesBool {
		if !v {
			primes = append(primes, k+2)
		}
	}
	return
}

// Некий предложенный алгоритм 1
func GetPrimes1(max int) (primes []int) {
	ch := make(chan int)
	go Generate(ch)
	for i := 0; ; i++ {
		prime := <-ch
		if prime > max {
			break
		}
		primes = append(primes, prime)
		ch1 := make(chan int)
		go Filter(ch, ch1, prime)
		ch = ch1
	}
	return
}

func Generate(ch chan<- int) {
	for i := 2; ; i++ {
		ch <- i
	}
}

func Filter(in <-chan int, out chan<- int, prime int) {
	for {
		i := <-in
		if i%prime != 0 {
			out <- i
		}
	}
}

// Некий предложенный алгоритм 2
func GetPrimes2(max int) (primes []int) {
	for i := 2; i <= max; i++ {
		if isPrime(i) {
			primes = append(primes, i)
		}
	}
	return
}

func isPrime(number int) bool {
	switch {
	case number < 2:
		return false
	case number == 2:
		return true
	case number%2 == 0:
		return false
	default:
		for i := 3; (i * i) <= number; i += 2 {
			if number%i == 0 {
				return false
			}
		}
		return true
	}
}
