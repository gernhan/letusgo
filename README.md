---
## Difference between authentication and authorization
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
---
## Http Basic Authentication
* Basic Authentication part of specification of http
    * Send username/password with every request
    * Use authorization header and keyword basic
        * Put username & password together
        * Convert them to base64
            * put generic binary data into printable form
            * base64 is reversible
              * Never use with http; only https
        * use basic authentication
---
## Storing password
* Never store passwords
* Instead, store one-way encryption "hash" values of password
* For added security
  * Hash on the client
  * Hash That again on the server
* Hashing algorithms
  * [Bcrypt](https://godoc.org/golang.org/x/crypto/bcrypt) - current choice
---
## Bearer tokens and Hmac
* Bearer tokens
  * Added to http spec with Oauth2
  * Uses authorization header and keyword 'Bearer'
* To prevent faked bearer tokens, use cryptographic "signing"
  * Cryptographic signing is a way the prove that the value was created by a certain person
  * [Hmac](https://godoc.org/crypto/hmac)
---
## Discussion
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
---
## Not yet discussed
* Encryption
  * Symmetric key
    * AES
  * Asymmetric key
    * RSA