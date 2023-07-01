package lr

import (
	"sort"
	"strings"
)

const (
	AcceptRule  = "#accept"
	StartMarker = "^"
	EndMarker   = "$"
	Wildcard    = "*"
)

// LR(1) production item of the form: A -> [a B] . [c] , x
type Item struct {
	*Term
	*Clause

	Dot int

	LookAhead string

	IsReduce bool
	IsKernel bool

	Key  string
	Next *Item
	Core *Item

	Pool *itemPool
}

func (item *Item) Shift() *Item {
	if item.Next == nil {
		item.Next = item.Pool.Get(
			item.Term,
			item.Clause,
			item.Dot+1,
			item.LookAhead)
	}

	return item.Next
}

func (item *Item) ReplaceLookAhead(symbol string) *Item {
	if symbol == item.LookAhead {
		return item
	}

	return item.Pool.Get(
		item.Term,
		item.Clause,
		item.Dot,
		symbol)
}

func (item *Item) String() string {
	if item.Key == "" {
		result := ""
		if item.LookAhead != "" {
			result = item.Core.String() + ", " + item.LookAhead
		} else {
			result += item.Term.Name + ":"

			for idx, term := range item.Clause.Bindings {
				if idx == item.Dot {
					result += "." + term.Name
				} else {
					result += " " + term.Name
				}
			}

			if item.Dot == len(item.Clause.Bindings) {
				result += "."
			} else {
				result += " "
			}
		}

		item.Key = result
	}

	return item.Key
}

// sort order:
//  kernel items before non-kernel items
//  clause sort id
//  dot position
//  lookahead string
func (this *Item) Compare(other *Item) int {
	if this.IsKernel {
		if !other.IsKernel {
			return -1
		}
	} else if other.IsKernel {
		return 1
	}

	if this.SortId != other.SortId {
		if this.SortId < other.SortId {
			return -1
		}
		return 1
	}

	if this.Dot != other.Dot {
		if this.Dot < other.Dot {
			return -1
		}
		return 1
	}

	if this.LookAhead == other.LookAhead {
		return 0
	}

	if this.LookAhead < other.LookAhead {
		return -1
	}

	return 1
}

type itemPoolKey struct {
	ClauseId  int
	Dot       int
	LookAhead string
}

type itemPool struct {
	items map[itemPoolKey]*Item
}

func newItemPool() *itemPool {
	return &itemPool{map[itemPoolKey]*Item{}}
}

func (pool *itemPool) Get(
	term *Term,
	clause *Clause,
	dot int,
	lookAhead string) *Item {

	key := itemPoolKey{clause.SortId, dot, lookAhead}
	item, ok := pool.items[key]
	if !ok {
		item = &Item{
			Term:      term,
			Clause:    clause,
			Dot:       dot,
			LookAhead: lookAhead,
			IsReduce:  dot == len(clause.Bindings),
			IsKernel:  dot != 0,
			Pool:      pool,
		}

		if lookAhead != "" {
			item.Core = pool.Get(term, clause, dot, "")
		} else {
			item.Core = item
		}

		pool.items[key] = item
	}

	return item
}

type Items []*Item

func (items Items) String() string {
	chunks := make([]string, 0, len(items))
	for _, item := range items {
		chunks = append(chunks, item.String())
	}

	return strings.Join(chunks, ";")
}

func (items Items) Len() int {
	return len(items)
}

func (items Items) Less(iIdx int, jIdx int) bool {
	return items[iIdx].Compare(items[jIdx]) < 0
}

func (items Items) Swap(i int, j int) {
	items[i], items[j] = items[j], items[i]
}

type stateAction struct {
	shift  *ItemSet
	reduce *Item
}

type ItemSet struct {
	Kernel      string
	KernelItems Items // sorted
	Items             // sorted

	StateNum int

	Goto map[string]*ItemSet

	Reduce map[string]Items

	ReduceReduceConflictSymbols []string
	ShiftReduceConflictSymbols  []string

	CodeGenConst  string
	CodeGenAction string
}

func newItemSet(kernelItems Items) *ItemSet {
	sort.Sort(kernelItems)
	return &ItemSet{
		Kernel:      kernelItems.String(),
		KernelItems: kernelItems,
		Items:       kernelItems,
		Goto:        map[string]*ItemSet{},
	}
}

func (set *ItemSet) canMergeFrom(other *ItemSet) bool {
	if len(set.ReduceReduceConflictSymbols) > 0 ||
		len(set.ShiftReduceConflictSymbols) > 0 {

		return false
	}

	if len(other.ReduceReduceConflictSymbols) > 0 ||
		len(other.ShiftReduceConflictSymbols) > 0 {

		return false
	}

	actions := make(map[string]*stateAction, len(set.Reduce)+len(set.Goto))
	for symbol, items := range set.Reduce {
		if len(items) > 1 {
			return false // don't merge state with reduce/reduce error
		}
		actions[symbol] = &stateAction{reduce: items[0]}
	}

	for symbol, next := range set.Goto {
		actions[symbol] = &stateAction{shift: next}
	}

	for symbol, items := range other.Reduce {
		if len(items) > 1 {
			return false // don't merge state with reduce/reduce error
		}

		action, ok := actions[symbol]
		if !ok {
			continue
		}

		if action.shift != nil {
			// This introduces new shift/reduce error.  In theory, this should
			// never happen
			return false
		}

		if action.reduce.Compare(items[0]) != 0 {
			return false // This introduces new reduce/reduce error
		}
	}

	for symbol, next := range set.Goto {
		action, ok := actions[symbol]
		if !ok {
			continue
		}

		if action.reduce != nil {
			// This introduces new shift/reduce error.  In theory, this should
			// never happen
			return false
		}

		if action.shift.StateNum != next.StateNum {
			// Don't merge different state machines
			return false
		}
	}

	return true
}

func (set *ItemSet) mergeFrom(other *ItemSet) {
	kernelCount := 0
	items := make(Items, 0, len(set.Items)+len(other.Items))
	thisIdx := 0
	otherIdx := 0
	for thisIdx < len(set.Items) && otherIdx < len(other.Items) {
		thisItem := set.Items[thisIdx]
		otherItem := other.Items[otherIdx]

		var toAdd *Item
		cmp := thisItem.Compare(otherItem)
		if cmp < 0 {
			toAdd = thisItem
			thisIdx += 1
		} else if cmp == 0 {
			toAdd = thisItem
			thisIdx += 1
			otherIdx += 1
		} else {
			toAdd = otherItem
			otherIdx += 1
		}

		if toAdd.IsKernel {
			kernelCount += 1
		}

		items = append(items, toAdd)
	}

	kernelItems := items[:kernelCount]
	set.Kernel = kernelItems.String()
	set.KernelItems = kernelItems
	set.Items = items

	for symbol, next := range other.Goto {
		set.Goto[symbol] = next
	}

	for symbol, items := range other.Reduce {
		set.Reduce[symbol] = items
	}
}

func (set *ItemSet) clone() *ItemSet {
	gotoMap := make(map[string]*ItemSet, len(set.Goto))
	for symbol, state := range set.Goto {
		gotoMap[symbol] = state
	}

	reduce := make(map[string]Items, len(set.Reduce))
	for symbol, items := range set.Reduce {
		reduce[symbol] = items
	}

	newItem := *set
	newItem.Goto = gotoMap
	newItem.Reduce = reduce

	return &newItem
}

func (set *ItemSet) computeConflictSymbols() {
	shiftReduce := []string{}
	reduceReduce := []string{}
	for symbol, items := range set.Reduce {
		_, ok := set.Goto[symbol]
		if ok {
			shiftReduce = append(shiftReduce, symbol)
		}

		if len(items) > 1 {
			// NOTE: In theory, we can handle reducing multiple productions
			// as long as neither production rule is a prefix of the other
			// production rule.  However, this require inspecting the stack
			// which is kind of painful to code gen.
			reduceReduce = append(reduceReduce, symbol)
		}
	}

	set.ShiftReduceConflictSymbols = shiftReduce
	set.ReduceReduceConflictSymbols = reduceReduce
}

func (set *ItemSet) compressShiftItemsAndSort() {
	if len(set.ShiftReduceConflictSymbols) > 0 ||
		len(set.ReduceReduceConflictSymbols) > 0 {

		return // Don't compress error state to output more debug info
	}

	added := map[string]struct{}{}
	kernelCount := 0
	items := make(Items, 0, len(set.Items))
	for _, item := range set.Items {
		toAdd := item
		if !item.IsReduce {
			toAdd = item.ReplaceLookAhead("")
		}

		_, ok := added[toAdd.String()]
		if ok {
			continue
		}
		added[toAdd.String()] = struct{}{}

		if toAdd.IsKernel {
			kernelCount += 1
		}

		items = append(items, toAdd)
	}

	// Note that the kernel items are already sorted
	sort.Sort(items[kernelCount:])

	kernelItems := items[:kernelCount]
	set.Kernel = kernelItems.String()
	set.KernelItems = kernelItems
	set.Items = items
}

func (set *ItemSet) compress() {
	if len(set.ShiftReduceConflictSymbols) > 0 ||
		len(set.ReduceReduceConflictSymbols) > 0 {

		return // Don't compress error state to output more debug info
	}

	counts := make(map[string]int, len(set.Items))
	for _, item := range set.Items {
		if item.IsReduce {
			counts[item.Core.String()] += 1
		}
	}

	max := 0
	maxKey := ""
	for key, count := range counts {
		if count > max {
			max = count
			maxKey = key
		}
	}

	if max == 1 {
		maxKey = ""
	}

	added := map[string]struct{}{}
	kernelCount := 0
	items := make(Items, 0, len(set.Items))
	reduce := make(map[string]Items, len(set.Reduce))
	for _, item := range set.Items {
		var toAdd *Item
		if item.IsReduce {
			toAdd = item
			if item.Core.String() == maxKey {
				toAdd = item.ReplaceLookAhead(Wildcard)
			}

		} else {
			// Shift production does not depend on look ahead symbols
			toAdd = item.ReplaceLookAhead("")
		}

		_, ok := added[toAdd.String()]
		if ok {
			continue
		}
		added[toAdd.String()] = struct{}{}

		if toAdd.IsKernel {
			kernelCount += 1
		}

		if toAdd.IsReduce {
			reduce[toAdd.LookAhead] = append(reduce[toAdd.LookAhead], toAdd)
		}

		items = append(items, toAdd)
	}

	kernelItems := items[:kernelCount]
	set.Kernel = kernelItems.String()
	set.KernelItems = kernelItems
	set.Items = items

	set.Reduce = reduce
}

type firstTermsEntry struct {
	hasNil    bool
	terminals []string
}

type LRStates struct {
	*Grammar

	FirstTerms map[string]*firstTermsEntry

	ItemPool *itemPool

	States        map[string]*ItemSet
	OrderedStates []*ItemSet

	ShiftReduceConflictsCount  int
	ReduceReduceConflictsCount int
}

func (states *LRStates) maybeAdd(state *ItemSet) (*ItemSet, bool) {
	if state.Kernel == "" {
		return nil, false
	}

	origState, ok := states.States[state.Kernel]
	if ok {
		return origState, false
	}

	states.States[state.Kernel] = state
	states.OrderedStates = append(states.OrderedStates, state)
	state.StateNum = len(states.OrderedStates)

	return state, true
}

func (states *LRStates) populateStartStates() {
	acceptTerm := &Term{
		Name: AcceptRule,
	}
	startTerm := &Term{
		Name:       StartMarker,
		IsTerminal: true,
	}

	for idx, start := range states.Starts {
		item := states.ItemPool.Get(
			acceptTerm,
			&Clause{
				SortId:   -(len(states.Starts)-idx - 1),
				Bindings: []*Term{startTerm, start},
			},
			1,
			EndMarker)
		states.maybeAdd(newItemSet(Items{item}))
	}
}

func (states *LRStates) generateStates() {
	symbols := make([]string, 0, len(states.Terms))
	for symbol, _ := range states.Terms {
		symbols = append(symbols, symbol)
	}
	sort.Strings(symbols)

	exploredIdx := 0
	for exploredIdx < len(states.OrderedStates) {
		for _, state := range states.OrderedStates[exploredIdx:] {
			states.populateClosure(state)

			shiftableItems := make(map[string]Items, len(state.Items))
			for _, item := range state.Items {
				if item.IsReduce {
					continue
				}

				nextSymbol := item.Clause.Bindings[item.Dot].Name
				shiftableItems[nextSymbol] = append(
					shiftableItems[nextSymbol],
					item)
			}

			// This is needed to keep runs deterministic
			for _, symbol := range symbols {
				items, ok := shiftableItems[symbol]
				if !ok {
					continue
				}

				gotoKernelItems := make(Items, 0, len(items))
				for _, item := range items {
					gotoKernelItems = append(gotoKernelItems, item.Shift())
				}

				gotoState, _ := states.maybeAdd(newItemSet(gotoKernelItems))
				if gotoState == nil {
					continue
				}

				nextState := state.Goto[symbol]
				if nextState == nil {
					state.Goto[symbol] = gotoState
				} else if nextState != gotoState {
					panic("This should never happen")
				}
			}

			state.computeConflictSymbols()

			// NOTE: it's safe to drop the lookahead symbol from shift items
			// once the goto states for the current states are initialized
			// since shifting does not depend on the lookahead symbol.
			state.compressShiftItemsAndSort()

			exploredIdx += 1
		}
	}
}

func (states *LRStates) populateClosure(state *ItemSet) {
	type checkedEntry struct {
		rule     string
		terminal string
	}
	checked := map[checkedEntry]struct{}{}
	added := map[string]struct{}{}

	var workspace []string

	toExplore := state.Items
	for len(toExplore) > 0 {
		nextToExplore := Items{}

		for _, item := range toExplore {
			if item.IsReduce {
				continue
			}

			rule := item.Clause.Bindings[item.Dot]
			if rule.IsTerminal {
				continue
			}

			terminals := states.firstTerminals(item, workspace)

			for _, terminal := range terminals {
				key := checkedEntry{rule.Name, terminal}
				_, ok := checked[key]
				if ok {
					continue
				}
				checked[key] = struct{}{}

				for _, clause := range rule.Clauses {
					item := states.ItemPool.Get(rule, clause, 0, terminal)

					_, ok := added[item.String()]
					if !ok {
						added[item.String()] = struct{}{}
						state.Items = append(state.Items, item)
						nextToExplore = append(nextToExplore, item)
					}
				}
			}

			if len(workspace) < len(terminals) {
				workspace = terminals
			}
		}

		toExplore = nextToExplore
	}

	reduce := map[string]Items{}
	for _, item := range state.Items {
		if item.IsReduce {
			reduce[item.LookAhead] = append(reduce[item.LookAhead], item)
		}
	}
	state.Reduce = reduce
}

// NOTE: This works well in practice, but in degenerated case, where all the
// terms have nil first terminal, the result list could be much longer than
// the equvialent set.
func (states *LRStates) firstTerminals(
	item *Item,
	workspace []string) []string {

	result := workspace[:0]

	for _, term := range item.Clause.Bindings[item.Dot+1:] {
		entry := states.FirstTerms[term.Name]
		if entry == nil {
			continue
		}

		result = append(result, entry.terminals...)
		if !entry.hasNil {
			return result
		}
	}

	entry := states.FirstTerms[item.LookAhead]
	if entry == nil || entry.hasNil {
		panic("Shouldn't reach here: " + item.String())
	}

	return append(result, entry.terminals...)
}

func (states *LRStates) computeFirstTerminals() {
	firstTerms := map[string]map[string]struct{}{
		"$": map[string]struct{}{
			"$": struct{}{},
		},
	}

	for _, term := range states.Terms {
		set := map[string]struct{}{}

		firstTerms[term.Name] = set
		if term.IsTerminal {
			set[term.Name] = struct{}{}
		} else {
			for _, clause := range term.Clauses {
				if len(clause.Bindings) == 0 {
					set[""] = struct{}{}
				}
			}
		}
	}

	modified := true
	for modified {
		modified = false

		for _, rule := range states.NonTerminals {
			set := firstTerms[rule.Name]
			add := func(symbol string) {
				_, ok := set[symbol]
				if !ok {
					modified = true
					set[symbol] = struct{}{}
				}
			}

			for _, clause := range rule.Clauses {
				nilTermCount := 0
				for _, term := range clause.Bindings {
					hasNil := false
					for symbol, _ := range firstTerms[term.Name] {
						if symbol != "" {
							add(symbol)
						} else {
							hasNil = true
							nilTermCount += 1
						}
					}

					if !hasNil {
						break
					}
				}

				if nilTermCount == len(clause.Bindings) {
					add("")
				}
			}
		}
	}

	states.FirstTerms = make(map[string]*firstTermsEntry, len(firstTerms))
	for symbol, set := range firstTerms {
		hasNil := false
		terms := make([]string, 0, len(set))
		for term, _ := range set {
			if term == "" {
				hasNil = true
			} else {
				terms = append(terms, term)
			}
		}

		states.FirstTerms[symbol] = &firstTermsEntry{hasNil, terms}
	}
}

func (states *LRStates) mergeStates() {
	modified := true
	for modified {
		modified = false

		coreKernels := make([]string, 0, len(states.OrderedStates))
		mergeCandidates := make(map[string][]*ItemSet, len(states.States))

		for _, state := range states.OrderedStates {
			added := map[string]struct{}{}
			kernelItems := make(Items, 0, len(state.KernelItems))

			for _, item := range state.KernelItems {
				_, ok := added[item.Core.String()]
				if ok {
					continue
				}
				added[item.Core.String()] = struct{}{}

				kernelItems = append(kernelItems, item.Core)
			}

			// NOTE: no need to sort since item.KernelItems is already sorted
			kernelString := kernelItems.String()

			_, ok := mergeCandidates[kernelString]
			if !ok {
				coreKernels = append(coreKernels, kernelString)
			}

			mergeCandidates[kernelString] = append(
				mergeCandidates[kernelString],
				state)
		}
		newStates := make([]*ItemSet, 0, len(states.OrderedStates))

		// lr state kernel -> merged state
		stateMapping := make(map[string]*ItemSet, len(states.States))

		// NOTE: iterate over coreKernels to preserve state ordering
		for _, coreKernel := range coreKernels {
			candidates := mergeCandidates[coreKernel]

			mergedStates := make([]*ItemSet, 0, len(candidates))
			for _, candidate := range candidates {
				merged := false
				for _, mergedState := range mergedStates {
					if mergedState.canMergeFrom(candidate) {
						mergedState.mergeFrom(candidate)

						stateMapping[candidate.Kernel] = mergedState
						merged = true
						break
					}
				}

				if !merged {
					newState := candidate.clone()

					mergedStates = append(mergedStates, newState)
					newStates = append(newStates, newState)

					stateMapping[candidate.Kernel] = newState
				}
			}
		}

		for idx, state := range newStates {
			newGoto := make(map[string]*ItemSet, len(state.Goto))
			for symbol, next := range state.Goto {
				newGoto[symbol] = stateMapping[next.Kernel]
			}
			state.Goto = newGoto

			state.StateNum = idx + 1
		}

		if len(states.OrderedStates) > len(newStates) {
			modified = true

			states.OrderedStates = newStates

			states.States = make(map[string]*ItemSet, len(newStates))
			for _, state := range newStates {
				states.States[state.Kernel] = state
			}
		}
	}
}

func (states *LRStates) shuffleAcceptStates() {
	orderedStates := states.OrderedStates[:len(states.Starts)]

	acceptStates := []*ItemSet{}
	otherStates := []*ItemSet{}
	for _, state := range states.OrderedStates[len(states.Starts):] {
		isAcceptState := false
		for _, item := range state.KernelItems {
			if item.Name == AcceptRule && item.LookAhead == EndMarker {
				isAcceptState = true
				break
			}
		}
		if isAcceptState {
			acceptStates = append(acceptStates, state)
		} else {
			otherStates = append(otherStates, state)
		}
	}

	orderedStates = append(orderedStates, acceptStates...)
	orderedStates = append(orderedStates, otherStates...)

	for idx, state := range orderedStates {
		state.StateNum = idx + 1
	}

	states.OrderedStates = orderedStates
}

func NewLRStates(grammar *Grammar) *LRStates {
	states := &LRStates{
		Grammar:  grammar,
		States:   map[string]*ItemSet{},
		ItemPool: newItemPool(),
	}

	states.computeFirstTerminals()
	states.populateStartStates()
	states.generateStates()

	states.mergeStates()

	states.shuffleAcceptStates()

	shiftReduceCount := 0
	reduceReduceCount := 0
	for _, state := range states.OrderedStates {
		state.compress()

		shiftReduceCount += len(state.ShiftReduceConflictSymbols)
		reduceReduceCount += len(state.ReduceReduceConflictSymbols)
	}

	states.ShiftReduceConflictsCount = shiftReduceCount
	states.ReduceReduceConflictsCount = reduceReduceCount

	return states
}
