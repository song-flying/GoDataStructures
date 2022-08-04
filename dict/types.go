package dict

type (
	Entry[K any, V any] interface {
		Key() K
		Value() V
	}

	Dict[K comparable, V comparable] interface {
		Get(key K) (V, bool)
		Put(key K, value V)
	}
)
