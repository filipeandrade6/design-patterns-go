## Proxy Design Pattern

[Refactoring Guru](https://refactoring.guru/design-patterns/proxy)

Proxy Design Pattern is a structural design pattern. This pattern suggests providing an extra layer of indirection for controlled and intelligent access to the main object.

In this pattern, a new proxy class is created that implements the same interface as the main object. This lets you execute some behavior before the actual logic of the main object. Let’s understand it more with an example

 1. A debit card is a proxy for your bank account. It follows the same interface as the bank account, and it is easier to use.
 1. A web server such as Nginx can act as a proxy for your application server. It provides
  - Controlled access to your application server – For example, it can do rate limiting
  - Additional behavior – For example, it can do some caching

Let’ss see the UML Diagram

### UML Diagram:

In below UML diagram

 - **Subject**: it represents the interface which proxy and the realSubject should follow
 - **Proxy**: The proxy class embeds the realSubject and passes on the request to the realSubject after it has finished its processing
 - **RealSubject**: This is the class that holds the main business logic and which is kept behind a proxy
 - **Client**: Can interact with both proxy and realSubject as they both follow the same interface.

![uml diagram with example](https://github.com/filipeandrade6/go-design-patterns/blob/master/structural/proxy/img/Proxy-Design-Pattern-1.jpg?raw=true)

Below is the corresponding mapping UML diagram with the practical example of nginx and application server which was described above.

![uml diagram with example](https://github.com/filipeandrade6/go-design-patterns/blob/master/structural/proxy/img/Proxy-Design-Pattern-2.jpg?raw=true)

### Mapping

The below table represents the mapping from the UML diagram actors to actual implementation actors in code

| UML | Example |
| - | - |
| subject | server.go |
| proxy | nginx.go |
| realSubject | application.go |
| client | main.go |