# pw-extract

## Use case

A password is saved to a password manager. It's unique and 30 characters long. It looks something like:

`#6tl%9dKYvTHOt43^gLD8G59qeDrMN`

Your bank, in the interests of "security", does not allow you to paste the entire string from your password manager.

Instead it requests the 5th, 16th and 24th character from the password.

```shell script
$ ./pw-extract 
Enter Password: Enter character numbers, space separated: 5 16 24
5: %
16: 3
24: 9
```
