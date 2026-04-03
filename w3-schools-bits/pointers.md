# THE POINTERS
pointers is the name of my group.
Pointers are the address in memory of a value. 
To get the address in memory of the value points which are (the address of operator and the dereference operator).
# The reasons for the use of pointers
1. it is for performance
2. it is for representing no value
3. it is for modification and also
4. for shared state

pointers are implemented for garbage collection and reduces risk and to save memory

# my illustration
func main() {
    age := 28
    ptr := &age
    fmt.Println("val of age:", age)
    fmt.Println("ptr to age:", ptr)
}
output will look like
value of age: 28
pointer to age: ox1400106020

This is what a pointer look like in code.