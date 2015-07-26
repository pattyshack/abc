import os

from buildutil.topo_sorter import TopoSorter


class AnalysisPass(object):
  def run(self, seed_target):
    raise NotImplemented


class BindDependencies(AnalysisPass):
  def __init__(self, pkgs):
    self.pkgs = pkgs

  def run(self, seed_target):
    frontier = [seed_target]
    while frontier:
      next_frontier = []
      for target in frontier:
        if target.deps_binded:
          continue

        deps = target.dependency_patterns.get_matching_targets(self.pkgs)

        # filter test rules from library / binary targets.
        if not target.is_test_rule():
          deps = [t for t in deps if not t.is_test_rule()]

        for t in deps:
          assert t.is_visible_to(target), (
              '%s is not visible to %s' % (t.full_name(), target.full_name()))
        target.dependencies = {t.full_name(): t for t in deps}
        target.deps_binded = True

        next_frontier.extend(deps)
      frontier = next_frontier


class CheckCycles(AnalysisPass):
  def __init__(self):
    pass

  def run(self, seed_target):
    stack = [seed_target]

    # NOTE: use DFS instead of topo sort for checking since I want to know the
    # exact cycle.
    while stack:
      for d in stack[-1].dependencies.values():
        assert d.in_cycle is not True
        if d.in_cycle is False:
          continue
        if d in stack:
          cycle = stack[stack.index(d):]
          for t in cycle:
            t.in_cycle = True
          assert False, ('Cycle detected: %s -> ...' %
              ' -> '.join([t.full_name() for t in cycle]))
        stack.append(d)
        break
      else:
        stack.pop().in_cycle = False

class BuildTargets(AnalysisPass):
  def __init__(self):
    self.sorter = TopoSorter()

  def run(self, seed_target):
    print 'Building', seed_target.full_name()
    print '=' * 80
    order = self.sorter.sort(seed_target)

    for i, target in enumerate(order):
      print 'Triggered', target.full_name(), '(%s of %s)' % (i, len(order) - 1)
      if self._should_build(target):
        print "SHIT", target.pkg_build_dir(), target.pkg_genfile_dir()
        # TODO make genfile / build dir
        # TODO clean up artifacts in genfile / build dir.
        succeeded = target.build()
        assert succeeded, 'Failed to build %s' % target.full_name()

        target.update_artifacts_max_mtime(verify_existence=True)
        target.has_modified = True
      else:
        print 'Target is up-to-date.'
        if target.has_modified is None:
          target.has_modified = False
      print '-' * 80

  def _should_build(self, target):
    if target.has_modified is not None:
      # No need to rebuild previously checked target.
      return False

    target.update_sources_max_mtime()
    target.update_artifacts_max_mtime(verify_existence=False)

    return target.should_build()

