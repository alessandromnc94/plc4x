/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package model

import (
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// S7PayloadUserDataItem is the data-structure of this message
type S7PayloadUserDataItem struct {
	ReturnCode    DataTransportErrorCode
	TransportSize DataTransportSize
	Child         IS7PayloadUserDataItemChild
}

// IS7PayloadUserDataItem is the corresponding interface of S7PayloadUserDataItem
type IS7PayloadUserDataItem interface {
	// GetCpuFunctionType returns CpuFunctionType (discriminator field)
	GetCpuFunctionType() uint8
	// GetCpuSubfunction returns CpuSubfunction (discriminator field)
	GetCpuSubfunction() uint8
	// GetDataLength returns DataLength (discriminator field)
	GetDataLength() uint16
	// GetReturnCode returns ReturnCode (property field)
	GetReturnCode() DataTransportErrorCode
	// GetTransportSize returns TransportSize (property field)
	GetTransportSize() DataTransportSize
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

type IS7PayloadUserDataItemParent interface {
	SerializeParent(writeBuffer utils.WriteBuffer, child IS7PayloadUserDataItem, serializeChildFunction func() error) error
	GetTypeName() string
}

type IS7PayloadUserDataItemChild interface {
	Serialize(writeBuffer utils.WriteBuffer) error
	InitializeParent(parent *S7PayloadUserDataItem, returnCode DataTransportErrorCode, transportSize DataTransportSize)
	GetParent() *S7PayloadUserDataItem

	GetTypeName() string
	IS7PayloadUserDataItem
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *S7PayloadUserDataItem) GetReturnCode() DataTransportErrorCode {
	return m.ReturnCode
}

func (m *S7PayloadUserDataItem) GetTransportSize() DataTransportSize {
	return m.TransportSize
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewS7PayloadUserDataItem factory function for S7PayloadUserDataItem
func NewS7PayloadUserDataItem(returnCode DataTransportErrorCode, transportSize DataTransportSize) *S7PayloadUserDataItem {
	return &S7PayloadUserDataItem{ReturnCode: returnCode, TransportSize: transportSize}
}

func CastS7PayloadUserDataItem(structType interface{}) *S7PayloadUserDataItem {
	if casted, ok := structType.(S7PayloadUserDataItem); ok {
		return &casted
	}
	if casted, ok := structType.(*S7PayloadUserDataItem); ok {
		return casted
	}
	if casted, ok := structType.(IS7PayloadUserDataItemChild); ok {
		return casted.GetParent()
	}
	return nil
}

func (m *S7PayloadUserDataItem) GetTypeName() string {
	return "S7PayloadUserDataItem"
}

func (m *S7PayloadUserDataItem) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *S7PayloadUserDataItem) GetLengthInBitsConditional(lastItem bool) uint16 {
	return m.Child.GetLengthInBits()
}

func (m *S7PayloadUserDataItem) GetParentLengthInBits() uint16 {
	lengthInBits := uint16(0)

	// Simple field (returnCode)
	lengthInBits += 8

	// Simple field (transportSize)
	lengthInBits += 8

	// Implicit Field (dataLength)
	lengthInBits += 16

	return lengthInBits
}

func (m *S7PayloadUserDataItem) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func S7PayloadUserDataItemParse(readBuffer utils.ReadBuffer, cpuFunctionType uint8, cpuSubfunction uint8) (*S7PayloadUserDataItem, error) {
	if pullErr := readBuffer.PullContext("S7PayloadUserDataItem"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Simple Field (returnCode)
	if pullErr := readBuffer.PullContext("returnCode"); pullErr != nil {
		return nil, pullErr
	}
	_returnCode, _returnCodeErr := DataTransportErrorCodeParse(readBuffer)
	if _returnCodeErr != nil {
		return nil, errors.Wrap(_returnCodeErr, "Error parsing 'returnCode' field")
	}
	returnCode := _returnCode
	if closeErr := readBuffer.CloseContext("returnCode"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (transportSize)
	if pullErr := readBuffer.PullContext("transportSize"); pullErr != nil {
		return nil, pullErr
	}
	_transportSize, _transportSizeErr := DataTransportSizeParse(readBuffer)
	if _transportSizeErr != nil {
		return nil, errors.Wrap(_transportSizeErr, "Error parsing 'transportSize' field")
	}
	transportSize := _transportSize
	if closeErr := readBuffer.CloseContext("transportSize"); closeErr != nil {
		return nil, closeErr
	}

	// Implicit Field (dataLength) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
	dataLength, _dataLengthErr := readBuffer.ReadUint16("dataLength", 16)
	_ = dataLength
	if _dataLengthErr != nil {
		return nil, errors.Wrap(_dataLengthErr, "Error parsing 'dataLength' field")
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type S7PayloadUserDataItemChild interface {
		InitializeParent(*S7PayloadUserDataItem, DataTransportErrorCode, DataTransportSize)
		GetParent() *S7PayloadUserDataItem
	}
	var _child S7PayloadUserDataItemChild
	var typeSwitchError error
	switch {
	case cpuFunctionType == 0x00 && cpuSubfunction == 0x03: // S7PayloadDiagnosticMessage
		_child, typeSwitchError = S7PayloadDiagnosticMessageParse(readBuffer, cpuFunctionType, cpuSubfunction)
	case cpuFunctionType == 0x00 && cpuSubfunction == 0x05: // S7PayloadAlarm8
		_child, typeSwitchError = S7PayloadAlarm8Parse(readBuffer, cpuFunctionType, cpuSubfunction)
	case cpuFunctionType == 0x00 && cpuSubfunction == 0x06: // S7PayloadNotify
		_child, typeSwitchError = S7PayloadNotifyParse(readBuffer, cpuFunctionType, cpuSubfunction)
	case cpuFunctionType == 0x00 && cpuSubfunction == 0x0c: // S7PayloadAlarmAckInd
		_child, typeSwitchError = S7PayloadAlarmAckIndParse(readBuffer, cpuFunctionType, cpuSubfunction)
	case cpuFunctionType == 0x00 && cpuSubfunction == 0x11: // S7PayloadAlarmSQ
		_child, typeSwitchError = S7PayloadAlarmSQParse(readBuffer, cpuFunctionType, cpuSubfunction)
	case cpuFunctionType == 0x00 && cpuSubfunction == 0x12: // S7PayloadAlarmS
		_child, typeSwitchError = S7PayloadAlarmSParse(readBuffer, cpuFunctionType, cpuSubfunction)
	case cpuFunctionType == 0x00 && cpuSubfunction == 0x13: // S7PayloadAlarmSC
		_child, typeSwitchError = S7PayloadAlarmSCParse(readBuffer, cpuFunctionType, cpuSubfunction)
	case cpuFunctionType == 0x00 && cpuSubfunction == 0x16: // S7PayloadNotify8
		_child, typeSwitchError = S7PayloadNotify8Parse(readBuffer, cpuFunctionType, cpuSubfunction)
	case cpuFunctionType == 0x04 && cpuSubfunction == 0x01: // S7PayloadUserDataItemCpuFunctionReadSzlRequest
		_child, typeSwitchError = S7PayloadUserDataItemCpuFunctionReadSzlRequestParse(readBuffer, cpuFunctionType, cpuSubfunction)
	case cpuFunctionType == 0x08 && cpuSubfunction == 0x01: // S7PayloadUserDataItemCpuFunctionReadSzlResponse
		_child, typeSwitchError = S7PayloadUserDataItemCpuFunctionReadSzlResponseParse(readBuffer, cpuFunctionType, cpuSubfunction)
	case cpuFunctionType == 0x04 && cpuSubfunction == 0x02: // S7PayloadUserDataItemCpuFunctionMsgSubscription
		_child, typeSwitchError = S7PayloadUserDataItemCpuFunctionMsgSubscriptionParse(readBuffer, cpuFunctionType, cpuSubfunction)
	case cpuFunctionType == 0x08 && cpuSubfunction == 0x02 && dataLength == 0x00: // S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse
		_child, typeSwitchError = S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponseParse(readBuffer, cpuFunctionType, cpuSubfunction)
	case cpuFunctionType == 0x08 && cpuSubfunction == 0x02 && dataLength == 0x02: // S7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponse
		_child, typeSwitchError = S7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponseParse(readBuffer, cpuFunctionType, cpuSubfunction)
	case cpuFunctionType == 0x08 && cpuSubfunction == 0x02 && dataLength == 0x05: // S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse
		_child, typeSwitchError = S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponseParse(readBuffer, cpuFunctionType, cpuSubfunction)
	case cpuFunctionType == 0x04 && cpuSubfunction == 0x0b: // S7PayloadUserDataItemCpuFunctionAlarmAck
		_child, typeSwitchError = S7PayloadUserDataItemCpuFunctionAlarmAckParse(readBuffer, cpuFunctionType, cpuSubfunction)
	case cpuFunctionType == 0x08 && cpuSubfunction == 0x0b: // S7PayloadUserDataItemCpuFunctionAlarmAckResponse
		_child, typeSwitchError = S7PayloadUserDataItemCpuFunctionAlarmAckResponseParse(readBuffer, cpuFunctionType, cpuSubfunction)
	case cpuFunctionType == 0x04 && cpuSubfunction == 0x13: // S7PayloadUserDataItemCpuFunctionAlarmQuery
		_child, typeSwitchError = S7PayloadUserDataItemCpuFunctionAlarmQueryParse(readBuffer, cpuFunctionType, cpuSubfunction)
	case cpuFunctionType == 0x08 && cpuSubfunction == 0x13: // S7PayloadUserDataItemCpuFunctionAlarmQueryResponse
		_child, typeSwitchError = S7PayloadUserDataItemCpuFunctionAlarmQueryResponseParse(readBuffer, cpuFunctionType, cpuSubfunction)
	default:
		// TODO: return actual type
		typeSwitchError = errors.New("Unmapped type")
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch.")
	}

	if closeErr := readBuffer.CloseContext("S7PayloadUserDataItem"); closeErr != nil {
		return nil, closeErr
	}

	// Finish initializing
	_child.InitializeParent(_child.GetParent(), returnCode, transportSize)
	return _child.GetParent(), nil
}

func (m *S7PayloadUserDataItem) Serialize(writeBuffer utils.WriteBuffer) error {
	return m.Child.Serialize(writeBuffer)
}

func (m *S7PayloadUserDataItem) SerializeParent(writeBuffer utils.WriteBuffer, child IS7PayloadUserDataItem, serializeChildFunction func() error) error {
	if pushErr := writeBuffer.PushContext("S7PayloadUserDataItem"); pushErr != nil {
		return pushErr
	}

	// Simple Field (returnCode)
	if pushErr := writeBuffer.PushContext("returnCode"); pushErr != nil {
		return pushErr
	}
	_returnCodeErr := m.ReturnCode.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("returnCode"); popErr != nil {
		return popErr
	}
	if _returnCodeErr != nil {
		return errors.Wrap(_returnCodeErr, "Error serializing 'returnCode' field")
	}

	// Simple Field (transportSize)
	if pushErr := writeBuffer.PushContext("transportSize"); pushErr != nil {
		return pushErr
	}
	_transportSizeErr := m.TransportSize.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("transportSize"); popErr != nil {
		return popErr
	}
	if _transportSizeErr != nil {
		return errors.Wrap(_transportSizeErr, "Error serializing 'transportSize' field")
	}

	// Implicit Field (dataLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	dataLength := uint16(uint16(uint16(m.GetLengthInBytes())) - uint16(uint16(4)))
	_dataLengthErr := writeBuffer.WriteUint16("dataLength", 16, (dataLength))
	if _dataLengthErr != nil {
		return errors.Wrap(_dataLengthErr, "Error serializing 'dataLength' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("S7PayloadUserDataItem"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *S7PayloadUserDataItem) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
