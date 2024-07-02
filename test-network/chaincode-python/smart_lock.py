import sys

def lock():
    print("Locking...")

def unlock():
    print("Unlocking...")

if __name__ == "__main__":
    if sys.argv[1] == "lock":
        lock()
    elif sys.argv[1] == "unlock":
        unlock()
