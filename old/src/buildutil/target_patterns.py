# target pattern (a subset of blaze syntax)
#
# unless blaze, target must be a rule.  Does not support file and :* syntax;
# use filegroup instead.
#
# relative to project root:
#
# //foo/bar:zzz a specific rule in foo/bar
# //foo/bar:all all rules in foo/bar
# //foo/bar/... all targets recursively beneath foo/bar
# //:all all targets in project root
# //... all targets
#
# relative to current directory.  for example assume we're in foo
#
# bar:zzz a specific
# bar:all
# bar/... all targets recursively beneath foo/bar
# :zzz a specific target in foo
# :all all targets in foo
# ... all targets recursively beneath foo
#
# subtractive pattern:
#
# //... -//foo/bar:all all targets except foo/bar
#
# NOTE: buildutil becomes a server, and persistently tracks
# content hash/size/mtime, then convert src files into implicit target rules
# and support file pattern.
import re

from buildutil.util import pkg_name_join, validate_target_pattern


class TargetPatterns(object):
  def __init__(
      self,
      current_pkg_path):
    self.current_pkg_path = current_pkg_path

    self.additive_patterns = []
    self.subtractive_patterns = []

  def set_patterns(self, pattern_strs):
    self.additive_patterns = []
    self.subtractive_patterns = []

    for pattern_str in pattern_strs:
      group = self.additive_patterns
      if pattern_str.startswith('-'):
        pattern_str = pattern_str[1:]
        group = self.subtractive_patterns

      assert validate_target_pattern(pattern_str), pattern_str
      pattern_str = pkg_name_join(self.current_pkg_path, pattern_str)

      if pattern_str == '...' or pattern_str == '//...':
        pattern = RecursiveTargetPattern(pattern_str[:-3])
      elif pattern_str.endswith('/...'):
        pattern = RecursiveTargetPattern(pattern_str[:-4])
      elif pattern_str.endswith(':all'):
        pattern = AllTargetPattern(pattern_str[:-4])
      else:
        pattern = SingleTargetPattern(pattern_str)
      group.append(pattern)

  def matches(self, target):
    for pattern in self.subtractive_patterns:
      if pattern.matches(target):
        return False

    for pattern in self.additive_patterns:
      if pattern.matches(target):
        return True

    return False

  def get_matching_targets(self, packages):
    candidates = []
    for pattern in self.additive_patterns:
      candidates.extend(pattern.get_matching_targets(packages))

    result = []
    for target in candidates:
      for pattern in self.subtractive_patterns:
        if pattern.matches(target):
          break
      else:
        result.append(target)

    return result


class SingleTargetPattern(object):
  def __init__(self, target_path):
    self.target_path = target_path

  def matches(self, target):
    return self.target_path == target.target_path()

  def get_matching_targets(self, packages):
    return [packages.get_or_load_target(self.target_path)]


class AllTargetPattern(object):
  def __init__(self, pkg_path):
    self.pkg_path = pkg_path

  def matches(self, target):
    return self.pkg_path == target.pkg_path()

  def get_matching_targets(self, packages):
    return packages.get_or_load_package(self.pkg_path).get_all_targets()


class RecursiveTargetPattern(object):
  def __init__(self, pkg_path):
    self.pkg_path = pkg_path

  def matches(self, target):
    if self.pkg_path == target.pkg_path():
      return True

    if self.pkg_path == '//':
      return True

    return target.pkg_path().startswith(self.pkg_path + '/')

  def get_matching_targets(self, packages):
    targets = []
    for pkg in packages.get_or_load_all_subpackages(self.pkg_path):
      targets.extend(pkg.get_all_targets())
    return targets

