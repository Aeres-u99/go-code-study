package main
import "fmt"

func main() {
    elements := map[string]map[string]string {
        "H" : map[string]string {
            "name": "Hydrogen",
            "state": "gas",
        },
        "He": map[string]string {
            "name": "Helium",
            "state": "gas",
        },
        "Li": map[string]string {
            "name": "Lithium",
            "state": "solid",
        },
    }
    if el, ok := elements["Li"]; ok {
        fmt.Println(el["name"], el["state"])
    }
    fmt.Println(elements)
}
