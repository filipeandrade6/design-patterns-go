## Observer Design Pattern

Observer Design Pattern is a behavioral design pattern. This pattern allows an instance **(called subject)** to publish events to other multiple instances **(called observers)**.  These **observers** subscribe to the **subject** and hence get notified by events in case of any change happening in the **subject**.

Let’s take an example. In the E-Commerce website, many items go out of stock. There can be customers, who are interested in a particular item that went out of stock. There are three solutions to this problem

 1. The customer keeps checking the availability of the item at some frequency.
 1. E-Commerce bombard customers with all new items available which are in stock
 1. The customer subscribes only to the particular item he is interested in and gets notified in the case that item is available. Also, multiple customers can subscribe to the same product

Option 3 is most viable, and this is what Observer Patter is all about. The major components of the observer pattern are:

 1. **Subject** – It is the instance to which publishes an event when anything changes.
 1. **Observer** – It subscribes to the subject and gets notified by the events.

Generally, **Subject** and **Observer** are implemented as an interface. Concrete implementation of both are used

### UML Diagram:

 ![uml diagram](https://github.com/filipeandrade6/go-design-patterns/blob/master/behavioural/observer/img/Observer-Design-Pattern-1.jpg?raw=true)

### Mapping:

The below table represents the mapping from the UML diagram actors to actual implementation actors in “Practical Example” below

| UML | Example |
| - | - |
| Subject | subject.go |
| Concrete Subject | item.go |
| observer | observer.go |
| Concrete Observer 1 | customer.go |
| Client | main.go |
