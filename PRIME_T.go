package main
import ("os"
        "fmt"   
        "time"
        "bufio")

func generatePrimes(n int, primes []int) []int {
	counter := primes[len(primes)-1] + 1
	for {
		if n < primes[len(primes)-1] {
			break
		}
		prime := true
		for _, elem := range primes {
			if counter % elem == 0 {
				prime = false
				break
			}
      if elem * elem >= counter {
        break
			}
		}
		if prime {
			primes = append(primes, counter)
		}
		counter++
	}
	return primes
}

func main() {
  start := time.Now()
  bio := bufio.NewWriter(os.Stdout)
  
  var primes = make([]int, 2, 10100)
  primes[0] = 2
  primes[1] = 3
  primes = generatePrimes(10100, primes)
  endgen := time.Since(start)
  var n, t int
  fmt.Scanf("%d", &t)
  var cache = make([]bool, 100100, 100100)
  for i:= 0; i < len(primes); i++ {
    cache[primes[i]] = true
  }
  
  for ;t > 0; t-- {
    fmt.Fscanf(os.Stdin, "%d", &n)
    if cache[n] {
      fmt.Fprintln(bio, "TAK")
    }else{
      fmt.Fprintln(bio, "NIE")
    }
  }
  end := time.Since(start)
  if 0 == 1 {
    fmt.Printf("gen prime: %v \n", endgen)
    fmt.Printf("total: %v \n", end)
  }
  bio.Flush()
}