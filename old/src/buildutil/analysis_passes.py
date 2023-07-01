import os
import sys

from buildutil.topo_sorter import TopoSorter


class AnalysisPass(object):
  def run(self, seed_targets):
    raise NotImplemented


class BindDependencies(AnalysisPass):
  def __init__(self, pkgs):
    self.pkgs = pkgs

  def run(self, seed_targets):
    frontier = seed_targets
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
          assert t.is_visible_to(target), '%s is not visible to %s' % (
              t.target_path(),
              target.target_path())
        target.dependencies = {t.target_path(): t for t in deps}
        target.deps_binded = True

        next_frontier.extend(deps)
      frontier = next_frontier


class CheckCycles(AnalysisPass):
  def __init__(self):
    pass

  def run(self, seed_targets):
    for target in seed_targets:
      self._check(target)

  def _check(self, seed_target):
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
              ' -> '.join([t.target_path() for t in cycle]))
        stack.append(d)
        break
      else:
        stack.pop().in_cycle = False


class BuildTargets(AnalysisPass):
  def __init__(self):
    self.sorter = TopoSorter()

  def run(self, seed_targets):
    order = self.sorter.sort(seed_targets)

    self._make_dirs([t.genfile_abs_path() for t in order])
    self._make_dirs([t.build_abs_path() for t in order])

    for i, target in enumerate(order):
      print 'Building', target.target_path(), '(%s of %s)' % (i, len(order) - 1)
      if self._should_build(target):
        succeeded = target.build()
        assert succeeded, 'Failed to build %s' % target.target_path()

        target.update_artifacts_max_mtime(verify_existence=True)
        target.has_modified = True

        print 'Done'
      else:
        print 'Target is up-to-date.'
        if target.has_modified is None:
          target.has_modified = False
      print '-' * 80

    print 'BUILD DONE'

  def _make_dirs(self, dirs):
    for d in set(dirs):
      try:
        os.makedirs(d)
      except OSError as e:
        if e.errno == 17:  # i.e., file exists
          assert os.path.isdir(d), d
        else:
          raise

  def _should_build(self, target):
    if target.has_modified is not None:
      # No need to rebuild previously checked target.
      return False

    target.update_sources_max_mtime()
    target.update_artifacts_max_mtime(verify_existence=False)

    return target.should_build()


class TestTargets(AnalysisPass):
  def __init__(self):
    pass

  def run(self, seed_targets):
    tests = [t for t in seed_targets if t.is_test_rule()]

    passed = 0
    for i, target in enumerate(tests):
      print 'Testing',  target.target_path(), '(%s of %s)' % (i, len(tests) - 1)
      succeeded = target.test()
      if succeeded:
        passed += 1
      print 'Done'
      print '-' * 80

    print '%s of %s test targets passed.' % (passed, len(tests))
    if passed != len(tests):
      print 'TEST FAILED'
      sys.exit(1)
    else:
      print 'TEST PASSED'


class PrintBuildOrder(AnalysisPass):
  def __init__(self):
    self.sorter = TopoSorter()

  def run(self, seed_targets):
    order = self.sorter.sort(seed_targets)

    print 'Build order:'
    for target in order:
      print ' ', target.target_path()

