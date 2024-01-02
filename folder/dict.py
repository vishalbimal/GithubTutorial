from collections import OrderedDict

dict = {"ravi": "1", "rajnish": "9", "sanjeev": "15", "ravi": "2", "suraj": "32"}
dict1 = OrderedDict(sorted(dict.items()))

print(dict1)
