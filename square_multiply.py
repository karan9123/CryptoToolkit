"""
This script calculates the result of a^p mod n using the square and multiply algorithm.
"""
import argparse


def square_multiply(a: int, pw: int, n: int) -> int:
    """
        Calculates a^p mod n using the square and multiply algorithm.
        Args:
            a: The base value (a) as an integer.
            pw: The power value (p) as an integer.
            n: The mod value (n) as an integer.

        Returns:
            The result of a^p mod n as an integer.
    """
    k = 1
    pw = bin(pw)
    for i in range(2, len(pw)):
        k = (k * k) % n
        if pw[i] == "1":
            k = (k * a) % n
    return k


if __name__ == '__main__':
    parser = argparse.ArgumentParser()
    parser.add_argument("base", help="Base Value", type=int)
    parser.add_argument("pow", help="Power", type=int)
    parser.add_argument("mod", help="Mod", type=int)
    args = parser.parse_args()
    print(square_multiply(args.base, args.pow, args.mod))
