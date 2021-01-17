package BloomFilter

import (
    "github.com/willf/bitset"
    "io"
)

//
type memBitSet struct {
    bs bitset.BitSet
}

func NewMemoryBitSet() BitSet {
    r := &memBitSet{}
    return r
}
func (self *memBitSet) set(offset int64) error {
    self.bs.Set(uint(offset))
    return nil
}

func (self *memBitSet) get(offset int64) (int64, error) {
    ok := self.bs.Test(uint(offset))
    if ok {
        return 1, nil
    }
    return 0, nil
}
func (self*memBitSet) del(at int64) error {
    self.bs.DeleteAt(uint(at))
    return nil
}
func (self*memBitSet) Load(r io.Reader) (int64, error) {
    return self.bs.ReadFrom(r)
}
func (self*memBitSet) Save(w io.Writer) (int64, error) {
    return self.bs.WriteTo(w)
}
