# Static declaration of list (mutable)
a = list[str](["",""])
a[0] = "Hello"
a[1] = "World"
print(a[0], a[1])
print(a)

# Dynamic declaration of tuple (non-mutable)
primes = (2, 3, 5, 7, 11, 13)
print(primes)
