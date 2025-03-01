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

// BACnetErrorRemovedReadPropertyConditional is the data-structure of this message
type BACnetErrorRemovedReadPropertyConditional struct {
	*BACnetError
}

// IBACnetErrorRemovedReadPropertyConditional is the corresponding interface of BACnetErrorRemovedReadPropertyConditional
type IBACnetErrorRemovedReadPropertyConditional interface {
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

func (m *BACnetErrorRemovedReadPropertyConditional) GetServiceChoice() uint8 {
	return 0x0D
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetErrorRemovedReadPropertyConditional) InitializeParent(parent *BACnetError, errorClass *BACnetApplicationTagEnumerated, errorCode *BACnetApplicationTagEnumerated) {
	m.BACnetError.ErrorClass = errorClass
	m.BACnetError.ErrorCode = errorCode
}

func (m *BACnetErrorRemovedReadPropertyConditional) GetParent() *BACnetError {
	return m.BACnetError
}

// NewBACnetErrorRemovedReadPropertyConditional factory function for BACnetErrorRemovedReadPropertyConditional
func NewBACnetErrorRemovedReadPropertyConditional(errorClass *BACnetApplicationTagEnumerated, errorCode *BACnetApplicationTagEnumerated) *BACnetErrorRemovedReadPropertyConditional {
	_result := &BACnetErrorRemovedReadPropertyConditional{
		BACnetError: NewBACnetError(errorClass, errorCode),
	}
	_result.Child = _result
	return _result
}

func CastBACnetErrorRemovedReadPropertyConditional(structType interface{}) *BACnetErrorRemovedReadPropertyConditional {
	if casted, ok := structType.(BACnetErrorRemovedReadPropertyConditional); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetErrorRemovedReadPropertyConditional); ok {
		return casted
	}
	if casted, ok := structType.(BACnetError); ok {
		return CastBACnetErrorRemovedReadPropertyConditional(casted.Child)
	}
	if casted, ok := structType.(*BACnetError); ok {
		return CastBACnetErrorRemovedReadPropertyConditional(casted.Child)
	}
	return nil
}

func (m *BACnetErrorRemovedReadPropertyConditional) GetTypeName() string {
	return "BACnetErrorRemovedReadPropertyConditional"
}

func (m *BACnetErrorRemovedReadPropertyConditional) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetErrorRemovedReadPropertyConditional) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *BACnetErrorRemovedReadPropertyConditional) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetErrorRemovedReadPropertyConditionalParse(readBuffer utils.ReadBuffer) (*BACnetErrorRemovedReadPropertyConditional, error) {
	if pullErr := readBuffer.PullContext("BACnetErrorRemovedReadPropertyConditional"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("BACnetErrorRemovedReadPropertyConditional"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetErrorRemovedReadPropertyConditional{
		BACnetError: &BACnetError{},
	}
	_child.BACnetError.Child = _child
	return _child, nil
}

func (m *BACnetErrorRemovedReadPropertyConditional) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetErrorRemovedReadPropertyConditional"); pushErr != nil {
			return pushErr
		}

		if popErr := writeBuffer.PopContext("BACnetErrorRemovedReadPropertyConditional"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetErrorRemovedReadPropertyConditional) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
