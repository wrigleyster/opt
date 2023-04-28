# Opt
a light weight 'optional' library for golang


## Usage
```go
func findElement(list []Element, name string) opt.Maybe[Element] {
    for _, element := range list {
        if element.name == name {
            return opt.Some(element)
        }
    }
    return opt.No[Element]()
}
func findYoungest(list []Element) opt.Maybe[Element] {
    return opt.First(sort.Sort(list))
}

john := findElement(userList, "john")
if john.Exists {
    register(john.Value)
}
cathy := findElement(userList, "cathy").orElse(Element{name: "cathy", age: 26})
register(cathy)

tinyTim := findYoungest(userList).orPanic()
register(tinyTim)
```
