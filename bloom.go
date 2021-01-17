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
    del(offset int64) error
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
    for _, fn := range self.fns {
        offset := fn(str)
        if err := self.bs.set(int64(offset)); err != nil {
            return err
        }
    }
    return nil
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
    return self.bs.set(offset)
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

func (self*Bloom)DelInt(offset int64) error {
    return self.bs.del(offset)
}

func (self*Bloom)DelString(str string) error {
    for _, fn := range self.fns {
        offset := fn(str)
        if err := self.bs.del(int64(offset)); err != nil {
            return err
        }
    }
    return nil
}