package easync

type FieldMask struct {
	mask uint64
}

func (m *FieldMask) Mark(n uint64) {
	m.mask |= n
}

func (m *FieldMask) Check(n uint64) bool {
	return m.mask | n != 0
}

func (m *FieldMask) Clear() {
	m.mask = 0
}

type MapMask[Key int32|int64|uint32|uint64|bool|string] struct {
	maskSet map[Key]struct{}
	maskDelete map[Key]struct{}
	maskClear bool
}

func (m *MapMask[Key]) MarkSet(key Key) {
	m.maskSet[key] = struct{}{}
	if m.maskClear {
		return
	}
	delete(m.maskDelete, key)
}

func (m *MapMask[Key]) MarkDelete(key Key) {
	delete(m.maskSet, key)
	if m.maskClear {
		return
	}
	m.maskDelete[key] = struct{}{}
}

func (m *MapMask[Key]) MarkClear() {
	clear(m.maskSet)
	if m.maskClear {
		return
	}
	clear(m.maskDelete)
	m.maskClear = true
}

func (m *MapMask[Key]) Clear() {
	clear(m.maskSet)
	clear(m.maskDelete)
	m.maskClear = false
}