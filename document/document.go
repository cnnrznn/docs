/* This package provides methods to represent a document, including
 * manipulation.
 */
package document

type Document struct {
	State   []byte
	Log     []Op
	Version int
}

type Op struct {
	Pos  int
	Char byte
	Type int // 0-insert, 1-delete
}

func New() *Document {
	return &Document{}
}

func (d *Document) String() string {
	return string(d.State)
}

func (d *Document) apply(op Op) {
	switch op.Type {
	case 0:
		d.State = append(d.State, ' ')
		for i := len(d.State) - 1; i > op.Pos; i-- {
			d.State[i] = d.State[i-1]
		}
		d.State[op.Pos] = op.Char
	case 1:
		for i := op.Pos; i < len(d.State)-1; i++ {
			d.State[i] = d.State[i+1]
		}
		d.State = d.State[:len(d.State)-1]
	}
}

func (d *Document) Operate(op Op, version int) Op {
	for i := version; i < len(d.Log); i++ {
		op.Transform(d.Log[i])
	}

	if op.Type != 2 {
		d.Log = append(d.Log, op)
		d.apply(op)
		d.Version++
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
