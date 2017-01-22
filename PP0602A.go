package main
import ("fmt"
        "log")

func solve(n []int, k int) {
  for i := 1; i <= n[0]; i++ {
    if i&1 == k {
      fmt.Printf("%d ", n[i])
    }
  }
}

func readInt() (n int) {
  if _, err := fmt.Scanf("%d", &n); err != nil {
    // try again
    if _, err = fmt.Scanf("%d", &n); err != nil {
      log.Fatal("Blad odczytu int")
    }
  }
  return n
}

func main(){
  t := readInt()
  for ; t > 0; t-- {
    n := make([]int, 101, 101)
    n[0] = readInt()
    for i := 1; i <= n[0]; i++ {
      n[i] = readInt()
    }
    solve(n, 0)
    solve(n, 1)
    fmt.Println()
  }
}