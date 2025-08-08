
import json

# Config format:
# {
#   "M": 5,             int
#   "L": 10,            int
#   "N": 100,           int
#   "Rc": 2.5,          float
#   "Size": "",         empty string
#   "method": "cell",   (cell|brute)    method used for finding neighbors
#   "loop": true        boolean -> periodic contour
# }



if __name__ == "__main__":
    args = []
    toReturn = ''
    config = json.load(open("config/config.json"))
    args = config.values()
    for arg in args:
        if arg != '':   #skip size value for now
            toReturn += str(arg) + ' '
    print(toReturn.strip())