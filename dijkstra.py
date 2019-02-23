class Node:
    def __init__(self, ID, connections):
        self.id = ID   # String
        self.cons = connections      # Array of connections

    def __repr__(self):
        return self.id

class Con:
    def __init__(self, other, weight):
        self.other = other
        self.weight = weight

def getNode(nodeID, graph):
    for node in graph:
        if node.id == nodeID:
            return node
    return None

def getMinDist(distances):
    lowest = [None, -1]
    for node, dist in distances.items():
        if dist < lowest[1] or lowest[1] == -1:
            lowest = [node, dist]

    return lowest

def dijkstra(graph):
    q = []
    prev = {}
    dist = {}
    for node in graph:
        q.append(node)
        prev[node] = -1
        dist[node] = -1

    dist[graph[0]] = 0

    while len(q) > 0:
        u = q.pop(0)

        for con in u.cons:
            alt = dist[u] + con.weight
            node = getNode(con.other, graph)
            if alt < dist[node] or dist[node] == -1:
                dist[node] = alt
                prev[node] = u

    return dist, prev


# Graph:
# A -> (D, 3), (B, 7)
# B -> (A, 7), (D, 2), (E, 6), (C, 3)
# C -> (B, 3), (E, 1), (D, 4)
# D -> (A, 3), (B, 2), (C, 4), (E, 7)
# E -> (D, 7), (B, 6), (C, 1)

if __name__ == "__main__":
    graph = [Node("A", [Con("D", 3), Con("B", 7)]),
             Node("B", [Con("A", 7), Con("D", 2), Con("E", 6), Con("C", 3)]),
             Node("C", [Con("B", 3), Con("E", 1), Con("D", 4)]),
             Node("D", [Con("A", 3), Con("B", 2), Con("C", 4), Con("E", 7)]),
             Node("E", [Con("D", 7), Con("B", 6), Con("C", 1)])]

    print(dijkstra(graph))
