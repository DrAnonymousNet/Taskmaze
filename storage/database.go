package storage

import "sync"


type DB struct{
	GlobalTasksMap map[int]Task
	mu sync.Mutex
}

var TaskDB *DB


func InitDB() *DB {
	GlobalTasksMap :=  make(map[int]Task)
	TaskDB := DB{GlobalTasksMap: GlobalTasksMap}
	return &TaskDB
}



