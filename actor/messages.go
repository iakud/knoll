package actor

type PoisonPill struct{}
type Started struct{}
type Stopped struct{}

var (
	poisonPill PoisonPill = PoisonPill{}
	started    Started    = Started{}
	stopped    Stopped    = Stopped{}
)
