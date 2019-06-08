package fsm

import "sync"

type StateType int
type ActionType int
type Handler func(action ActionType, args ...interface{}) (nextState StateType)

type FSM struct {
	lock        sync.Mutex
	CurState    StateType             // 当前状态
	handlers    map[StateType]Handler // 每个状态对应的函数
	endHandlers map[StateType]bool    // 记录结束状态
}

func New(startState StateType) *FSM {
	return &FSM{
		CurState:    startState,
		handlers:    make(map[StateType]Handler),
		endHandlers: make(map[StateType]bool),
	}
}

func (me *FSM) AddState(state StateType, handler Handler) {
	me.lock.Lock()
	defer me.lock.Unlock()
	me.handlers[state] = handler
}

func (me *FSM) Do(action ActionType, args ...interface{}) (finished bool) {
	me.lock.Lock()
	defer me.lock.Unlock()

	h, ok := me.handlers[me.CurState]
	if !ok {
		return
	}
	finished, _ = me.endHandlers[me.CurState]
	if s := h(action, args...); s != 0 {
		me.CurState = s
	}
	return
}
