package BloomFilter

import (
    "context"
    "github.com/go-redis/redis/v8"
    "io"
    "io/ioutil"
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
func (self *redisBitSet) set(offset uint64) error {
    cmd := self.conn.SetBit(context.Background(), self.key, int64(offset), 1)
    return cmd.Err()
}

func (self *redisBitSet) get(offset uint64) (uint64, error) {
    cmd := self.conn.GetBit(context.Background(), self.key, int64(offset))
    return uint64(cmd.Val()), cmd.Err()
}
func (self *redisBitSet) del(offset uint64) error {
    cmd := self.conn.SetBit(context.Background(), self.key, int64(offset), 0)
    return cmd.Err()
}
func (self *redisBitSet) Load(r io.Reader) (int64, error) {
    data, err := ioutil.ReadAll(r)
    if err != nil {
        return 0, err
    }
    cmd := self.conn.Set(context.Background(), self.key, data, 0)
    return 0, cmd.Err()
}
func (self *redisBitSet) Save(w io.Writer) (int64, error) {
    cmd := self.conn.Get(context.Background(), self.key)
    if cmd.Err() != nil {
        return 0, cmd.Err()
    }
    data, err := cmd.Bytes()
    if err != nil {
        return 0, err
    }
    n, err := w.Write(data)
    return int64(n), err
}
