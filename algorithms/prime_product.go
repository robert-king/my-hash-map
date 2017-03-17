package algorithms

var primes []uint64

func init() {
	primes = append(make([]uint64, 'a'), []uint64{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101}...)
}

func PrimeProduct(s string) uint64 {
	hash := uint64(1)
	for _, c := range s {
		hash *= primes[c]
	}
	return hash
}
