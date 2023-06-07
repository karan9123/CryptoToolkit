# 1. Square and multiply
### Run with parameters: base power mod

`` python3 square_multiply.py 3 4 11``

#### &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;  3<sup>4</sup> mod 11 = 4

 The square and multiply algorithm is a method to compute modular exponentiation, i.e., calculating __a^p mod n__ for given integers __a, p__, and __n__. It is a more efficient algorithm for performing modular exponentiation compared to the naive method of computing a^p first and then taking the modulus with n.

The square and multiply algorithm reduces the number of multiplications needed to compute **a^p mod n**. It does so by breaking down the exponent p into its binary representation and performing repeated squaring and modular multiplications based on the bits of the binary representation.
