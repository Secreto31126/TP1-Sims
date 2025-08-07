import matplotlib.pyplot as plt
from classes import Particle

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

def plot_particles(particles, focused_neighbors, focused_id=0):
    """Visualizes particles with optional focus on one particle"""
    plt.figure(figsize=(8, 8))
    
    # Plot all particles
    for particle in particles:
        if(particle.id == focused_id):
            particle_color = 'red'
        elif(particle.id in focused_neighbors):
            particle_color = 'blue'
        else:
            particle_color = 'grey'

        circle = plt.Circle((particle.x, particle.y), particle.radius, 
                          fill=False, color=particle_color)
        plt.gca().add_patch(circle)
        
    plt.title('Particle Playground')
    plt.xlabel('X position')
    plt.ylabel('Y position')
    plt.axis('equal')
    plt.grid(True)
    plt.tight_layout()
    plt.show()

# Example usage:
if __name__ == "__main__":
    # Read all data from single file in alternating format
    path = input('Enter the path of the file relative to the project root: ')
    print(path)
    particles, neighbors = read_particle_data(path)
    print(f"Read {len(particles)} particles from {path}")
    # Visualize
    focused_id = int(input('Enter the ID of the focused particle: '))
    plot_particles(particles,neighbors[focused_id], focused_id)