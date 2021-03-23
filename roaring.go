package BloomFilter

import (
    roaring "github.com/RoaringBitmap/roaring/roaring64"
    "io"
)

type smallMemBitSet struct {
    bs roaring.Bitmap
}

func NewSmallMemoryBitSet() BitSet {
    r := &smallMemBitSet{}
    return r
}
func (self *smallMemBitSet) set(offset uint64) error {
    self.bs.Add(offset)
    return nil
}
func (self *smallMemBitSet) get(offset uint64) (uint64, error) {
    ok := self.bs.Contains(offset)
    if ok {
        return 1, nil
    }
    return 0, nil
}
func (self*smallMemBitSet) del(offset uint64) error {
    self.bs.Remove(offset)
    return ErrNotSupport
}
func (self*smallMemBitSet) count() uint64 {
    return self.bs.GetCardinality()
}
func (self *smallMemBitSet) IsUint64() bool {
    return true
}
func (self*smallMemBitSet) Load(r io.Reader) (int64, error) {
    return self.bs.ReadFrom(r)
}
func (self*smallMemBitSet) Save(w io.Writer) (int64, error) {
    return self.bs.WriteTo(w)
}