## Bridge Design Pattern

[Refactoring Guru](https://refactoring.guru/design-patterns/bridge)

Bridge design pattern is a structural design pattern that allows the separation of abstraction from its implementation. Sounds confusing? Don’t worry, it will be more clear as we go along.

This pattern suggests dividing a large class into two separate hierarchy

 - Abstraction – It is an interface and children of the **Abstraction** are referred to as **Refined Abstraction**. The abstraction contains a reference to the implementation.
 - Implementation – It is also an interface and children’s of the **Implementation** are referred to as **Concrete Implementation**

Abstraction hierarchy is being referred to by clients without worrying about the implementation. Let’s take an example. Assume you have two types of computer **mac** and **windows**. Also, let’s say two types of printer **epson** and **hp** . Both computers and printers needs to work with each other in any combination.  The client will only access the computer without worrying about how print is happening. Instead of creating four structs for the 2*2 combination, we create two hierarchies

 - Abstraction Hierarchy
 - Implementation Hierarchy

See the below figure. These two hierarchies communicate with each other via a bridge where **Abstraction** (computer here) contains a reference to the **Implementation**(printer here). Both the abstraction and implementation can continue to develop independently without affecting each other.  Notice how **win** and **mac** embed the reference to **printer**. We can change the **Abstraction’sImplementation** (i.e., computer’s printer) at run time as abstraction refers to implementation via the interface. On calling **mac.print()** or **windows.print()** it dispatches the request to **printer.printFile()**. This acts as a bridge and provides a loose coupling between the two.

![uml diagram with example](https://github.com/filipeandrade6/go-design-patterns/blob/master/structural/bridge/img/Bridge-Design-Pattern-1.jpg?raw=true)

### UML Diagram:

![uml diagram with example](https://github.com/filipeandrade6/go-design-patterns/blob/master/structural/bridge/img/Bridge-Design-Pattern-2.jpg?raw=true)

### Mapping

The below table represents the mapping from the UML diagram actors to actual implementation actors in “Practical Example” below

| UML | Example |
| - | - |
| Abstraction | computer.go |
| Refined Abstraction 1 | win.go |
| Refined Abstraction 2 | mac.go |
| Implementation | printer.go |
| Concrete Implementation 1 | epson.go |
| Concrete Implementation 2 | hp.go |
| Client | main.go |