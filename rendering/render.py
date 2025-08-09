import matplotlib.pyplot as plt
from classes import Particle
import numpy as np
import sys
path = sys.argv[1]
M = int(sys.argv[2])
L = int(sys.argv[3])
def read_particle_data(filename):
    """Reads particle data in alternating format (properties then neighbors)"""
    particles = []
    neighbors = {}
    
    with open(filename, 'r') as file:

        #skips empty lines
        lines = [line.strip() if line.strip() != '' else '' for line in file]
        
        # Process pairs of lines (particle properties + neighbors)
        for i in range(0, len(lines), 2):
            if i+1 >= len(lines):
                break  # In case file has odd number of lines
            
            # Parse particle properties
            prop_line = lines[i]
            radius, x, y = map(float, prop_line.split())
            particle_id = i // 2  # ID based on position in file
            particles.append(Particle(x, y, radius, id=particle_id))
            
            # Parse neighbors
            neighbor_line = lines[i+1]
            neighbor_ids = list(map(int, neighbor_line.split())) if neighbor_line else []
            neighbors[particle_id] = neighbor_ids

    
    return particles, neighbors

def plot_particles(particles, focused_neighbors, focused_id=0, square_size=5, L=10):
    """Redrawable particle visualization"""
    plt.clf()  # Clear previous plot instead of making new windows
    
    # Set up figure if it doesn't exist
    if not plt.get_fignums():
        plt.figure(figsize=(13, 13))
        plt.ion()  # Turn on interactive mode
        plt.show()
    ax = plt.gca()
    grid_positions = np.linspace(0, L, M + 1)
    for x in grid_positions:
        ax.plot([x, x], [0, L], color='grey', linestyle='-', linewidth=1)
        ax.plot([0, L], [x, x], color='grey', linestyle='-', linewidth=1)
    
    # Plot all particles
    for particle in particles:
        if particle.id == focused_id:
            color = 'red' 
            Rc = 1+particle.radius
            ax.add_patch(plt.Circle((particle.x, particle.y), 
                              Rc, fill=False, color='green', linewidth=1, linestyle='--'))
        elif particle.id in focused_neighbors:
            color = 'blue'
        else:
            color ='grey'
        circle = plt.Circle((particle.x, particle.y), particle.radius,
                          fill=False, color=color, linewidth=1)
        ax.add_patch(circle)
    
    plt.title(f'Focused: {focused_id} | Neighbors: {len(focused_neighbors)}')
    plt.xlabel('X position')
    plt.ylabel('Y position')
    plt.axis('equal')
    # plt.grid(True)
    # plt.tight_layout()
    plt.draw()
    plt.pause(0.1)  # Allow time for GUI update
    plt.show(block=True)

if __name__ == "__main__":
    plt.ion()  # Enable interactive mode early
    
    # Data loading (unchanged)
    

    particles, neighbors = read_particle_data(path)
    print(f"Read {len(particles)} particles")
    
    try:
        focused_id = int(sys.argv[4])
        plot_particles(particles, neighbors.get(focused_id, []), focused_id, np.floor(L/M), L)
    except ValueError:
        print("Please enter a valid number or 'exit'")
    # plt.close()