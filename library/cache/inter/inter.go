package inter

type Cache interface {
	Save(key, value[]byte)error
	Get(key []byte) []byte
	Has(key []byte) bool
}


