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
## Storing password
* Never store passwords
* Instead, store one-way encryption "hash" values of password
* For added security
  * Hash on the client
  * Hash That again on the server
* Hashing algorithms
  * Bcrypt
  * Scrypt - new choice
## Bearer tokens and Hmac
* Bearer tokens
  * Added to http spec with Oauth2
  * Uses authorization header and keyword 'Bearer'
* To prevent faked bearer tokens, use cryptographic "signing"
  * Cryptographic signing is a way the prove that the value was created by a certain person
  * [Hmac](https://godoc.org/crypto/hmac)

1. Click **Source** on the left side.
2. Click the README.md link from the list of files.
3. Click the **Edit** button.
4. Delete the following text: *Delete this line to make a change to the README from Bitbucket.*
5. After making your change, click **Commit** and then **Commit** again in the dialog. The commit page will open and you’ll see the change you just made.
6. Go back to the **Source** page.

---

## Create a file

Next, you’ll add a new file to this repository.

1. Click the **New file** button at the top of the **Source** page.
2. Give the file a filename of **contributors.txt**.
3. Enter your name in the empty file space.
4. Click **Commit** and then **Commit** again in the dialog.
5. Go back to the **Source** page.

Before you move on, go ahead and explore the repository. You've already seen the **Source** page, but check out the **Commits**, **Branches**, and **Settings** pages.

---

## Clone a repository

Use these steps to clone from SourceTree, our client for using the repository command-line free. Cloning allows you to work on your files locally. If you don't yet have SourceTree, [download and install first](https://www.sourcetreeapp.com/). If you prefer to clone from the command line, see [Clone a repository](https://confluence.atlassian.com/x/4whODQ).

1. You’ll see the clone button under the **Source** heading. Click that button.
2. Now click **Check out in SourceTree**. You may need to create a SourceTree account or log in.
3. When you see the **Clone New** dialog in SourceTree, update the destination path and name if you’d like to and then click **Clone**.
4. Open the directory you just created to see your repository’s files.

Now that you're more familiar with your Bitbucket repository, go ahead and add a new file locally. You can [push your change back to Bitbucket with SourceTree](https://confluence.atlassian.com/x/iqyBMg), or you can [add, commit,](https://confluence.atlassian.com/x/8QhODQ) and [push from the command line](https://confluence.atlassian.com/x/NQ0zDQ).