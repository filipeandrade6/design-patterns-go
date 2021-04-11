## Factory Design Pattern

[Refactoring Guru](https://refactoring.guru/design-patterns/factory-method)

Factory design pattern is a creational design pattern and it is also one of the most commonly used pattern. This pattern provides a way to hide the creation logic of the instances being created.

The client only interacts with a factory struct and tells the kind of instances that needs to be created. The factory class interacts with the corresponding concrete structs and returns the correct instance back.

In below example

 - We have **iGun** interface which defines all methods a gun should have
 - There is **gun** struct that implements the **iGun** interface.
 - Two concrete guns **ak47** and **maverick**. Both embed gun struct and hence also indirectly implement all methods of **iGun** and hence are of **iGun** type
 - We have a gunFactory struct which creates the gun of type **ak47** or **maverick**.
 - The **main.go** acts as a client and instead of directly interacting with **ak47** or **maverick**, it relies on **gunFactory** to create instances of **ak47** and **maverick**

### UML Diagram

![uml diagram](https://github.com/filipeandrade6/go-design-patterns/blob/master/creational/factory/img/Factory-Design-Pattern-1.jpg?raw=true)

UML Diagram with the example given above

![uml diagram with example](https://github.com/filipeandrade6/go-design-patterns/blob/master/creational/factory/img/Factory-Design-Pattern-2.jpg?raw=true)

### Mapping:

The below table represents the mapping from the UML diagram actors to actual implementation actors in **“Example”** below

| UML | Example |
| - | - |
| ProductFactory | gunFactory.go |
| iProduct | iGun.go |
| Product | gun.go |
| Concrete iProduct 1 | ak47go |
| Concrete iProduct 1 | maverick.go |
| Client | main.go |