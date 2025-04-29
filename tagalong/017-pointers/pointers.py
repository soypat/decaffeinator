# The simplest case where one can observe
# pointer action in Python is with lists.
def doSomething(thelist):
    thelist.append(3)

l = [0, 1]

doSomething(l)
print(l) # prints [0, 1, 3]
# This happens because lists in Python
# have a pointer which is passed to functions
# allowing functions to modify data.