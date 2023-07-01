import glob
import os.path

from buildutil.config import (
    LOCATE_IN_BUILD_DIR,
    LOCATE_IN_GENFILE_DIR,
    LOCATE_IN_SRC_DIR,
    )
from buildutil.target_rule import TargetRule


GO_SECTION = 'go_rules'

DEFAULT_GO = '/usr/local/go/bin/go'


class GoPkgTargetRule(TargetRule):
  def __init__(
      self,
      config,
      pkg_path,
      name='go',
      deps=()):
    super(GoPkgTargetRule, self).__init__(
        config,
        pkg_path,
        name,
        sources=(),
        dependencies=deps,
        artifacts=(),
        visibility_set=['//...'])

    self._globbed = False
    self._srcs = []
    self._genfiles = []

  def locate_source_order(self):
    return [LOCATE_IN_SRC_DIR, LOCATE_IN_GENFILE_DIR]

  def locate_artifact_order(self):
    return [LOCATE_IN_GENFILE_DIR, LOCATE_IN_BUILD_DIR]

  def sources(self):
    self.maybe_glob_files()
    return self._sources

  def artifacts(self):
    self.maybe_glob_files()
    return self._artifacts

  @classmethod
  def rule_name(cls):
    return "go_pkg"

  def is_test_rule(self):
    return True

  def maybe_glob_files(self):
    if self._globbed:
      return

    builds = glob.glob(os.path.join(self.build_abs_path(), '*.go'))
    assert not builds, (
        '.go files found in build directory (target: %s)' % self.target_path())

    genfiles = glob.glob(os.path.join(self.genfile_abs_path(), '*.go'))
    srcs = glob.glob(os.path.join(self.src_abs_path(), '*.go'))

    self._genfiles = [os.path.split(f)[1] for f in genfiles]
    self._srcs = [os.path.split(f)[1] for f in srcs]

    sources = list(set(self._genfiles).union(set(self._srcs)))
    self._sources = sources
    self._artifacts = sources + ['__go_build_ok__']

    assert sources, (
        'No .go file found in genfile / src directories (target: %s)' %
            self.target_path())

    self._globbed = True

  def build(self):
    self.make_genfile_links()

    gopkg_abs_dir = os.path.join(self.config.build_dir_abs_path, '.gopkg')
    self.execute_cmd('mkdir -p %s' % gopkg_abs_dir)

    gobuild_abs_dir = self.build_abs_path('.gobuild')
    if not os.path.exists(gobuild_abs_dir):
      self.execute_cmd('mkdir -p %s' % gobuild_abs_dir)
      self.execute_cmd('ln -s %s %s' % (
          gopkg_abs_dir,
          os.path.join(gobuild_abs_dir, 'pkg')))
      self.execute_cmd('ln -s %s %s' % (
          self.config.genfile_dir_abs_path,
          os.path.join(gobuild_abs_dir, 'src')))
      self.execute_cmd('ln -s %s %s' % (
          self.build_abs_path(),
          os.path.join(gobuild_abs_dir, 'bin')))

    pkg_path = self.pkg_path()
    assert pkg_path != '//'

    go = self.config.get(GO_SECTION, 'go_location', DEFAULT_GO)
    self.execute_cmd(
        '%s install %s' % (go, pkg_path[2:]),
        additional_env={'GOPATH': gobuild_abs_dir})

    self.execute_cmd('touch %s' % self.genfile_abs_path('__go_build_ok__'))
    return True

  def make_genfile_links(self):
    assert self._globbed

    for f in self._sources:
      genfile = self.genfile_abs_path(f)

      if f in self._genfiles:
        if f not in self._srcs:  # A generated file
          assert os.path.isfile(genfile)
          continue

        assert os.path.islink(genfile)
        self.execute_cmd('rm %s' % genfile)

      assert f in self._srcs
      self.execute_cmd('ln -s %s %s' % (self.src_abs_path(f), genfile))

  def test(self):
    pkg_path = self.pkg_path()
    assert pkg_path != '//'

    go = self.config.get(GO_SECTION, 'go_location', DEFAULT_GO)
    self.execute_cmd(
        '%s test %s' % (go, pkg_path[2:]),
        additional_env={'GOPATH': self.build_abs_path('.gobuild')})

    return True
