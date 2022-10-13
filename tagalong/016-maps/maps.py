ages = {
    "Sarah":        32,
    "Billy":        12,
    "Jeremiah":     99,
    "John Baptist": 47,
}
print(ages["Sarah"])
billyAge, billyPresent = 0, False
if "Billy" in ages:
    billyAge, billyPresent = ages["Billy"], True

print(billyAge, billyPresent)

x13Age, x13Present = 0, False
if "x13" in ages:
    x13Age, x13Present = ages["x13"], True

print(x13Age, x13Present)

ages["Faustus"] = 66
print(ages)