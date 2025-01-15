package actor

const LocalAddress = "local"

const pidSeparator = "/"

// Process ID
func NewPID(address, id string) *PID {
	p := &PID{
		Address: address,
		ID:      id,
	}
	return p
}

func (pid *PID) Equals(other *PID) bool {
	return pid.Address == other.Address && pid.ID == other.ID
}

func (pid *PID) Child(id string) *PID {
	childID := pid.ID + pidSeparator + id
	return NewPID(pid.Address, childID)
}
