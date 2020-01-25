package tuple

// Tuple hold data for each row
type Tuple struct {
	Header     Header
	Attributes map[string]interface{}
}

// Header hold meta-data for tuples
type Header struct {
	Visibility int
}

// Write data
func Write() bool {
}

// Read data
func Read() []Tuple {

}
