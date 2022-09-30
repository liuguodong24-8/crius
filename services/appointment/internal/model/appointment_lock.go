package model

// AppointmentLock 预约锁
type AppointmentLock struct {
	BranchID        string
	RoomGroupID     string
	Way             int8
	AppointmentDate string
	AppointmentAt   int32
}

const (
	// AppointmentHashLockKey 锁hash结构redis key
	AppointmentHashLockKey = "appointment:hash:lock:id:%s"
	// AppointmentZsetLockKey 锁zset结构redis key
	AppointmentZsetLockKey = "appointment:zset:lock"
)
