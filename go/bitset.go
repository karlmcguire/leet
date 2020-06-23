type bitset []uint64

func newBitset(size uint64) bitset { return make(bitset, (size+63)/64) }
func (s bitset) set(b int)         { s[b/64] |= 1 << (b & 63) }
func (s bitset) get(b int) bool    { return s[b/64]&(1<<(b&63)) != 0 }
