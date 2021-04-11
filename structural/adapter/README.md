## Adapter Design Pattern

[Refactoring Guru](https://refactoring.guru/design-patterns/adapter)

This design pattern is a Structural Design Pattern. The patter is best understood with an example. Let’s say you have two laptops

 1. MacBook Pro
 1. Windows Laptop

MacBook Pro has a USB port that is square in shape and Windows have a USB port that is circular in shape. You as a client have a USB cable that is square in shape so it can only be inserted in the mac laptop. So you see the problem here

#### Problem:

 - We have a class (Client) that is expecting some features of an object (square USB port here), but we have another object called adaptee (Windows Laptop here) which offers the same functionality but through a different interface( circular port)

This is where Adapter Pattern comes into the picture. We create a class known as Adapter that will

 - Adhere to the same interface which client expects ( Square USB port here)
 - Translate the request from the client to the adaptee in the form that adaptee expects. Basically, in our example act as an adapter that accepts USB in square port and then inserts into circular port in windows laptop.

### When to Use

 - Use this design pattern when the objects implement a different interface as required by the client.

### UML Diagram

![uml diagram with example](https://github.com/filipeandrade6/go-design-patterns/blob/master/structural/adapter/img/Adapter-Design-Pattern-1.jpg?raw=true)

Below is the corresponding mapping UML diagram with the example given above

![uml diagram with example](https://github.com/filipeandrade6/go-design-patterns/blob/master/structural/adapter/img/Adapter-Design-Pattern-2.jpg?raw=true)

### Mapping

The below table represents the mapping from the UML diagram actors to actual implementation actors in code.

| UML | Example |
| - | - |
| Target | computer.go |
| Concrete Prototype 1 | mac.go |
| Concrete Prototype 2 (Adapter) | windowsAdapter.go |
| adaptee | windows.go |
| client | client.go |