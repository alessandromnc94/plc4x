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
	"io"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetBinaryPV is the data-structure of this message
type BACnetBinaryPV struct {
	RawData *BACnetContextTagEnumerated

	// Arguments.
	TagNumber uint8
}

// IBACnetBinaryPV is the corresponding interface of BACnetBinaryPV
type IBACnetBinaryPV interface {
	// GetRawData returns RawData (property field)
	GetRawData() *BACnetContextTagEnumerated
	// GetIsInactive returns IsInactive (virtual field)
	GetIsInactive() bool
	// GetIsActive returns IsActive (virtual field)
	GetIsActive() bool
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

func (m *BACnetBinaryPV) GetRawData() *BACnetContextTagEnumerated {
	return m.RawData
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *BACnetBinaryPV) GetIsInactive() bool {
	rawData := m.RawData
	_ = rawData
	return bool(bool(bool((m.GetRawData()) != (nil))) && bool(bool(((*m.GetRawData()).GetPayload().GetActualValue()) == (0))))
}

func (m *BACnetBinaryPV) GetIsActive() bool {
	rawData := m.RawData
	_ = rawData
	return bool(bool(bool((m.GetRawData()) != (nil))) && bool(bool(((*m.GetRawData()).GetPayload().GetActualValue()) == (1))))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetBinaryPV factory function for BACnetBinaryPV
func NewBACnetBinaryPV(rawData *BACnetContextTagEnumerated, tagNumber uint8) *BACnetBinaryPV {
	return &BACnetBinaryPV{RawData: rawData, TagNumber: tagNumber}
}

func CastBACnetBinaryPV(structType interface{}) *BACnetBinaryPV {
	if casted, ok := structType.(BACnetBinaryPV); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetBinaryPV); ok {
		return casted
	}
	return nil
}

func (m *BACnetBinaryPV) GetTypeName() string {
	return "BACnetBinaryPV"
}

func (m *BACnetBinaryPV) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetBinaryPV) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Optional Field (rawData)
	if m.RawData != nil {
		lengthInBits += (*m.RawData).GetLengthInBits()
	}

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *BACnetBinaryPV) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetBinaryPVParse(readBuffer utils.ReadBuffer, tagNumber uint8) (*BACnetBinaryPV, error) {
	if pullErr := readBuffer.PullContext("BACnetBinaryPV"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Optional Field (rawData) (Can be skipped, if a given expression evaluates to false)
	var rawData *BACnetContextTagEnumerated = nil
	{
		currentPos = readBuffer.GetPos()
		if pullErr := readBuffer.PullContext("rawData"); pullErr != nil {
			return nil, pullErr
		}
		_val, _err := BACnetContextTagParse(readBuffer, tagNumber, BACnetDataType_ENUMERATED)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'rawData' field")
		default:
			rawData = CastBACnetContextTagEnumerated(_val)
			if closeErr := readBuffer.CloseContext("rawData"); closeErr != nil {
				return nil, closeErr
			}
		}
	}

	// Virtual field
	_isInactive := bool(bool((rawData) != (nil))) && bool(bool(((*rawData).GetPayload().GetActualValue()) == (0)))
	isInactive := bool(_isInactive)
	_ = isInactive

	// Virtual field
	_isActive := bool(bool((rawData) != (nil))) && bool(bool(((*rawData).GetPayload().GetActualValue()) == (1)))
	isActive := bool(_isActive)
	_ = isActive

	if closeErr := readBuffer.CloseContext("BACnetBinaryPV"); closeErr != nil {
		return nil, closeErr
	}

	// Create the instance
	return NewBACnetBinaryPV(rawData, tagNumber), nil
}

func (m *BACnetBinaryPV) Serialize(writeBuffer utils.WriteBuffer) error {
	if pushErr := writeBuffer.PushContext("BACnetBinaryPV"); pushErr != nil {
		return pushErr
	}

	// Optional Field (rawData) (Can be skipped, if the value is null)
	var rawData *BACnetContextTagEnumerated = nil
	if m.RawData != nil {
		if pushErr := writeBuffer.PushContext("rawData"); pushErr != nil {
			return pushErr
		}
		rawData = m.RawData
		_rawDataErr := rawData.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("rawData"); popErr != nil {
			return popErr
		}
		if _rawDataErr != nil {
			return errors.Wrap(_rawDataErr, "Error serializing 'rawData' field")
		}
	}
	// Virtual field
	if _isInactiveErr := writeBuffer.WriteVirtual("isInactive", m.GetIsInactive()); _isInactiveErr != nil {
		return errors.Wrap(_isInactiveErr, "Error serializing 'isInactive' field")
	}
	// Virtual field
	if _isActiveErr := writeBuffer.WriteVirtual("isActive", m.GetIsActive()); _isActiveErr != nil {
		return errors.Wrap(_isActiveErr, "Error serializing 'isActive' field")
	}

	if popErr := writeBuffer.PopContext("BACnetBinaryPV"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *BACnetBinaryPV) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
