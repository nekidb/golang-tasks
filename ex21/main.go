package main

// Есть интерфейс, с которым работает клиентское приложение.
type StorageInterface interface {
	Get(key int) (val string)
}

// Есть сервис, с которым клиент хотел бы взаимодействовать, но интерфейс сервиса клиенту не понятен.
// При этом мы не хотим менять исходный код сервиса.
type Storage struct {
	data map[int]string
}

func (s *Storage) GetValueByKey(key int) string {
	return s.data[key]
}

// Поэтому мы добавляем адаптер. Который реализует пользовательский интерфейс и сам обращается к сервису.
type StorageAdapter struct {
	s *Storage
}
// Конструктор адаптера
func NewStorageAdapter(s *Storage) StorageInterface {
	return &StorageAdapter{s}
}
// Реализация интерфейса
func (sa *StorageAdapter) Get(key int) (val string) {
	val = sa.s.GetValueByKey(key)
	return
}

func main() {
	s := &Storage{data: make(map[int]string)}
	sa := NewStorageAdapter(s)

	sa.Get(1)
}
