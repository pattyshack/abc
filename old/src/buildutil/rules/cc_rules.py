import os.path

from buildutil.target_rule import TargetRule


HDR_EXTS = ('.h', '.hh')
SRC_EXTS = ('.c', '.cc', '.cpp')

CC_SECTION = 'cc_rules'

DEFAULT_CC = 'g++'
DEFAULT_CFLAGS = '-Wall -pthread'
DEFAULT_LFLAGS = '-pthread'


class CcLibraryTargetRule(TargetRule):
  def __init__(
      self,
      config,
      pkg_path,
      name,
      srcs=(),
      hdrs=(),
      deps=(),
      visibility=None):

    sources = []
    artifacts = []

    for hdr in set(hdrs):
      _, ext = os.path.splitext(hdr)
      assert ext in HDR_EXTS, (
          'Unexpected header extension: %s (target: %s:%s)' % (
              hdr,
              pkg_path,
              name))

      sources.append(hdr)
      artifacts.append(hdr)

    for src in set(srcs):
      src_name, ext = os.path.splitext(src)
      assert ext in SRC_EXTS, (
          'Unexpected src extension: %s (target: %s:%s)' % (
              src,
              pkg_path,
              name))

      sources.append(src)
      artifacts.append(src_name + '.o')

    super(CcLibraryTargetRule, self).__init__(
        config,
        pkg_path,
        name,
        sources=sources,
        dependencies=deps,
        artifacts=artifacts,
        visibility_set=visibility)

  @classmethod
  def rule_name(cls):
    return "cc_library"

  def build(self):
    cc = self.config.get(CC_SECTION, 'cc_location', DEFAULT_CC)
    cflags = self.config.get(CC_SECTION, 'cflags', DEFAULT_CFLAGS)
    hdr_dirs = list(self.config.get(CC_SECTION, 'hdr_dirs', '').split(':'))
    hdr_dirs.append(self.config.src_dir_abs_path)
    hdr_dirs.append(self.config.genfile_dir_abs_path)

    for src in self.sources():
      name, ext = os.path.splitext(src)
      if ext in HDR_EXTS:
        continue

      self.execute_cmd(
          '%s -c %s %s -o %s %s' % (
              cc,
              cflags,
              ' '.join('-I' + d for d in hdr_dirs),
              self.build_abs_path(name=name + '.o'),
              self.locate_file(src)))

    return True


class CcBinaryTargetRule(TargetRule):
  def __init__(
      self,
      config,
      pkg_path,
      name,
      visibility=None):
    super(CcBinaryTargetRule, self).__init__(
        config,
        pkg_path,
        name,
        sources=(),
        dependencies=[':%s.objs' % name],
        artifacts=[name],
        visibility_set=visibility)

  @classmethod
  def generate_targets(
      cls,
      targets_accumulator,
      config,
      current_pkg_path,
      **kwargs):

    name = kwargs['name']
    kwargs['name'] = name + '.objs'

    visibility = None
    if 'visibility' in kwargs:
      visibility = kwargs['visibility']

    targets_accumulator.append(
        CcLibraryTargetRule(config, current_pkg_path, **kwargs))

    targets_accumulator.append(
        cls(config,
            current_pkg_path,
            name,
            visibility=visibility))

  @classmethod
  def rule_name(cls):
    return "cc_binary"

  @classmethod
  def include_dependencies_artifacts(cls):
    return False

  def build(self):
    cc = self.config.get(CC_SECTION, 'cc_location', DEFAULT_CC)
    lflags = self.config.get(CC_SECTION, 'lflags', DEFAULT_CFLAGS)

    # TODO handle static / dynamical libs (-L & -l options)
    obj_files = []
    for artifact in self.list_dependencies_artifacts():
      if not artifact.endswith('.o'):
        continue

      abs_path = self.config.locate_file(
          artifact,
          locate_order=self.locate_artifact_order())
      assert abs_path
      obj_files.append(abs_path)

    self.execute_cmd(
        '%s %s -o %s %s' % (
            cc,
            lflags,
            self.build_abs_path(name=self.name),
            ' '.join(obj_files)))

    return True


class CcTestTargetRule(CcBinaryTargetRule):
  def __init__(
      self,
      config,
      pkg_path,
      name,
      visibility=None):
    super(CcTestTargetRule, self).__init__(
        config,
        pkg_path,
        name,
        visibility=visibility)

  @classmethod
  def rule_name(cls):
    return "cc_test"

  def is_test_rule(self):
    return True

