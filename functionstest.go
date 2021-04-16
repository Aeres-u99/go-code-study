package main
import "fmt"

func alternatingaddsub(elem []float64) float64 {
    total := 0.0
    var flag bool
    flag = true
    for _, v := range elem {
        if flag {
            total+=v
            flag = false
        } else {
            total-=v
        }
    }
    return total
}

func average(xs []float64) float64 {
    total := 0.0
    for _, v := range xs {
        total += v
    }
    return total / float64(len(xs))
}
func vairadicAverage(args ...float64) float64 {
    total := 0.0
    for _, v := range args {
        total += v
    }
    return total
}
func makeEvenGenerator() func() uint {
    i := uint(0)
    return func() (ret uint) {
        ret = i
        i += 2
        return
    }
}
func makeOddGenerator() func() uint {
    i := uint(1)
    return func() (ret uint) {
        ret = i
        i += 2
        return
    }
}

func factorial(x uint) uint {
    if x == 0 {
        return 1
    }
    return x * factorial(x-1)
}

func main() {
    x := 0
    add := func(x,y int) int {
        return x + y
    }
    increment := func() int {
        x++
        return x
    }
    nextEven := makeEvenGenerator()
    nextOdd := makeOddGenerator()

    xs := []float64{98,93,77,82,83,54,232,12}
    fmt.Println(average(xs))
    fmt.Println(alternatingaddsub(xs))
    fmt.Println(vairadicAverage(1.1,1.2,1.2,1.3,1.4,1.5))
    fmt.Println(add(1,1))
    fmt.Println(increment())
    fmt.Println(increment())
    fmt.Println(nextEven())
    fmt.Println(nextEven())
    fmt.Println(nextOdd())
    fmt.Println(nextOdd())
    fmt.Println(factorial(12))
    fmt.Println(factorial(10))
}
