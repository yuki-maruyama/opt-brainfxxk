package interpreter

import (
	"context"
	"errors"
	"fmt"
	"io"

	"github.com/rosylilly/brainfxxk/ast"
	"github.com/rosylilly/brainfxxk/optimizer"
	"github.com/rosylilly/brainfxxk/parser"
)

var (
	ErrInputFinished  = fmt.Errorf("input finished")
	ErrMemoryOverflow = fmt.Errorf("memory overflow")
)

type Interpreter struct {
	Program *ast.Program
	Config  *Config
	Memory  []byte
	Pointer int
}

func Run(ctx context.Context, s io.Reader, c *Config) (int, error) {
	p, err := parser.Parse(s)
	if err != nil {
		return 0, err
	}

	return NewInterpreter(p, c).Run(ctx)
}

func NewInterpreter(p *ast.Program, c *Config) *Interpreter {
	return &Interpreter{
		Program: p,
		Config:  c,
		Memory:  make([]byte, c.MemorySize),
		Pointer: 0,
	}
}

func (i *Interpreter) Run(ctx context.Context) (int, error) {
	p, err := optimizer.NewOptimizer().Optimize(i.Program)
	if err != nil {
		return 0, err
	}

	count, err := i.runExpressions(ctx, p.Expressions)
	if errors.Is(err, ErrInputFinished) && !i.Config.RaiseErrorOnEOF {
		return count, nil
	}
	return count, err
}

func (i *Interpreter) runExpressions(ctx context.Context, exprs []ast.Expression) (int, error) {
	count := 0
	for _, expr := range exprs {
		exprCount, err := i.runExpression(ctx, expr)
		if err != nil {
			return count, err
		}
		count += exprCount
	}
	return count, nil
}

func (i *Interpreter) runExpression(ctx context.Context, expr ast.Expression) (int, error) {
	select {
	case <-ctx.Done():
		return 0, ctx.Err()
	default:
	}

	count := 1

	switch e := expr.(type) {
	case *ast.PointerIncrementExpression:
		if i.Pointer == len(i.Memory)-1 && i.Config.RaiseErrorOnOverflow {
			return count, fmt.Errorf("%w: %d to pointer overflow, on %d:%d", ErrMemoryOverflow, i.Pointer, e.StartPos(), e.EndPos())
		}
		i.Pointer += 1
	case *ast.MultiplePointerIncrementExpression:
		if i.Pointer == len(i.Memory)-1 && i.Config.RaiseErrorOnOverflow {
			return count, fmt.Errorf("%w: %d to pointer overflow, on %d:%d", ErrMemoryOverflow, i.Pointer, e.StartPos(), e.EndPos())
		}
		i.Pointer += e.Count
	case *ast.PointerDecrementExpression:
		if i.Pointer == 0 && i.Config.RaiseErrorOnOverflow {
			return count, fmt.Errorf("%w: %d to pointer underflow, on %d:%d", ErrMemoryOverflow, i.Pointer, e.StartPos(), e.EndPos())
		}
		i.Pointer -= 1
	case *ast.MultiplePointerDecrementExpression:
		if i.Pointer == len(i.Memory)-1 && i.Config.RaiseErrorOnOverflow {
			return count, fmt.Errorf("%w: %d to pointer overflow, on %d:%d", ErrMemoryOverflow, i.Pointer, e.StartPos(), e.EndPos())
		}
		i.Pointer -= e.Count
	case *ast.ValueIncrementExpression:
		if i.Memory[i.Pointer] == 255 && i.Config.RaiseErrorOnOverflow {
			return count, fmt.Errorf("%w: %d to memory overflow, on %d:%d", ErrMemoryOverflow, i.Pointer, e.StartPos(), e.EndPos())
		}
		i.Memory[i.Pointer] += 1
	case *ast.ValueDecrementExpression:
		if i.Memory[i.Pointer] == 0 && i.Config.RaiseErrorOnOverflow {
			return count, fmt.Errorf("%w: %d to memory underflow, on %d:%d", ErrMemoryOverflow, i.Pointer, e.StartPos(), e.EndPos())
		}
		i.Memory[i.Pointer] -= 1
	case *ast.OutputExpression:
		if _, err := i.Config.Writer.Write([]byte{i.Memory[i.Pointer]}); err != nil {
			return count, err
		}
	case *ast.InputExpression:
		b := make([]byte, 1)
		if _, err := i.Config.Reader.Read(b); err != nil {
			if errors.Is(err, io.EOF) {
				return count, ErrInputFinished
			}
			return count, err
		}
		i.Memory[i.Pointer] = b[0]
	case *ast.WhileExpression:
		for i.Memory[i.Pointer] != 0 {
			bodyCount, err := i.runExpressions(ctx, e.Body)
			if err != nil {
				return count, err
			}
			count += bodyCount
		}
	}
	return count, nil
}
