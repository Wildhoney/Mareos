package main

import "fmt"

/**
 * @struct Mareos
 * @return void
 */
type Mareos struct {
    length, width int
}

/**
 * @method Map
 * @return void
 */
func (m Mareos) Map() string {
	return "Map"
}

/**
 * @method Group
 * @return void
 */
func (m Mareos) Group() string {
	return "Group"
}

/**
 * @method Reduce
 * @return void
 */
func (m Mareos) Reduce() string {
	return "Reduce"
}

func main() {
    m := Mareos{}
    fmt.Println("Map: ", m.Map())
}