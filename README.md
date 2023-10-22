<!--toc:start-->
  - [Difference between authentication and authorization](#difference-between-authentication-and-authorization)
  - [HTTP Basic Authentication](#http-basic-authentication)
  - [Storing password](#storing-password)
  - [Bearer tokens and HMAC](#bearer-tokens-and-hmac)
  - [Discussion](#discussion)
  - [Not yet discussed](#not-yet-discussed)
- [Go routines](#go-routines)
<!--toc:end-->

### Difference between authentication and authorization
* Authentication
  * Determine who you are
  * Verifies that no-one is impersonating you
  * Three ways to authenticate
    * Who you are (biometrics)
    * What you have (eg, atm card, key, phone)
    * What you know (username, password)
  * Two-factor authentication
* Authorization
    * Says what you are allowed to do
    * The name of http header used for authorization
### HTTP Basic Authentication
* Part of specification of HTTP protocol
    * Send username/password with every request
    * Use authorization header and keyword basic
        * Put username & password together
        * Convert them to base64
            * put generic binary data into printable form
            * base64 is reversible
              * Never use with http; only https
        * use basic authentication
### Storing password
* Never store passwords
* Instead, store one-way encryption "hash" values of password
* For added security
  * Hash on the client
  * Hash That again on the server
* Hashing algorithms
  * [BCRYPT](https://godoc.org/golang.org/x/crypto/bcrypt) - current choice
### Bearer tokens and HMAC
* Bearer tokens
  * Added to HTTP spec with Oauth2
  * Uses authorization header and keyword 'Bearer'
* To prevent faked bearer tokens, use cryptographic "signing"
  * Cryptographic signing is a way to prove that the value was created by a certain person
  * [HMAC](https://godoc.org/crypto/hmac)
### Discussion
* Cryptography
  * <u>Large field</u>
* Hashing
  * MD5 - ***don't use***
  * SHA
  * BCrypt
  * SCrypt
* Signing
  * Symmetric key (*same key to sign (encrypt) / verify (decrypt)*)
    * HMAC
  * Asymmetric key
    * RSA
    * ECDSA - **better than RSA**; faster; smaller keys
    * Private key to sign (encrypt) / public key to verify (decrypt)
  * JWT

### Not yet discussed

* Encryption
  * Symmetric key
    * AES
  * Asymmetric key
    * RSA

## Go routines
1. Lightweight: Goroutines are lightweight compared to traditional operating system threads. The Go runtime manages goroutines, and they are more efficient in terms of memory and CPU usage.

2. Concurrent Execution: Go scheduler automatically manages the execution of goroutines, distributing them across multiple CPU cores. The scheduler uses a technique called "work-stealing" to efficiently schedule go routines.

3. Channel Communication: Goroutines communicate with each other using channels. Channels are used to send and receive data between goroutines in a safe and synchronized way.

4. Blocking Calls: When a goroutine makes a blocking call (e.g., I/O operation or waiting for channel data), the Go scheduler switches to another ready goroutine, allowing the program to utilize other available go routines efficiently.

5. Synchronization: Go provides synchronization primitives like sync.Mutex and sync.WaitGroup to synchronize access to shared resources and coordinate the execution of goroutines.

6. Graceful Concurrency: Go encourages the use of channels and synchronization primitives to avoid data races and ensure that goroutines coordinate their actions correctly. This helps in writing concurrent code that is less prone to race conditions.

7. Goroutine Stacks: Each goroutine has its own small initial stack, which grows and shrinks as needed. The size of the initial stack is small, so Go can create many goroutines without consuming a lot of memory.

8. Garbage Collection: Go has a garbage collector that manages memory automatically, so you don't need to worry about memory management when using goroutines.

9. Error Handling: It's essential to handle errors properly when using goroutines. Unhandled errors in goroutines can lead to unexpected behavior or panics in the program.

10. Overall, Go routines are a powerful and elegant way to implement concurrent programming in Go. They make it easy to write scalable and efficient concurrent applications with minimal complexity. However, it's essential to understand the implications of concurrent programming and use synchronization mechanisms to ensure the correctness of your code.
