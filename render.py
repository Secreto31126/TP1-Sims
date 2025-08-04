import matplotlib.pyplot as plt
import csv

def read_data(filename):
    x = []
    y = []
    with open(filename, 'r') as file:
        print('inside file ;D')
        reader = csv.DictReader(file)
        for row in reader:
            x.append(float(row['x']))
            y.append(float(row['y']))
    return x, y

def plot_data(x, y):
    plt.figure(figsize=(8, 5))
    plt.plot(x, y, marker='o', linestyle='-', color='blue')
    plt.title('Plot from CSV File')
    plt.xlabel('X values')
    plt.ylabel('Y values')
    plt.grid(True)
    plt.tight_layout()
    plt.show()


x_vals, y_vals = read_data("./test.csv")
plot_data(x_vals, y_vals)