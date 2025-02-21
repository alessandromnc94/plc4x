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

// KnxGroupAddress2Level is the data-structure of this message
type KnxGroupAddress2Level struct {
	*KnxGroupAddress
	MainGroup uint8
	SubGroup  uint16
}

// IKnxGroupAddress2Level is the corresponding interface of KnxGroupAddress2Level
type IKnxGroupAddress2Level interface {
	IKnxGroupAddress
	// GetMainGroup returns MainGroup (property field)
	GetMainGroup() uint8
	// GetSubGroup returns SubGroup (property field)
	GetSubGroup() uint16
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

func (m *KnxGroupAddress2Level) GetNumLevels() uint8 {
	return uint8(2)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *KnxGroupAddress2Level) InitializeParent(parent *KnxGroupAddress) {}

func (m *KnxGroupAddress2Level) GetParent() *KnxGroupAddress {
	return m.KnxGroupAddress
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *KnxGroupAddress2Level) GetMainGroup() uint8 {
	return m.MainGroup
}

func (m *KnxGroupAddress2Level) GetSubGroup() uint16 {
	return m.SubGroup
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewKnxGroupAddress2Level factory function for KnxGroupAddress2Level
func NewKnxGroupAddress2Level(mainGroup uint8, subGroup uint16) *KnxGroupAddress2Level {
	_result := &KnxGroupAddress2Level{
		MainGroup:       mainGroup,
		SubGroup:        subGroup,
		KnxGroupAddress: NewKnxGroupAddress(),
	}
	_result.Child = _result
	return _result
}

func CastKnxGroupAddress2Level(structType interface{}) *KnxGroupAddress2Level {
	if casted, ok := structType.(KnxGroupAddress2Level); ok {
		return &casted
	}
	if casted, ok := structType.(*KnxGroupAddress2Level); ok {
		return casted
	}
	if casted, ok := structType.(KnxGroupAddress); ok {
		return CastKnxGroupAddress2Level(casted.Child)
	}
	if casted, ok := structType.(*KnxGroupAddress); ok {
		return CastKnxGroupAddress2Level(casted.Child)
	}
	return nil
}

func (m *KnxGroupAddress2Level) GetTypeName() string {
	return "KnxGroupAddress2Level"
}

func (m *KnxGroupAddress2Level) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *KnxGroupAddress2Level) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (mainGroup)
	lengthInBits += 5

	// Simple field (subGroup)
	lengthInBits += 11

	return lengthInBits
}

func (m *KnxGroupAddress2Level) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func KnxGroupAddress2LevelParse(readBuffer utils.ReadBuffer, numLevels uint8) (*KnxGroupAddress2Level, error) {
	if pullErr := readBuffer.PullContext("KnxGroupAddress2Level"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Simple Field (mainGroup)
	_mainGroup, _mainGroupErr := readBuffer.ReadUint8("mainGroup", 5)
	if _mainGroupErr != nil {
		return nil, errors.Wrap(_mainGroupErr, "Error parsing 'mainGroup' field")
	}
	mainGroup := _mainGroup

	// Simple Field (subGroup)
	_subGroup, _subGroupErr := readBuffer.ReadUint16("subGroup", 11)
	if _subGroupErr != nil {
		return nil, errors.Wrap(_subGroupErr, "Error parsing 'subGroup' field")
	}
	subGroup := _subGroup

	if closeErr := readBuffer.CloseContext("KnxGroupAddress2Level"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &KnxGroupAddress2Level{
		MainGroup:       mainGroup,
		SubGroup:        subGroup,
		KnxGroupAddress: &KnxGroupAddress{},
	}
	_child.KnxGroupAddress.Child = _child
	return _child, nil
}

func (m *KnxGroupAddress2Level) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("KnxGroupAddress2Level"); pushErr != nil {
			return pushErr
		}

		// Simple Field (mainGroup)
		mainGroup := uint8(m.MainGroup)
		_mainGroupErr := writeBuffer.WriteUint8("mainGroup", 5, (mainGroup))
		if _mainGroupErr != nil {
			return errors.Wrap(_mainGroupErr, "Error serializing 'mainGroup' field")
		}

		// Simple Field (subGroup)
		subGroup := uint16(m.SubGroup)
		_subGroupErr := writeBuffer.WriteUint16("subGroup", 11, (subGroup))
		if _subGroupErr != nil {
			return errors.Wrap(_subGroupErr, "Error serializing 'subGroup' field")
		}

		if popErr := writeBuffer.PopContext("KnxGroupAddress2Level"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *KnxGroupAddress2Level) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
