import ConfigParser
import optparse
import os
import os.path
import subprocess

from buildutil.analysis_passes import (
    BindDependencies,
    BuildTargets,
    CheckCycles,
    PrintBuildOrder,
    TestTargets,
    )
from buildutil.config import Config
from buildutil.package import PackageSet
from buildutil.target_patterns import TargetPatterns
from buildutil.util import validate_target_pattern


def sanitize_pattern_strs(config, cwd_abs_path, args):
  pattern_strs = []
  for orig_arg in args:
    if orig_arg.startswith('//') or orig_arg.startswith('-//'):
      # target name must be valid if given in full form
      assert validate_target_pattern(orig_arg), (
          'Invalid target pattern: %s' % orig_arg)
      pattern_strs.append(orig_arg)
      continue

    if orig_arg.startswith('-'):
      subtractive = True
      arg = orig_arg[1:]
    else:
      subtractive = False
      arg = orig_arg

    if arg.endswith('...'):
      pkg_name = arg[:-3]
      target_name = '...'
    else:
      if ':' in arg:
        pkg_name, _, target_name = arg.rpartition(':')
      else:
        pkg_name = arg
        target_name = 'all'

    pkg_abs_path = os.path.normpath(os.path.join(cwd_abs_path, pkg_name))

    if pkg_abs_path != config.src_dir_abs_path:
      assert pkg_abs_path.startswith(config.src_dir_abs_path + '/'), (
          'Invalid target pattern: %s' % orig_arg)

    pkg_path = '//' + pkg_abs_path[(len(config.src_dir_abs_path) + 1):]

    if target_name == '...':
      if pkg_path == '//':
        pattern_str = '//...'
      else:
        pattern_str = pkg_path + '/...'
    else:
      pattern_str = pkg_path + ':' + target_name

    if subtractive:
      pattern_str = '-' + pattern_str

    assert validate_target_pattern(pattern_str), (
        'Invalid target pattern: %s (expansion: %s)' % (
            orig_arg,
            pattern_str))
    pattern_strs.append(pattern_str)

  return pattern_strs


def main():
  usage = 'USAGE: %prog <analyze|build|test|clean> [options] [targets]*'
  args_parser = optparse.OptionParser(usage)

  options, args = args_parser.parse_args()
  if not args:
    args_parser.error('Must specify command: analyze|build|test|clean')

  ini_config = ConfigParser.ConfigParser()

  cfg_abs_path = os.path.expanduser('~/.buildutil_config.ini')
  if os.path.isfile(cfg_abs_path):
    ini_config.read(cfg_abs_path)

  cwd_abs_path = os.getcwd()

  config = Config(cwd_abs_path, ini_config)
  pkgs = PackageSet(config)

  cmd = args[0]
  args = args[1:]
  if cmd == 'clean':
    if args:
      args_parser.error('Should not specify targets for clean command')

    def remove(x):
      print 'Removing', x
      r = subprocess.call('rm -rf %s' % x, shell=True)
      assert r == 0

    remove(config.build_dir_abs_path)
    remove(config.genfile_dir_abs_path)
    print 'Done'
    return
  elif cmd == 'analyze':
    passes = [
        BindDependencies(pkgs),
        CheckCycles(),
        PrintBuildOrder(),
        ]
  elif cmd == 'build':
    passes = [
        BindDependencies(pkgs),
        CheckCycles(),
        BuildTargets(),
        ]
  elif cmd == 'test':
    passes = [
        BindDependencies(pkgs),
        CheckCycles(),
        BuildTargets(),
        TestTargets(),
        ]
  else:
    args_parser.error(
        'Invalid commend (%s).  Must be: analyze|build|test|clean')

  if not args:
    args_parser.error('Must specify targets for %s command' % cmd)

  current_pkg_path = '//'
  if cwd_abs_path.startswith(config.src_dir_abs_path + '/'):
    current_pkg_path = config.src_abs_path_to_pkg_path(cwd_abs_path)

  patterns = TargetPatterns(current_pkg_path)

  pattern_strs = sanitize_pattern_strs(config, cwd_abs_path, args)
  patterns.set_patterns(pattern_strs)

  targets = patterns.get_matching_targets(pkgs)
  for p in passes:
    p.run(targets)


if __name__ == '__main__':
  main()
