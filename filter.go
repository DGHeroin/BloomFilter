package BloomFilter

import (
    "errors"
    "io"
)

var (
    ErrNotSupport = errors.New("not support")
)

type Filter struct {
    bs  BitSet
    fns []HashFunc
}

type BitSet interface {
    set(offset uint64) error
    get(offset uint64) (uint64, error)
    del(offset uint64) error
    Load(r io.Reader) (int64, error)
    Save(w io.Writer) (int64, error)
}

func NewFilter(bs BitSet) *Filter {
    return &Filter{
        bs:  bs,
        fns: NewHashFunc(),
    }
}

func (self *Filter) Add(str string) error {
    for _, fn := range self.fns {
        offset := fn(str)
        if err := self.bs.set(offset); err != nil {
            return err
        }
    }
    return nil
}

func (self *Filter) Exists(str string) bool {
    var a uint64 = 1
    for _, fn := range self.fns {
        offset := fn(str)
        val, err := self.bs.get(offset)
        if err != nil {
            return false
        }
        if val != a {
            return false
        }
    }
    return true
}

func (self *Filter) AddInt(offset uint64) error {
    return self.bs.set(offset)
}

func (self *Filter) ExistsInt(offset uint64) bool {
    var a uint64 = 1
    val, err := self.bs.get(offset)
    if err != nil {
        return false
    }
    if val != a {
        return false
    }
    return true
}

func (self *Filter) DelInt(offset uint64) error {
    return self.bs.del(offset)
}

func (self *Filter) DelString(str string) error {
    for _, fn := range self.fns {
        offset := fn(str)
        if err := self.bs.del(offset); err != nil {
            return err
        }
    }
    return nil
}
