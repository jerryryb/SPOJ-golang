package main
import "fmt"

func factor(n int64) int64 {
  result := int64(1)
  for i := int64(1); i <= n; i++ {
    result *= i
  }
  return result
}

func main(){
  var d int
  var n int64
  fmt.Scanf("%d", &d)
  for ; d > 0; d-- {
    fmt.Scanf("%d", &n)
    if n < 10 {
      n = factor(n)
      fmt.Printf("%d %d\n", (n/10)%10, n%10)
    }else {
      fmt.Println("0 0")
    }
  }
}