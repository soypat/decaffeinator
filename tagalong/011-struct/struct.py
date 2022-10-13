class Vertex:
    def __init__(self, x:int=0, y:int=0):
        self.x= int(x)
        self.y= int(y)

    def __str__(self):
        return f"Vertex({self.x}, {self.y})"

    def __setattr__(self, __name: str, __value: any) -> None:
        super().__setattr__(__name, int(__value))

v0 = Vertex()
print(v0)
a = 0

v1 = Vertex(1, 2)
print(v1.x)

# This is a bug with no __setattr__ method.
# I didn't look for a spurious case, this was taken as is from
# https://go.dev/tour/moretypes/4
v1.x = 1e9
print(v1)
