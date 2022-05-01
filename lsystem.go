// Package lsystemstr generates strings for L-systems (Lindenmayer system)

package lsystemstr

import (
	"errors"
	"fmt"
	"strings"
)

type Rule struct {
	p string
	s string
}

func (r Rule) String() string {
	return fmt.Sprintf("%s → %s", r.p, r.s)
}

// Predecessor returns the string to be replaced.
func (r *Rule) Predecessor() string {
	return r.p
}

// Successor returns the string that will replace the Predecessor.
func (r *Rule) Successor() string {
	return r.s
}

type Lsystem struct {
	axiom    string
	rules    []Rule
	sentence string
	r        *strings.Replacer
}

// New returns a new L-system with the starting axiom as the parameter.
func New(axiom string) Lsystem {
	return Lsystem{axiom: axiom, sentence: axiom}
}

// NewRule returns a new Rule to add to the L-System.
func NewRule(predecessor string, successor string) Rule {
	return Rule{p: predecessor, s: successor}
}

// CurrentRules returns a slice of the rules currently added to the L-system
func (l *Lsystem) CurrentRules() []Rule {
	return l.rules
}

// Sentence returns the current generation sentence.
func (l *Lsystem) Sentence() string {
	return l.sentence
}

// updateReplacer is a helper function to be called after modifying the rules
// of the L-system to a new strings.Replacer.
func (l *Lsystem) updateReplacer() {
	ruleStrings := make([]string, len(l.rules)*2)

	for i := 0; i < len(l.rules); i++ {
		ruleStrings[i*2] = l.rules[i].p
		ruleStrings[i*2+1] = l.rules[i].s
	}
	l.r = strings.NewReplacer(ruleStrings...)
}

// AddRules adds rules to the L-system.
func (l *Lsystem) AddRules(r ...Rule) {
	l.rules = append(l.rules, r...)
	l.updateReplacer()
}

// ReplaceRules overwrites the current rules in the L-system with new rules.
func (l *Lsystem) ReplaceRules(r ...Rule) {
	l.rules = r
	l.updateReplacer()
}

// AddRulesStr adds rules to the L-system from strings pairs.
// AddRulesStr errors if given an odd number of arguments.
func (l *Lsystem) AddRulesStr(s ...string) error {
	if len(s)%2 == 1 {
		return errors.New("odd argument count")
	}

	r := make([]Rule, 0, len(s)/2)

	for i := 0; i < len(s); i += 2 {
		r = append(r, NewRule(s[i], s[i+1]))
	}

	l.AddRules(r...)
	return nil
}

// Increment increments the L-system generation by 1.
// Alternative to calling Iterate(1)
func (l *Lsystem) Increment() {
	l.sentence = l.r.Replace(l.sentence)
}

// GetSuccessor returns the replacement string according to the systems rules for the given string.
// This does not change the internal state.
func (l *Lsystem) GetSuccessor(predecessor string) string {
	return l.r.Replace(predecessor)
}

// Iterate advances the L-system generation by n times.
func (l *Lsystem) Iterate(n uint) {
	for i := uint(0); i < n; i++ {
		l.Increment()
	}
}
