import functools


# @functools.lru_cache(None)
def fib(n):
    if n < 2:
        return n
    return fib(n-1) + fib(n-2)


i = 0
while True:
    i += 1
    print(f"{i}: {fib(i)}")
