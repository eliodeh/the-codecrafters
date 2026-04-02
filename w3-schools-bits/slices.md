# A detailed summary of slices
In golang, slices handles the length tha grows and shrins as fit.
slices are created in various ways such as using the datatype values, creating a slice from an array and using the make() function.

slice := []int, in slice there are functions that is used to re: it returns the length of the slice of element
2. cap function return capacity of numbers of elements that can grow or shrink  
# my illustration 
func main()  {
    slice1 := []int{}
    fmt.println(len(slice1))
}  

slice in Go have access elements, change elements, appendelements