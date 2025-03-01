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

// BACnetSegmentation is the data-structure of this message
type BACnetSegmentation struct {
	RawData *BACnetApplicationTagEnumerated
}

// IBACnetSegmentation is the corresponding interface of BACnetSegmentation
type IBACnetSegmentation interface {
	// GetRawData returns RawData (property field)
	GetRawData() *BACnetApplicationTagEnumerated
	// GetIsSegmentedBoth returns IsSegmentedBoth (virtual field)
	GetIsSegmentedBoth() bool
	// GetIsSegmentedTransmit returns IsSegmentedTransmit (virtual field)
	GetIsSegmentedTransmit() bool
	// GetIsSegmentedReceive returns IsSegmentedReceive (virtual field)
	GetIsSegmentedReceive() bool
	// GetIsNoSegmentation returns IsNoSegmentation (virtual field)
	GetIsNoSegmentation() bool
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

func (m *BACnetSegmentation) GetRawData() *BACnetApplicationTagEnumerated {
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

func (m *BACnetSegmentation) GetIsSegmentedBoth() bool {
	return bool(bool(bool((m.GetRawData()) != (nil))) && bool(bool((m.GetRawData().GetPayload().GetActualValue()) == (0))))
}

func (m *BACnetSegmentation) GetIsSegmentedTransmit() bool {
	return bool(bool(bool((m.GetRawData()) != (nil))) && bool(bool((m.GetRawData().GetPayload().GetActualValue()) == (1))))
}

func (m *BACnetSegmentation) GetIsSegmentedReceive() bool {
	return bool(bool(bool((m.GetRawData()) != (nil))) && bool(bool((m.GetRawData().GetPayload().GetActualValue()) == (3))))
}

func (m *BACnetSegmentation) GetIsNoSegmentation() bool {
	return bool(bool(bool((m.GetRawData()) != (nil))) && bool(bool((m.GetRawData().GetPayload().GetActualValue()) == (4))))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetSegmentation factory function for BACnetSegmentation
func NewBACnetSegmentation(rawData *BACnetApplicationTagEnumerated) *BACnetSegmentation {
	return &BACnetSegmentation{RawData: rawData}
}

func CastBACnetSegmentation(structType interface{}) *BACnetSegmentation {
	if casted, ok := structType.(BACnetSegmentation); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetSegmentation); ok {
		return casted
	}
	return nil
}

func (m *BACnetSegmentation) GetTypeName() string {
	return "BACnetSegmentation"
}

func (m *BACnetSegmentation) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetSegmentation) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (rawData)
	lengthInBits += m.RawData.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *BACnetSegmentation) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetSegmentationParse(readBuffer utils.ReadBuffer) (*BACnetSegmentation, error) {
	if pullErr := readBuffer.PullContext("BACnetSegmentation"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Simple Field (rawData)
	if pullErr := readBuffer.PullContext("rawData"); pullErr != nil {
		return nil, pullErr
	}
	_rawData, _rawDataErr := BACnetApplicationTagParse(readBuffer)
	if _rawDataErr != nil {
		return nil, errors.Wrap(_rawDataErr, "Error parsing 'rawData' field")
	}
	rawData := CastBACnetApplicationTagEnumerated(_rawData)
	if closeErr := readBuffer.CloseContext("rawData"); closeErr != nil {
		return nil, closeErr
	}

	// Virtual field
	_isSegmentedBoth := bool(bool((rawData) != (nil))) && bool(bool((rawData.GetPayload().GetActualValue()) == (0)))
	isSegmentedBoth := bool(_isSegmentedBoth)
	_ = isSegmentedBoth

	// Virtual field
	_isSegmentedTransmit := bool(bool((rawData) != (nil))) && bool(bool((rawData.GetPayload().GetActualValue()) == (1)))
	isSegmentedTransmit := bool(_isSegmentedTransmit)
	_ = isSegmentedTransmit

	// Virtual field
	_isSegmentedReceive := bool(bool((rawData) != (nil))) && bool(bool((rawData.GetPayload().GetActualValue()) == (3)))
	isSegmentedReceive := bool(_isSegmentedReceive)
	_ = isSegmentedReceive

	// Virtual field
	_isNoSegmentation := bool(bool((rawData) != (nil))) && bool(bool((rawData.GetPayload().GetActualValue()) == (4)))
	isNoSegmentation := bool(_isNoSegmentation)
	_ = isNoSegmentation

	if closeErr := readBuffer.CloseContext("BACnetSegmentation"); closeErr != nil {
		return nil, closeErr
	}

	// Create the instance
	return NewBACnetSegmentation(rawData), nil
}

func (m *BACnetSegmentation) Serialize(writeBuffer utils.WriteBuffer) error {
	if pushErr := writeBuffer.PushContext("BACnetSegmentation"); pushErr != nil {
		return pushErr
	}

	// Simple Field (rawData)
	if pushErr := writeBuffer.PushContext("rawData"); pushErr != nil {
		return pushErr
	}
	_rawDataErr := m.RawData.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("rawData"); popErr != nil {
		return popErr
	}
	if _rawDataErr != nil {
		return errors.Wrap(_rawDataErr, "Error serializing 'rawData' field")
	}
	// Virtual field
	if _isSegmentedBothErr := writeBuffer.WriteVirtual("isSegmentedBoth", m.GetIsSegmentedBoth()); _isSegmentedBothErr != nil {
		return errors.Wrap(_isSegmentedBothErr, "Error serializing 'isSegmentedBoth' field")
	}
	// Virtual field
	if _isSegmentedTransmitErr := writeBuffer.WriteVirtual("isSegmentedTransmit", m.GetIsSegmentedTransmit()); _isSegmentedTransmitErr != nil {
		return errors.Wrap(_isSegmentedTransmitErr, "Error serializing 'isSegmentedTransmit' field")
	}
	// Virtual field
	if _isSegmentedReceiveErr := writeBuffer.WriteVirtual("isSegmentedReceive", m.GetIsSegmentedReceive()); _isSegmentedReceiveErr != nil {
		return errors.Wrap(_isSegmentedReceiveErr, "Error serializing 'isSegmentedReceive' field")
	}
	// Virtual field
	if _isNoSegmentationErr := writeBuffer.WriteVirtual("isNoSegmentation", m.GetIsNoSegmentation()); _isNoSegmentationErr != nil {
		return errors.Wrap(_isNoSegmentationErr, "Error serializing 'isNoSegmentation' field")
	}

	if popErr := writeBuffer.PopContext("BACnetSegmentation"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *BACnetSegmentation) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
