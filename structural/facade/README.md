## Facade Design Pattern

[Refactoring Guru](https://refactoring.guru/design-patterns/facade)

Facade Pattern is classified as a structural design pattern. This design pattern is meant to hide the complexities of the underlying system and provide a simple interface to the client. It provides a unified interface to underlying many interfaces in the system so that from the client perspective it is easier to use. Basically it provides a higher level abstraction over a complicated system.

The term **Facade** itself means **the principal front of a building, that faces on to a street or open space**

Only the front face of the building is shown all the underlying complexity is hidden behind.

Let’s understand the Facade Design Pattern with a simple example. In this era of the digital wallet, when someone actually does a wallet debit/credit there are a lot of things that are happening in the background which the client may not be aware of. Below list illustrates some of the activities which happen during the credit/debit process

 - Check Account
 - Check Security Pin
 - Credit/Debit Balance
 - Make Ledger Entry
 - Send Notification

As can be noticed, there are a lot of things that happen for a single debit/credit operation. This is where the Facade pattern comes into picture. As a client one only needs to enter the Wallet Number, Security Pin, Amount and specify the type of operation. The rest of the things are taken care of in the background. Here we create a **WalletFacade** which provides a simple interface to the client and which takes care of dealing with all underlying operations.

![uml diagram with example](https://github.com/filipeandrade6/go-design-patterns/blob/master/structural/facade/img/Facade-Design-Pattern-Behind-Scenes.jpg?raw=true)

### Problem Statement:

 - In order to use the complex system, the client had to know the underlying details. Need to provide a simple interface to the client so that they can use a complex system without knowing any of its inner complex details.

### When to Use:

 - When you want to expose a complex system in a simplified way.

Like in the above example of credit/debit wallet they need to know only one interface and the rest of the things should be taken care of by that interface.

### UML Diagram:

![uml diagram with example](https://github.com/filipeandrade6/go-design-patterns/blob/master/structural/facade/img/Facade-Design-Pattern-1.jpg?raw=true)

Below is the corresponding mapping UML diagram with the practical example given below

![uml diagram with example](https://github.com/filipeandrade6/go-design-patterns/blob/master/structural/facade/img/Facade-Design-Pattern-2.jpg?raw=true)

### Mapping

The below table represents the mapping from the UML diagram actors to actual implementation actors in code.

| UML | Example |
| - | - |
| Wallet Facade | walletFacade.go |
| account | account.go |
| securityCode | securityCode.go |
| wallet | wallet.go |
| ledger | ledger.go |
| notification | notification.go |
| Client | main.go |