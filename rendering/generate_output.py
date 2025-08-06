import random
import math

# Configuration
PARTICLE_COUNT = 1000
MAX_RADIUS = 0.5
WORLD_SIZE = 20.0
MAX_NEIGHBORS = 8
NEIGHBOR_RADIUS = 2.5  # Distance threshold for neighbors

def generate_particles():
    """Generate random particle data"""
    particles = []
    for i in range(PARTICLE_COUNT):
        radius = random.uniform(0.1, MAX_RADIUS)
        x = random.uniform(0, WORLD_SIZE)
        y = random.uniform(0, WORLD_SIZE)
        particles.append({
            'id': i,
            'x': x,
            'y': y,
            'radius': radius
        })
    return particles

def find_neighbors(particles):
    """Find nearby particles within NEIGHBOR_RADIUS"""
    for i, p1 in enumerate(particles):
        neighbors = []
        for j, p2 in enumerate(particles):
            if i == j:
                continue  # Skip self
            distance = math.sqrt((p1['x']-p2['x'])**2 + (p1['y']-p2['y'])**2)
            if distance < NEIGHBOR_RADIUS:
                neighbors.append(j)
        # Limit to random subset to make more realistic
        if len(neighbors) > MAX_NEIGHBORS:
            neighbors = random.sample(neighbors, MAX_NEIGHBORS)
        p1['neighbors'] = neighbors
    return particles

def write_output_file(particles, filename="output.txt"):
    """Write data in the specified format"""
    with open(filename, 'w') as f:
        # Write particle properties
        for p in particles:
            f.write(f"{p['radius']} {p['x']} {p['y']}\n")
        
        # Write neighbor lists
        for p in particles:
            neighbors = ' '.join(map(str, p['neighbors']))
            f.write(f"{p['id']} {neighbors}\n")

if __name__ == "__main__":
    print("Generating test data...")
    particles = generate_particles()
    particles = find_neighbors(particles)
    write_output_file(particles)
    print(f"Generated {PARTICLE_COUNT} particles in output.txt")
    print(f"Example first particle: {particles[0]}")