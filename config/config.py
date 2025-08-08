
import json

# Config format:
# {
#   "M": 5,             int
#   "L": 10,            int
#   "N": 100,           int
#   "Rc": 2.5,          float
#   "Size": "",         empty string
#   "method": "cell",   (cell|brute)    method used for finding neighbors
#   "loop": "loop"        string -> periodic contour, loop is true, empty if false
# }



if __name__ == "__main__":
    args = []
    config = json.load(open("config/config.json"))
    for key, value in config.items():
        if key == 'Size' and value == '':
            continue  # skip empty Size as per your original logic
        else:
            args.append(str(value))
    print(' '.join(args))