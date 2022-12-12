# Passwd-cracker
A  simple Go program that can be used to test the strength of user-defined passwords:

This program uses a brute-force attack to guess the password by trying all possible combinations of characters in the password. This type of attack can be effective for short passwords, but it becomes increasingly impractical as the password length increases.

For a more effective password cracking tool, you could use a dictionary attack, which involves trying common words and phrases as passwords. You could also use a precomputed hash table, which allows you to quickly look up the hash values of common passwords and compare them to the hash of the password you are trying to crack.

Installation: 
```
git clone https://github.com/whoami-anoint/Passwd-cracker
cd Passwd-cracker
go run passwd-cracker.go 
```





```Note that this program is intended for educational purposes only, and should not be used to actually crack passwords without the owner's permission. Breaking into someone else's account without their permission is illegal and can result in severe penalties.```
