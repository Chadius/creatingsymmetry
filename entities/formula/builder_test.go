package formula_test

import (
	"github.com/Chadius/creating-symmetry/entities/formula"
	. "gopkg.in/check.v1"
	"reflect"
)

type BuilderTest struct {}

var _ = Suite(&BuilderTest{})

func (b *BuilderTest) TestIdentityFormula(checker *C) {
	identityFormula := formula.NewBuilder().Build()
	checker.Assert(reflect.TypeOf(identityFormula).String(), Equals, "*formula.Identity")
}