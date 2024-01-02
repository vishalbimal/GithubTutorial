# myDict= {'yash':2,'ravi':10,'rajnish':9,'sanjeev':15,'suraj':32}
# myKeys=list(myDict.keys())
# myKeys.sort()

# sorted_dict={i:myDict[i] for i in myKeys}

# print(sorted_dict)

# DISPLAYING THE KEYS IN SORTED ORDER

def dictionary(): #function calling
    #declare hash funstion
    key_value={}

   #initializing value
    key_value[2]=56 
    key_value[1]=5 
    key_value[3]=9
    key_value[4]=20
    key_value[5]=33

    print("Task 1 :- \n")

    print("key_value",key_value)
  #iterkeys() returns an iteration over the dictionary's keys.
    for i in sorted(key_value.keys()):
        print(i, end="")

def main():
    dictionary()

if __name__ == "__main__":
    main()


