## Template Method Design Pattern

[Refactoring Guru](https://refactoring.guru/design-patterns/template-method)

Template Method Design Pattern is a behavioral design pattern that lets you define a template or algorithm for a particular operation.  Let’s understand the template design pattern with an example.

Consider the example of One Time Password or OTP. There are different types of OTP that can be triggered for eg. OTP can be **SMS** OTP or **EMAIL** OTP.  But irrespective it is an **SMS** OTP or **EMAIL** OTP,  the entire steps of the OTP process are the same.  The steps are

 - Generate a random n digit number.
 - Save this number in the cache for later verification.
 - Prepare the content
 - Send the notification
 - Publish the metrics

Even in the future let’s say a push notification **OTP** is introduced but still it will go through the above steps.

In such scenarios when the steps of a particular operation are the same but the steps of the operations can be implemented in a different way by different implementors , then that becomes a candidate for the Template Method Design Pattern. We define a template or algorithm which comprises of a fixed number of methods. The implementer of the operation overrides the methods of the template.

Now check out the below code example.

 - **iOtp** represents an interface that defines the set of methods that any otp type should implement
 - **sms** and **email** are the implementors of **iOtp** interface
 - **otp** is the struct which defines the template method **genAndSendOTP()**. **otp** embeds **iOtp** interface.

Important: The combination of iOtp interface and otp struct provides the capabilities of Abstract Class in go. For reference see [Abstract Class in GO: Complete Guide](https://golangbyexample.com/go-abstract-class/)