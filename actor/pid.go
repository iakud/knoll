package actor

const LocalAddress = "local"

// Process ID
type PID struct {
	Address string
	ID      string
}

const pidSeparator = "/"

func NewPID(address, id string) *PID {
	p := &PID{
		Address: address,
		ID:      id,
	}
	return p
}

func (pid *PID) String() string {
	return pid.Address + pidSeparator + pid.ID
}

func (pid *PID) Equals(other *PID) bool {
	return pid.Address == other.Address && pid.ID == other.ID
}

func (pid *PID) Child(id string) *PID {
	childID := pid.ID + pidSeparator + id
	return NewPID(pid.Address, childID)
}
