package BloomFilter

import (
    "errors"
    "io"
)

var (
    ErrNotSupport = errors.New("not support")
)

type Bloom struct {
    bs  BitSet
    fns []HashFunc
}

type BitSet interface {
    set(offset int64) error
    get(offset int64) (int64, error)
    Load(r io.Reader) (int64, error)
    Save(w io.Writer) (int64, error)
}

func NewBloom(bs BitSet) *Bloom {
    return &Bloom{
        bs:  bs,
        fns: NewHashFunc(),
    }
}

func (self *Bloom) Add(str string) error {
    var err error
    for _, fn := range self.fns {
        offset := fn(str)
        self.bs.set(int64(offset))
    }
    return err
}

func (self *Bloom) Exists(str string) bool {
    var a int64 = 1
    for _, fn := range self.fns {
        offset := fn(str)
        val, err := self.bs.get(int64(offset))
        if err != nil {
            return false
        }
        if val != a {
            return false
        }
    }
    return true
}

func (self *Bloom) AddInt(offset int64) error {
    self.bs.set(offset)
    return nil
}

func (self *Bloom) ExistsInt(offset int64) bool {
    var a int64 = 1
    val, err := self.bs.get(offset)
    if err != nil {
        return false
    }
    if val != a {
        return false
    }
    return true
}
