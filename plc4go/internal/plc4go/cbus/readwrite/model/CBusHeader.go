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
	"github.com/rs/zerolog/log"
)

// Code generated by code-generation. DO NOT EDIT.

// CBusHeader is the data-structure of this message
type CBusHeader struct {
	PriorityClass          PriorityClass
	DestinationAddressType DestinationAddressType
}

// ICBusHeader is the corresponding interface of CBusHeader
type ICBusHeader interface {
	// GetPriorityClass returns PriorityClass (property field)
	GetPriorityClass() PriorityClass
	// GetDestinationAddressType returns DestinationAddressType (property field)
	GetDestinationAddressType() DestinationAddressType
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *CBusHeader) GetPriorityClass() PriorityClass {
	return m.PriorityClass
}

func (m *CBusHeader) GetDestinationAddressType() DestinationAddressType {
	return m.DestinationAddressType
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCBusHeader factory function for CBusHeader
func NewCBusHeader(priorityClass PriorityClass, destinationAddressType DestinationAddressType) *CBusHeader {
	return &CBusHeader{PriorityClass: priorityClass, DestinationAddressType: destinationAddressType}
}

func CastCBusHeader(structType interface{}) *CBusHeader {
	if casted, ok := structType.(CBusHeader); ok {
		return &casted
	}
	if casted, ok := structType.(*CBusHeader); ok {
		return casted
	}
	return nil
}

func (m *CBusHeader) GetTypeName() string {
	return "CBusHeader"
}

func (m *CBusHeader) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *CBusHeader) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (priorityClass)
	lengthInBits += 2

	// Reserved Field (reserved)
	lengthInBits += 1

	// Reserved Field (reserved)
	lengthInBits += 2

	// Simple field (destinationAddressType)
	lengthInBits += 3

	return lengthInBits
}

func (m *CBusHeader) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func CBusHeaderParse(readBuffer utils.ReadBuffer) (*CBusHeader, error) {
	if pullErr := readBuffer.PullContext("CBusHeader"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Simple Field (priorityClass)
	if pullErr := readBuffer.PullContext("priorityClass"); pullErr != nil {
		return nil, pullErr
	}
	_priorityClass, _priorityClassErr := PriorityClassParse(readBuffer)
	if _priorityClassErr != nil {
		return nil, errors.Wrap(_priorityClassErr, "Error parsing 'priorityClass' field")
	}
	priorityClass := _priorityClass
	if closeErr := readBuffer.CloseContext("priorityClass"); closeErr != nil {
		return nil, closeErr
	}

	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadBit("reserved")
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field")
		}
		if reserved != bool(false) {
			log.Info().Fields(map[string]interface{}{
				"expected value": bool(false),
				"got value":      reserved,
			}).Msg("Got unexpected response.")
		}
	}

	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint8("reserved", 2)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field")
		}
		if reserved != uint8(0) {
			log.Info().Fields(map[string]interface{}{
				"expected value": uint8(0),
				"got value":      reserved,
			}).Msg("Got unexpected response.")
		}
	}

	// Simple Field (destinationAddressType)
	if pullErr := readBuffer.PullContext("destinationAddressType"); pullErr != nil {
		return nil, pullErr
	}
	_destinationAddressType, _destinationAddressTypeErr := DestinationAddressTypeParse(readBuffer)
	if _destinationAddressTypeErr != nil {
		return nil, errors.Wrap(_destinationAddressTypeErr, "Error parsing 'destinationAddressType' field")
	}
	destinationAddressType := _destinationAddressType
	if closeErr := readBuffer.CloseContext("destinationAddressType"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("CBusHeader"); closeErr != nil {
		return nil, closeErr
	}

	// Create the instance
	return NewCBusHeader(priorityClass, destinationAddressType), nil
}

func (m *CBusHeader) Serialize(writeBuffer utils.WriteBuffer) error {
	if pushErr := writeBuffer.PushContext("CBusHeader"); pushErr != nil {
		return pushErr
	}

	// Simple Field (priorityClass)
	if pushErr := writeBuffer.PushContext("priorityClass"); pushErr != nil {
		return pushErr
	}
	_priorityClassErr := m.PriorityClass.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("priorityClass"); popErr != nil {
		return popErr
	}
	if _priorityClassErr != nil {
		return errors.Wrap(_priorityClassErr, "Error serializing 'priorityClass' field")
	}

	// Reserved Field (reserved)
	{
		_err := writeBuffer.WriteBit("reserved", bool(false))
		if _err != nil {
			return errors.Wrap(_err, "Error serializing 'reserved' field")
		}
	}

	// Reserved Field (reserved)
	{
		_err := writeBuffer.WriteUint8("reserved", 2, uint8(0))
		if _err != nil {
			return errors.Wrap(_err, "Error serializing 'reserved' field")
		}
	}

	// Simple Field (destinationAddressType)
	if pushErr := writeBuffer.PushContext("destinationAddressType"); pushErr != nil {
		return pushErr
	}
	_destinationAddressTypeErr := m.DestinationAddressType.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("destinationAddressType"); popErr != nil {
		return popErr
	}
	if _destinationAddressTypeErr != nil {
		return errors.Wrap(_destinationAddressTypeErr, "Error serializing 'destinationAddressType' field")
	}

	if popErr := writeBuffer.PopContext("CBusHeader"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *CBusHeader) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
