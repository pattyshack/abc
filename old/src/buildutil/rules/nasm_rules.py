import os.path

from buildutil.target_rule import TargetRule


NASM_SECTION = 'nasm_rules'
DEFAULT_NASM = '/usr/bin/nasm'

INCLUDES_EXTS = ('.asm')

OUTPUT_FORMAT_EXTS = {
    'bin': '.bin',
    'elf': '.o',
    }

class NasmTargetRule(TargetRule):
  def __init__(
      self,
      config,
      pkg_path,
      name,
      src=None,
      fmt='bin',
      includes=(),
      deps=(),
      visibility=None):
    file_name, ext = os.path.splitext(src)
    assert ext == '.asm'

    assert fmt in OUTPUT_FORMAT_EXTS
    output_ext = OUTPUT_FORMAT_EXTS[fmt]

    sources = [src]
    for f in includes:
      _, ext = os.path.splitext(f)
      assert ext in INCLUDES_EXTS, (
          'Unexpected include extension: %s (target: %s:%s)' % (
              f,
              pkg_path,
              name))

      sources.append(f)

    super(NasmTargetRule, self).__init__(
        config=config,
        pkg_path=pkg_path,
        name=name,
        sources=sources,
        dependencies=deps,
        artifacts=[file_name + output_ext],
        visibility_set=visibility)

    self.output_format = fmt
    self.output_ext = output_ext

  @classmethod
  def rule_name(cls):
    return "gen_nasm"

  def build(self):
    nasm = self.config.get(NASM_SECTION, 'nasm_location', DEFAULT_NASM)

    src = self.sources()[0]

    self.execute_cmd('%s -w+orphan-labels %s -f %s -o %s' % (
        nasm,
        self.locate_file(src),
        self.output_format,
        self.build_abs_path(name=self.name + self.output_ext)))

    return True

  def list_artifacts(self):
    # Don't propagate artifacts
    return []
