package greenspun

import(
	"fmt"
	"testing"
)

func TestPairString(t *testing.T) {
	ConfirmString := func(c *Pair, r string) {
		if s := c.String(); s != r {
			t.Fatalf("%v.String() should be %v", s, r)
		}
	}

	ConfirmString(List(), "()")
	ConfirmString(Cons(nil, nil), "(nil)")
	ConfirmString(Cons(0, nil), "(0)")
	ConfirmString(Cons(0, 1), "(0 . 1)")
	ConfirmString(Cons(0, Cons(1, 2)), "(0 1 . 2)")
	ConfirmString(Cons(0, Cons(1, Cons(2, 3))), "(0 1 2 . 3)")

	ConfirmString(List(0, 1, 2), "(0 1 2)")
	ConfirmString(List(0, 1, 2, 3), "(0 1 2 3)")
	ConfirmString(List(0, 1, 2, Cons(3, 4)), "(0 1 2 (3 . 4))")
	ConfirmString(Cons(List(0, 1, 2), Cons(3, 4)), "((0 1 2) 3 . 4)")
	l := List(0, 1, 2)
	l.End().Rplacd(Cons(3, 4))
	ConfirmString(l, "(0 1 2 3 . 4)")
	ConfirmString(List(0, List(1, 2, 3), List(2, 3), 3), "(0 (1 2 3) (2 3) 3)")
}

func TestPairEqual(t *testing.T) {
	ConfirmEqual := func(l *Pair, r interface{}, ok bool) {
		if x := l.Equal(r); x != ok {
			t.Fatalf("%v.Equal(%v) should be %v but is %v", l, r, ok, x)
		}
	}

	ConfirmEqual(&Pair{ &Pair{ 1, 1 }, nil }, &Pair{ nil, nil }, false)
	ConfirmEqual(&Pair{ nil, &Pair{ 1, 1 } }, &Pair{ nil, nil }, false)
	ConfirmEqual(&Pair{ nil, nil }, &Pair{ &Pair{ 1, 1 }, nil }, false)
	ConfirmEqual(&Pair{ nil, nil }, &Pair{ nil, &Pair{ 1, 1 } }, false)
	ConfirmEqual(&Pair{ &Pair{ 1, 1 }, nil }, &Pair{ &Pair{ 1, 1 }, nil }, true)
	ConfirmEqual(&Pair{ nil, &Pair{ 1, 1 } }, &Pair{ nil, &Pair{ 1, 1 } }, true)
	ConfirmEqual(&Pair{ &Pair{ 1, 1 }, &Pair{ 1, 1 } }, &Pair{ &Pair{ 1, 1 }, &Pair{ 1, 1 } }, true)

	ConfirmEqual(Cons(nil, nil), Pair{ nil, nil }, true)
	ConfirmEqual(Cons(nil, nil), &Pair{ nil, nil }, true)
	ConfirmEqual(Cons(nil, nil), Cons(nil, nil), true)

	ConfirmEqual(Cons(1, nil), Pair{ 1, nil }, true)
	ConfirmEqual(Cons(1, nil), &Pair{ 1, nil }, true)
	ConfirmEqual(Cons(1, nil), Cons(1, nil), true)

	ConfirmEqual(Cons(nil, 1), Pair{ nil, 1 }, true)
	ConfirmEqual(Cons(nil, 1), &Pair{ nil, 1 }, true)
	ConfirmEqual(Cons(nil, 1), Cons(nil, 1), true)

	ConfirmEqual(Cons(1, nil), Pair{ nil, nil }, false)
	ConfirmEqual(Cons(1, nil), &Pair{ nil, nil }, false)
	ConfirmEqual(Cons(1, nil), Cons(nil, nil), false)

	ConfirmEqual(Cons(nil, 1), Pair{ nil, nil }, false)
	ConfirmEqual(Cons(nil, 1), &Pair{ nil, nil }, false)
	ConfirmEqual(Cons(nil, 1), Cons(nil, nil), false)

	ConfirmEqual(Cons(nil, 1), Pair{ 1, nil }, false)
	ConfirmEqual(Cons(nil, 1), &Pair{ 1, nil }, false)
	ConfirmEqual(Cons(nil, 1), Cons(1, nil), false)

	ConfirmEqual(Cons(1, nil), Pair{ nil, 1 }, false)
	ConfirmEqual(Cons(1, nil), &Pair{ nil, 1 }, false)
	ConfirmEqual(Cons(1, nil), Cons(nil, 1), false)

	ConfirmEqual(Cons(nil, 1), Pair{ nil, 1 }, true)
	ConfirmEqual(Cons(nil, 1), &Pair{ nil, 1 }, true)
	ConfirmEqual(Cons(nil, 1), Cons(nil, 1), true)

	ConfirmEqual(Cons(Cons(0, 1), 2), Pair{ &Pair{ 0, 1 }, 2 }, true)
	ConfirmEqual(Cons(Cons(0, 1), 2), &Pair{ &Pair{ 0, 1 }, 2 }, true)
	ConfirmEqual(Cons(Cons(0, 1), 2), Cons( &Pair{ 0, 1 }, 2 ), true)

	ConfirmEqual(Cons(Cons(1, 1), nil), Cons(nil, nil), false)
	ConfirmEqual(Cons(nil, Cons(1, 1)), Cons(nil, nil), false)
	ConfirmEqual(Cons(nil, nil), Cons(Cons(1, 1), nil), false)
	ConfirmEqual(Cons(nil, nil), Cons(nil, Cons(1, 1)), false)
	ConfirmEqual(Cons(Cons(1, 1), nil), Cons(Cons(1, 1), nil), true)
	ConfirmEqual(Cons(nil, Cons(1, 1)), Cons(nil, Cons(1, 1)), true)
	ConfirmEqual(Cons(Cons(1, 1), Cons(1, 1)), Cons(Cons(1, 1), Cons(1, 1)), true)

	ConfirmEqual(List(1), &Pair{ nil, nil }, false)
	ConfirmEqual(List(1), &Pair{ 1, nil }, true)
	ConfirmEqual(List(1), Cons(1, nil), true)

	ConfirmEqual(List(1), List(), false)
	ConfirmEqual(List(1), nil, false)
	ConfirmEqual(List(1), Cons(nil, nil), false)

	ConfirmEqual(List(nil, 1), &Pair{ nil, &Pair{ 1, nil } }, true)
	ConfirmEqual(List(nil, 1), Cons(nil, Cons(1, nil)), true)

	ConfirmEqual(List(nil, 1), &Pair{ nil, nil }, false)
	ConfirmEqual(List(nil, 1), Cons(nil, nil), false)

	ConfirmEqual(List(Cons(0, 1), 2), &Pair{ &Pair{ 0, 1 }, 2 }, false)
	ConfirmEqual(List(Cons(0, 1), 2), Cons(&Pair{ 0, 1 }, 2), false)
	ConfirmEqual(List(Cons(0, 1), 2), Cons(Cons(0, 1), 2), false)
	ConfirmEqual(List(Cons(0, 1), 2), List(Cons(0, 1), 2), true)
	ConfirmEqual(List(Cons(0, 1), 2), Cons(Cons(0, 1), Cons(2, nil)), true)


	ConfirmEqual(List(nil, 1), &Pair{ 1, nil }, false)
	ConfirmEqual(List(nil, 1), &Pair{ nil, 1 }, false)
	ConfirmEqual(List(nil, 1), Cons(nil, 1), false)
	ConfirmEqual(List(nil, 1), &Pair{ nil, &Pair{ 1, nil} }, true)

	ConfirmEqual(List(nil, 1), &Pair{ nil, 1 }, false)
	ConfirmEqual(List(nil, 1), &Pair{ nil, &Pair{ 1, nil } }, true)
	ConfirmEqual(List(nil, 1), &Pair{ nil, &Pair{ 1, nil } }, true)

	ConfirmEqual(List(nil, 1), &Pair{ nil, &Pair{ 1, nil } }, true)
	ConfirmEqual(List(1, 1), &Pair{ 1, &Pair{ 1, nil } }, true)
}

func TestPairLen(t *testing.T) {
	ConfirmLen := func(c *Pair, r int) {
		if l := c.Len(); l != r {
			t.Fatalf("%v.Len() should be %v but is %v", c, r, l)
		}
	}

	ConfirmLen(nil, 0)
	ConfirmLen(List(), 0)
	ConfirmLen(&Pair{}, 1)
	ConfirmLen(Cons(nil, nil), 1)
	ConfirmLen(Cons(0, nil), 1)
	ConfirmLen(List(0), 1)
	ConfirmLen(Cons(0, 1), 1)
	ConfirmLen(List(0, 1), 2)
	ConfirmLen(List(0, 1, 2), 3)
}

func TestPairIsNil(t *testing.T) {
	ConfirmIsNil := func(c *Pair, r bool) {
		if n := c.IsNil(); n != r {
			t.Fatalf("%v.IsNil() should be %v but is %v", c, r, n)
		}
	}

	ConfirmIsNil(nil, true)
	ConfirmIsNil(Cons(nil, nil), false)
	ConfirmIsNil(Cons(1, nil), false)
	ConfirmIsNil(List(), true)
	ConfirmIsNil(List(nil), false)
}

func TestPairPush(t *testing.T) {
	ConfirmPush := func(c *Pair, v interface{}, r *Pair) {
		if x := c.Push(v); !x.Equal(r) {
			t.Fatalf("%v.Push(%v) should be %v but is %v", c, v, r, x)
		}
	}

	ConfirmPush(nil, nil, List(nil))
	ConfirmPush(nil, List(), List(List()))
	ConfirmPush(nil, 1, List(1))

	ConfirmPush(List(), nil, List(nil))
	ConfirmPush(List(0), 1, List(1, 0))
}

func TestPairPop(t *testing.T) {
	RefutePop := func(c *Pair) {
		vc := fmt.Sprintf("%v", c)
		defer ConfirmPanic(t, "%v.Pop() should panic", vc)()
		c.Pop()
	}

	ConfirmPop := func(c *Pair, rv interface{}, r *Pair) {
		switch v, x := c.Pop(); {
		case !r.Equal(x):
			t.Fatalf("1. %v.Pop() should be %v, %v but is %v, %v", c, rv, r, v, x)
		case rv != v:
			t.Fatalf("2. %v.Pop() should be %v, %v but is %v, %v", c, rv, r, v, x)
		}
	}

	RefutePop(nil)
	RefutePop(List())
	ConfirmPop(List(1), 1, nil)
	ConfirmPop(List(1, 2), 1, List(2))
}

func TestIntPair(t *testing.T) {
	ConfirmIntPair := func(c *Pair, l, r int) {
		if x, y := c.IntPair(); x != l || y != r {
			t.Fatalf("%v.IntPair() should be (%v, %v) but is (%v, %v)", c, l, r, x, y)
		}
	}

	ConfirmIntPair(Cons(0, 0), 0, 0)
	ConfirmIntPair(Cons(0, 1), 0, 1)
	ConfirmIntPair(Cons(1, 0), 1, 0)
}

func TestPairPair(t *testing.T) {
	ConfirmPairPair := func(c, l, r *Pair) {
		if x, y := c.PairPair(); !x.Equal(l) || !y.Equal(r) {
			t.Fatalf("%v.IntPair() should be (%v, %v) but is (%v, %v)", c, l, r, x, y)
		}
	}

	ConfirmPairPair(Cons(List(), List()), List(), List())
	ConfirmPairPair(Cons(List(0, 1), List(2, 3)), List(0, 1), List(2, 3))
}

func TestPairNext(t *testing.T) {
	RefuteNext := func(c *Pair) {
		vc := fmt.Sprintf("%v", c)
		defer ConfirmPanic(t, "%v.Next() should panic", vc)()
		c.Next()
	}

	ConfirmNext := func(c *Pair, r *Pair) {
		if x := c.Next(); !x.Equal(r) {
			t.Fatalf("%v.Next() should be %v but is %v", c, r, x)
		}
	}

	RefuteNext(nil)
	RefuteNext(List())
	ConfirmNext(List(0, 1), List(1))
}

func TestPairCar(t *testing.T) {
	RefuteCar := func(c *Pair) {
		vc := fmt.Sprintf("%v", c)
		defer ConfirmPanic(t, "%v.Car() should panic", vc)()
		c.Car()
	}

	ConfirmCar := func(c *Pair, r interface{}) {
		if car, ok := c.Car().(Equatable); ok {
			if !car.Equal(r) {
				t.Fatalf("%v.Car() should be %v but is %v", c, r, car)
			}
		} else if r, ok := r.(Equatable); ok {
			if !r.Equal(car) {
				t.Fatalf("%v.Car() should be %v but is %v", c, r, car)
			}
		} else {
			if car != r {
				t.Fatalf("%v.Car() should be %v but is %v", c, r, car)
			}
		}
	}

	RefuteCar(nil)
	RefuteCar(List())
	ConfirmCar(Cons(0, nil), 0)
	ConfirmCar(Cons(1, 0), 1)
	ConfirmCar(Cons(List(1), 0), Cons(1, nil))
	ConfirmCar(Cons(Cons(1, nil), 0), Cons(1, nil))
	ConfirmCar(Cons(Cons(2, 1), 0), Cons(2, 1))
	ConfirmCar(Cons(List(1, nil, nil), 0), Cons(1, Cons(nil, Cons(nil, nil))))
	ConfirmCar(Cons(List(1, nil, nil), 0), List(1, nil, nil))
}

func TestPairCdr(t *testing.T) {
	RefuteCdr := func(c *Pair) {
		vc := fmt.Sprintf("%v", c)
		defer ConfirmPanic(t, "%v.Cdr() should panic", vc)()
		c.Cdr()
	}

	ConfirmCdr := func(c *Pair, r interface{}) {
		if cdr, ok := c.Cdr().(Equatable); ok {
			if !cdr.Equal(r) {
				t.Fatalf("%v.Cdr() should be %v but is %v", c, r, cdr)
			}
		} else if r, ok := r.(Equatable); ok {
			if !r.Equal(cdr) {
				t.Fatalf("%v.Cdr() should be %v but is %v", c, r, cdr)
			}
		} else {
			if cdr != r {
				t.Fatalf("%v.Cdr() should be %v but is %v", c, r, cdr)
			}
		}
	}

	RefuteCdr(nil)
	RefuteCdr(List())
	ConfirmCdr(Cons(0, nil), nil)
	ConfirmCdr(Cons(0, 1), 1)
	ConfirmCdr(Cons(0, Cons(1, nil)), Cons(1, nil))
	ConfirmCdr(Cons(0, Cons(1, 2)), Cons(1, 2))
}

func TestPairCaar(t *testing.T) {
	RefuteCaar := func(c *Pair) {
		vc := fmt.Sprintf("%v", c)
		defer ConfirmPanic(t, "%v.Caar() should panic", vc)()
		c.Caar()
	}

	ConfirmCaar := func(c *Pair, r interface{}) {
		if caar, ok := c.Caar().(Equatable); ok {
			if !caar.Equal(r) {
				t.Fatalf("%v.Caar() should be %v but is %v", c, r, caar)
			}
		} else if r, ok := r.(Equatable); ok {
			if !r.Equal(caar) {
				t.Fatalf("%v.Caar() should be %v but is %v", c, r, caar)
			}
		} else {
			if caar != r {
				t.Fatalf("%v.Caar() should be %v but is %v", c, r, caar)
			}
		}
	}

	RefuteCaar(nil)
	RefuteCaar(Cons(0, nil))
	RefuteCaar(Cons(0, 1))
	ConfirmCaar(Cons(Cons(0, 1), nil), 0)
}

func TestPairCadr(t *testing.T) {
	RefuteCadr := func(c *Pair) {
		vc := fmt.Sprintf("%v", c)
		defer ConfirmPanic(t, "%v.Cadr() should panic", vc)()
		c.Cadr()
	}

	ConfirmCadr := func(c *Pair, r interface{}) {
		if cadr, ok := c.Cadr().(Equatable); ok {
			if !cadr.Equal(r) {
				t.Fatalf("%v.Cadr() should be %v but is %v", c, r, cadr)
			}
		} else if r, ok := r.(Equatable); ok {
			if !r.Equal(cadr) {
				t.Fatalf("%v.Cadr() should be %v but is %v", c, r, cadr)
			}
		} else {
			if cadr != r {
				t.Fatalf("%v.Cadr() should be %v but is %v", c, r, cadr)
			}
		}
	}

	RefuteCadr(nil)
	RefuteCadr(Cons(0, nil))
	RefuteCadr(Cons(0, 1))
	ConfirmCadr(Cons(Cons(0, 1), nil), 1)
}

func TestPairCdar(t *testing.T) {
	RefuteCdar := func(c *Pair) {
		vc := fmt.Sprintf("%v", c)
		defer ConfirmPanic(t, "%v.Cdar() should panic", vc)()
		c.Cdar()
	}

	ConfirmCdar := func(c *Pair, r interface{}) {
		if cdar, ok := c.Cdar().(Equatable); ok {
			if !cdar.Equal(r) {
				t.Fatalf("%v.Cdar() should be %v but is %v", c, r, cdar)
			}
		} else if r, ok := r.(Equatable); ok {
			if !r.Equal(cdar) {
				t.Fatalf("%v.Cdar() should be %v but is %v", c, r, cdar)
			}
		} else {
			if cdar != r {
				t.Fatalf("%v.Cdar() should be %v but is %v", c, r, cdar)
			}
		}
	}

	RefuteCdar(nil)
	RefuteCdar(Cons(0, nil))
	ConfirmCdar(Cons(0, Cons(1, nil)), 1)
	ConfirmCdar(Cons(0, Cons(1, 2)), 1)
}

func TestPairCddr(t *testing.T) {
	RefuteCddr := func(c *Pair) {
		vc := fmt.Sprintf("%v", c)
		defer ConfirmPanic(t, "%v.Cddr() should panic", vc)()
		c.Cddr()
	}

	ConfirmCddr := func(c *Pair, r interface{}) {
		if cddr, ok := c.Cddr().(Equatable); ok {
			if !cddr.Equal(r) {
				t.Fatalf("%v.Cddr() should be %v but is %v", c, r, cddr)
			}
		} else if r, ok := r.(Equatable); ok {
			if !r.Equal(cddr) {
				t.Fatalf("%v.Cddr() should be %v but is %v", c, r, cddr)
			}
		} else {
			if cddr != r {
				t.Fatalf("%v.Cddr() should be %v but is %v", c, r, cddr)
			}
		}
	}

	RefuteCddr(nil)
	RefuteCddr(Cons(0, nil))
	ConfirmCddr(Cons(0, Cons(1, nil)), nil)
	ConfirmCddr(Cons(0, Cons(1, 2)), 2)
}

func TestPairRplaca(t *testing.T) {
	ConfirmRplaca := func(c *Pair, v interface{}, r *Pair) {
		cs := c.String()
		c.Rplaca(v)
		if x := c.Equal(r); !x {
			t.Fatalf("%v.Rplaca(%v) should be %v but is %v", cs, v, r, c)
		}
	}

	ConfirmRplaca(Cons(nil, nil), 1, Cons(1, nil))
	ConfirmRplaca(Cons(Cons(0, 1), 2), 1, Cons(1, 2))
}

func TestPairRplacd(t *testing.T) {
	ConfirmRplacd := func(c *Pair, v interface{}, r *Pair) {
		cs := c.String()
		c.Rplacd(v)
		if x := c.Equal(r); !x {
			t.Fatalf("%v.Rplacd(%v) should be %v but is %v", cs, v, r, c)
		}
	}

	ConfirmRplacd(Cons(nil, nil), 1, Cons(nil, 1))
	ConfirmRplacd(Cons(0, Cons(1, 2)), 1, Cons(0, 1))
	ConfirmRplacd(Cons(Cons(0, 1), 2), 1, Cons(Cons(0, 1), 1))
}

func TestPairMove(t *testing.T) {
	RefuteMove := func(c *Pair, i int) {
		vc := fmt.Sprintf("%v", c)
		defer ConfirmPanic(t, "%v.Move(%v) should panic", vc, i)()
		c.Move(i)
	}

	ConfirmMove := func(c *Pair, i int, r interface{}) {
		if offset := c.Move(i); !offset.Equal(r) {
			t.Fatalf("%v.Move(%v) should be %v but is %v", c, i, r, offset)
		}
	}

	RefuteMove(nil, -1)
	RefuteMove(nil, 0)
	RefuteMove(nil, 1)

	RefuteMove(Cons(0, nil), -1)
	ConfirmMove(Cons(0, nil), 0, Cons(0, nil))
	RefuteMove(Cons(0, nil), 1)

	RefuteMove(Cons(0, Cons(1, nil)), -1)
	ConfirmMove(Cons(0, Cons(1, nil)), 0, Cons(0, Cons(1, nil)))
	ConfirmMove(Cons(0, Cons(1, nil)), 1, Cons(1, nil))
	RefuteMove(Cons(0, Cons(1, nil)), 2)

	RefuteMove(Cons(0, Cons(1, 2)), -1)
	ConfirmMove(Cons(0, Cons(1, 2)), 0, Cons(0, Cons(1, 2)))
	ConfirmMove(Cons(0, Cons(1, 2)), 1, Cons(1, 2))
	RefuteMove(Cons(0, Cons(1, 2)), 2)

	RefuteMove(Cons(0, Cons(1, Cons(2, 3))), -1)
	ConfirmMove(Cons(0, Cons(1, Cons(2, 3))), 0, Cons(0, Cons(1, Cons(2, 3))))
	ConfirmMove(Cons(0, Cons(1, Cons(2, 3))), 1, Cons(1, Cons(2, 3)))
	ConfirmMove(Cons(0, Cons(1, Cons(2, 3))), 2, Cons(2, 3))
	RefuteMove(Cons(0, Cons(1, Cons(2, 3))), 3)
}

func TestPairEnd(t *testing.T) {
	RefuteEnd := func(c *Pair) {
		vc := fmt.Sprintf("%v", c)
		defer ConfirmPanic(t, "%v.End() should panic", vc)()
		c.End()
	}

	ConfirmEnd := func(c, o *Pair) {
		end := c.End()
		if ok := end.Equal(o); !ok {
			t.Fatalf("%v.End() should be %v but is %v", c, o, end)
		}
	}

	RefuteEnd(nil)
	ConfirmEnd(Cons(0, nil), Cons(0, nil))
	ConfirmEnd(Cons(0, 1), Cons(0, 1))
	ConfirmEnd(List(0, 1, 2), Cons(2, nil))
	ConfirmEnd(List(0, 1, 2, 3), Cons(3, nil))
	ConfirmEnd(Cons(0, Cons(1, Cons(2, Cons(3, nil)))), Cons(3, nil))
}

func TestPairAppend(t *testing.T) {
	ConfirmAppend := func(c *Pair, v interface{}, r interface{}) {
		cs := fmt.Sprintf("%v", c)
		if x := c.Append(v); !x.Equal(r) {
			t.Fatalf("%v.Append(%v) should be %v but is %v", cs, v, r, x)
		}
	}

	ConfirmAppend(List(), 1, List(1))
	ConfirmAppend(List(), List(1), List(1))
	ConfirmAppend(List(), List(1, 2), List(1, 2))
	ConfirmAppend(List(), List(1, 2, 3), List(1, 2, 3))
	ConfirmAppend(List(1), 2, List(1, 2))
	ConfirmAppend(List(1), List(2), List(1, 2))
	ConfirmAppend(List(1), List(2, 3), List(1, 2, 3))

	ConfirmMultipleAppend := func(c *Pair, r interface{}, v... interface{}) {
		call := fmt.Sprintf("%v.Append(%v)", c, v)
		if x := c.Append(v...); !x.Equal(r) {
			t.Fatalf("%v should be %v but is %v", call, r, x)
		}
	}

	ConfirmMultipleAppend(List(), List(1, 2, 3), List(1), List(2, 3))
	ConfirmMultipleAppend(List(), List(1, 2, 3, 4), List(1, 2), 3, List(4))
	ConfirmMultipleAppend(List(), List(1, 2, 3), List(1, 2, 3))
	ConfirmMultipleAppend(List(1), List(1, 2), 2)
	ConfirmMultipleAppend(List(1), List(1, 2, 3), 2, 3)
	ConfirmMultipleAppend(List(1), List(1, 2), List(2))
	ConfirmMultipleAppend(List(1), List(1, 2, 3), List(2, 3))
}

func TestPairEach(t *testing.T) {
	list := List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	count := 0

	ConfirmEach := func(c *Pair, f interface{}) {
		count = 0
		c.Each(f)
		if l := c.Len(); l != count {
			t.Fatalf("%v.Each() should have iterated %v times not %v times", c, l, count)
		}
	}

	ConfirmEach(list, func(i interface{}) {
		if i != count {
			t.Fatalf("1: %v.Each() element %v erroneously reported as %v", list, count, i)
		}
		count++
	})

	ConfirmEach(list, func(index int, i interface{}) {
		if i != index {
			t.Fatalf("2: %v.Each() element %v erroneously reported as %v", list, index, i)
		}
		count++
	})

	ConfirmEach(list, func(key, i interface{}) {
		if i.(int) != key.(int) {
			t.Fatalf("3: %v.Each() element %v erroneously reported as %v", list, key, i)
		}
		count++
	})

	list = List()
	ConfirmEach(list, func(i interface{}) {
		if i != count {
			t.Fatalf("4: %v.Each() element %v erroneously reported as %v", list, count, i)
		}
		count++
	})

	list = nil
	ConfirmEach(list, func(i interface{}) {
		if i != count {
			t.Fatalf("5: %v.Each() element %v erroneously reported as %v", list, count, i)
		}
		count++
	})
}


func TestPairStep(t *testing.T) {
	list := List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	count := 0

	ConfirmStep := func(c *Pair, start, step, expected int, f interface{}) {
		count = 0
		c.Step(start, step, f)
		if expected != count {
			t.Fatalf("%v.Step(%v, %v) should have iterated %v times not %v times", c, start, step, expected, count)
		}
	}

	f := func(index int, i interface{}) {
		count++
	}

	ConfirmStep(list, 0, 1, 10, f)
	ConfirmStep(list, 0, 2, 5, f)
	ConfirmStep(list, 0, 3, 4, f)
	ConfirmStep(list, 0, 4, 3, f)
	ConfirmStep(list, 0, 5, 2, f)
	ConfirmStep(list, 0, 6, 2, f)

	ConfirmStep(list, 1, 1, 9, f)
	ConfirmStep(list, 1, 2, 5, f)
	ConfirmStep(list, 1, 3, 3, f)
	ConfirmStep(list, 1, 4, 3, f)
	ConfirmStep(list, 1, 5, 2, f)
	ConfirmStep(list, 1, 6, 2, f)

	ConfirmStep(list, 2, 1, 8, f)
	ConfirmStep(list, 2, 2, 4, f)
	ConfirmStep(list, 2, 3, 3, f)
	ConfirmStep(list, 2, 4, 2, f)
	ConfirmStep(list, 2, 5, 2, f)
	ConfirmStep(list, 2, 6, 2, f)

	ConfirmStep(list, 3, 1, 7, f)
	ConfirmStep(list, 3, 2, 4, f)
	ConfirmStep(list, 3, 3, 3, f)
	ConfirmStep(list, 3, 4, 2, f)
	ConfirmStep(list, 3, 5, 2, f)
	ConfirmStep(list, 3, 6, 2, f)
}

func TestPairMap(t *testing.T) {
	list := List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	doubles := List(0, 2, 4, 6, 8, 10, 12, 14, 16, 18)
	ConfirmMap := func(c, r *Pair, f interface{}) {
		switch m := c.Map(f); {
		case c.Len() != m.Len():
			t.Fatalf("%v.Map() should have iterated %v times not %v times", c, c.Len(), m.Len())
		case !m.Equal(r):
			t.Fatalf("%v.Map() should be %v but is %v", c, r, m)
		}
	}

	ConfirmMap(list, doubles, func(i interface{}) interface{} {
		return i.(int) * 2
	})

	ConfirmMap(list, doubles, func(index int, i interface{}) interface{} {
		return index + i.(int)
	})

	ConfirmMap(list, doubles, func(key, i interface{}) interface{} {
		return key.(int) + i.(int)
	})
}

func TestPairReduce(t *testing.T) {
	ConfirmReduce := func(c *Pair, r, seed, f interface{}) {
		if x, ok := c.Reduce(seed, f).(Equatable); ok {
			if !x.Equal(r) {
				t.Fatalf("%v.Reduce() should be %v but is %v", c, r, x)
			}
		} else if r, ok := r.(Equatable); ok {
			if !r.Equal(x) {
				t.Fatalf("%v.Reduce() should be %v but is %v", c, r, x)
			}
		} else {
			if x != r {
				t.Fatalf("%v.Reduce() should be %v but is %v", c, r, x)
			}
		}
		
	}

	ConfirmReduce(List(1, 2, 3), 6, 1, func(seed, value interface{}) interface{} {
		return seed.(int) * value.(int)
	})

	ConfirmReduce(List(1, 2, 3), 6, 1, func(index int, seed, value interface{}) interface{} {
		return seed.(int) * value.(int)
	})

	ConfirmReduce(List(1, 2, 3), 6, 1, func(key, seed, value interface{}) interface{} {
		return seed.(int) * value.(int)
	})

	ConfirmReduce(List(1, 2, 3), 7, 1, func(seed, value interface{}) interface{} {
		return seed.(int) + value.(int)
	})

	ConfirmReduce(List(1, 2, 3), 7, 1, func(index int, seed, value interface{}) interface{} {
		return seed.(int) + value.(int)
	})

	ConfirmReduce(List(1, 2, 3), 7, 1, func(key, seed, value interface{}) interface{} {
		return seed.(int) + value.(int)
	})

	ConfirmReduce(List("A", "B", "C"), "ABC", "", func(seed, value interface{}) interface{} {
		return seed.(string) + value.(string)
	})

	ConfirmReduce(List("A", "B", "C"), "ABC", "", func(index int, seed, value interface{}) interface{} {
		return seed.(string) + value.(string)
	})

	ConfirmReduce(List("A", "B", "C"), "ABC", "", func(key, seed, value interface{}) interface{} {
		return seed.(string) + value.(string)
	})
}

func TestPairWhile(t *testing.T) {
	list := List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	ConfirmLimit := func(condition bool, c *Pair, l int, f interface{}) {
		if count := c.While(condition, f); count != l {
			t.Fatalf("%v.While(%v, %v) should have iterated %v times not %v times", c, condition, l, l, count)
		}
	}

	limit := 5
	ConfirmLimit(true, list, limit, func(i interface{}) bool {
		return i != limit
	})

	limit = 6
	ConfirmLimit(true, list, limit, func(index int, i interface{}) bool {
		return index != limit
	})

	limit = 7
	ConfirmLimit(true, list, limit, func(key, i interface{}) bool {
		return key != limit
	})

	limit = 5
	ConfirmLimit(false, list, limit, func(i interface{}) bool {
		return i == limit
	})

	limit = 6
	ConfirmLimit(false, list, limit, func(index int, i interface{}) bool {
		return index == limit
	})

	limit = 7
	ConfirmLimit(false, list, limit, func(key, i interface{}) bool {
		return key == limit
	})
}

func TestPairPartition(t *testing.T) {
	RefutePartition := func(c *Pair, offset int) {
		vc := fmt.Sprintf("%v", c)
		defer ConfirmPanic(t, "%v.Partition(%v) should panic", vc, offset)()
		c.Partition(offset)
	}

	ConfirmPartition := func(l *Pair, offset int, x, y *Pair) {
		ls := fmt.Sprintf("%v", l)
		switch c1, c2 := l.Partition(offset); {
		case !c1.Equal(x):
			t.Fatalf("1: %v.Partition(%v) should be (%v, %v) but is (%v, %v)", ls, offset, x, y, c1, c2)
		case !c2.Equal(c2):
			t.Fatalf("2: %v.Partition(%v) should be (%v, %v) but is (%v, %v)", ls, offset, x, y, c1, c2)
		}
	}

	RefutePartition(nil, -1)

	ConfirmPartition(nil, 0, nil, nil)
	ConfirmPartition(nil, 1, nil, nil)

	ConfirmPartition(List(), 0, nil, nil)
	ConfirmPartition(List(), 1, nil, nil)

	ConfirmPartition(List(0), 0, List(0), List())
	ConfirmPartition(List(0), 1, List(0), List())

	ConfirmPartition(List(0, 1), 0, List(0), List(1))
	ConfirmPartition(List(0, 1), 1, List(0, 1), List())
	ConfirmPartition(List(0, 1), 2, List(0, 1), List())

	ConfirmPartition(List(0, 1, 2), 0, List(0), List(1, 2))
	ConfirmPartition(List(0, 1, 2), 1, List(0, 1), List(2))
	ConfirmPartition(List(0, 1, 2), 2, List(0, 1, 2), List())
	ConfirmPartition(List(0, 1, 2), 3, List(0, 1, 2), List())
}

func TestPairReverse(t *testing.T) {
	ConfirmReverse := func(l, r *Pair) {
		if c := l.Reverse(); !c.Equal(r) {
			t.Fatalf("%v.Reverse() should be %v but is %v", l, r, c)
		}
	}

	ConfirmReverse(List(), List())
	ConfirmReverse(List(0), List(0))
	ConfirmReverse(List(0, 1), List(1, 0))
	ConfirmReverse(List(0, 1, 2), List(2, 1, 0))
	ConfirmReverse(List(0, 1, List(2, 3), 4), List(4, List(2, 3), 1, 0))
	ConfirmReverse(List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9), List(9, 8, 7, 6, 5, 4, 3, 2, 1, 0))
}

func TestPairCopy(t *testing.T) {
	ConfirmCopy := func(l, r *Pair) {
		if c := l.Copy(); !c.Equal(r) {
			t.Fatalf("%v.Copy() should be %v but is %v", l, r, c)
		}
	}

	ConfirmCopy(List(), List())
	ConfirmCopy(List(0), List(0))
	ConfirmCopy(List(0, 1), List(0, 1))
	ConfirmCopy(List(0, 1, List(2, 3), 4), List(0, 1, List(2, 3), 4))
	ConfirmCopy(List(0, 1, List(2, List(3, 4, 5)), 6), List(0, 1, List(2, List(3, 4, 5)), 6))
}

func TestPairRepeat(t *testing.T) {
	ConfirmRepeat := func(l *Pair, count int, r *Pair) {
		if c := l.Repeat(count); !c.Equal(r) {
			t.Fatalf("%v.Repeat(%v) should be %v but is %v", l, count, r, c)
		}
	}

	ConfirmRepeat(List(), 0, List())
	ConfirmRepeat(List(), 1, List())
	ConfirmRepeat(List(), 2, List())
	ConfirmRepeat(List(), 3, List())

	ConfirmRepeat(List(0), 0, List())
	ConfirmRepeat(List(0), 1, List(0))
	ConfirmRepeat(List(0), 2, List(0, 0))
	ConfirmRepeat(List(0), 3, List(0, 0, 0))

	ConfirmRepeat(List(0, 1), 0, List())
	ConfirmRepeat(List(0, 1), 1, List(0, 1))
	ConfirmRepeat(List(0, 1), 2, List(0, 1, 0, 1))
	ConfirmRepeat(List(0, 1), 3, List(0, 1, 0, 1, 0, 1))

	ConfirmRepeat(List(0, 1, List(2, 3), 4), 0, List())
	ConfirmRepeat(List(0, 1, List(2, 3), 4), 1, List(0, 1, List(2, 3), 4))
	ConfirmRepeat(List(0, 1, List(2, 3), 4), 2, List(0, 1, List(2, 3), 4, 0, 1, List(2, 3), 4))
	ConfirmRepeat(List(0, 1, List(2, 3), 4), 3, List(0, 1, List(2, 3), 4, 0, 1, List(2, 3), 4, 0, 1, List(2, 3), 4))
}

func TestPairZip(t *testing.T) {
	ConfirmZip := func(x, y, r *Pair) {
		if c := x.Zip(y); !c.Equal(r) {
			t.Fatalf("%v.Zip(%v) should be %v but is %v", x, y, r, c)
		}
	}

	ConfirmZip(nil, nil, nil)
	ConfirmZip(List(), List(), List())
	ConfirmZip(Cons(nil, nil), Cons(nil, nil), Cons(Cons(nil, nil), nil))
	ConfirmZip(Cons(1, nil), Cons(2, nil), Cons(Cons(1, 2), nil))
	ConfirmZip(Cons(2, nil), Cons(1, nil), Cons(Cons(2, 1), nil))
	ConfirmZip(List(1, 2), List(3, 4), List(Cons(1, 3), Cons(2, 4)))
	ConfirmZip(List(1, 2, 3, 4, 5), List(5, 4, 3, 2, 1), List(Cons(1, 5), Cons(2, 4), Cons(3, 3), Cons(4, 2), Cons(5, 1)))
}