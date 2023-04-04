"""
This script calculates the result of a^p mod n using the square and multiply algorithm.

"""


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


def main() -> None:
    """
        Main function to get user input and call square_multiply function.
        Prints the result of a^p mod n.
    """
    print("a^p mod n")
    a = int(input("Enter base value (a): "))
    pw = int(input("Enter Power(p): "))
    n = int(input("Enter Mod value(n): "))
    print(square_multiply(a, pw, n))


if __name__ == '__main__':
    main()
