package math

import (
	"lisp"
	"lisp/parser"
	"math"
)

var (
	Sin, Sinh, Asin, Asinh lisp.Gfac
	Cos, Cosh, Acos, Acosh lisp.Gfac
	Tan, Tanh, Atan, Atanh lisp.Gfac
	Exp, Log               lisp.Gfac
	Sqrt                   lisp.Gfac
	Pow                    lisp.Gfac
)

func Int(t []lisp.Token, p *lisp.Lisp) (lisp.Token, error) {
	if len(t) != 1 {
		return lisp.None, lisp.ErrParaNum
	}
	u, err := p.Exec(t[0])
	if err != nil {
		return lisp.None, err
	}
	switch u.Kind {
	case lisp.Int:
		return u, nil
	case lisp.Float:
		return lisp.Token{lisp.Int, int64(u.Text.(float64))}, nil
	case lisp.String:
		a, b := parser.ParseInt([]byte(u.Text.(string)))
		if b == 0 {
			return lisp.None, lisp.ErrNotConv
		}
		return lisp.Token{lisp.Int, a}, nil
	}
	return lisp.None, lisp.ErrFitType
}

func Float(t []lisp.Token, p *lisp.Lisp) (lisp.Token, error) {
	if len(t) != 1 {
		return lisp.None, lisp.ErrParaNum
	}
	u, err := p.Exec(t[0])
	if err != nil {
		return lisp.None, err
	}
	switch u.Kind {
	case lisp.Int:
		return lisp.Token{lisp.Float, float64(u.Text.(int64))}, nil
	case lisp.Float:
		return u, nil
	case lisp.String:
		a, b := parser.ParseFloat([]byte(u.Text.(string)))
		if b == 0 {
			return lisp.None, lisp.ErrNotConv
		}
		return lisp.Token{lisp.Float, a}, nil
	}
	return lisp.None, lisp.ErrFitType
}

func Wrap1(f func(float64) float64) func([]lisp.Token, *lisp.Lisp) (lisp.Token, error) {
	return func(t []lisp.Token, p *lisp.Lisp) (lisp.Token, error) {
		if len(t) != 1 {
			return lisp.None, lisp.ErrParaNum
		}
		u, err := p.Exec(t[0])
		if err != nil {
			return lisp.None, err
		}
		if u.Kind != lisp.Float {
			return lisp.None, lisp.ErrFitType
		}
		return lisp.Token{lisp.Float, f(u.Text.(float64))}, nil
	}
}

func Wrap2(f func(float64, float64) float64) func([]lisp.Token, *lisp.Lisp) (lisp.Token, error) {
	return func(t []lisp.Token, p *lisp.Lisp) (lisp.Token, error) {
		if len(t) != 2 {
			return lisp.None, lisp.ErrParaNum
		}
		u, err := p.Exec(t[0])
		if err != nil {
			return lisp.None, err
		}
		if u.Kind != lisp.Float {
			return lisp.None, lisp.ErrFitType
		}
		v, err := p.Exec(t[1])
		if err != nil {
			return lisp.None, err
		}
		if v.Kind != lisp.Float {
			return lisp.None, lisp.ErrFitType
		}
		return lisp.Token{lisp.Float, f(u.Text.(float64), v.Text.(float64))}, nil
	}
}

func init() {
	Sin, Sinh, Asin, Asinh = Wrap1(math.Sin), Wrap1(math.Sinh), Wrap1(math.Asin), Wrap1(math.Asinh)
	Cos, Cosh, Acos, Acosh = Wrap1(math.Cos), Wrap1(math.Cosh), Wrap1(math.Acos), Wrap1(math.Acosh)
	Tan, Tanh, Atan, Atanh = Wrap1(math.Tan), Wrap1(math.Tanh), Wrap1(math.Atan), Wrap1(math.Atanh)
	Exp, Log = Wrap1(math.Exp), Wrap1(math.Log)
	Sqrt = Wrap1(math.Sqrt)
	Pow = Wrap2(math.Pow)
}