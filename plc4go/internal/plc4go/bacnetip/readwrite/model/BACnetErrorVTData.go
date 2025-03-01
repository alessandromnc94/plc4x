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
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetErrorVTData is the data-structure of this message
type BACnetErrorVTData struct {
	*BACnetError
}

// IBACnetErrorVTData is the corresponding interface of BACnetErrorVTData
type IBACnetErrorVTData interface {
	IBACnetError
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

func (m *BACnetErrorVTData) GetServiceChoice() uint8 {
	return 0x17
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetErrorVTData) InitializeParent(parent *BACnetError, errorClass *BACnetApplicationTagEnumerated, errorCode *BACnetApplicationTagEnumerated) {
	m.BACnetError.ErrorClass = errorClass
	m.BACnetError.ErrorCode = errorCode
}

func (m *BACnetErrorVTData) GetParent() *BACnetError {
	return m.BACnetError
}

// NewBACnetErrorVTData factory function for BACnetErrorVTData
func NewBACnetErrorVTData(errorClass *BACnetApplicationTagEnumerated, errorCode *BACnetApplicationTagEnumerated) *BACnetErrorVTData {
	_result := &BACnetErrorVTData{
		BACnetError: NewBACnetError(errorClass, errorCode),
	}
	_result.Child = _result
	return _result
}

func CastBACnetErrorVTData(structType interface{}) *BACnetErrorVTData {
	if casted, ok := structType.(BACnetErrorVTData); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetErrorVTData); ok {
		return casted
	}
	if casted, ok := structType.(BACnetError); ok {
		return CastBACnetErrorVTData(casted.Child)
	}
	if casted, ok := structType.(*BACnetError); ok {
		return CastBACnetErrorVTData(casted.Child)
	}
	return nil
}

func (m *BACnetErrorVTData) GetTypeName() string {
	return "BACnetErrorVTData"
}

func (m *BACnetErrorVTData) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetErrorVTData) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *BACnetErrorVTData) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetErrorVTDataParse(readBuffer utils.ReadBuffer) (*BACnetErrorVTData, error) {
	if pullErr := readBuffer.PullContext("BACnetErrorVTData"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("BACnetErrorVTData"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetErrorVTData{
		BACnetError: &BACnetError{},
	}
	_child.BACnetError.Child = _child
	return _child, nil
}

func (m *BACnetErrorVTData) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetErrorVTData"); pushErr != nil {
			return pushErr
		}

		if popErr := writeBuffer.PopContext("BACnetErrorVTData"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetErrorVTData) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
