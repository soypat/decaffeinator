def sqrt(x:int) -> str:
    if x < 0:
        return sqrt(-x) + "i"
    elif x == 0:
        return "0"
    return str(x**0.5)

print(sqrt(2), sqrt(-4))