Template compiler for go, with template syntax closer to golang


Template Body Directives:
-------------
[[# ... ]]
 - template comment.  The comment is not emmitted to the generated code.

[[$ <content> ]]
 - copy content directly into the generated template code.  used for var decl,
   assignment, etc

[[for <loop text>]] <body> [[end]]
 - the loop text is copied verbatim

[[switch <switch text> ]] [[case <case text>]] <body> [[default]] <body> [[end]]
 - the switch text and case texts are copied verbatim
 - note that the whitespaces between [[switch]] and the first [[case]] is
   always trimmed from the generated output.  Non-whitespace text between
   [[switch]] and the first [[case]] will result in error.

[[if <predicate>]] <body> [[else if <predicate>]] <body> [[else]] <body> [[end]]
 - the predicates are copied verbatim

[[continue]]

[[break]]
 - Note that this does not support break label.  An escape hatch is to use
   [[$ ... ]] to define the loop label and to break to it

[[return]]
 - early normal exit

[[error <err>]]
 - error exit

[[embed <expr>]]
 - Embed a sub-template.  Expr must implement io.WriterTo

$<var>
 - output var.  var must be a valid output value type

$(<expr>)
 - output evaluated expr.  expr must evaluate to a valid output value type

$$
 - output $


Valid Output Value Type:
------------------------
fmt.Stringer        output as string formatted (%s)
string              output as string formatted (%s)
[]byte              output as string formatted (%s)
bool                output as bool formatted (%t)
uints, ints         output as int formatted (%d)
floats, complex     output as float/complex formatted (%g)

NOTE: since byte is alias for uint8 and rune is alias for int32, these are
output as integers.


Whitespace Triming:
-------------------
[[ ]] style directives may optionally specify '-' to trim leading / trailing
whitespaces:

[[- Trim leading whitespaces on the same line as the directive, as well as the
    previous line's '\n' if it's adjacent to those leading whitespaces.

-]] Trim trailing whitespaces on the same line as the directive, potentially
    up to and including the current line's '\n'


File Format:
------------
package <name>

// In the header section (everything before the section marker), both //
// and /* */ style comments are allowed

// Optional import.  The imports are copied verbatim.
import (
    ...
)

// Note:
// 1. unlike golang struct, this requires explicit argument name
// 2. arg and type must be on the same line (type can span multiple lines)
//
// XXX: maybe add support to specify directive start/end markers' characters.
template <name> {
    arg1 type1
    arg2 type2
    ...
}

// Everything after the section maker is part of the template
%%
<template>
