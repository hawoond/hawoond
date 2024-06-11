package constant

import "sync"

var (
	instances = make(map[string]any)
	onceMap   = make(map[string]*sync.Once)
	mutex     sync.Mutex
)

// SetInstance: 인스턴스 최초 설정
func SetInstance(key string, value any) {
	mutex.Lock()
	defer mutex.Unlock()

	if _, exists := onceMap[key]; !exists {
		onceMap[key] = &sync.Once{}
	}

	onceMap[key].Do(func() {
		instances[key] = value
	})
}

// GetInstance: 인스턴스 반환
func GetInstance(key string) (any, bool) {
	mutex.Lock()
	defer mutex.Unlock()

	value, exists := instances[key]
	return value, exists
}
