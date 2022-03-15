package objectpool

type iPoolObject interface {
	getID() string
}
