import datetime
import time

today = datetime.date.today()

# This control structure is poorly supported
# in VSCode as of October 2022, Python 3.10. Quite hard to write.
match today.weekday():
    case 5:
        print("Today.")
    case 4:
        print("Tomorrow.")
    case 5:
        print("In two days.")
    case _:
        print("Too far away.")