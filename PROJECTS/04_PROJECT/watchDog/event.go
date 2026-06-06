package watchdog

import "time"

type Op uint32

const (
    Created  Op = 1 << iota // 1
    Modified                // 2
    Deleted                 // 4
    Renamed                 // 8
)

func (op Op) String() string {
    switch op {
    case Created:
        return "created"
    case Modified:
        return "modified"
    case Deleted:
        return "deleted"
    case Renamed:
        return "renamed"
    default:
        return "unknown"
    }
}

type Event struct {
    Path string
    Op   Op
    Time time.Time
    Size int64
}