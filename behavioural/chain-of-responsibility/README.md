## Chain of Responsibility Design Pattern

Chain of Responsibility Design Pattern is a behavioral design pattern. It lets you create a chain of request handlers. For every incoming request, it is passed through the chain and each of the handler:

 1. Processes the request or skips the processing.
 1. Decides whether to pass the request to the next handler in the chain or not

Chain of Responsibility Design pattern will be best understood with an example. Let’s take the case of a hospital. A hospital has multiple departments such as:

 1. Reception
 1. Doctor
 1. Medicine Room
 1. Cashier

Whenever any patient arrives he first goes to **Reception** then to **Doctor** then to **Medicine Room** and then to **Cashier** and so on. In a way, a patient is sent into a chain of departments which when done, sends the patient to further departments. This is where the Chain of Responsibility pattern comes into the picture.

### When to Use:

 - The pattern is applicable when there are multiple candidates to process the same request.
 - When you don’t want the client to choose the receiver as multiple objects can handle the request. Also, you want to decouple the client from receivers. The Client only needs to know the first element in the chain.

As in the example of the hospital, a patient first goes to the reception and then reception based upon a patient’s current status sends up to the next handler in the chain.

UML Diagram:

[](![uml diagram](https://github.com/filipeandrade6/go-design-patterns/blob/master/behavioural/chain-of-responsibility/img/Chain-of-Responsibility-Design-Pattern-1.jpg?raw=true)

Below is the corresponding mapping UML diagram with the practical example given below

[](![uml diagram](https://github.com/filipeandrade6/go-design-patterns/blob/master/behavioural/chain-of-responsibility/img/Chain-of-Responsibility-Design-Pattern-2.jpg?raw=true)

### Mapping

| UML | Example |
| - | - |
| handler | department.go |
| Concrete Handler 1 | account.go |
| Concrete Handler 2 | doctor.go |
| Concrete Handler 3 | medical.go |
| Concrete Handler 4 | cashier.go |
| Client | main.go |