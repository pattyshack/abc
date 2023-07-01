from buildutil.rules.cc_rules import (
    CcBinaryTargetRule,
    CcLibraryTargetRule,
    CcTestTargetRule,
    )
from buildutil.rules.go_rules import (
    GoPkgTargetRule,
    )
from buildutil.rules.lex_yacc_rules import (
    LexTargetRule,
    YaccTargetRule,
    )
from buildutil.rules.nasm_rules import (
    NasmTargetRule,
    )
from buildutil.rules.py_rules import (
    PyBinaryTargetRule,
    PyLibraryTargetRule,
    PyTestTargetRule,
    )

# list of TargetRule subclasses.
RULES = [
    CcBinaryTargetRule,
    CcLibraryTargetRule,
    CcTestTargetRule,

    GoPkgTargetRule,

    LexTargetRule,
    YaccTargetRule,

    NasmTargetRule,

    PyBinaryTargetRule,
    PyLibraryTargetRule,
    PyTestTargetRule,
    ]
