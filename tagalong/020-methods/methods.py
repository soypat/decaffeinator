class Rectangle:
    def __init__(self, width: float, height: float) -> None:
        self.width = width
        self.height = height

    def area(self) -> float:
        return self.width * self.height

    def perim(self) -> float:
        return 2 * (self.width + self.height)

    def scale(self, factor: float) -> None:
        self.width *= factor
        self.height *= factor


if __name__ == "__main__":
    r = Rectangle(12.7, 10.0)
    print("area [mm²]:", r.area())
    print("perimeter [mm]:", r.perim())

    r.scale(1 / 25.4)
    print("area [in²]:", r.area())
    print("perimeter [in]:", r.perim())
