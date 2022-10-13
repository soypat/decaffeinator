def collatz(a:int) -> tuple[int, int]:
    down = a//2
    up = 3*a+1
    return down, up

v = 60
down, up = collatz(v)
print(f"empezando en {v} hay que saber subir {up}, y bajar {down}")