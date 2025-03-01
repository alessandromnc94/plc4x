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

// COTPParameterTpduSize is the data-structure of this message
type COTPParameterTpduSize struct {
	*COTPParameter
	TpduSize COTPTpduSize

	// Arguments.
	Rest uint8
}

// ICOTPParameterTpduSize is the corresponding interface of COTPParameterTpduSize
type ICOTPParameterTpduSize interface {
	ICOTPParameter
	// GetTpduSize returns TpduSize (property field)
	GetTpduSize() COTPTpduSize
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

func (m *COTPParameterTpduSize) GetParameterType() uint8 {
	return 0xC0
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *COTPParameterTpduSize) InitializeParent(parent *COTPParameter) {}

func (m *COTPParameterTpduSize) GetParent() *COTPParameter {
	return m.COTPParameter
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *COTPParameterTpduSize) GetTpduSize() COTPTpduSize {
	return m.TpduSize
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCOTPParameterTpduSize factory function for COTPParameterTpduSize
func NewCOTPParameterTpduSize(tpduSize COTPTpduSize, rest uint8) *COTPParameterTpduSize {
	_result := &COTPParameterTpduSize{
		TpduSize:      tpduSize,
		COTPParameter: NewCOTPParameter(rest),
	}
	_result.Child = _result
	return _result
}

func CastCOTPParameterTpduSize(structType interface{}) *COTPParameterTpduSize {
	if casted, ok := structType.(COTPParameterTpduSize); ok {
		return &casted
	}
	if casted, ok := structType.(*COTPParameterTpduSize); ok {
		return casted
	}
	if casted, ok := structType.(COTPParameter); ok {
		return CastCOTPParameterTpduSize(casted.Child)
	}
	if casted, ok := structType.(*COTPParameter); ok {
		return CastCOTPParameterTpduSize(casted.Child)
	}
	return nil
}

func (m *COTPParameterTpduSize) GetTypeName() string {
	return "COTPParameterTpduSize"
}

func (m *COTPParameterTpduSize) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *COTPParameterTpduSize) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (tpduSize)
	lengthInBits += 8

	return lengthInBits
}

func (m *COTPParameterTpduSize) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func COTPParameterTpduSizeParse(readBuffer utils.ReadBuffer, rest uint8) (*COTPParameterTpduSize, error) {
	if pullErr := readBuffer.PullContext("COTPParameterTpduSize"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Simple Field (tpduSize)
	if pullErr := readBuffer.PullContext("tpduSize"); pullErr != nil {
		return nil, pullErr
	}
	_tpduSize, _tpduSizeErr := COTPTpduSizeParse(readBuffer)
	if _tpduSizeErr != nil {
		return nil, errors.Wrap(_tpduSizeErr, "Error parsing 'tpduSize' field")
	}
	tpduSize := _tpduSize
	if closeErr := readBuffer.CloseContext("tpduSize"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("COTPParameterTpduSize"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &COTPParameterTpduSize{
		TpduSize:      tpduSize,
		COTPParameter: &COTPParameter{},
	}
	_child.COTPParameter.Child = _child
	return _child, nil
}

func (m *COTPParameterTpduSize) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("COTPParameterTpduSize"); pushErr != nil {
			return pushErr
		}

		// Simple Field (tpduSize)
		if pushErr := writeBuffer.PushContext("tpduSize"); pushErr != nil {
			return pushErr
		}
		_tpduSizeErr := m.TpduSize.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("tpduSize"); popErr != nil {
			return popErr
		}
		if _tpduSizeErr != nil {
			return errors.Wrap(_tpduSizeErr, "Error serializing 'tpduSize' field")
		}

		if popErr := writeBuffer.PopContext("COTPParameterTpduSize"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *COTPParameterTpduSize) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
