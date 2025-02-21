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

// ApduDataMemoryWrite is the data-structure of this message
type ApduDataMemoryWrite struct {
	*ApduData

	// Arguments.
	DataLength uint8
}

// IApduDataMemoryWrite is the corresponding interface of ApduDataMemoryWrite
type IApduDataMemoryWrite interface {
	IApduData
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

func (m *ApduDataMemoryWrite) GetApciType() uint8 {
	return 0xA
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *ApduDataMemoryWrite) InitializeParent(parent *ApduData) {}

func (m *ApduDataMemoryWrite) GetParent() *ApduData {
	return m.ApduData
}

// NewApduDataMemoryWrite factory function for ApduDataMemoryWrite
func NewApduDataMemoryWrite(dataLength uint8) *ApduDataMemoryWrite {
	_result := &ApduDataMemoryWrite{
		ApduData: NewApduData(dataLength),
	}
	_result.Child = _result
	return _result
}

func CastApduDataMemoryWrite(structType interface{}) *ApduDataMemoryWrite {
	if casted, ok := structType.(ApduDataMemoryWrite); ok {
		return &casted
	}
	if casted, ok := structType.(*ApduDataMemoryWrite); ok {
		return casted
	}
	if casted, ok := structType.(ApduData); ok {
		return CastApduDataMemoryWrite(casted.Child)
	}
	if casted, ok := structType.(*ApduData); ok {
		return CastApduDataMemoryWrite(casted.Child)
	}
	return nil
}

func (m *ApduDataMemoryWrite) GetTypeName() string {
	return "ApduDataMemoryWrite"
}

func (m *ApduDataMemoryWrite) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *ApduDataMemoryWrite) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *ApduDataMemoryWrite) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ApduDataMemoryWriteParse(readBuffer utils.ReadBuffer, dataLength uint8) (*ApduDataMemoryWrite, error) {
	if pullErr := readBuffer.PullContext("ApduDataMemoryWrite"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("ApduDataMemoryWrite"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &ApduDataMemoryWrite{
		ApduData: &ApduData{},
	}
	_child.ApduData.Child = _child
	return _child, nil
}

func (m *ApduDataMemoryWrite) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataMemoryWrite"); pushErr != nil {
			return pushErr
		}

		if popErr := writeBuffer.PopContext("ApduDataMemoryWrite"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *ApduDataMemoryWrite) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
