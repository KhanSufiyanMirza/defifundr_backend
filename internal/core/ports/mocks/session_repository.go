// Code generated by counterfeiter. DO NOT EDIT.
package mocks

import (
	"context"
	"sync"

	"github.com/demola234/defifundr/internal/core/domain"
	"github.com/demola234/defifundr/internal/core/ports"
	"github.com/google/uuid"
)

type FakeSessionRepository struct {
	BlockAllUserSessionsStub        func(context.Context, uuid.UUID) error
	blockAllUserSessionsMutex       sync.RWMutex
	blockAllUserSessionsArgsForCall []struct {
		arg1 context.Context
		arg2 uuid.UUID
	}
	blockAllUserSessionsReturns struct {
		result1 error
	}
	blockAllUserSessionsReturnsOnCall map[int]struct {
		result1 error
	}
	BlockSessionStub        func(context.Context, uuid.UUID) error
	blockSessionMutex       sync.RWMutex
	blockSessionArgsForCall []struct {
		arg1 context.Context
		arg2 uuid.UUID
	}
	blockSessionReturns struct {
		result1 error
	}
	blockSessionReturnsOnCall map[int]struct {
		result1 error
	}
	CreateSessionStub        func(context.Context, domain.Session) (*domain.Session, error)
	createSessionMutex       sync.RWMutex
	createSessionArgsForCall []struct {
		arg1 context.Context
		arg2 domain.Session
	}
	createSessionReturns struct {
		result1 *domain.Session
		result2 error
	}
	createSessionReturnsOnCall map[int]struct {
		result1 *domain.Session
		result2 error
	}
	DeleteSessionStub        func(context.Context, uuid.UUID) error
	deleteSessionMutex       sync.RWMutex
	deleteSessionArgsForCall []struct {
		arg1 context.Context
		arg2 uuid.UUID
	}
	deleteSessionReturns struct {
		result1 error
	}
	deleteSessionReturnsOnCall map[int]struct {
		result1 error
	}
	GetActiveSessionsByUserIDStub        func(context.Context, uuid.UUID) ([]domain.Session, error)
	getActiveSessionsByUserIDMutex       sync.RWMutex
	getActiveSessionsByUserIDArgsForCall []struct {
		arg1 context.Context
		arg2 uuid.UUID
	}
	getActiveSessionsByUserIDReturns struct {
		result1 []domain.Session
		result2 error
	}
	getActiveSessionsByUserIDReturnsOnCall map[int]struct {
		result1 []domain.Session
		result2 error
	}
	GetSessionByIDStub        func(context.Context, uuid.UUID) (*domain.Session, error)
	getSessionByIDMutex       sync.RWMutex
	getSessionByIDArgsForCall []struct {
		arg1 context.Context
		arg2 uuid.UUID
	}
	getSessionByIDReturns struct {
		result1 *domain.Session
		result2 error
	}
	getSessionByIDReturnsOnCall map[int]struct {
		result1 *domain.Session
		result2 error
	}
	GetSessionByRefreshTokenStub        func(context.Context, string) (*domain.Session, error)
	getSessionByRefreshTokenMutex       sync.RWMutex
	getSessionByRefreshTokenArgsForCall []struct {
		arg1 context.Context
		arg2 string
	}
	getSessionByRefreshTokenReturns struct {
		result1 *domain.Session
		result2 error
	}
	getSessionByRefreshTokenReturnsOnCall map[int]struct {
		result1 *domain.Session
		result2 error
	}
	UpdateRefreshTokenStub        func(context.Context, uuid.UUID, string) (*domain.Session, error)
	updateRefreshTokenMutex       sync.RWMutex
	updateRefreshTokenArgsForCall []struct {
		arg1 context.Context
		arg2 uuid.UUID
		arg3 string
	}
	updateRefreshTokenReturns struct {
		result1 *domain.Session
		result2 error
	}
	updateRefreshTokenReturnsOnCall map[int]struct {
		result1 *domain.Session
		result2 error
	}
	UpdateSessionStub        func(context.Context, domain.Session) error
	updateSessionMutex       sync.RWMutex
	updateSessionArgsForCall []struct {
		arg1 context.Context
		arg2 domain.Session
	}
	updateSessionReturns struct {
		result1 error
	}
	updateSessionReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeSessionRepository) BlockAllUserSessions(arg1 context.Context, arg2 uuid.UUID) error {
	fake.blockAllUserSessionsMutex.Lock()
	ret, specificReturn := fake.blockAllUserSessionsReturnsOnCall[len(fake.blockAllUserSessionsArgsForCall)]
	fake.blockAllUserSessionsArgsForCall = append(fake.blockAllUserSessionsArgsForCall, struct {
		arg1 context.Context
		arg2 uuid.UUID
	}{arg1, arg2})
	stub := fake.BlockAllUserSessionsStub
	fakeReturns := fake.blockAllUserSessionsReturns
	fake.recordInvocation("BlockAllUserSessions", []interface{}{arg1, arg2})
	fake.blockAllUserSessionsMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeSessionRepository) BlockAllUserSessionsCallCount() int {
	fake.blockAllUserSessionsMutex.RLock()
	defer fake.blockAllUserSessionsMutex.RUnlock()
	return len(fake.blockAllUserSessionsArgsForCall)
}

func (fake *FakeSessionRepository) BlockAllUserSessionsCalls(stub func(context.Context, uuid.UUID) error) {
	fake.blockAllUserSessionsMutex.Lock()
	defer fake.blockAllUserSessionsMutex.Unlock()
	fake.BlockAllUserSessionsStub = stub
}

func (fake *FakeSessionRepository) BlockAllUserSessionsArgsForCall(i int) (context.Context, uuid.UUID) {
	fake.blockAllUserSessionsMutex.RLock()
	defer fake.blockAllUserSessionsMutex.RUnlock()
	argsForCall := fake.blockAllUserSessionsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeSessionRepository) BlockAllUserSessionsReturns(result1 error) {
	fake.blockAllUserSessionsMutex.Lock()
	defer fake.blockAllUserSessionsMutex.Unlock()
	fake.BlockAllUserSessionsStub = nil
	fake.blockAllUserSessionsReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSessionRepository) BlockAllUserSessionsReturnsOnCall(i int, result1 error) {
	fake.blockAllUserSessionsMutex.Lock()
	defer fake.blockAllUserSessionsMutex.Unlock()
	fake.BlockAllUserSessionsStub = nil
	if fake.blockAllUserSessionsReturnsOnCall == nil {
		fake.blockAllUserSessionsReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.blockAllUserSessionsReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeSessionRepository) BlockSession(arg1 context.Context, arg2 uuid.UUID) error {
	fake.blockSessionMutex.Lock()
	ret, specificReturn := fake.blockSessionReturnsOnCall[len(fake.blockSessionArgsForCall)]
	fake.blockSessionArgsForCall = append(fake.blockSessionArgsForCall, struct {
		arg1 context.Context
		arg2 uuid.UUID
	}{arg1, arg2})
	stub := fake.BlockSessionStub
	fakeReturns := fake.blockSessionReturns
	fake.recordInvocation("BlockSession", []interface{}{arg1, arg2})
	fake.blockSessionMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeSessionRepository) BlockSessionCallCount() int {
	fake.blockSessionMutex.RLock()
	defer fake.blockSessionMutex.RUnlock()
	return len(fake.blockSessionArgsForCall)
}

func (fake *FakeSessionRepository) BlockSessionCalls(stub func(context.Context, uuid.UUID) error) {
	fake.blockSessionMutex.Lock()
	defer fake.blockSessionMutex.Unlock()
	fake.BlockSessionStub = stub
}

func (fake *FakeSessionRepository) BlockSessionArgsForCall(i int) (context.Context, uuid.UUID) {
	fake.blockSessionMutex.RLock()
	defer fake.blockSessionMutex.RUnlock()
	argsForCall := fake.blockSessionArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeSessionRepository) BlockSessionReturns(result1 error) {
	fake.blockSessionMutex.Lock()
	defer fake.blockSessionMutex.Unlock()
	fake.BlockSessionStub = nil
	fake.blockSessionReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSessionRepository) BlockSessionReturnsOnCall(i int, result1 error) {
	fake.blockSessionMutex.Lock()
	defer fake.blockSessionMutex.Unlock()
	fake.BlockSessionStub = nil
	if fake.blockSessionReturnsOnCall == nil {
		fake.blockSessionReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.blockSessionReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeSessionRepository) CreateSession(arg1 context.Context, arg2 domain.Session) (*domain.Session, error) {
	fake.createSessionMutex.Lock()
	ret, specificReturn := fake.createSessionReturnsOnCall[len(fake.createSessionArgsForCall)]
	fake.createSessionArgsForCall = append(fake.createSessionArgsForCall, struct {
		arg1 context.Context
		arg2 domain.Session
	}{arg1, arg2})
	stub := fake.CreateSessionStub
	fakeReturns := fake.createSessionReturns
	fake.recordInvocation("CreateSession", []interface{}{arg1, arg2})
	fake.createSessionMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeSessionRepository) CreateSessionCallCount() int {
	fake.createSessionMutex.RLock()
	defer fake.createSessionMutex.RUnlock()
	return len(fake.createSessionArgsForCall)
}

func (fake *FakeSessionRepository) CreateSessionCalls(stub func(context.Context, domain.Session) (*domain.Session, error)) {
	fake.createSessionMutex.Lock()
	defer fake.createSessionMutex.Unlock()
	fake.CreateSessionStub = stub
}

func (fake *FakeSessionRepository) CreateSessionArgsForCall(i int) (context.Context, domain.Session) {
	fake.createSessionMutex.RLock()
	defer fake.createSessionMutex.RUnlock()
	argsForCall := fake.createSessionArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeSessionRepository) CreateSessionReturns(result1 *domain.Session, result2 error) {
	fake.createSessionMutex.Lock()
	defer fake.createSessionMutex.Unlock()
	fake.CreateSessionStub = nil
	fake.createSessionReturns = struct {
		result1 *domain.Session
		result2 error
	}{result1, result2}
}

func (fake *FakeSessionRepository) CreateSessionReturnsOnCall(i int, result1 *domain.Session, result2 error) {
	fake.createSessionMutex.Lock()
	defer fake.createSessionMutex.Unlock()
	fake.CreateSessionStub = nil
	if fake.createSessionReturnsOnCall == nil {
		fake.createSessionReturnsOnCall = make(map[int]struct {
			result1 *domain.Session
			result2 error
		})
	}
	fake.createSessionReturnsOnCall[i] = struct {
		result1 *domain.Session
		result2 error
	}{result1, result2}
}

func (fake *FakeSessionRepository) DeleteSession(arg1 context.Context, arg2 uuid.UUID) error {
	fake.deleteSessionMutex.Lock()
	ret, specificReturn := fake.deleteSessionReturnsOnCall[len(fake.deleteSessionArgsForCall)]
	fake.deleteSessionArgsForCall = append(fake.deleteSessionArgsForCall, struct {
		arg1 context.Context
		arg2 uuid.UUID
	}{arg1, arg2})
	stub := fake.DeleteSessionStub
	fakeReturns := fake.deleteSessionReturns
	fake.recordInvocation("DeleteSession", []interface{}{arg1, arg2})
	fake.deleteSessionMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeSessionRepository) DeleteSessionCallCount() int {
	fake.deleteSessionMutex.RLock()
	defer fake.deleteSessionMutex.RUnlock()
	return len(fake.deleteSessionArgsForCall)
}

func (fake *FakeSessionRepository) DeleteSessionCalls(stub func(context.Context, uuid.UUID) error) {
	fake.deleteSessionMutex.Lock()
	defer fake.deleteSessionMutex.Unlock()
	fake.DeleteSessionStub = stub
}

func (fake *FakeSessionRepository) DeleteSessionArgsForCall(i int) (context.Context, uuid.UUID) {
	fake.deleteSessionMutex.RLock()
	defer fake.deleteSessionMutex.RUnlock()
	argsForCall := fake.deleteSessionArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeSessionRepository) DeleteSessionReturns(result1 error) {
	fake.deleteSessionMutex.Lock()
	defer fake.deleteSessionMutex.Unlock()
	fake.DeleteSessionStub = nil
	fake.deleteSessionReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSessionRepository) DeleteSessionReturnsOnCall(i int, result1 error) {
	fake.deleteSessionMutex.Lock()
	defer fake.deleteSessionMutex.Unlock()
	fake.DeleteSessionStub = nil
	if fake.deleteSessionReturnsOnCall == nil {
		fake.deleteSessionReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteSessionReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeSessionRepository) GetActiveSessionsByUserID(arg1 context.Context, arg2 uuid.UUID) ([]domain.Session, error) {
	fake.getActiveSessionsByUserIDMutex.Lock()
	ret, specificReturn := fake.getActiveSessionsByUserIDReturnsOnCall[len(fake.getActiveSessionsByUserIDArgsForCall)]
	fake.getActiveSessionsByUserIDArgsForCall = append(fake.getActiveSessionsByUserIDArgsForCall, struct {
		arg1 context.Context
		arg2 uuid.UUID
	}{arg1, arg2})
	stub := fake.GetActiveSessionsByUserIDStub
	fakeReturns := fake.getActiveSessionsByUserIDReturns
	fake.recordInvocation("GetActiveSessionsByUserID", []interface{}{arg1, arg2})
	fake.getActiveSessionsByUserIDMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeSessionRepository) GetActiveSessionsByUserIDCallCount() int {
	fake.getActiveSessionsByUserIDMutex.RLock()
	defer fake.getActiveSessionsByUserIDMutex.RUnlock()
	return len(fake.getActiveSessionsByUserIDArgsForCall)
}

func (fake *FakeSessionRepository) GetActiveSessionsByUserIDCalls(stub func(context.Context, uuid.UUID) ([]domain.Session, error)) {
	fake.getActiveSessionsByUserIDMutex.Lock()
	defer fake.getActiveSessionsByUserIDMutex.Unlock()
	fake.GetActiveSessionsByUserIDStub = stub
}

func (fake *FakeSessionRepository) GetActiveSessionsByUserIDArgsForCall(i int) (context.Context, uuid.UUID) {
	fake.getActiveSessionsByUserIDMutex.RLock()
	defer fake.getActiveSessionsByUserIDMutex.RUnlock()
	argsForCall := fake.getActiveSessionsByUserIDArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeSessionRepository) GetActiveSessionsByUserIDReturns(result1 []domain.Session, result2 error) {
	fake.getActiveSessionsByUserIDMutex.Lock()
	defer fake.getActiveSessionsByUserIDMutex.Unlock()
	fake.GetActiveSessionsByUserIDStub = nil
	fake.getActiveSessionsByUserIDReturns = struct {
		result1 []domain.Session
		result2 error
	}{result1, result2}
}

func (fake *FakeSessionRepository) GetActiveSessionsByUserIDReturnsOnCall(i int, result1 []domain.Session, result2 error) {
	fake.getActiveSessionsByUserIDMutex.Lock()
	defer fake.getActiveSessionsByUserIDMutex.Unlock()
	fake.GetActiveSessionsByUserIDStub = nil
	if fake.getActiveSessionsByUserIDReturnsOnCall == nil {
		fake.getActiveSessionsByUserIDReturnsOnCall = make(map[int]struct {
			result1 []domain.Session
			result2 error
		})
	}
	fake.getActiveSessionsByUserIDReturnsOnCall[i] = struct {
		result1 []domain.Session
		result2 error
	}{result1, result2}
}

func (fake *FakeSessionRepository) GetSessionByID(arg1 context.Context, arg2 uuid.UUID) (*domain.Session, error) {
	fake.getSessionByIDMutex.Lock()
	ret, specificReturn := fake.getSessionByIDReturnsOnCall[len(fake.getSessionByIDArgsForCall)]
	fake.getSessionByIDArgsForCall = append(fake.getSessionByIDArgsForCall, struct {
		arg1 context.Context
		arg2 uuid.UUID
	}{arg1, arg2})
	stub := fake.GetSessionByIDStub
	fakeReturns := fake.getSessionByIDReturns
	fake.recordInvocation("GetSessionByID", []interface{}{arg1, arg2})
	fake.getSessionByIDMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeSessionRepository) GetSessionByIDCallCount() int {
	fake.getSessionByIDMutex.RLock()
	defer fake.getSessionByIDMutex.RUnlock()
	return len(fake.getSessionByIDArgsForCall)
}

func (fake *FakeSessionRepository) GetSessionByIDCalls(stub func(context.Context, uuid.UUID) (*domain.Session, error)) {
	fake.getSessionByIDMutex.Lock()
	defer fake.getSessionByIDMutex.Unlock()
	fake.GetSessionByIDStub = stub
}

func (fake *FakeSessionRepository) GetSessionByIDArgsForCall(i int) (context.Context, uuid.UUID) {
	fake.getSessionByIDMutex.RLock()
	defer fake.getSessionByIDMutex.RUnlock()
	argsForCall := fake.getSessionByIDArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeSessionRepository) GetSessionByIDReturns(result1 *domain.Session, result2 error) {
	fake.getSessionByIDMutex.Lock()
	defer fake.getSessionByIDMutex.Unlock()
	fake.GetSessionByIDStub = nil
	fake.getSessionByIDReturns = struct {
		result1 *domain.Session
		result2 error
	}{result1, result2}
}

func (fake *FakeSessionRepository) GetSessionByIDReturnsOnCall(i int, result1 *domain.Session, result2 error) {
	fake.getSessionByIDMutex.Lock()
	defer fake.getSessionByIDMutex.Unlock()
	fake.GetSessionByIDStub = nil
	if fake.getSessionByIDReturnsOnCall == nil {
		fake.getSessionByIDReturnsOnCall = make(map[int]struct {
			result1 *domain.Session
			result2 error
		})
	}
	fake.getSessionByIDReturnsOnCall[i] = struct {
		result1 *domain.Session
		result2 error
	}{result1, result2}
}

func (fake *FakeSessionRepository) GetSessionByRefreshToken(arg1 context.Context, arg2 string) (*domain.Session, error) {
	fake.getSessionByRefreshTokenMutex.Lock()
	ret, specificReturn := fake.getSessionByRefreshTokenReturnsOnCall[len(fake.getSessionByRefreshTokenArgsForCall)]
	fake.getSessionByRefreshTokenArgsForCall = append(fake.getSessionByRefreshTokenArgsForCall, struct {
		arg1 context.Context
		arg2 string
	}{arg1, arg2})
	stub := fake.GetSessionByRefreshTokenStub
	fakeReturns := fake.getSessionByRefreshTokenReturns
	fake.recordInvocation("GetSessionByRefreshToken", []interface{}{arg1, arg2})
	fake.getSessionByRefreshTokenMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeSessionRepository) GetSessionByRefreshTokenCallCount() int {
	fake.getSessionByRefreshTokenMutex.RLock()
	defer fake.getSessionByRefreshTokenMutex.RUnlock()
	return len(fake.getSessionByRefreshTokenArgsForCall)
}

func (fake *FakeSessionRepository) GetSessionByRefreshTokenCalls(stub func(context.Context, string) (*domain.Session, error)) {
	fake.getSessionByRefreshTokenMutex.Lock()
	defer fake.getSessionByRefreshTokenMutex.Unlock()
	fake.GetSessionByRefreshTokenStub = stub
}

func (fake *FakeSessionRepository) GetSessionByRefreshTokenArgsForCall(i int) (context.Context, string) {
	fake.getSessionByRefreshTokenMutex.RLock()
	defer fake.getSessionByRefreshTokenMutex.RUnlock()
	argsForCall := fake.getSessionByRefreshTokenArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeSessionRepository) GetSessionByRefreshTokenReturns(result1 *domain.Session, result2 error) {
	fake.getSessionByRefreshTokenMutex.Lock()
	defer fake.getSessionByRefreshTokenMutex.Unlock()
	fake.GetSessionByRefreshTokenStub = nil
	fake.getSessionByRefreshTokenReturns = struct {
		result1 *domain.Session
		result2 error
	}{result1, result2}
}

func (fake *FakeSessionRepository) GetSessionByRefreshTokenReturnsOnCall(i int, result1 *domain.Session, result2 error) {
	fake.getSessionByRefreshTokenMutex.Lock()
	defer fake.getSessionByRefreshTokenMutex.Unlock()
	fake.GetSessionByRefreshTokenStub = nil
	if fake.getSessionByRefreshTokenReturnsOnCall == nil {
		fake.getSessionByRefreshTokenReturnsOnCall = make(map[int]struct {
			result1 *domain.Session
			result2 error
		})
	}
	fake.getSessionByRefreshTokenReturnsOnCall[i] = struct {
		result1 *domain.Session
		result2 error
	}{result1, result2}
}

func (fake *FakeSessionRepository) UpdateRefreshToken(arg1 context.Context, arg2 uuid.UUID, arg3 string) (*domain.Session, error) {
	fake.updateRefreshTokenMutex.Lock()
	ret, specificReturn := fake.updateRefreshTokenReturnsOnCall[len(fake.updateRefreshTokenArgsForCall)]
	fake.updateRefreshTokenArgsForCall = append(fake.updateRefreshTokenArgsForCall, struct {
		arg1 context.Context
		arg2 uuid.UUID
		arg3 string
	}{arg1, arg2, arg3})
	stub := fake.UpdateRefreshTokenStub
	fakeReturns := fake.updateRefreshTokenReturns
	fake.recordInvocation("UpdateRefreshToken", []interface{}{arg1, arg2, arg3})
	fake.updateRefreshTokenMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeSessionRepository) UpdateRefreshTokenCallCount() int {
	fake.updateRefreshTokenMutex.RLock()
	defer fake.updateRefreshTokenMutex.RUnlock()
	return len(fake.updateRefreshTokenArgsForCall)
}

func (fake *FakeSessionRepository) UpdateRefreshTokenCalls(stub func(context.Context, uuid.UUID, string) (*domain.Session, error)) {
	fake.updateRefreshTokenMutex.Lock()
	defer fake.updateRefreshTokenMutex.Unlock()
	fake.UpdateRefreshTokenStub = stub
}

func (fake *FakeSessionRepository) UpdateRefreshTokenArgsForCall(i int) (context.Context, uuid.UUID, string) {
	fake.updateRefreshTokenMutex.RLock()
	defer fake.updateRefreshTokenMutex.RUnlock()
	argsForCall := fake.updateRefreshTokenArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeSessionRepository) UpdateRefreshTokenReturns(result1 *domain.Session, result2 error) {
	fake.updateRefreshTokenMutex.Lock()
	defer fake.updateRefreshTokenMutex.Unlock()
	fake.UpdateRefreshTokenStub = nil
	fake.updateRefreshTokenReturns = struct {
		result1 *domain.Session
		result2 error
	}{result1, result2}
}

func (fake *FakeSessionRepository) UpdateRefreshTokenReturnsOnCall(i int, result1 *domain.Session, result2 error) {
	fake.updateRefreshTokenMutex.Lock()
	defer fake.updateRefreshTokenMutex.Unlock()
	fake.UpdateRefreshTokenStub = nil
	if fake.updateRefreshTokenReturnsOnCall == nil {
		fake.updateRefreshTokenReturnsOnCall = make(map[int]struct {
			result1 *domain.Session
			result2 error
		})
	}
	fake.updateRefreshTokenReturnsOnCall[i] = struct {
		result1 *domain.Session
		result2 error
	}{result1, result2}
}

func (fake *FakeSessionRepository) UpdateSession(arg1 context.Context, arg2 domain.Session) error {
	fake.updateSessionMutex.Lock()
	ret, specificReturn := fake.updateSessionReturnsOnCall[len(fake.updateSessionArgsForCall)]
	fake.updateSessionArgsForCall = append(fake.updateSessionArgsForCall, struct {
		arg1 context.Context
		arg2 domain.Session
	}{arg1, arg2})
	stub := fake.UpdateSessionStub
	fakeReturns := fake.updateSessionReturns
	fake.recordInvocation("UpdateSession", []interface{}{arg1, arg2})
	fake.updateSessionMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeSessionRepository) UpdateSessionCallCount() int {
	fake.updateSessionMutex.RLock()
	defer fake.updateSessionMutex.RUnlock()
	return len(fake.updateSessionArgsForCall)
}

func (fake *FakeSessionRepository) UpdateSessionCalls(stub func(context.Context, domain.Session) error) {
	fake.updateSessionMutex.Lock()
	defer fake.updateSessionMutex.Unlock()
	fake.UpdateSessionStub = stub
}

func (fake *FakeSessionRepository) UpdateSessionArgsForCall(i int) (context.Context, domain.Session) {
	fake.updateSessionMutex.RLock()
	defer fake.updateSessionMutex.RUnlock()
	argsForCall := fake.updateSessionArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeSessionRepository) UpdateSessionReturns(result1 error) {
	fake.updateSessionMutex.Lock()
	defer fake.updateSessionMutex.Unlock()
	fake.UpdateSessionStub = nil
	fake.updateSessionReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSessionRepository) UpdateSessionReturnsOnCall(i int, result1 error) {
	fake.updateSessionMutex.Lock()
	defer fake.updateSessionMutex.Unlock()
	fake.UpdateSessionStub = nil
	if fake.updateSessionReturnsOnCall == nil {
		fake.updateSessionReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.updateSessionReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeSessionRepository) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.blockAllUserSessionsMutex.RLock()
	defer fake.blockAllUserSessionsMutex.RUnlock()
	fake.blockSessionMutex.RLock()
	defer fake.blockSessionMutex.RUnlock()
	fake.createSessionMutex.RLock()
	defer fake.createSessionMutex.RUnlock()
	fake.deleteSessionMutex.RLock()
	defer fake.deleteSessionMutex.RUnlock()
	fake.getActiveSessionsByUserIDMutex.RLock()
	defer fake.getActiveSessionsByUserIDMutex.RUnlock()
	fake.getSessionByIDMutex.RLock()
	defer fake.getSessionByIDMutex.RUnlock()
	fake.getSessionByRefreshTokenMutex.RLock()
	defer fake.getSessionByRefreshTokenMutex.RUnlock()
	fake.updateRefreshTokenMutex.RLock()
	defer fake.updateRefreshTokenMutex.RUnlock()
	fake.updateSessionMutex.RLock()
	defer fake.updateSessionMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeSessionRepository) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ ports.SessionRepository = new(FakeSessionRepository)
