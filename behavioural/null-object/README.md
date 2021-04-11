## Null Object Design Pattern

Null Object Design Pattern is a behavioral design pattern. It is useful when the client code relies upon some dependency which can be null. Using this design pattern prevents clients from having to put null checks on the result of these dependencies. With that said, it should also be noted that client behavior is also fine with such null dependencies.

Main components of the Null Object Design Pattern are:

 1. **Entity** – it is the interface which defines primitive operations that child structs have to implement
 1. **ConcreteEntity** – it implements the entity interface
 1. **NullEntity** – it represents the null object. It also implements the entity interface but has null properties
 1. **Client** – the client gets the implementation of the entity interface and uses it. It doesn’t really care if the implementation is of ConcreteEntity or NullEntity. It treats both of them as same.

Let’s consider an example. Assume we have a college with many departments with each department having some number of professors.

**department** is represented by an interface

```go
type department interface {
    getNumberOfProfessors() int
    getName() string
}
```

where as **college** is represented as

```go
type college struct {
    departments []department
}
```

Now let’s say that there is an agency that wants to calculate the total number of professors for a particular college for specific departments only.  We will use the Null Object Design Pattern for this use case where a college will return a nullDepartment object (see **nullDepartment.go**) if a department doesn’t exists in college.
Notice the code in agency.go

 - **agency.go** doesn’t even care whether a particular department exists in **college** or not. college returns a null department object if that department doesn’t exist in the college.
 - It treats nullDepartment and real Department as same so null checks are avoided. It calls **getNumberOfProfessors()** on both the objects.

Above are two advantages that we get by using nullObject design pattern for this case. See the below code

