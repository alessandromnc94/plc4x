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

// BACnetTagPayloadOctetString is the data-structure of this message
type BACnetTagPayloadOctetString struct {
	Value string

	// Arguments.
	ActualLength uint32
}

// IBACnetTagPayloadOctetString is the corresponding interface of BACnetTagPayloadOctetString
type IBACnetTagPayloadOctetString interface {
	// GetValue returns Value (property field)
	GetValue() string
	// GetActualLengthInBit returns ActualLengthInBit (virtual field)
	GetActualLengthInBit() uint16
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

func (m *BACnetTagPayloadOctetString) GetValue() string {
	return m.Value
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *BACnetTagPayloadOctetString) GetActualLengthInBit() uint16 {
	return uint16(uint16(m.ActualLength) * uint16(uint16(8)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetTagPayloadOctetString factory function for BACnetTagPayloadOctetString
func NewBACnetTagPayloadOctetString(value string, actualLength uint32) *BACnetTagPayloadOctetString {
	return &BACnetTagPayloadOctetString{Value: value, ActualLength: actualLength}
}

func CastBACnetTagPayloadOctetString(structType interface{}) *BACnetTagPayloadOctetString {
	if casted, ok := structType.(BACnetTagPayloadOctetString); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetTagPayloadOctetString); ok {
		return casted
	}
	return nil
}

func (m *BACnetTagPayloadOctetString) GetTypeName() string {
	return "BACnetTagPayloadOctetString"
}

func (m *BACnetTagPayloadOctetString) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetTagPayloadOctetString) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// A virtual field doesn't have any in- or output.

	// Simple field (value)
	lengthInBits += uint16(m.GetActualLengthInBit())

	return lengthInBits
}

func (m *BACnetTagPayloadOctetString) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetTagPayloadOctetStringParse(readBuffer utils.ReadBuffer, actualLength uint32) (*BACnetTagPayloadOctetString, error) {
	if pullErr := readBuffer.PullContext("BACnetTagPayloadOctetString"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Virtual field
	_actualLengthInBit := uint16(actualLength) * uint16(uint16(8))
	actualLengthInBit := uint16(_actualLengthInBit)
	_ = actualLengthInBit

	// Simple Field (value)
	_value, _valueErr := readBuffer.ReadString("value", uint32(actualLengthInBit))
	if _valueErr != nil {
		return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
	}
	value := _value

	if closeErr := readBuffer.CloseContext("BACnetTagPayloadOctetString"); closeErr != nil {
		return nil, closeErr
	}

	// Create the instance
	return NewBACnetTagPayloadOctetString(value, actualLength), nil
}

func (m *BACnetTagPayloadOctetString) Serialize(writeBuffer utils.WriteBuffer) error {
	if pushErr := writeBuffer.PushContext("BACnetTagPayloadOctetString"); pushErr != nil {
		return pushErr
	}
	// Virtual field
	if _actualLengthInBitErr := writeBuffer.WriteVirtual("actualLengthInBit", m.GetActualLengthInBit()); _actualLengthInBitErr != nil {
		return errors.Wrap(_actualLengthInBitErr, "Error serializing 'actualLengthInBit' field")
	}

	// Simple Field (value)
	value := string(m.Value)
	_valueErr := writeBuffer.WriteString("value", uint32(m.GetActualLengthInBit()), "ASCII", (value))
	if _valueErr != nil {
		return errors.Wrap(_valueErr, "Error serializing 'value' field")
	}

	if popErr := writeBuffer.PopContext("BACnetTagPayloadOctetString"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *BACnetTagPayloadOctetString) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
