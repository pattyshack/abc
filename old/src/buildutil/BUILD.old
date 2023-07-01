py_library(
    name='analysis_passes',
    srcs=['analysis_passes.py'],
    deps=[
        ':config',
        ':package',
        ':topo_sorter',
        ])

py_library(
    name='config',
    srcs=['config.py'])

py_library(
    name='package',
    srcs=['package.py'],
    deps=[
        ':config',
        ':target_patterns',
        ':target_rule',
        ':util',
        'rules:rules',
        ])

py_library(
    name='target_patterns',
    srcs=['target_patterns.py'],
    deps=[':util'],
    visibility=['//buildutil/...'])

py_library(
    name='target_rule',
    srcs=['target_rule.py'],
    deps=[
        ':config',
        ':target_patterns',
        ':util',
        ],
    visibility=['//buildutil/...'])

py_library(
    name='topo_sorter',
    srcs=['topo_sorter.py'],
    deps=[':target_rule'])

py_library(
    name='util',
    srcs=['util.py'])

py_binary(
    name='buildutil',
    srcs=['main.py'],
    deps=[
        ':analysis_passes',
        ':package',
        ':target_patterns',
        ':target_rule',
        ':topo_sorter',
        ])

py_test(
    name='target_patterns_test',
    srcs=['target_patterns_test.py'],
    deps=[':target_patterns'])
