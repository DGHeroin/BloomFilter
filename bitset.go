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
func (self *memBitSet) set(offset uint64) error {
    self.bs.Set(uint(offset))
    return nil
}

func (self *memBitSet) get(offset uint64) (uint64, error) {
    ok := self.bs.Test(uint(offset))
    if ok {
        return 1, nil
    }
    return 0, nil
}
func (self*memBitSet) del(at uint64) error {
    self.bs.DeleteAt(uint(at))
    return nil
}
func (self*memBitSet) count() uint64 {
    return uint64(self.bs.Count())
}
func (self *memBitSet) IsUint64() bool {
    return false
}
func (self*memBitSet) Load(r io.Reader) (int64, error) {
    return self.bs.ReadFrom(r)
}
func (self*memBitSet) Save(w io.Writer) (int64, error) {
    return self.bs.WriteTo(w)
}
