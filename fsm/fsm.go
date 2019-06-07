package fsm

type StateType int
type ActionType int
type handler func(action ActionType, args ...interface{}) (nextState StateType)

type FSM struct {
	curState    StateType             // 当前状态
	handlers    map[StateType]handler // 每个状态对应的函数
	endHandlers map[StateType]bool    // 记录结束状态
}

func New(startState StateType) *FSM {
	return &FSM{
		curState:    startState,
		handlers:    make(map[StateType]handler),
		endHandlers: make(map[StateType]bool),
	}
}

func (me *FSM) AddState(state StateType, handler handler) {
	me.handlers[state] = handler
}

// 直接设置状态机为指定状态，不触发回调函数
func (me *FSM) SetState(stateType StateType) {
	me.curState = stateType
}

func (me *FSM) Do(action ActionType, args ...interface{}) (finished bool) {
	h, ok := me.handlers[me.curState]
	if !ok {
		return
	}
	finished, _ = me.endHandlers[me.curState]
	if s:=  h(action, args...);s!=0{
		me.curState = s
	}
	return
}
