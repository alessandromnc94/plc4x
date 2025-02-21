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

// ModbusPDUReadHoldingRegistersRequest is the data-structure of this message
type ModbusPDUReadHoldingRegistersRequest struct {
	*ModbusPDU
	StartingAddress uint16
	Quantity        uint16
}

// IModbusPDUReadHoldingRegistersRequest is the corresponding interface of ModbusPDUReadHoldingRegistersRequest
type IModbusPDUReadHoldingRegistersRequest interface {
	IModbusPDU
	// GetStartingAddress returns StartingAddress (property field)
	GetStartingAddress() uint16
	// GetQuantity returns Quantity (property field)
	GetQuantity() uint16
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *ModbusPDUReadHoldingRegistersRequest) GetErrorFlag() bool {
	return bool(false)
}

func (m *ModbusPDUReadHoldingRegistersRequest) GetFunctionFlag() uint8 {
	return 0x03
}

func (m *ModbusPDUReadHoldingRegistersRequest) GetResponse() bool {
	return bool(false)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *ModbusPDUReadHoldingRegistersRequest) InitializeParent(parent *ModbusPDU) {}

func (m *ModbusPDUReadHoldingRegistersRequest) GetParent() *ModbusPDU {
	return m.ModbusPDU
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *ModbusPDUReadHoldingRegistersRequest) GetStartingAddress() uint16 {
	return m.StartingAddress
}

func (m *ModbusPDUReadHoldingRegistersRequest) GetQuantity() uint16 {
	return m.Quantity
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewModbusPDUReadHoldingRegistersRequest factory function for ModbusPDUReadHoldingRegistersRequest
func NewModbusPDUReadHoldingRegistersRequest(startingAddress uint16, quantity uint16) *ModbusPDUReadHoldingRegistersRequest {
	_result := &ModbusPDUReadHoldingRegistersRequest{
		StartingAddress: startingAddress,
		Quantity:        quantity,
		ModbusPDU:       NewModbusPDU(),
	}
	_result.Child = _result
	return _result
}

func CastModbusPDUReadHoldingRegistersRequest(structType interface{}) *ModbusPDUReadHoldingRegistersRequest {
	if casted, ok := structType.(ModbusPDUReadHoldingRegistersRequest); ok {
		return &casted
	}
	if casted, ok := structType.(*ModbusPDUReadHoldingRegistersRequest); ok {
		return casted
	}
	if casted, ok := structType.(ModbusPDU); ok {
		return CastModbusPDUReadHoldingRegistersRequest(casted.Child)
	}
	if casted, ok := structType.(*ModbusPDU); ok {
		return CastModbusPDUReadHoldingRegistersRequest(casted.Child)
	}
	return nil
}

func (m *ModbusPDUReadHoldingRegistersRequest) GetTypeName() string {
	return "ModbusPDUReadHoldingRegistersRequest"
}

func (m *ModbusPDUReadHoldingRegistersRequest) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *ModbusPDUReadHoldingRegistersRequest) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (startingAddress)
	lengthInBits += 16

	// Simple field (quantity)
	lengthInBits += 16

	return lengthInBits
}

func (m *ModbusPDUReadHoldingRegistersRequest) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ModbusPDUReadHoldingRegistersRequestParse(readBuffer utils.ReadBuffer, response bool) (*ModbusPDUReadHoldingRegistersRequest, error) {
	if pullErr := readBuffer.PullContext("ModbusPDUReadHoldingRegistersRequest"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Simple Field (startingAddress)
	_startingAddress, _startingAddressErr := readBuffer.ReadUint16("startingAddress", 16)
	if _startingAddressErr != nil {
		return nil, errors.Wrap(_startingAddressErr, "Error parsing 'startingAddress' field")
	}
	startingAddress := _startingAddress

	// Simple Field (quantity)
	_quantity, _quantityErr := readBuffer.ReadUint16("quantity", 16)
	if _quantityErr != nil {
		return nil, errors.Wrap(_quantityErr, "Error parsing 'quantity' field")
	}
	quantity := _quantity

	if closeErr := readBuffer.CloseContext("ModbusPDUReadHoldingRegistersRequest"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &ModbusPDUReadHoldingRegistersRequest{
		StartingAddress: startingAddress,
		Quantity:        quantity,
		ModbusPDU:       &ModbusPDU{},
	}
	_child.ModbusPDU.Child = _child
	return _child, nil
}

func (m *ModbusPDUReadHoldingRegistersRequest) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ModbusPDUReadHoldingRegistersRequest"); pushErr != nil {
			return pushErr
		}

		// Simple Field (startingAddress)
		startingAddress := uint16(m.StartingAddress)
		_startingAddressErr := writeBuffer.WriteUint16("startingAddress", 16, (startingAddress))
		if _startingAddressErr != nil {
			return errors.Wrap(_startingAddressErr, "Error serializing 'startingAddress' field")
		}

		// Simple Field (quantity)
		quantity := uint16(m.Quantity)
		_quantityErr := writeBuffer.WriteUint16("quantity", 16, (quantity))
		if _quantityErr != nil {
			return errors.Wrap(_quantityErr, "Error serializing 'quantity' field")
		}

		if popErr := writeBuffer.PopContext("ModbusPDUReadHoldingRegistersRequest"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *ModbusPDUReadHoldingRegistersRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
