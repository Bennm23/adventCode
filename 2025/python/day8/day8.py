import time
from typing import List, Set, Tuple
from lib import read_to_input

Junction = Tuple[int, int, int]

class Connection:
    def __init__(self, from_junction: Junction, to_junction: Junction, distance: int):
        self.from_junction = from_junction
        self.to_junction = to_junction
        self.distance = distance
    def __str__(self):
        return f"Connection: {self.from_junction} <=> {self.to_junction}"
    def __eq__(self, other):
        return (self.from_junction == other.from_junction and self.to_junction == other.to_junction) or \
               (self.from_junction == other.to_junction and self.to_junction == other.from_junction)
    def __hash__(self):
        return hash((min(self.from_junction, self.to_junction), max(self.from_junction, self.to_junction)))


lines = read_to_input("day8/input")
junctions = []
for line in lines:
    split = line.strip().split(",")
    junctions.append((int(split[0]), int(split[1]), int(split[2])))

def sorted_connections() -> List[Connection]:
    connections: Set[Connection] = set()
    # Load all connections
    for i in range(len(junctions)):
        for j in range(i + 1, len(junctions)):
            junc1 = junctions[i]
            junc2 = junctions[j]
            distance = ((junc1[0] - junc2[0]) ** 2 + (junc1[1] - junc2[1]) ** 2 + (junc1[2] - junc2[2]) ** 2) ** 0.5
            connection = Connection(junc1, junc2, distance)
            connections.add(connection)
    
    connections = list(connections)
    connections.sort(key=lambda conn: conn.distance)
    return connections
    

ordered_connections = sorted_connections()

def expand_connection(conn: Connection, all_connections: Set[Connection], visited: Set[Junction], junctions: Set[Junction]):
    """
    Get a list of all connections that can be made from this connection's endpoints
    """
    visited.add(conn.from_junction)
    visited.add(conn.to_junction)
    junctions.add(conn.from_junction)
    junctions.add(conn.to_junction)

    for endpoint in [conn.from_junction, conn.to_junction]:
        for other in all_connections:
            if other.from_junction == endpoint and other.to_junction not in visited:
                expand_connection(other, all_connections, visited, junctions)

            elif other.to_junction == endpoint and other.from_junction not in visited:
                expand_connection(other, all_connections, visited, junctions)


def p1() -> int:
    connections: Set[Connection] = set()
    for i in range(0, 1000):
        connections.add(ordered_connections[i])

    conn_lengths = []
    expanded_connections: Set[Connection] = set()
    for conn in connections:
        junctions_in_conn: Set[Junction] = set()
        expand_connection(conn, connections, expanded_connections, junctions_in_conn)
        length = len(junctions_in_conn)
        conn_lengths.append(length)


    conn_lengths.sort(reverse=True)
    p1 = 1
    for i in range(0, min(3, len(conn_lengths))):
        print("Top Conn Length:", conn_lengths[i])
        p1 *= conn_lengths[i]

    return p1

def p2(stop_index: int = 4941) -> int:
    # Manually searched for part 2 cutoff, connections get to 1000 at 4941
    connections: Set[Connection] = set()
    print("Using Connections :", stop_index)
    for i in range(0, stop_index):
        connections.add(ordered_connections[i])
    print("Last Conn Used :", ordered_connections[stop_index - 1])
    p2 = ordered_connections[stop_index-1].to_junction[0] * ordered_connections[stop_index-1].from_junction[0]

    return p2

start = time.time()
print("------- Part 1 -------")
print("Part 1:", p1()) # 47040
print("------- Part 2 -------")
print("Part 2:", p2()) # 4884971896

end = time.time()
print("Time taken:", int((end - start) * 1000), "ms")