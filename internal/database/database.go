package database

type IDatabase interface {
	Init()
	Save(int64, string)
	FindShort(string) (int64, bool)
	FindFull(int64) (string, bool)
	Close()
}