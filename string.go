
package main

import "fmt"
import "strings"

func main() {

p := fmt.Println

p(strings.Count("test","t"))

p(strings.Index("test","t"))

p(strings.Repeat("a", 5))

p(strings.Split("a,b,c,d,", ","))

p(strings.Join([]string{"a","b"}, "--"))

}
