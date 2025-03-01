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
	"fmt"
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// Constant values.
const CEMIAdditionalInformationRelativeTimestamp_LEN uint8 = uint8(2)

// CEMIAdditionalInformationRelativeTimestamp is the data-structure of this message
type CEMIAdditionalInformationRelativeTimestamp struct {
	*CEMIAdditionalInformation
	RelativeTimestamp *RelativeTimestamp
}

// ICEMIAdditionalInformationRelativeTimestamp is the corresponding interface of CEMIAdditionalInformationRelativeTimestamp
type ICEMIAdditionalInformationRelativeTimestamp interface {
	ICEMIAdditionalInformation
	// GetRelativeTimestamp returns RelativeTimestamp (property field)
	GetRelativeTimestamp() *RelativeTimestamp
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

func (m *CEMIAdditionalInformationRelativeTimestamp) GetAdditionalInformationType() uint8 {
	return 0x04
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *CEMIAdditionalInformationRelativeTimestamp) InitializeParent(parent *CEMIAdditionalInformation) {
}

func (m *CEMIAdditionalInformationRelativeTimestamp) GetParent() *CEMIAdditionalInformation {
	return m.CEMIAdditionalInformation
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *CEMIAdditionalInformationRelativeTimestamp) GetRelativeTimestamp() *RelativeTimestamp {
	return m.RelativeTimestamp
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for const fields.
///////////////////////

func (m *CEMIAdditionalInformationRelativeTimestamp) GetLen() uint8 {
	return CEMIAdditionalInformationRelativeTimestamp_LEN
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCEMIAdditionalInformationRelativeTimestamp factory function for CEMIAdditionalInformationRelativeTimestamp
func NewCEMIAdditionalInformationRelativeTimestamp(relativeTimestamp *RelativeTimestamp) *CEMIAdditionalInformationRelativeTimestamp {
	_result := &CEMIAdditionalInformationRelativeTimestamp{
		RelativeTimestamp:         relativeTimestamp,
		CEMIAdditionalInformation: NewCEMIAdditionalInformation(),
	}
	_result.Child = _result
	return _result
}

func CastCEMIAdditionalInformationRelativeTimestamp(structType interface{}) *CEMIAdditionalInformationRelativeTimestamp {
	if casted, ok := structType.(CEMIAdditionalInformationRelativeTimestamp); ok {
		return &casted
	}
	if casted, ok := structType.(*CEMIAdditionalInformationRelativeTimestamp); ok {
		return casted
	}
	if casted, ok := structType.(CEMIAdditionalInformation); ok {
		return CastCEMIAdditionalInformationRelativeTimestamp(casted.Child)
	}
	if casted, ok := structType.(*CEMIAdditionalInformation); ok {
		return CastCEMIAdditionalInformationRelativeTimestamp(casted.Child)
	}
	return nil
}

func (m *CEMIAdditionalInformationRelativeTimestamp) GetTypeName() string {
	return "CEMIAdditionalInformationRelativeTimestamp"
}

func (m *CEMIAdditionalInformationRelativeTimestamp) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *CEMIAdditionalInformationRelativeTimestamp) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Const Field (len)
	lengthInBits += 8

	// Simple field (relativeTimestamp)
	lengthInBits += m.RelativeTimestamp.GetLengthInBits()

	return lengthInBits
}

func (m *CEMIAdditionalInformationRelativeTimestamp) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func CEMIAdditionalInformationRelativeTimestampParse(readBuffer utils.ReadBuffer) (*CEMIAdditionalInformationRelativeTimestamp, error) {
	if pullErr := readBuffer.PullContext("CEMIAdditionalInformationRelativeTimestamp"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Const Field (len)
	len, _lenErr := readBuffer.ReadUint8("len", 8)
	if _lenErr != nil {
		return nil, errors.Wrap(_lenErr, "Error parsing 'len' field")
	}
	if len != CEMIAdditionalInformationRelativeTimestamp_LEN {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", CEMIAdditionalInformationRelativeTimestamp_LEN) + " but got " + fmt.Sprintf("%d", len))
	}

	// Simple Field (relativeTimestamp)
	if pullErr := readBuffer.PullContext("relativeTimestamp"); pullErr != nil {
		return nil, pullErr
	}
	_relativeTimestamp, _relativeTimestampErr := RelativeTimestampParse(readBuffer)
	if _relativeTimestampErr != nil {
		return nil, errors.Wrap(_relativeTimestampErr, "Error parsing 'relativeTimestamp' field")
	}
	relativeTimestamp := CastRelativeTimestamp(_relativeTimestamp)
	if closeErr := readBuffer.CloseContext("relativeTimestamp"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("CEMIAdditionalInformationRelativeTimestamp"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &CEMIAdditionalInformationRelativeTimestamp{
		RelativeTimestamp:         CastRelativeTimestamp(relativeTimestamp),
		CEMIAdditionalInformation: &CEMIAdditionalInformation{},
	}
	_child.CEMIAdditionalInformation.Child = _child
	return _child, nil
}

func (m *CEMIAdditionalInformationRelativeTimestamp) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CEMIAdditionalInformationRelativeTimestamp"); pushErr != nil {
			return pushErr
		}

		// Const Field (len)
		_lenErr := writeBuffer.WriteUint8("len", 8, 2)
		if _lenErr != nil {
			return errors.Wrap(_lenErr, "Error serializing 'len' field")
		}

		// Simple Field (relativeTimestamp)
		if pushErr := writeBuffer.PushContext("relativeTimestamp"); pushErr != nil {
			return pushErr
		}
		_relativeTimestampErr := m.RelativeTimestamp.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("relativeTimestamp"); popErr != nil {
			return popErr
		}
		if _relativeTimestampErr != nil {
			return errors.Wrap(_relativeTimestampErr, "Error serializing 'relativeTimestamp' field")
		}

		if popErr := writeBuffer.PopContext("CEMIAdditionalInformationRelativeTimestamp"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *CEMIAdditionalInformationRelativeTimestamp) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
