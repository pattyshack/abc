class Node(object):
  def __init__(self, target):
    self.target = target
    self.edges = set(target.dependencies)

class TopoSorter(object):
  def __init__(self):
    pass

  def sort(self, seed_targets):
    nodes = self._collect_nodes(seed_targets)

    result = []

    while nodes:
      no_edges = []
      next_nodes = {}
      for name, node in nodes.iteritems():
        if not node.edges:
          no_edges.append(node.target)
        else:
          next_nodes[name] = node

      assert no_edges, 'Cycle found ...'
      result.extend(no_edges)

      for node in next_nodes.values():
        for target in no_edges:
          node.edges.discard(target.target_path())

      nodes = next_nodes

    return result

  def _collect_nodes(self, seed_targets):
    nodes = {}

    frontier = seed_targets
    while frontier:
      next_frontier = []
      for t in frontier:
        if t.target_path() in nodes:
          continue

        nodes[t.target_path()] = Node(t)
        next_frontier.extend(t.dependencies.values())

      frontier = next_frontier

    return nodes

