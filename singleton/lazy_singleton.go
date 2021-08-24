package singleton

import "sync"

var (
	lazySingleton *Singleton
	once          = &sync.Once{} // sync.once ensures only one instance would be created
)

func GetLazyInstance() *Singleton {
	if lazySingleton == nil {
		once.Do(func() {
			lazySingleton = &Singleton{}
		})
	}
	return lazySingleton
}
