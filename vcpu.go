/**
Once the go program starts,  go runtime will launch OS threads equivalent to the number of number of logical CPUs usable by the current process.  There is one logical CPU per virtual core where virtual core means

virtual_cores = x*number_of_physical_cores
where x=number of hardware threads per core

The runtime.Numcpus function can be used to get the the number of logical processors available to the GO program.

**/
package main

import (
    "fmt"
    "runtime"
)

func main() {
    fmt.Println(runtime.NumCPU())
}
