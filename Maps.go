package main
import "fmt"

func main(){
    /*
    var x map[string]int
    var y map[int]string
    var z map[string]string
    */
    // Just ^ would give a runtime error, Maps need to be initialized
    // like v
    // Basically use this, and dont use ^
    x := make(map[string]int)
    y := make(map[int]string)
    z := make(map[string]string)
    x["key"] = 10
    y[10] = "key"
    z["hi"] = "hello"
    fmt.Println(x["key"])
    fmt.Println(y[10])
    fmt.Println(z["hi"])
    elements := make(map[string]string)
    elements["H"] = "Hydrogen"
    elements["He"] = "Helium"
    elements["Li"] = "Lithium"
    elements["Be"] = "Beryllium"
    fmt.Println(elements)
    if name, ok := elements["Un"]; ok {
        fmt.Println(name, ok)
    }
}
