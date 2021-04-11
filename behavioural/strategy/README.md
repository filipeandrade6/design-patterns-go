## Strategy Design Pattern

Strategy design pattern is a behavioral design pattern. This design pattern allows you to change the behavior of an object at run time without any change in the class of that object.

Let’s understand the strategy pattern with an example. Suppose you are building an In-Memory-Cache. Since it is an In-Memory-Cache it is of limited size. Whenever it reaches its maximum size that some old entries from the cache need to be evicted. This eviction can happen via several algorithms. Some of the popular algorithms are

 - **LRU** – Least Recently Used: Remove the entry which has been used least recently.
 - **FIFO** – First In First Out: Remove the entry, which was created first.
 - **LFU** – Least Frequently Used: Remove the entry which was least frequently used.

Now the problem is how to decouple our Cache class with the algorithm such that we should be able to change the algorithm at run time. Also Cache class should not change when a new algorithm is being added. This is were Strategy Pattern comes into the picture. The strategy pattern suggests creating a family of the algorithm with each algorithm having its own class. Each of these classes follows the same interface and this makes the algorithm interchangeable within the family. Let’s say the common interface name is **evictionAlgo**.

Now our main **Cache** class will embed **evictionAlgo** interface. Instead of implementing all types of eviction algorithms in itself, our Cache class will delegate all it to the **evictionAlgo** interface. Since evictionAlgo is an interface, we can run time change the algorithm to either be LRU, FIFO, LFU without any change in Cache class.

### When to Use

 - When an object needs to support different behavior and you want to change the behavior at run time.
 - When you want to avoid a lot of conditionals of choosing the runtime behavior.
 - When you have different algorithms that are similar and they only differ in the way they execute some behavior.

### UML Diagram

Notice below UML diagram, Context (Cache) embeds the strategy (evictionAlgo) interface.

[](![uml diagram](https://github.com/filipeandrade6/go-design-patterns/blob/master/behavioural/strategy/img/Strategy-Design-Pattern-1?raw=true)

Below is the corresponding mapping UML diagram with the example given above

[](![uml diagram](https://github.com/filipeandrade6/go-design-patterns/blob/master/behavioural/strategy/img/Strategy-Design-Pattern-2?raw=true)

### Mapping

The below table represents the mapping from the UML diagram actors to actual implementation actors in code.

| UML | Example |
| - | - |
| Context | cache.go |
| Strategy | evictionAlgo.go |
| Concrete Strategy Object 1 | lfu.go |
| Concrete Strategy Object 2 | lru.go |
| Concrete Strategy Object 3 | fifo.go |
| Client | main.go |