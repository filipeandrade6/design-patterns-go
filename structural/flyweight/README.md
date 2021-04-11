## Flyweight

[Refactoring Guru](https://refactoring.guru/design-patterns/flyweight)

It is a structural design pattern. This pattern is used when a large number of similar objects need to be created. These objects are called flyweight objects and are immutable.

Let’s first see an example. Flyweight Pattern will be clear after this example.

In a game of Counter-Strike, Terrorist and Counter-Terrorist have a different type of dress. For simplicity, let’s assume that both Terrorist and Counter-Terrorists have one dress type each. The dress object is embedded in the player object as below

Below is the struct for a player, we can see the dress object is embedded in player struct

```go
type player struct {
    dress      dress
    playerType string //Can be T or CT
    lat        int
    long       int
}
```

Let’s say there are 5 Terrorists and 5 Counter-Terrorist, so a total of 10 players. Now there are two options with respect to dress

 1. Each of the 10 player objects creates a different dress object and embed them. Total 10 dress objects will be created
 1. We create two dress object
  - Single Terrorist Dress Object: This will be shared across 5 Terrorist
  - Single Counter-Terrorist Dress Object: This will be shared across 5 Counter-Terrorist

As you can that in Approach 1, total of 10 dress objects are created while in approach 2 only 2 dress objects are created. The second approach is what we follow in the Flyweight design pattern. The two dress objects which we created are called the flyweight objects. Flyweight pattern takes out the common parts and creates flyweight objects. These flyweight objects (dress here)  can then be shared among multiple objects (player here). This drastically reduces the number of dress objects and the good part is that even if you create more players, still only two dress objects will be sufficient.

In the flyweight pattern, we store the flyweight objects in the map.  Whenever the other objects which share the flyweight objects are created, then flyweight objects are fetched from the map.

**Intrinsic and Extrinsic States**

 - **Intrinsic State** – Dress in the intrinsic state as it can be shared across multiple Terrorist and Counter-Terrorist Objects
 - **Extrinsic State** – Player location and the player weapon is an extrinsic state as it is different for every object.

### When to Use:

 - When the objects have some intrinsic properties which can be shared.
  - As in the above example, dress is the intrinsic property that was taken out and shared.
 - Use flyweight when a large number of objects needs to be created which can cause memory issue. In case figure out all the common or intrinsic state and create flyweight objects for that.

### UML Diagram:

![uml diagram with example](https://github.com/filipeandrade6/go-design-patterns/blob/master/structural/flyweight/img/Flyweight-Design-Pattern-1-1.jpg?raw=true)

Below is the corresponding mapping UML diagram with the example given above

![uml diagram with example](https://github.com/filipeandrade6/go-design-patterns/blob/master/structural/flyweight/img/Flyweight-Design-Pattern-2-1.jpg?raw=true)

### Mapping:

The below table represents the mapping from the UML diagram actors to actual implementation actors in code.

| UML | Example |
| - | - |
| Flyweight Factory | dressFactory.go |
| Flyweight Interface | dress.go |
| Concrete Flyweight Object 1 | terroristDress.go |
| Concrete Flyweight Object 1 | counterTerroristDress.go |
| Context | player.go |
| Client | main.go |