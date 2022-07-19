package repository

import "sync"

type Storage struct {
	mutex sync.RWMutex
	Data  map[int64]int64
}

func (s *Storage) GetSortPointIdByDstOfficeId(dstOfficeId int64) (int64, bool) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	val, ok := s.Data[dstOfficeId]
	return val, ok
}

func (s *Storage) SaveSortPointIdForDstOffice(dstOfficeId int64, sortPointId int64) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.Data[dstOfficeId] = sortPointId
}
