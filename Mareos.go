package main

import (
	"fmt"
	"io/ioutil"
)

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

/**
 * @method main
 * @return 0
 */
func main() {

    m := Mareos{}
    fmt.Println("Map: ", m.Map())

    buffer, _ := ioutil.ReadFile("Facebook.json")
    contents := string(buffer)
    fmt.Println(contents)

}