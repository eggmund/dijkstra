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


def dijkstra(graph):
    priority = [(graph[0], 0)]   # Priority queue holding (node, distance) tuples
    priority = priority + [(node, -1) for node in graph[1:]]
    
    while len(priority) > 0:
        curr = priority.pop()
        for con in curr.cons:
            other = getNode(con.other, priority)
            if other == None:
                raise LookupError("Couldn't find conn node.")

        


# Graph:
# A -> (D, 3), (B, 7)
# B -> (A, 7), (D, 2), (E, 6), (C, 3)
# C -> (B, 3), (E, 1), (D, 4)
# D -> (A, 3), (B, 2), (C, 4), (E, 7)

graph = [Node("A", [Con("D", 3), Con("B", 7)]),
         Node("B", [Con("A", 7), Con("D", 2), Con("E", 6), Con("C", 3)]),
         Node("C", [Con("B", 3), Con("E", 1), Con("D", 4)]),
         Node("D", [Con("A", 3), Con("B", 2), Con("C", 4), Con("E", 7)])]

dijkstra(graph)
