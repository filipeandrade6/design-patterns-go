## Abstract Factory Design Pattern

[Refactoring Guru](https://refactoring.guru/design-patterns/abstract-factory)

Abstract Factory Design Pattern is a creational design pattern that lets you create a family of related objects. It is an abstraction over the factory pattern. It is best explained with an example. Let’s say we have two factories

 - **nike**
 - **adidas**

Imagine you need to buy a sports kit which has a **shoe** and **short**. Preferably most of the time you would want to buy a full sports kit of a similar factory i.e either **nike** or **adidas**. This is where the abstract factory comes into the picture as concrete products that you want is **shoe** and a **short** and these products will be created by the abstract factory of **nike** and **adidas**.

Both these two factories – **nike** and **adidas** implement iSportsFactory interface.

We have two product interfaces.

 - **iShoe** – this interface is implemented by **nikeShoe** and **adidasShoe** concrete product.
 - **iShort** – this interface is implemented by **nikeShort** and **adidasShort** concrete product.