package internal

import (
	"math/rand"
	"reflect"
	"testing/quick"

	"github.com/google/go-cmp/cmp"

	"testing"
)

func chain(tokens []*Token) *Token {
	head := &Token{}
	curt := head
	for _, tok := range tokens {
		curt.Next = tok
		curt = curt.Next
	}
	return head.Next
}

func TestTokenize(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect []*Token
	}{
		{"lrb",
			"(",
			[]*Token{
				{Kind: Lrb, Position: Position{StartedAt: 0, Line: 0}},
				{Kind: Eof, Position: Position{StartedAt: 1, Line: 0}},
			},
		},
		{"rb",
			"()",
			[]*Token{
				{Kind: Lrb, Position: Position{StartedAt: 0, Line: 0}},
				{Kind: Rrb, Position: Position{StartedAt: 1, Line: 0}},
				{Kind: Eof, Position: Position{StartedAt: 2, Line: 0}},
			},
		},
		{"comment",
			";hello",
			[]*Token{
				{Kind: Comment, Raw: []rune(";hello"), Position: Position{StartedAt: 0, Line: 0}},
				{Kind: Eof, Position: Position{StartedAt: len(";hello"), Line: 0}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tok, err := Tokenize([]rune(tt.input))
			if err != nil {
				t.Error(err)
			}
			if diff := cmp.Diff(chain(tt.expect), tok); diff != "" {
				t.Errorf("diff: %s", diff)
			}
		})
	}
}

func TestTokenize_RoundTrip(t *testing.T) {
	f := func(s string) bool {
		tokens, err := Tokenize([]rune(s))
		if err != nil {
			return false
		}
		last := tokens
		for last.Next != nil {
			last = last.Next
		}
		return last.Kind == Eof
	}

	corpus := []string{"(", ")", ";hi", "()", "(;x)", "", "\n", "((()))", "123", "-123", "+9-12098*898123"}

	cfg := &quick.Config{
		MaxCount: 300,
		Rand:     rand.New(rand.NewSource(3)),
		Values: func(values []reflect.Value, r *rand.Rand) {
			s := corpus[r.Intn(len(corpus))]
			values[0] = reflect.ValueOf(s)
		},
	}
	if err := quick.Check(f, cfg); err != nil {
		t.Error(err)
	}
}
