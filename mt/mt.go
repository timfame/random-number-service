package mt

const (
	w             uint64 = 64
	n             uint64 = 312
	m             uint64 = 156
	r             uint64 = 31
	a             uint64 = 0xB5026F5AA96619E9
	u             uint64 = 29
	d             uint64 = 0x5555555555555555
	s             uint64 = 17
	b             uint64 = 0x71D67FFFEDA60000
	t             uint64 = 37
	c             uint64 = 0xFFF7EEE000000000
	l             uint64 = 43
	f             uint64 = 6364136223846793005
	defaultSeed   uint64 = 5489
	lowerBitsMask uint64 = (1 << r) - 1
	upperBitsMask        = ^lowerBitsMask
)

type MT struct {
	numbers [n]uint64
	step    uint64
}

func New(opts ...Option) *MT {
	config := &Config{}
	for _, opt := range opts {
		opt(config)
	}
	if config.Seed == 0 {
		config.Seed = defaultSeed
	}

	result := &MT{
		step: n - 1,
	}

	result.numbers[0] = config.Seed
	for i := uint64(1); i < n; i++ {
		result.numbers[i] = f*(result.numbers[i-1]^(result.numbers[i]>>(w-2))) + i
	}
	return result
}

func (mt *MT) Next() uint64 {
	mt.step++
	if mt.step >= n {
		mt.twist()
	}
	return tempering(mt.numbers[mt.step])
}

func (mt *MT) twist() {
	for i := uint64(0); i < n-1; i++ {
		x := (mt.numbers[i] & upperBitsMask) | (mt.numbers[i+1] & lowerBitsMask)
		xA := x >> 1
		if (x & 1) == 1 {
			xA ^= a
		}
		mt.numbers[i] = mt.numbers[(i+m)%n] ^ xA
	}
	x := (mt.numbers[n-1] & upperBitsMask) | (mt.numbers[0] & lowerBitsMask)
	xA := x >> 1
	if (x & 1) == 1 {
		xA ^= a
	}
	mt.numbers[n-1] = mt.numbers[m-1] ^ xA
	mt.step = 0
}

func tempering(x uint64) uint64 {
	x ^= (x >> u) & d
	x ^= (x << s) & b
	x ^= (x << t) & c
	x ^= x >> l
	return x
}
