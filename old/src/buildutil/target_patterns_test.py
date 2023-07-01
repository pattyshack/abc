import unittest

from buildutil.target_patterns import validate_target_pattern


class TargetPatternTestCase(unittest.TestCase):
  def test_target_regex(self):
    is_valid = lambda x: self.assertTrue(validate_target_pattern(x))
    not_valid = lambda x: self.assertFalse(validate_target_pattern(x))

    is_valid('...')
    is_valid('bar/...')
    is_valid('foo/bar/...')
    is_valid('//...')
    is_valid('//foo/...')
    is_valid('//foo/bar/...')

    is_valid(':all')
    is_valid('foo:all')
    is_valid('foo/bar:all')
    is_valid('//:all')
    is_valid('//foo:all')
    is_valid('//foo/bar:zzz')

    not_valid('//asdf/../foo')
    not_valid('/...')
    not_valid('/bar/...')
    not_valid('/bar:all')

if __name__ == '__main__':
    unittest.main()
