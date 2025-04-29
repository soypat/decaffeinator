import math
from typing import Protocol, Tuple

class Geometry(Protocol):
    def area(self) -> float: ...
    def perim(self) -> float: ...

class Rectangle:
    def __init__(self, width: float, height: float) -> None:
        self.width = width
        self.height = height
    def area(self) -> float:
        return self.width * self.height
    def perim(self) -> float:
        return 2 * (self.width + self.height)

class Circle:
    def __init__(self, radius: float) -> None:
        self.radius = radius
    def area(self) -> float:
        return math.pi * self.radius ** 2
    def perim(self) -> float:
        return 2 * math.pi * self.radius

def shape_efficiency(g: Geometry) -> float:
    return g.area() / g.perim()

def detect_circle(g: Geometry) -> Tuple[float, bool]:
    if isinstance(g, Circle):
        return g.radius, True
    return 0.0, False

if __name__ == "__main__":
    rect = Rectangle(1, 4)
    square = Rectangle(4, 4)
    c = Circle(4)

    print("1x4 rectangle efficiency:", shape_efficiency(rect))
    print("square efficiency:", shape_efficiency(square))
    print("circle efficiency:", shape_efficiency(c))

    print(detect_circle(rect))
    print(detect_circle(c))
