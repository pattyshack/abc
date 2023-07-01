cc_library( name='lib1',
    srcs=[
        'lib1a.cc',
        'lib1b.cc',
        ],
    hdrs=['lib1.h'])


cc_library(
    name='lib2',
    srcs=['lib2.cc'],
    hdrs=['lib2.h'],
    deps=[
        ':lib1',
        ':tmpl',
        ])

cc_library(
    name='tmpl',
    hdrs=['tmpl_inl.h'])

cc_binary(
    name='hello_world',
    srcs=[
        'hello.cc',
        'world.cc',
        ],
    hdrs=['hello_world.h'],
    deps=[':lib2'])

cc_test(
    name='foo_unittest',
    srcs=['foo_unittest.cc'],
    deps=['//gtest:gtest'])

gen_lex(
    name='basic_scanner.l',
    srcs=['basic_scanner.l'])

cc_binary(
    name='basic_scanner',
    srcs=['basic_scanner.yy.c'],
    deps=[':basic_scanner.l'])

gen_yacc(
    name='test_parser.y',
    srcs=['test_parser.y'])

gen_lex(
    name='test_scanner.l',
    srcs=['test_scanner.l'],
    deps=[':test_parser.y'])

cc_binary(
    name='test_parser',
    srcs=[
        'test_scanner.yy.c',
        'test_parser.tab.cc',
        'test_parser.cc',
        ],
    hdrs=[
        'location.hh',
        'stack.hh',
        'position.hh',
        'test_parser.tab.hh',
        'test_scanner.h'
        ],
    deps=[
        ':test_parser.y',
        ':test_scanner.l',
        ])

gen_nasm(
    name='boot',
    src='boot.asm',
    includes=['print_string.asm'])
