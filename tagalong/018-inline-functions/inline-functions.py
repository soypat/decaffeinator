def SuperRandom(normalRandom:lambda:int) -> int:
    """SuperRandom returns a super random number using a
    not so random source function "normalRandom".
    (this is not actually more random!)"""
    superrand = 12345678
    for _ in range(3):
        rand = normalRandom()
        superrand = superrand*7 + rand*31
    return superrand

# Python's inline function creation uses
# the lambda keyword, though its limited to
# one-liners so we can't really use it here.
# We do with a simple def.
a = 287117
def notSoRand()->int:
    global a
    a = a * 7
    return a

superrand1 = SuperRandom(notSoRand)
superrand2 = SuperRandom(notSoRand)
superrand3 = SuperRandom(notSoRand)

print(superrand1, superrand2, superrand3)




