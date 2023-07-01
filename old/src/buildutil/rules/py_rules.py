import os
import os.path
import shutil
import tempfile

from buildutil.config import LOCATE_IN_GENFILE_DIR, LOCATE_IN_SRC_DIR
from buildutil.target_rule import TargetRule
from buildutil.target_patterns import TargetPatterns


PY_SECTION = 'py_rules'

DEFAULT_BASH_ABS_PATH = '/bin/bash'
DEFAULT_PYTHON_ABS_PATH = '/usr/bin/python'
DEFAULT_UNZIP_ABS_PATH = '/usr/bin/unzip'
DEFAULT_PAR_EXTRACTION_DIR = '/tmp'

RUNNER_TEMPLATE = """#!%(bash)s
env PYTHONPATH=%(runtime_dir)s %(python)s %(runtime_dir)s/%(main_py)s $@
"""

SELF_EXTRACTOR_TEMPLATE = """#!%(bash)s

TEMP_DIR=`mktemp -d %(extract_dir)s/%(target_name)s.XXXXXXXX`

# Unzipping instead of running using python zipimport because zipimport
# does not properly handle non-py files (e.g., cmodules).
%(unzip)s -d $TEMP_DIR $0 > /dev/null

env PYTHONPATH=$TEMP_DIR/%(target_name)s.runtime %(python)s $TEMP_DIR/%(target_name)s.runtime/%(main_py)s $@

EXIT_CODE=$?

rm -rf $TEMP_DIR

exit $EXIT_CODE

____ARCHIVE_BELOW____
"""


class PyInitTargetRule(TargetRule):
  def __init__(self, config, pkg_path):
    super(PyInitTargetRule, self).__init__(
        config=config,
        pkg_path=pkg_path,
        name='__init__.py',
        sources=[],
        dependencies=[],
        artifacts=['__init__.py'],
        visibility_set=['//...'])  # public to all

  @classmethod
  def is_unique_target(cls):
    return False

  @classmethod
  def generate_targets(
      cls,
      targets_accumulator,
      config,
      current_pkg_path,
      **kwargs):
    for pkg_path in PyInitTargetRule.list_dep_pkg_paths(current_pkg_path):
      targets_accumulator.append(
          PyInitTargetRule(config=config, pkg_path=pkg_path))

  @classmethod
  def include_in_all_targets(cls):
    return False

  @classmethod
  def rule_name(cls):
    assert False, 'Should never be called directly'

  def should_build(self):
    init = '__init__.py'
    src_init = self.src_abs_path(name=init)
    genfile_init = self.genfile_abs_path(name=init)
    build_init = self.build_abs_path(name=init)

    assert not os.path.exists(build_init)

    assert not os.path.isfile(src_init), (
        'Cannot have __init__.py in src directory. pkg: %s' % self.pkg_path())
    if os.path.isfile(genfile_init):
      return False

    return True

  def build(self):
    init = '__init__.py'
    src_init = self.src_abs_path(name=init)
    genfile_init = self.genfile_abs_path(name=init)

    assert not os.path.isfile(src_init)

    self.execute_cmd('touch %s' % genfile_init)
    return True

  @staticmethod
  def list_dep_pkg_paths(pkg_path):
    result = []

    while True:
      assert pkg_path.startswith('//')
      result.append(pkg_path)

      if pkg_path == '//':
        return result

      pkg_path, _ = os.path.split(pkg_path)

  @staticmethod
  def list_dep_target_paths(pkg_path):
    result = []
    for path in PyInitTargetRule.list_dep_pkg_paths(pkg_path):
      result.append(path + ':__init__.py')

    return result


class PyLibraryTargetRule(TargetRule):
  def __init__(
      self,
      config,
      pkg_path,
      name,
      srcs=(),
      deps=(),
      visibility=None):

    deps = list(deps) + PyInitTargetRule.list_dep_target_paths(pkg_path)

    super(PyLibraryTargetRule, self).__init__(
        config=config,
        pkg_path=pkg_path,
        name=name,
        sources=srcs,
        dependencies=deps,
        artifacts=srcs,
        visibility_set=visibility)

  @classmethod
  def generate_targets(
      cls,
      targets_accumulator,
      config,
      current_pkg_path,
      **kwargs):
    PyInitTargetRule.generate_targets(
        targets_accumulator=targets_accumulator,
        config=config,
        current_pkg_path=current_pkg_path)
    targets_accumulator.append(
        cls(config=config, pkg_path=current_pkg_path, **kwargs))

  @classmethod
  def rule_name(cls):
    return "py_library"

  def build(self):
    for src_file_name in self.sources():
      abs_path = self.locate_file(src_file_name)
      assert abs_path

      self.execute_cmd('touch %s' % abs_path)

    return True


class PyBinaryTargetRule(TargetRule):
  def __init__(
      self,
      config,
      pkg_path,
      name,
      srcs=(),
      deps=(),
      visibility=None):

    assert len(srcs) == 1, 'Must have exactly one .py src file'

    deps = list(deps) + PyInitTargetRule.list_dep_target_paths(pkg_path)

    super(PyBinaryTargetRule, self).__init__(
        config=config,
        pkg_path=pkg_path,
        name=name,
        sources=srcs,
        dependencies=deps,
        artifacts=[],  # compute dynamically
        visibility_set=visibility)

  @classmethod
  def generate_targets(
      cls,
      targets_accumulator,
      config,
      current_pkg_path,
      **kwargs):
    PyInitTargetRule.generate_targets(
        targets_accumulator,
        config,
        current_pkg_path)

    bin_target = cls(config=config, pkg_path=current_pkg_path, **kwargs)
    targets_accumulator.append(bin_target)

    name = kwargs['name']
    visibility = None
    if 'visibility' in kwargs:
      visibility = kwargs['visibility']

    par_target = PyParTargetRule(
        config=config,
        pkg_path=current_pkg_path,
        name=name,
        visibility=visibility)
    par_target.is_test_par = True

    targets_accumulator.append(par_target)

  @classmethod
  def rule_name(cls):
    return "py_binary"

  def artifacts(self):
    assert self.deps_binded

    result = set()
    for artifact in self.list_dependencies_artifacts():
      result.add(
        os.path.join(self.name + '.runtime', artifact[2:]))

    for src in self.sources():
      result.add(
          os.path.join(self.name + '.runtime', self.pkg_path(src)[2:]))

    result.add(self.name)

    return result

  @classmethod
  def include_dependencies_artifacts(cls):
    return False

  def build(self):
    src_pkg_paths = self.list_dependencies_artifacts()
    for name in self.sources():
      src_pkg_paths.add(self.pkg_path(name=name))

    runtime_abs_path = self.build_abs_path(name=self.name + '.runtime')
    self.execute_cmd('rm -rf %s' % runtime_abs_path)

    dir_abs_paths = set()
    for pkg_path in src_pkg_paths:
      pkg_path, _ = os.path.split(pkg_path)
      assert pkg_path.startswith('//')
      dir_abs_paths.add(os.path.join(runtime_abs_path, pkg_path[2:]))

    for abs_path in dir_abs_paths:
      self.execute_cmd('mkdir -p %s' % abs_path)

    for pkg_path in src_pkg_paths:
      abs_path = self.config.locate_file(
          pkg_path,
          locate_order=(LOCATE_IN_SRC_DIR, LOCATE_IN_GENFILE_DIR))
      assert abs_path, 'Failed to locate: %s' % pkg_path

      self.execute_cmd('ln -s %s %s' % (
          abs_path,
          os.path.join(runtime_abs_path, pkg_path[2:])))

    return self.write_runner_script()

  def write_runner_script(self):
    script_abs_path = self.build_abs_path(name=self.name)

    tmpl_vals ={
        'bash': self.config.get(
            PY_SECTION,
            'bash_location',
            DEFAULT_BASH_ABS_PATH),
        'runtime_dir': script_abs_path + '.runtime',
        'python': self.config.get(
            PY_SECTION,
            'python_location',
            DEFAULT_PYTHON_ABS_PATH),
        'main_py': self.pkg_path(name=self.sources()[0])[2:],
        }

    print 'Writing:', script_abs_path
    with open(script_abs_path, 'w') as f:
      f.write(RUNNER_TEMPLATE % tmpl_vals)

    os.chmod(script_abs_path, 0755)
    return True


class PyParTargetRule(TargetRule):
  def __init__(
      self,
      config,
      pkg_path,
      name,
      visibility=None):
    super(PyParTargetRule, self).__init__(
        config=config,
        pkg_path=pkg_path,
        name=name + '.par',
        sources=(),
        dependencies=[':%s' % name],
        artifacts=[name + '.par'],
        visibility_set=visibility)

    self.is_test_par = False
    self.original_target_name = name

  @classmethod
  def rule_name(cls):
    assert False, 'Should never be called directly'

  @classmethod
  def include_in_all_targets(cls):
    return False

  def is_test_rule(self):
    return self.is_test_par

  @classmethod
  def include_dependencies_artifacts(cls):
    return False

  def build(self):
    tmp_dir = tempfile.mkdtemp(
        prefix=self.name + '.tmp.',
        dir=self.build_abs_path())
    print 'Created temp dir:', tmp_dir

    runtime_dir_name = self.original_target_name + '.runtime'
    self.execute_cmd('ln -s %s %s' % (
        self.build_abs_path(name=runtime_dir_name),
        os.path.join(tmp_dir, runtime_dir_name)))

    file_pkg_paths = self.list_dependencies_artifacts()
    runner_script_pkg_path = self.pkg_path(name=self.original_target_name)

    file_relative_paths = []
    for f in file_pkg_paths:
      if f == runner_script_pkg_path:
        continue

      _, _, f = f.partition(runtime_dir_name)
      assert f
      file_relative_paths.append(runtime_dir_name + f)

    self.execute_cmd('cd %s ; zip pkg.zip %s' % (
        tmp_dir,
        ' '.join(file_relative_paths)))

    extractor_path = os.path.join(tmp_dir, 'extractor.sh')
    print 'Writing extractor script:', extractor_path
    self.write_extractor_script(extractor_path)

    self.execute_cmd('cd %s ; cat extractor.sh pkg.zip >> %s' % (
        tmp_dir,
        self.name))

    par_tmp_path = os.path.join(tmp_dir, self.name)

    self.execute_cmd('zip -A %s' % par_tmp_path)

    self.execute_cmd('chmod 755 %s' % par_tmp_path)

    self.execute_cmd('mv %s %s' % (
        par_tmp_path,
        self.build_abs_path(name=self.name)))

    print 'Removing temp dir:', tmp_dir
    shutil.rmtree(tmp_dir)

    return True

  def write_extractor_script(self, extractor_path):
    assert len(self.dependencies) == 1
    dep = self.dependencies[self.pkg_path() + ':' + self.original_target_name]
    assert isinstance(dep, PyBinaryTargetRule)
    assert len(dep.sources()) == 1

    tmpl_vals ={
        'bash': self.config.get(
            PY_SECTION,
            'bash_location',
            DEFAULT_BASH_ABS_PATH),
        'extract_dir': self.config.get(
            PY_SECTION,
            'par_extraction_location',
            DEFAULT_PAR_EXTRACTION_DIR),
        'target_name': self.original_target_name,
        'unzip': self.config.get(
            PY_SECTION,
            'unzip_location',
            DEFAULT_UNZIP_ABS_PATH),
        'python': self.config.get(
            PY_SECTION,
            'python_location',
            DEFAULT_PYTHON_ABS_PATH),
        'main_py': self.pkg_path(name=dep.sources()[0])[2:],
        }

    with open(extractor_path, 'w') as f:
      f.write(SELF_EXTRACTOR_TEMPLATE % tmpl_vals)


# XXX maybe don't subclass ByBinary cuz it's limits to a single src file ...
class PyTestTargetRule(PyBinaryTargetRule):
  def __init__(
      self,
      config,
      pkg_path,
      name,
      srcs=(),
      deps=(),
      visibility=None):
    super(PyTestTargetRule, self).__init__(
        config=config,
        pkg_path=pkg_path,
        name=name,
        srcs=srcs,
        deps=deps,
        visibility=visibility)

  @classmethod
  def rule_name(cls):
    return "py_test"

  def is_test_rule(self):
    return True

