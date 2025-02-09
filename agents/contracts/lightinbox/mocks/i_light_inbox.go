// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	big "math/big"

	bind "github.com/ethereum/go-ethereum/accounts/abi/bind"
	common "github.com/ethereum/go-ethereum/common"

	event "github.com/ethereum/go-ethereum/event"

	lightinbox "github.com/synapsecns/sanguine/agents/contracts/lightinbox"

	mock "github.com/stretchr/testify/mock"

	types "github.com/ethereum/go-ethereum/core/types"
)

// ILightInbox is an autogenerated mock type for the ILightInbox type
type ILightInbox struct {
	mock.Mock
}

// Address provides a mock function with given fields:
func (_m *ILightInbox) Address() common.Address {
	ret := _m.Called()

	var r0 common.Address
	if rf, ok := ret.Get(0).(func() common.Address); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Address)
		}
	}

	return r0
}

// AgentManager provides a mock function with given fields: opts
func (_m *ILightInbox) AgentManager(opts *bind.CallOpts) (common.Address, error) {
	ret := _m.Called(opts)

	var r0 common.Address
	if rf, ok := ret.Get(0).(func(*bind.CallOpts) common.Address); ok {
		r0 = rf(opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Address)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.CallOpts) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Destination provides a mock function with given fields: opts
func (_m *ILightInbox) Destination(opts *bind.CallOpts) (common.Address, error) {
	ret := _m.Called(opts)

	var r0 common.Address
	if rf, ok := ret.Get(0).(func(*bind.CallOpts) common.Address); ok {
		r0 = rf(opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Address)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.CallOpts) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FilterAttestationAccepted provides a mock function with given fields: opts
func (_m *ILightInbox) FilterAttestationAccepted(opts *bind.FilterOpts) (*lightinbox.LightInboxAttestationAcceptedIterator, error) {
	ret := _m.Called(opts)

	var r0 *lightinbox.LightInboxAttestationAcceptedIterator
	if rf, ok := ret.Get(0).(func(*bind.FilterOpts) *lightinbox.LightInboxAttestationAcceptedIterator); ok {
		r0 = rf(opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*lightinbox.LightInboxAttestationAcceptedIterator)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.FilterOpts) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FilterInitialized provides a mock function with given fields: opts
func (_m *ILightInbox) FilterInitialized(opts *bind.FilterOpts) (*lightinbox.LightInboxInitializedIterator, error) {
	ret := _m.Called(opts)

	var r0 *lightinbox.LightInboxInitializedIterator
	if rf, ok := ret.Get(0).(func(*bind.FilterOpts) *lightinbox.LightInboxInitializedIterator); ok {
		r0 = rf(opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*lightinbox.LightInboxInitializedIterator)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.FilterOpts) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FilterInvalidReceipt provides a mock function with given fields: opts
func (_m *ILightInbox) FilterInvalidReceipt(opts *bind.FilterOpts) (*lightinbox.LightInboxInvalidReceiptIterator, error) {
	ret := _m.Called(opts)

	var r0 *lightinbox.LightInboxInvalidReceiptIterator
	if rf, ok := ret.Get(0).(func(*bind.FilterOpts) *lightinbox.LightInboxInvalidReceiptIterator); ok {
		r0 = rf(opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*lightinbox.LightInboxInvalidReceiptIterator)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.FilterOpts) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FilterInvalidReceiptReport provides a mock function with given fields: opts
func (_m *ILightInbox) FilterInvalidReceiptReport(opts *bind.FilterOpts) (*lightinbox.LightInboxInvalidReceiptReportIterator, error) {
	ret := _m.Called(opts)

	var r0 *lightinbox.LightInboxInvalidReceiptReportIterator
	if rf, ok := ret.Get(0).(func(*bind.FilterOpts) *lightinbox.LightInboxInvalidReceiptReportIterator); ok {
		r0 = rf(opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*lightinbox.LightInboxInvalidReceiptReportIterator)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.FilterOpts) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FilterInvalidStateReport provides a mock function with given fields: opts
func (_m *ILightInbox) FilterInvalidStateReport(opts *bind.FilterOpts) (*lightinbox.LightInboxInvalidStateReportIterator, error) {
	ret := _m.Called(opts)

	var r0 *lightinbox.LightInboxInvalidStateReportIterator
	if rf, ok := ret.Get(0).(func(*bind.FilterOpts) *lightinbox.LightInboxInvalidStateReportIterator); ok {
		r0 = rf(opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*lightinbox.LightInboxInvalidStateReportIterator)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.FilterOpts) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FilterInvalidStateWithAttestation provides a mock function with given fields: opts
func (_m *ILightInbox) FilterInvalidStateWithAttestation(opts *bind.FilterOpts) (*lightinbox.LightInboxInvalidStateWithAttestationIterator, error) {
	ret := _m.Called(opts)

	var r0 *lightinbox.LightInboxInvalidStateWithAttestationIterator
	if rf, ok := ret.Get(0).(func(*bind.FilterOpts) *lightinbox.LightInboxInvalidStateWithAttestationIterator); ok {
		r0 = rf(opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*lightinbox.LightInboxInvalidStateWithAttestationIterator)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.FilterOpts) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FilterInvalidStateWithSnapshot provides a mock function with given fields: opts
func (_m *ILightInbox) FilterInvalidStateWithSnapshot(opts *bind.FilterOpts) (*lightinbox.LightInboxInvalidStateWithSnapshotIterator, error) {
	ret := _m.Called(opts)

	var r0 *lightinbox.LightInboxInvalidStateWithSnapshotIterator
	if rf, ok := ret.Get(0).(func(*bind.FilterOpts) *lightinbox.LightInboxInvalidStateWithSnapshotIterator); ok {
		r0 = rf(opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*lightinbox.LightInboxInvalidStateWithSnapshotIterator)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.FilterOpts) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FilterOwnershipTransferred provides a mock function with given fields: opts, previousOwner, newOwner
func (_m *ILightInbox) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*lightinbox.LightInboxOwnershipTransferredIterator, error) {
	ret := _m.Called(opts, previousOwner, newOwner)

	var r0 *lightinbox.LightInboxOwnershipTransferredIterator
	if rf, ok := ret.Get(0).(func(*bind.FilterOpts, []common.Address, []common.Address) *lightinbox.LightInboxOwnershipTransferredIterator); ok {
		r0 = rf(opts, previousOwner, newOwner)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*lightinbox.LightInboxOwnershipTransferredIterator)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.FilterOpts, []common.Address, []common.Address) error); ok {
		r1 = rf(opts, previousOwner, newOwner)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetGuardReport provides a mock function with given fields: opts, index
func (_m *ILightInbox) GetGuardReport(opts *bind.CallOpts, index *big.Int) (struct {
	StatementPayload []byte
	ReportSignature  []byte
}, error) {
	ret := _m.Called(opts, index)

	var r0 struct {
		StatementPayload []byte
		ReportSignature  []byte
	}
	if rf, ok := ret.Get(0).(func(*bind.CallOpts, *big.Int) struct {
		StatementPayload []byte
		ReportSignature  []byte
	}); ok {
		r0 = rf(opts, index)
	} else {
		r0 = ret.Get(0).(struct {
			StatementPayload []byte
			ReportSignature  []byte
		})
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.CallOpts, *big.Int) error); ok {
		r1 = rf(opts, index)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetReportsAmount provides a mock function with given fields: opts
func (_m *ILightInbox) GetReportsAmount(opts *bind.CallOpts) (*big.Int, error) {
	ret := _m.Called(opts)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(*bind.CallOpts) *big.Int); ok {
		r0 = rf(opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.CallOpts) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStoredSignature provides a mock function with given fields: opts, index
func (_m *ILightInbox) GetStoredSignature(opts *bind.CallOpts, index *big.Int) ([]byte, error) {
	ret := _m.Called(opts, index)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(*bind.CallOpts, *big.Int) []byte); ok {
		r0 = rf(opts, index)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.CallOpts, *big.Int) error); ok {
		r1 = rf(opts, index)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Initialize provides a mock function with given fields: opts, agentManager_, origin_, destination_
func (_m *ILightInbox) Initialize(opts *bind.TransactOpts, agentManager_ common.Address, origin_ common.Address, destination_ common.Address) (*types.Transaction, error) {
	ret := _m.Called(opts, agentManager_, origin_, destination_)

	var r0 *types.Transaction
	if rf, ok := ret.Get(0).(func(*bind.TransactOpts, common.Address, common.Address, common.Address) *types.Transaction); ok {
		r0 = rf(opts, agentManager_, origin_, destination_)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.TransactOpts, common.Address, common.Address, common.Address) error); ok {
		r1 = rf(opts, agentManager_, origin_, destination_)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LocalDomain provides a mock function with given fields: opts
func (_m *ILightInbox) LocalDomain(opts *bind.CallOpts) (uint32, error) {
	ret := _m.Called(opts)

	var r0 uint32
	if rf, ok := ret.Get(0).(func(*bind.CallOpts) uint32); ok {
		r0 = rf(opts)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.CallOpts) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Multicall provides a mock function with given fields: opts, calls
func (_m *ILightInbox) Multicall(opts *bind.TransactOpts, calls []lightinbox.MultiCallableCall) (*types.Transaction, error) {
	ret := _m.Called(opts, calls)

	var r0 *types.Transaction
	if rf, ok := ret.Get(0).(func(*bind.TransactOpts, []lightinbox.MultiCallableCall) *types.Transaction); ok {
		r0 = rf(opts, calls)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.TransactOpts, []lightinbox.MultiCallableCall) error); ok {
		r1 = rf(opts, calls)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Origin provides a mock function with given fields: opts
func (_m *ILightInbox) Origin(opts *bind.CallOpts) (common.Address, error) {
	ret := _m.Called(opts)

	var r0 common.Address
	if rf, ok := ret.Get(0).(func(*bind.CallOpts) common.Address); ok {
		r0 = rf(opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Address)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.CallOpts) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Owner provides a mock function with given fields: opts
func (_m *ILightInbox) Owner(opts *bind.CallOpts) (common.Address, error) {
	ret := _m.Called(opts)

	var r0 common.Address
	if rf, ok := ret.Get(0).(func(*bind.CallOpts) common.Address); ok {
		r0 = rf(opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Address)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.CallOpts) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ParseAttestationAccepted provides a mock function with given fields: log
func (_m *ILightInbox) ParseAttestationAccepted(log types.Log) (*lightinbox.LightInboxAttestationAccepted, error) {
	ret := _m.Called(log)

	var r0 *lightinbox.LightInboxAttestationAccepted
	if rf, ok := ret.Get(0).(func(types.Log) *lightinbox.LightInboxAttestationAccepted); ok {
		r0 = rf(log)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*lightinbox.LightInboxAttestationAccepted)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(types.Log) error); ok {
		r1 = rf(log)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ParseInitialized provides a mock function with given fields: log
func (_m *ILightInbox) ParseInitialized(log types.Log) (*lightinbox.LightInboxInitialized, error) {
	ret := _m.Called(log)

	var r0 *lightinbox.LightInboxInitialized
	if rf, ok := ret.Get(0).(func(types.Log) *lightinbox.LightInboxInitialized); ok {
		r0 = rf(log)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*lightinbox.LightInboxInitialized)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(types.Log) error); ok {
		r1 = rf(log)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ParseInvalidReceipt provides a mock function with given fields: log
func (_m *ILightInbox) ParseInvalidReceipt(log types.Log) (*lightinbox.LightInboxInvalidReceipt, error) {
	ret := _m.Called(log)

	var r0 *lightinbox.LightInboxInvalidReceipt
	if rf, ok := ret.Get(0).(func(types.Log) *lightinbox.LightInboxInvalidReceipt); ok {
		r0 = rf(log)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*lightinbox.LightInboxInvalidReceipt)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(types.Log) error); ok {
		r1 = rf(log)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ParseInvalidReceiptReport provides a mock function with given fields: log
func (_m *ILightInbox) ParseInvalidReceiptReport(log types.Log) (*lightinbox.LightInboxInvalidReceiptReport, error) {
	ret := _m.Called(log)

	var r0 *lightinbox.LightInboxInvalidReceiptReport
	if rf, ok := ret.Get(0).(func(types.Log) *lightinbox.LightInboxInvalidReceiptReport); ok {
		r0 = rf(log)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*lightinbox.LightInboxInvalidReceiptReport)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(types.Log) error); ok {
		r1 = rf(log)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ParseInvalidStateReport provides a mock function with given fields: log
func (_m *ILightInbox) ParseInvalidStateReport(log types.Log) (*lightinbox.LightInboxInvalidStateReport, error) {
	ret := _m.Called(log)

	var r0 *lightinbox.LightInboxInvalidStateReport
	if rf, ok := ret.Get(0).(func(types.Log) *lightinbox.LightInboxInvalidStateReport); ok {
		r0 = rf(log)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*lightinbox.LightInboxInvalidStateReport)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(types.Log) error); ok {
		r1 = rf(log)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ParseInvalidStateWithAttestation provides a mock function with given fields: log
func (_m *ILightInbox) ParseInvalidStateWithAttestation(log types.Log) (*lightinbox.LightInboxInvalidStateWithAttestation, error) {
	ret := _m.Called(log)

	var r0 *lightinbox.LightInboxInvalidStateWithAttestation
	if rf, ok := ret.Get(0).(func(types.Log) *lightinbox.LightInboxInvalidStateWithAttestation); ok {
		r0 = rf(log)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*lightinbox.LightInboxInvalidStateWithAttestation)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(types.Log) error); ok {
		r1 = rf(log)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ParseInvalidStateWithSnapshot provides a mock function with given fields: log
func (_m *ILightInbox) ParseInvalidStateWithSnapshot(log types.Log) (*lightinbox.LightInboxInvalidStateWithSnapshot, error) {
	ret := _m.Called(log)

	var r0 *lightinbox.LightInboxInvalidStateWithSnapshot
	if rf, ok := ret.Get(0).(func(types.Log) *lightinbox.LightInboxInvalidStateWithSnapshot); ok {
		r0 = rf(log)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*lightinbox.LightInboxInvalidStateWithSnapshot)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(types.Log) error); ok {
		r1 = rf(log)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ParseOwnershipTransferred provides a mock function with given fields: log
func (_m *ILightInbox) ParseOwnershipTransferred(log types.Log) (*lightinbox.LightInboxOwnershipTransferred, error) {
	ret := _m.Called(log)

	var r0 *lightinbox.LightInboxOwnershipTransferred
	if rf, ok := ret.Get(0).(func(types.Log) *lightinbox.LightInboxOwnershipTransferred); ok {
		r0 = rf(log)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*lightinbox.LightInboxOwnershipTransferred)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(types.Log) error); ok {
		r1 = rf(log)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RenounceOwnership provides a mock function with given fields: opts
func (_m *ILightInbox) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	ret := _m.Called(opts)

	var r0 *types.Transaction
	if rf, ok := ret.Get(0).(func(*bind.TransactOpts) *types.Transaction); ok {
		r0 = rf(opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.TransactOpts) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SubmitAttestation provides a mock function with given fields: opts, attPayload, attSignature, agentRoot_, snapGas_
func (_m *ILightInbox) SubmitAttestation(opts *bind.TransactOpts, attPayload []byte, attSignature []byte, agentRoot_ [32]byte, snapGas_ []*big.Int) (*types.Transaction, error) {
	ret := _m.Called(opts, attPayload, attSignature, agentRoot_, snapGas_)

	var r0 *types.Transaction
	if rf, ok := ret.Get(0).(func(*bind.TransactOpts, []byte, []byte, [32]byte, []*big.Int) *types.Transaction); ok {
		r0 = rf(opts, attPayload, attSignature, agentRoot_, snapGas_)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.TransactOpts, []byte, []byte, [32]byte, []*big.Int) error); ok {
		r1 = rf(opts, attPayload, attSignature, agentRoot_, snapGas_)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SubmitAttestationReport provides a mock function with given fields: opts, attPayload, arSignature, attSignature
func (_m *ILightInbox) SubmitAttestationReport(opts *bind.TransactOpts, attPayload []byte, arSignature []byte, attSignature []byte) (*types.Transaction, error) {
	ret := _m.Called(opts, attPayload, arSignature, attSignature)

	var r0 *types.Transaction
	if rf, ok := ret.Get(0).(func(*bind.TransactOpts, []byte, []byte, []byte) *types.Transaction); ok {
		r0 = rf(opts, attPayload, arSignature, attSignature)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.TransactOpts, []byte, []byte, []byte) error); ok {
		r1 = rf(opts, attPayload, arSignature, attSignature)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SubmitStateReportWithAttestation provides a mock function with given fields: opts, stateIndex, srSignature, snapPayload, attPayload, attSignature
func (_m *ILightInbox) SubmitStateReportWithAttestation(opts *bind.TransactOpts, stateIndex *big.Int, srSignature []byte, snapPayload []byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	ret := _m.Called(opts, stateIndex, srSignature, snapPayload, attPayload, attSignature)

	var r0 *types.Transaction
	if rf, ok := ret.Get(0).(func(*bind.TransactOpts, *big.Int, []byte, []byte, []byte, []byte) *types.Transaction); ok {
		r0 = rf(opts, stateIndex, srSignature, snapPayload, attPayload, attSignature)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.TransactOpts, *big.Int, []byte, []byte, []byte, []byte) error); ok {
		r1 = rf(opts, stateIndex, srSignature, snapPayload, attPayload, attSignature)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SubmitStateReportWithSnapshot provides a mock function with given fields: opts, stateIndex, srSignature, snapPayload, snapSignature
func (_m *ILightInbox) SubmitStateReportWithSnapshot(opts *bind.TransactOpts, stateIndex *big.Int, srSignature []byte, snapPayload []byte, snapSignature []byte) (*types.Transaction, error) {
	ret := _m.Called(opts, stateIndex, srSignature, snapPayload, snapSignature)

	var r0 *types.Transaction
	if rf, ok := ret.Get(0).(func(*bind.TransactOpts, *big.Int, []byte, []byte, []byte) *types.Transaction); ok {
		r0 = rf(opts, stateIndex, srSignature, snapPayload, snapSignature)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.TransactOpts, *big.Int, []byte, []byte, []byte) error); ok {
		r1 = rf(opts, stateIndex, srSignature, snapPayload, snapSignature)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SubmitStateReportWithSnapshotProof provides a mock function with given fields: opts, stateIndex, statePayload, srSignature, snapProof, attPayload, attSignature
func (_m *ILightInbox) SubmitStateReportWithSnapshotProof(opts *bind.TransactOpts, stateIndex *big.Int, statePayload []byte, srSignature []byte, snapProof [][32]byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	ret := _m.Called(opts, stateIndex, statePayload, srSignature, snapProof, attPayload, attSignature)

	var r0 *types.Transaction
	if rf, ok := ret.Get(0).(func(*bind.TransactOpts, *big.Int, []byte, []byte, [][32]byte, []byte, []byte) *types.Transaction); ok {
		r0 = rf(opts, stateIndex, statePayload, srSignature, snapProof, attPayload, attSignature)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.TransactOpts, *big.Int, []byte, []byte, [][32]byte, []byte, []byte) error); ok {
		r1 = rf(opts, stateIndex, statePayload, srSignature, snapProof, attPayload, attSignature)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SynapseDomain provides a mock function with given fields: opts
func (_m *ILightInbox) SynapseDomain(opts *bind.CallOpts) (uint32, error) {
	ret := _m.Called(opts)

	var r0 uint32
	if rf, ok := ret.Get(0).(func(*bind.CallOpts) uint32); ok {
		r0 = rf(opts)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.CallOpts) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TransferOwnership provides a mock function with given fields: opts, newOwner
func (_m *ILightInbox) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	ret := _m.Called(opts, newOwner)

	var r0 *types.Transaction
	if rf, ok := ret.Get(0).(func(*bind.TransactOpts, common.Address) *types.Transaction); ok {
		r0 = rf(opts, newOwner)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.TransactOpts, common.Address) error); ok {
		r1 = rf(opts, newOwner)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// VerifyReceipt provides a mock function with given fields: opts, rcptPayload, rcptSignature
func (_m *ILightInbox) VerifyReceipt(opts *bind.TransactOpts, rcptPayload []byte, rcptSignature []byte) (*types.Transaction, error) {
	ret := _m.Called(opts, rcptPayload, rcptSignature)

	var r0 *types.Transaction
	if rf, ok := ret.Get(0).(func(*bind.TransactOpts, []byte, []byte) *types.Transaction); ok {
		r0 = rf(opts, rcptPayload, rcptSignature)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.TransactOpts, []byte, []byte) error); ok {
		r1 = rf(opts, rcptPayload, rcptSignature)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// VerifyReceiptReport provides a mock function with given fields: opts, rcptPayload, rrSignature
func (_m *ILightInbox) VerifyReceiptReport(opts *bind.TransactOpts, rcptPayload []byte, rrSignature []byte) (*types.Transaction, error) {
	ret := _m.Called(opts, rcptPayload, rrSignature)

	var r0 *types.Transaction
	if rf, ok := ret.Get(0).(func(*bind.TransactOpts, []byte, []byte) *types.Transaction); ok {
		r0 = rf(opts, rcptPayload, rrSignature)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.TransactOpts, []byte, []byte) error); ok {
		r1 = rf(opts, rcptPayload, rrSignature)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// VerifyStateReport provides a mock function with given fields: opts, statePayload, srSignature
func (_m *ILightInbox) VerifyStateReport(opts *bind.TransactOpts, statePayload []byte, srSignature []byte) (*types.Transaction, error) {
	ret := _m.Called(opts, statePayload, srSignature)

	var r0 *types.Transaction
	if rf, ok := ret.Get(0).(func(*bind.TransactOpts, []byte, []byte) *types.Transaction); ok {
		r0 = rf(opts, statePayload, srSignature)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.TransactOpts, []byte, []byte) error); ok {
		r1 = rf(opts, statePayload, srSignature)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// VerifyStateWithAttestation provides a mock function with given fields: opts, stateIndex, snapPayload, attPayload, attSignature
func (_m *ILightInbox) VerifyStateWithAttestation(opts *bind.TransactOpts, stateIndex *big.Int, snapPayload []byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	ret := _m.Called(opts, stateIndex, snapPayload, attPayload, attSignature)

	var r0 *types.Transaction
	if rf, ok := ret.Get(0).(func(*bind.TransactOpts, *big.Int, []byte, []byte, []byte) *types.Transaction); ok {
		r0 = rf(opts, stateIndex, snapPayload, attPayload, attSignature)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.TransactOpts, *big.Int, []byte, []byte, []byte) error); ok {
		r1 = rf(opts, stateIndex, snapPayload, attPayload, attSignature)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// VerifyStateWithSnapshot provides a mock function with given fields: opts, stateIndex, snapPayload, snapSignature
func (_m *ILightInbox) VerifyStateWithSnapshot(opts *bind.TransactOpts, stateIndex *big.Int, snapPayload []byte, snapSignature []byte) (*types.Transaction, error) {
	ret := _m.Called(opts, stateIndex, snapPayload, snapSignature)

	var r0 *types.Transaction
	if rf, ok := ret.Get(0).(func(*bind.TransactOpts, *big.Int, []byte, []byte) *types.Transaction); ok {
		r0 = rf(opts, stateIndex, snapPayload, snapSignature)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.TransactOpts, *big.Int, []byte, []byte) error); ok {
		r1 = rf(opts, stateIndex, snapPayload, snapSignature)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// VerifyStateWithSnapshotProof provides a mock function with given fields: opts, stateIndex, statePayload, snapProof, attPayload, attSignature
func (_m *ILightInbox) VerifyStateWithSnapshotProof(opts *bind.TransactOpts, stateIndex *big.Int, statePayload []byte, snapProof [][32]byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	ret := _m.Called(opts, stateIndex, statePayload, snapProof, attPayload, attSignature)

	var r0 *types.Transaction
	if rf, ok := ret.Get(0).(func(*bind.TransactOpts, *big.Int, []byte, [][32]byte, []byte, []byte) *types.Transaction); ok {
		r0 = rf(opts, stateIndex, statePayload, snapProof, attPayload, attSignature)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.TransactOpts, *big.Int, []byte, [][32]byte, []byte, []byte) error); ok {
		r1 = rf(opts, stateIndex, statePayload, snapProof, attPayload, attSignature)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Version provides a mock function with given fields: opts
func (_m *ILightInbox) Version(opts *bind.CallOpts) (string, error) {
	ret := _m.Called(opts)

	var r0 string
	if rf, ok := ret.Get(0).(func(*bind.CallOpts) string); ok {
		r0 = rf(opts)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.CallOpts) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WatchAttestationAccepted provides a mock function with given fields: opts, sink
func (_m *ILightInbox) WatchAttestationAccepted(opts *bind.WatchOpts, sink chan<- *lightinbox.LightInboxAttestationAccepted) (event.Subscription, error) {
	ret := _m.Called(opts, sink)

	var r0 event.Subscription
	if rf, ok := ret.Get(0).(func(*bind.WatchOpts, chan<- *lightinbox.LightInboxAttestationAccepted) event.Subscription); ok {
		r0 = rf(opts, sink)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(event.Subscription)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.WatchOpts, chan<- *lightinbox.LightInboxAttestationAccepted) error); ok {
		r1 = rf(opts, sink)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WatchInitialized provides a mock function with given fields: opts, sink
func (_m *ILightInbox) WatchInitialized(opts *bind.WatchOpts, sink chan<- *lightinbox.LightInboxInitialized) (event.Subscription, error) {
	ret := _m.Called(opts, sink)

	var r0 event.Subscription
	if rf, ok := ret.Get(0).(func(*bind.WatchOpts, chan<- *lightinbox.LightInboxInitialized) event.Subscription); ok {
		r0 = rf(opts, sink)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(event.Subscription)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.WatchOpts, chan<- *lightinbox.LightInboxInitialized) error); ok {
		r1 = rf(opts, sink)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WatchInvalidReceipt provides a mock function with given fields: opts, sink
func (_m *ILightInbox) WatchInvalidReceipt(opts *bind.WatchOpts, sink chan<- *lightinbox.LightInboxInvalidReceipt) (event.Subscription, error) {
	ret := _m.Called(opts, sink)

	var r0 event.Subscription
	if rf, ok := ret.Get(0).(func(*bind.WatchOpts, chan<- *lightinbox.LightInboxInvalidReceipt) event.Subscription); ok {
		r0 = rf(opts, sink)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(event.Subscription)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.WatchOpts, chan<- *lightinbox.LightInboxInvalidReceipt) error); ok {
		r1 = rf(opts, sink)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WatchInvalidReceiptReport provides a mock function with given fields: opts, sink
func (_m *ILightInbox) WatchInvalidReceiptReport(opts *bind.WatchOpts, sink chan<- *lightinbox.LightInboxInvalidReceiptReport) (event.Subscription, error) {
	ret := _m.Called(opts, sink)

	var r0 event.Subscription
	if rf, ok := ret.Get(0).(func(*bind.WatchOpts, chan<- *lightinbox.LightInboxInvalidReceiptReport) event.Subscription); ok {
		r0 = rf(opts, sink)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(event.Subscription)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.WatchOpts, chan<- *lightinbox.LightInboxInvalidReceiptReport) error); ok {
		r1 = rf(opts, sink)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WatchInvalidStateReport provides a mock function with given fields: opts, sink
func (_m *ILightInbox) WatchInvalidStateReport(opts *bind.WatchOpts, sink chan<- *lightinbox.LightInboxInvalidStateReport) (event.Subscription, error) {
	ret := _m.Called(opts, sink)

	var r0 event.Subscription
	if rf, ok := ret.Get(0).(func(*bind.WatchOpts, chan<- *lightinbox.LightInboxInvalidStateReport) event.Subscription); ok {
		r0 = rf(opts, sink)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(event.Subscription)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.WatchOpts, chan<- *lightinbox.LightInboxInvalidStateReport) error); ok {
		r1 = rf(opts, sink)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WatchInvalidStateWithAttestation provides a mock function with given fields: opts, sink
func (_m *ILightInbox) WatchInvalidStateWithAttestation(opts *bind.WatchOpts, sink chan<- *lightinbox.LightInboxInvalidStateWithAttestation) (event.Subscription, error) {
	ret := _m.Called(opts, sink)

	var r0 event.Subscription
	if rf, ok := ret.Get(0).(func(*bind.WatchOpts, chan<- *lightinbox.LightInboxInvalidStateWithAttestation) event.Subscription); ok {
		r0 = rf(opts, sink)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(event.Subscription)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.WatchOpts, chan<- *lightinbox.LightInboxInvalidStateWithAttestation) error); ok {
		r1 = rf(opts, sink)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WatchInvalidStateWithSnapshot provides a mock function with given fields: opts, sink
func (_m *ILightInbox) WatchInvalidStateWithSnapshot(opts *bind.WatchOpts, sink chan<- *lightinbox.LightInboxInvalidStateWithSnapshot) (event.Subscription, error) {
	ret := _m.Called(opts, sink)

	var r0 event.Subscription
	if rf, ok := ret.Get(0).(func(*bind.WatchOpts, chan<- *lightinbox.LightInboxInvalidStateWithSnapshot) event.Subscription); ok {
		r0 = rf(opts, sink)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(event.Subscription)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.WatchOpts, chan<- *lightinbox.LightInboxInvalidStateWithSnapshot) error); ok {
		r1 = rf(opts, sink)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WatchOwnershipTransferred provides a mock function with given fields: opts, sink, previousOwner, newOwner
func (_m *ILightInbox) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *lightinbox.LightInboxOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {
	ret := _m.Called(opts, sink, previousOwner, newOwner)

	var r0 event.Subscription
	if rf, ok := ret.Get(0).(func(*bind.WatchOpts, chan<- *lightinbox.LightInboxOwnershipTransferred, []common.Address, []common.Address) event.Subscription); ok {
		r0 = rf(opts, sink, previousOwner, newOwner)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(event.Subscription)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.WatchOpts, chan<- *lightinbox.LightInboxOwnershipTransferred, []common.Address, []common.Address) error); ok {
		r1 = rf(opts, sink, previousOwner, newOwner)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewILightInbox interface {
	mock.TestingT
	Cleanup(func())
}

// NewILightInbox creates a new instance of ILightInbox. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewILightInbox(t mockConstructorTestingTNewILightInbox) *ILightInbox {
	mock := &ILightInbox{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
