## State Design Pattern

State design pattern is a behavioral design pattern that is based on Finite State Machine. We will explain the State Design Pattern in the context of an example of a Vending Machine. For simplicity, let’s assume that vending machine only has one type of item or product. Also for simplicity lets assume that a Vending Machine can be in 4 different states

 1. **hasItem**
 1. **noItem**
 1. **itemRequested**
 1. **hasMoney**

A vending machine will also have different actions. Again for simplicity lets assume that there are only four actions:

 1. **Select the item**
 1. **Add the item**
 1. **Insert Money**
 1. **Dispense Item**

### When To Use

 - **Use the State design pattern when the object can be in many different states. Depending upon current request the object needs to change its current state**
  - In the above example, Vending Machine can be in many different states. A Vending Machine will move from one state to another. Let’s say Vending Machine is in **itemRequested** then it will move to **hasMoney** state once the action “Insert Money” is done
 - Use when an object will have different responses to the same request depending upon the current state. Using state design pattern here will prevent a lot of conditional statements
  - For example in the case of Vending Machine, if a user is want to purchase an item then the machine will proceed if it is **hasItemState** or it will reject if it is in **noItemState**. If you notice here that the Vending Machine on the request of purchase of an item gives two different responses depending upon whether it is in **hasItemState** or **noItemState**.  Do notice the vendingMachine.go file below, it doesn’t have any kind of conditional statement. All the logic is being handled by concrete state implementations.

### UML Diagram

 ![uml diagram](https://github.com/filipeandrade6/go-design-patterns/blob/master/behavioural/state/img/State-Design-Pattern-1.jpg?raw=true)

### Mapping

The below table represents the mapping from the UML diagram actors to actual implementation actors in code.

| UML | Example |
| - | - |
| Context | vendingMachine.go |
| State Interface | state.go |
| Concrete State 1 | noItemState.go |
| Concrete State 2 | hasItemState.go |
| Concrete State 3 | itemRequestedState.go |
| Concrete State 4 | hasMoneyState.go |

### Explanation:
 - We have an interface “State” which defines signatures of functions that represents action in the context of Vending Machine. Below are the actions function signatures
   1. addItem(int) error
   1. requestItem() error
   1. insertMoney(money int) error
   1. dispenseItem() error
 - Each of the concrete state implementations implements all 4 above function and either move to another state on these actions or gives some response.
 - Each of the concrete state also embeds a pointer to current Vending Machine object so that state transition can happen on that object.