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

// BACnetErrorConfirmedEventNotification is the data-structure of this message
type BACnetErrorConfirmedEventNotification struct {
	*BACnetError
}

// IBACnetErrorConfirmedEventNotification is the corresponding interface of BACnetErrorConfirmedEventNotification
type IBACnetErrorConfirmedEventNotification interface {
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

func (m *BACnetErrorConfirmedEventNotification) GetServiceChoice() uint8 {
	return 0x02
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetErrorConfirmedEventNotification) InitializeParent(parent *BACnetError, errorClass *BACnetApplicationTagEnumerated, errorCode *BACnetApplicationTagEnumerated) {
	m.BACnetError.ErrorClass = errorClass
	m.BACnetError.ErrorCode = errorCode
}

func (m *BACnetErrorConfirmedEventNotification) GetParent() *BACnetError {
	return m.BACnetError
}

// NewBACnetErrorConfirmedEventNotification factory function for BACnetErrorConfirmedEventNotification
func NewBACnetErrorConfirmedEventNotification(errorClass *BACnetApplicationTagEnumerated, errorCode *BACnetApplicationTagEnumerated) *BACnetErrorConfirmedEventNotification {
	_result := &BACnetErrorConfirmedEventNotification{
		BACnetError: NewBACnetError(errorClass, errorCode),
	}
	_result.Child = _result
	return _result
}

func CastBACnetErrorConfirmedEventNotification(structType interface{}) *BACnetErrorConfirmedEventNotification {
	if casted, ok := structType.(BACnetErrorConfirmedEventNotification); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetErrorConfirmedEventNotification); ok {
		return casted
	}
	if casted, ok := structType.(BACnetError); ok {
		return CastBACnetErrorConfirmedEventNotification(casted.Child)
	}
	if casted, ok := structType.(*BACnetError); ok {
		return CastBACnetErrorConfirmedEventNotification(casted.Child)
	}
	return nil
}

func (m *BACnetErrorConfirmedEventNotification) GetTypeName() string {
	return "BACnetErrorConfirmedEventNotification"
}

func (m *BACnetErrorConfirmedEventNotification) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetErrorConfirmedEventNotification) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *BACnetErrorConfirmedEventNotification) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetErrorConfirmedEventNotificationParse(readBuffer utils.ReadBuffer) (*BACnetErrorConfirmedEventNotification, error) {
	if pullErr := readBuffer.PullContext("BACnetErrorConfirmedEventNotification"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("BACnetErrorConfirmedEventNotification"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetErrorConfirmedEventNotification{
		BACnetError: &BACnetError{},
	}
	_child.BACnetError.Child = _child
	return _child, nil
}

func (m *BACnetErrorConfirmedEventNotification) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetErrorConfirmedEventNotification"); pushErr != nil {
			return pushErr
		}

		if popErr := writeBuffer.PopContext("BACnetErrorConfirmedEventNotification"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetErrorConfirmedEventNotification) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
