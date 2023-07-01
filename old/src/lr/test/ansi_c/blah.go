package ansi_c

type Symbol string

func (Symbol) Id() CSymbolId { return 0 }

func (Symbol) Loc() CLocation { return CLocation{} }
