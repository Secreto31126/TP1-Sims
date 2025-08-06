import matplotlib.pyplot as plt
from classes import Particle

def read_static_data(filename):
    """Reads particle properties (radius, x, y)"""
    particles = []
    with open(filename, 'r') as file:
        for line in file:
            # Skip empty lines
            if not line.strip():
                continue
            # Parse radius, x, y
            radius, x, y = map(float, line.strip().split())
            particles.append(Particle(x, y, radius))
    return particles

def read_neighbors(filename):
    """Reads neighbor lists (id, neighbor1, neighbor2, ...)"""
    neighbors = {}
    with open(filename, 'r') as file:
        for line in file:
            # Skip empty lines
            if not line.strip():
                continue
            # Parse id and neighbors
            parts = list(map(int, line.strip().split()))
            particle_id = parts[0]
            neighbor_ids = parts[1:]
            neighbors[particle_id] = neighbor_ids
    return neighbors

def read_dynamic_data(filename, particles):
    """Updates particles with neighbor information"""
    neighbors = read_neighbors(filename)
    for particle_id, neighbor_ids in neighbors.items():
        if 0 <= particle_id < len(particles):
            particles[particle_id].id = particle_id
            particles[particle_id].neighbors = neighbor_ids
    return particles

def plot_particles(particles, focused_id=None):
    """Visualizes particles with optional focus on one particle"""
    plt.figure(figsize=(10, 10))
    
    # Plot all particles
    for particle in particles:
        circle = plt.Circle((particle.x, particle.y), particle.radius, 
                          fill=False, color='blue')
        plt.gca().add_patch(circle)
        plt.text(particle.x, particle.y, str(particle.id), 
                ha='center', va='center')
    
    # Highlight focused particle if specified
    if focused_id is not None and 0 <= focused_id < len(particles):
        focused = particles[focused_id]
        circle = plt.Circle((focused.x, focused.y), focused.radius, 
                          fill=True, color='red', alpha=0.3)
        plt.gca().add_patch(circle)
        
        # Draw lines to neighbors
        if hasattr(focused, 'neighbors'):
            for neighbor_id in focused.neighbors:
                if 0 <= neighbor_id < len(particles):
                    neighbor = particles[neighbor_id]
                    plt.plot([focused.x, neighbor.x], 
                            [focused.y, neighbor.y], 
                            'r--', alpha=0.3)
    
    plt.title('Particle System Visualization')
    plt.xlabel('X position')
    plt.ylabel('Y position')
    plt.axis('equal')
    plt.grid(True)
    plt.tight_layout()
    plt.show()

# Example usage:
if __name__ == "__main__":
    # Read particle properties (static data)
    particles = read_static_data('particles.txt')
    
    # Read and assign neighbor information (dynamic data)
    particles = read_dynamic_data('neighbors.txt', particles)
    
    # Visualize
    focused_id = int(input('Enter the ID of the focused particle: '))
    plot_particles(particles, focused_id)