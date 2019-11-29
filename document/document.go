/* This package provides methods to represent a document, including
 * manipulation.
 */
package document

import (
    "fmt"
    "log"

    pb "github.com/cnnrznn/docs/editor"
)

type Document struct {
	State   []byte
	Log     []Op
	Version int
}

type Op struct {
    Sender int
	Type    int // 0-insert, 1-delete
	Version int
	Pos     int
	Char    byte
}

const (
    INSERT int = 0
    DELETE int = 1
    NULLOP int = 2
)

func (op Op) String() string {
    return fmt.Sprintf("{V:%v, T:%v, P:%v, C:%v}", op.Version, op.Type, op.Pos, op.Char)
}

func OpConvDoc(op Op) pb.Op {
    return pb.Op{Version: int64(op.Version),
                 Sender: int64(op.Sender),
                 Type: int32(op.Type),
                 Char: []byte{op.Char},
                 Pos: int64(op.Pos)}
}

func OpConvPB(op pb.Op) Op {
    return Op{Version: int(op.Version),
                       Sender: int(op.Sender),
                       Type: int(op.Type),
                       Char: op.Char[0],
                       Pos: int(op.Pos)}
}

func (d *Document) String() string {
	return string(d.State)
}

func (d *Document) Apply(op Op) {
    if d.Version != op.Version {
        log.Panic("Document version does not match update")
    }
    d.Version++

	if op.Pos < 0 { return }

    if op.Pos > len(d.State) {
		newState := make([]byte, op.Pos + 1)
        copy(newState, d.State)
        d.State = newState
	}

	switch op.Type {
	case 0:
		d.State = append(d.State, ' ')
		for i := len(d.State) - 1; i > op.Pos; i-- {
			d.State[i] = d.State[i-1]
		}
		d.State[op.Pos] = op.Char
	case 1:
		if op.Pos > len(d.State)-1 {
			return
		}

		for i := op.Pos; i < len(d.State)-1; i++ {
			d.State[i] = d.State[i+1]
		}
		d.State = d.State[:len(d.State)-1]
	}
}

func (d *Document) Operate(op Op) Op {
	for i := op.Version; i < len(d.Log); i++ {
		op.Transform(d.Log[i])
	}

	if op.Type != 2 {
        op.Version = d.Version
		d.Log = append(d.Log, op)
		d.Apply(op)
	}

	return op
}

func (op *Op) Transform(other Op) {
	switch {
	case op.Type == 0 && other.Type == 0:
		if other.Pos <= op.Pos {
			op.Pos++
		}
	case op.Type == 0 && other.Type == 1:
		if other.Pos < op.Pos {
			op.Pos--
		}
	case op.Type == 1 && other.Type == 0:
		if other.Pos <= op.Pos {
			op.Pos++
		}
	case op.Type == 1 && other.Type == 1:
		if other.Pos < op.Pos {
			op.Pos--
		} else if other.Pos == op.Pos {
			op.Type = 2
		}
	}
}
