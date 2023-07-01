import os.path

from buildutil.target_rule import TargetRule


LEX_YACC_SECTION = 'lex_yacc_rules'
DEFAULT_LEX = '/usr/bin/flex'
DEFAULT_YACC = '/usr/bin/bison'

class LexTargetRule(TargetRule):
  def __init__(
      self,
      config,
      pkg_path,
      name,
      srcs=(),
      deps=(),
      visibility=None):
    assert len(srcs) == 1
    file_name, ext = os.path.splitext(srcs[0])
    assert ext == '.l'

    super(LexTargetRule, self).__init__(
        config=config,
        pkg_path=pkg_path,
        name=name,
        sources=srcs,
        dependencies=deps,
        artifacts=[file_name + '.yy.c'],
        visibility_set=visibility)

  @classmethod
  def rule_name(cls):
    return "gen_lex"

  def build(self):
    lex = self.config.get(LEX_YACC_SECTION, 'lex_location', DEFAULT_LEX)

    src = self.sources()[0]
    name, _ = os.path.splitext(src)

    self.execute_cmd('%s -o %s %s' % (
        lex,
        self.genfile_abs_path(name + '.yy.c'),
        self.locate_file(src)))

    return True

  def list_artifacts(self):
    # Don't propagate artifacts.
    return []


class YaccTargetRule(TargetRule):
  def __init__(
      self,
      config,
      pkg_path,
      name,
      srcs=(),
      deps=(),
      visibility=None):
    assert len(srcs) == 1
    file_name, ext = os.path.splitext(srcs[0])
    assert ext == '.y'

    # NOTE: additional files (e.g., location.hh, position.hh, stack.hh)
    # may also be generated.  However, since we don't propagate the
    # artifacts up the dependency china, tracking .tab.c and .tab.h should
    # be sufficient to determine if the artifacts are out of date.
    artifacts = [file_name + '.tab.cc', file_name + '.tab.hh']

    super(YaccTargetRule, self).__init__(
        config=config,
        pkg_path=pkg_path,
        name=name,
        sources=srcs,
        dependencies=deps,
        artifacts=artifacts,
        visibility_set=visibility)

  @classmethod
  def rule_name(cls):
    return "gen_yacc"

  def build(self):
    yacc = self.config.get(LEX_YACC_SECTION, 'yacc_location', DEFAULT_YACC)

    src = self.sources()[0]
    name, _ = os.path.splitext(src)

    self.execute_cmd('%s -v -o %s %s' % (
        yacc,
        self.genfile_abs_path(name + '.tab.cc'),
        self.locate_file(src)))

    return True

  def list_artifacts(self):
    # Don't propagate artifacts.
    return []

