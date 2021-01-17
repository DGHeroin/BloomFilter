package BloomFilter

import (
    "context"
    "github.com/go-redis/redis/v8"
    "io"
)

type redisBitSet struct {
    key  string
    conn *redis.Conn
}

func NewRedisBitSet(key string, conn *redis.Conn) BitSet {
    r := &redisBitSet{
        key:  key,
        conn: conn,
    }
    return r
}
func (self *redisBitSet) set(offset int64) error {
    cmd := self.conn.SetBit(context.Background(), self.key, int64(offset), 1)
    return cmd.Err()
}

func (self *redisBitSet) get(offset int64) (int64, error) {
    cmd := self.conn.GetBit(context.Background(), self.key, int64(offset))
    return cmd.Val(), cmd.Err()
}

func (self *redisBitSet) Load(r io.Reader) (int64, error) {
    return 0, nil
}
func (self *redisBitSet) Save(w io.Writer) (int64, error) {
    return 0, nil
}
