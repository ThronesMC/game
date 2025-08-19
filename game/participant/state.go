package participant

var (
	StateAlive              = alive{}
	StateEliminated         = eliminated{}
	StateRespawning         = respawning{}
	StateTemporarySpectator = spectator{IsPermanent: false}
	StatePermanentSpectator = spectator{IsPermanent: true}
)

type State interface {
	Permanent() bool
}

type alive struct {
	State
}

func (alive) Permanent() bool {
	return false
}

type eliminated struct {
	State
}

func (eliminated) Permanent() bool {
	return true
}

type respawning struct {
	State
}

func (respawning) Permanent() bool {
	return false
}

type spectator struct {
	State
	IsPermanent bool
}

func (s spectator) Permanent() bool {
	return s.IsPermanent
}
