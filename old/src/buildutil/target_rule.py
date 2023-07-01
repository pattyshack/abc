import os
import subprocess

from buildutil.config import (
    DEFAULT_LOCATE_ORDER,
    LOCATE_ARTIFACT_ORDER,
    LOCATE_SOURCE_ORDER,
    )
from buildutil.target_patterns import TargetPatterns
from buildutil.util import validate_pkg_path, validate_target_name


class TargetRule(object):
  def __init__(
      self,
      config,
      pkg_path,
      name,
      sources=(),
      dependencies=(),
      artifacts=(),
      visibility_set=None):
    """
    sources: list of files (in relative path form) in the current directory
      (may also be in sub directories).
    dependencies: list of target path.
    artifacts: list of "output" files (in relative path form).
    visibility_set: list of visibility targets (None means use package default)
    """
    assert validate_pkg_path(pkg_path), 'Invalid package path: %s' % pkg_path
    assert validate_target_name(name), (
        'Invalid target name: %s (pkg: %s)' % (name, pkg_path))

    self._pkg_path = pkg_path
    self.name = name
    self.config = config
    self._sources = sources
    self.dependency_patterns = TargetPatterns(pkg_path)
    self.dependency_patterns.set_patterns(dependencies)
    self._artifacts = artifacts

    if visibility_set is not None:
      self.visibility_patterns = TargetPatterns(pkg_path)
      self.visibility_patterns.set_patterns(visibility_set)
    else:
      # bind when the target is added to a package during registration.
      self.visibility_patterns = None

    # The following are initialized in later analysis passes.

    self.deps_binded = False
    self.dependencies = {}

    # None for not checked, False for no cycle, True for cycle.
    self.in_cycle = None

    # NOTE: If this was implemented as a server (like blaze), then we should
    # check source content hashes/size instead of sources/artifacts mtime
    self.soruces_max_mtime = None
    self.artifacts_max_mtime = None
    # None for not checked, True if has build, False if build was unnecessary.
    self.has_modified = None


  def target_path(self):
    """DO NOT OVERRIDE"""
    return self._pkg_path + ':' + self.name

  def pkg_path(self, name=''):
    if not name:
      return self._pkg_path

    return os.path.join(self._pkg_path, name)

  def src_abs_path(self, name=''):
    """DO NOT OVERRIDE"""
    return self.config.pkg_path_to_src_abs_path(self.pkg_path(name=name))

  def genfile_abs_path(self, name=''):
    """DO NOT OVERRIDE"""
    return self.config.pkg_path_to_genfile_abs_path(self.pkg_path(name=name))

  def build_abs_path(self, name=''):
    """DO NOT OVERRIDE"""
    return self.config.pkg_path_to_build_abs_path(self.pkg_path(name=name))

  def is_visible_to(self, target):
    """DO NOT OVERRIDE"""
    if self.pkg_path() == target.pkg_path():
      return True

    return self.visibility_patterns.matches(target)

  def _get_max_mtime(
      self,
      files,
      locate_order,
      verify_existence=False):
    """DO NOT OVERRIDE"""
    max_mtime = None
    for f in files:
      abs_path = self.locate_file(f, locate_order=locate_order)
      if abs_path is None:
        assert not verify_existence, (
            'Failed to locate: %s (target: %s)' % (f, self.target_path()))
        return None

      mtime = os.lstat(abs_path).st_mtime  # Don't follow links
      if max_mtime is None or max_mtime < mtime:
        max_mtime = mtime

    return max_mtime

  def update_sources_max_mtime(self):
    """DO NOT OVERRIDE"""
    self.sources_max_mtime = self._get_max_mtime(
        self.sources(),
        self.locate_source_order(),
        verify_existence=True)

  def update_artifacts_max_mtime(self, verify_existence=True):
    """DO NOT OVERRIDE"""
    artifacts = self.artifacts()

    assert artifacts, 'Target must have at least one artifact: %s (pkg: %s)' % (
        name,
        pkg.pkg_path)

    self.artifacts_max_mtime = self._get_max_mtime(
        artifacts,
        self.locate_artifact_order(),
        verify_existence=verify_existence)

  def execute_cmd(self, cmd_str, additional_env=None):
    """DO NOT OVERRIDE.  Use this for shelling out commands for building /
    testing."""
    env = {
      'PROJECT_ROOT_DIR' : self.config.project_dir_abs_path,
      'SRC_DIR' : self.config.src_dir_abs_path,
      'GENFILE_DIR' : self.config.genfile_dir_abs_path,
      'BUILD_DIR': self.config.build_dir_abs_path,
      'PACKAGE': self.pkg_path()[2:],
      'TARGET': self.name,
    }
    if additional_env:
      env.update(additional_env)

    print 'Executing:', cmd_str
    p = subprocess.Popen(
        cmd_str,
        shell=True,
        cwd=self.config.project_dir_abs_path,
        env=env)

    r = p.wait()
    assert r == 0, 'Failed to execute: %s' % cmd_str

  def locate_file(
      self,
      file_name,
      locate_order=DEFAULT_LOCATE_ORDER):
    """DO NOT OVERRIDE"""
    return self.config.locate_file(
        self.pkg_path(name=file_name),
        locate_order=locate_order)

  def locate_source_order(self):
    return LOCATE_SOURCE_ORDER

  def locate_artifact_order(self):
    return LOCATE_ARTIFACT_ORDER

  @classmethod
  def is_unique_target(cls):
    """Return false if the target can be register multiple times into the
    same package (The first entry is used; the rest are ignored)."""
    return True

  @classmethod
  def generate_targets(
      cls,
      targets_accumulator,
      config,
      current_pkg_path,
      **kwargs):
    """Override to customize target registration (see PyBinaryTargetRule for
    example)."""
    targets_accumulator.append(
        cls(config=config, pkg_path=current_pkg_path, **kwargs))

  @classmethod
  def rule_name(cls):
    """The function name used in BUILD file, e.g., cc_library"""
    raise NotImplemented

  @classmethod
  def include_in_all_targets(cls):
    """When true, return the target as part of pkg.get_all_targets(), which
    is used for :all or ... target expansion"""
    return True

  def is_test_rule(self):
    """When building libraries / binaries, test targets are ignored.  When
    true, must implement test"""
    return False

  def sources(self):
    """Override if the source list cannot be computed statically during
    initialization.  Can assume dependencies are binded and built when
    overriding"""
    return self._sources

  def artifacts(self):
    """Override if the artifact list cannot be computed statically during
    initialization. Can assume dependencies are binded and built when
    overriding."""
    return self._artifacts

  def list_dependencies_artifacts(self):
    result = set()
    for d in self.dependencies.values():
      result = result.union(d.list_artifacts())

    return result

  @classmethod
  def include_dependencies_artifacts(cls):
    """Controls list_artifacts default behavior.  Useful for stopping artifacts
    from propagating beyond a certain target."""
    return True

  def list_artifacts(self):
    result = set()
    for name in self.artifacts():
      result.add(self.pkg_path(name=name))

    if self.include_dependencies_artifacts():
      result = result.union(self.list_dependencies_artifacts())

    return result

  def should_build(self):
    # Artifacts are created without source and dependencies.
    if not self.sources() and not self.dependencies:
      return True

    # First time building the artifacts.
    if self.artifacts_max_mtime is None:
      return True

    if self.sources():
      assert self.sources_max_mtime

      # Sources are newer than the artifacts.
      if self.artifacts_max_mtime < self.sources_max_mtime:
        return True

    for dep in self.dependencies.values():
      assert dep.has_modified is not None
      # A dependency changed within the same session (i.e., when multiple
      # targets are specified in the same build command).  This is more
      # accurate than the mtime check.
      if dep.has_modified:
        return True

      assert dep.artifacts_max_mtime is not None
      # The dependency changed from a previous session.
      if self.artifacts_max_mtime < dep.artifacts_max_mtime:
        return True

    return False

  def build(self):
    """How the target should be build.  Returns true if build succeeded, false
    otherwise."""
    print 'BUILD', self.name
    #raise NotImplemented
    return True

  def test(self):
    """How the target should be tested.  Returns true if test succeeded, false
    otherwise."""
    try:
      self.execute_cmd(self.build_abs_path(self.name))
    except AssertionError:
      return False
    return True

