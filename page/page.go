package page

import (
	"fmt"
	"jumbosushi/simpledb/tuple"
)

// Page hold the info for each tuple
type Page struct {
	Header  Header
	Slots   []byte
	Tuplues []Tuple
}

// Header contains meta-data for a page
type Header struct {
	PageSize int
	Checkum  int
	Version  int
}

// Database is data type for database
type Database struct {
	Name   string
	Tables []Table
}

// Table is data type for table
type Table struct {
	Name string
}

// Write tuple
func Write() bool {
}

// Read tuple
func Read() []Tuple {

}

// Hello is test method for now
func Hello() {
	fmt.Println("Page imported")
}
