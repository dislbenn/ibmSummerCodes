import ctypes
import os

def main():
    # Path to shared object file
    path = os.path.join(os.path.abspath("."), "int_input.so")
    # Load shared object file
    lib = ctypes.CDLL(path)
    # Call shared object method
    user_input = int(input("Enter an integer: "))
    print("Answer is: ")
    print("---------------------")
    lib.int_input(user_input)
    print("---------------------\n\n")

if __name__ == "__main__":
    main()