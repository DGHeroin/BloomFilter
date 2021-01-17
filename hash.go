package BloomFilter

type HashFunc func(string) uint64

func NewHashFunc() []HashFunc {
    return []HashFunc{mBKDRHash, mSDBMHash, mDJBHash}
}

func mBKDRHash(str string) uint64 {
    seed := uint64(131) // 31 131 1313 13131 131313 etc..
    hash := uint64(0)
    for i := 0; i < len(str); i++ {
        hash = (hash * seed) + uint64(str[i])
    }
    return hash & 0x7FFFFFFF
}
func mSDBMHash(str string) uint64 {
    hash := uint64(0)
    for i := 0; i < len(str); i++ {
        hash = uint64(str[i]) + (hash << 6) + (hash << 16) - hash
    }
    return hash & 0x7FFFFFFF
}
func mDJBHash(str string) uint64 {
    hash := uint64(0)
    for i := 0; i < len(str); i++ {
        hash = ((hash << 5) + hash) + uint64(str[i])
    }
    return hash & 0x7FFFFFFF
}
