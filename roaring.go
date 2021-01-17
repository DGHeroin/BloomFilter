package BloomFilter

import (
    "github.com/RoaringBitmap/roaring"
    "io"
)

type smallMemBitSet struct {
    bs roaring.Bitmap
}

func NewSmallMemoryBitSet() BitSet {
    r := &smallMemBitSet{}
    return r
}
func (self *smallMemBitSet) set(offset int64) error {
    self.bs.Add(uint32(offset))
    return nil
}
func (self *smallMemBitSet) get(offset int64) (int64, error) {
    ok := self.bs.Contains(uint32(offset))
    if ok {
        return 1, nil
    }
    return 0, nil
}
func (self*smallMemBitSet) del(offset int64) error {
    self.bs.Remove(uint32(offset))
    return ErrNotSupport
}
func (self*smallMemBitSet) Load(r io.Reader) (int64, error) {
    return self.bs.ReadFrom(r)
}
func (self*smallMemBitSet) Save(w io.Writer) (int64, error) {
    return self.bs.WriteTo(w)
}