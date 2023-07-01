Naive / (mostly) canonical lr parser generator.

This deviates slightly from the dragon book's LR(1) presentation:

1. The start production is item is [#accept -> ^ . S, $] instead of
   [#accept -> . S, $].  This avoids the need to special case the start
   state's kernel.

2. The token stream act as pseudo symbol stack.  Instead of relying on
   a secondary GOTO table for reduction, the goto entries are encoded as
   shift actions in the ACTION table.  The reduction action will pop
   n symbols from state stack, and push the reduced symbol to the front
   of the pseudo symbol stack.  The parser then perform a second lookup
   to shift the reduced symbol.

3. This uses the LALR "core union" idea to merge states.  However, unlike
   LALR, this merges two states iff the merge does not introduce any new
   (shift/reduce, reduce/reduce) conflicts, i.e., we allow multiple copies of
   states with the same core.

4. For each (conflict-free) state, we replace the most used core production
   items of the form [A -> B . , x] with a single [A -> B . , *] entry.  The
   * indicates the look ahead symbol can be anything. Internally, an ACTION
   table lookup will first attempt to look up the look ahead specific action,
   and if that fails, fall back to looking up the wildcard action.

5. For convenience, we allow %start to accept a list of production names
   instead of a single production name.  The generated state graph will
   have one start state for each specified production.

The generator's parser is self bootstrapped from ./internal/parser/grammar.lr

---------

Note on parser state size:

With the optimizations listed above, it looks like we can generate a fully
LR(1) compliant parser using a roughly the same number of states as LALR, but
at the expense of more shift action entries.

For comparison, the following are some stats generated from the ansi c's grammar
(in ./test/ansi_c)

Canonical LR1:
Number of states: 1572
Number of shift actions: 14481
Number of reduce actions: 15182

With Core Union (#3 above):
Number of states: 351
Number of shift actions: 2989
Number of reduce actions: 4084

With Core Union And Reduction Action Merging (#3 and #4 above):
Number of states: 351
Number of shift actions: 2989
Number of reduce actions: 246

YACC / LALR:
Number of states: 349
Number of shift + goto actions: 1702 + 238 = 1940
Number of reduce actions: ?
