## Singleton Design Patern

[Refactoring Guru](https://refactoring.guru/design-patterns/singleton)

Singleton Design Pattern is a creational design pattern and also one of the most commonly used design pattern. This pattern is used when only a single instance of the struct should exist. This single instance is called a singleton object. Some of the cases where the singleton object is applicable:

 1. **DB instance** – we only want to create only one instance of DB object and that instance will be used throughout the application.
 2. **Logger instance** – again only one instance of the logger should be created and it should be used throughout the application.

The singleton instance is created when the struct is first initialized.  Usually, there is getInstance() method defined on the struct for which only one instance needs to be created. Once created then the same singleton instance is returned every time by the **getInstance()**.

In GO we have goroutines. Hence the singleton struct should return the same instance whenever multiple goroutines are trying to access that instance. It is very easy to get a singleton design pattern wrong. The below code illustrates the right way to create a singleton object.

```go
var lock = &sync.Mutex{}

type single struct {
}

var singleInstance *single

func getInstance() *single {
    if singleInstance == nil {
        lock.Lock()
        defer lock.Unlock()
        if singleInstance == nil {
            fmt.Println("Creting Single Instance Now")
            singleInstance = &single{}
        } else {
            fmt.Println("Single Instance already created-1")
        }
    } else {
        fmt.Println("Single Instance already created-2")
    }
    return singleInstance
}
```

Above code ensures that only one instance of the single struct is created. Some point worth noting.

 1. There is a check at the start for nil singleInstance. This is to prevent the expensive lock operations every time getinstance() method is called. If this check fails then it means that singleInstance is already created
 2. The singleInstance is created inside the lock.
 3. There is another check for nil singleIinstance after the lock is acquired. This is to make sure that if more than one goroutine bypass the first check then only one goroutine is able to create the singleton instance otherwise each of the goroutine will create its own instance of the single struct.
