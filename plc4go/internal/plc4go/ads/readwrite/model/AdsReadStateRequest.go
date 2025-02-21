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

// AdsReadStateRequest is the data-structure of this message
type AdsReadStateRequest struct {
	*AdsData
}

// IAdsReadStateRequest is the corresponding interface of AdsReadStateRequest
type IAdsReadStateRequest interface {
	IAdsData
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

func (m *AdsReadStateRequest) GetCommandId() CommandId {
	return CommandId_ADS_READ_STATE
}

func (m *AdsReadStateRequest) GetResponse() bool {
	return bool(false)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *AdsReadStateRequest) InitializeParent(parent *AdsData) {}

func (m *AdsReadStateRequest) GetParent() *AdsData {
	return m.AdsData
}

// NewAdsReadStateRequest factory function for AdsReadStateRequest
func NewAdsReadStateRequest() *AdsReadStateRequest {
	_result := &AdsReadStateRequest{
		AdsData: NewAdsData(),
	}
	_result.Child = _result
	return _result
}

func CastAdsReadStateRequest(structType interface{}) *AdsReadStateRequest {
	if casted, ok := structType.(AdsReadStateRequest); ok {
		return &casted
	}
	if casted, ok := structType.(*AdsReadStateRequest); ok {
		return casted
	}
	if casted, ok := structType.(AdsData); ok {
		return CastAdsReadStateRequest(casted.Child)
	}
	if casted, ok := structType.(*AdsData); ok {
		return CastAdsReadStateRequest(casted.Child)
	}
	return nil
}

func (m *AdsReadStateRequest) GetTypeName() string {
	return "AdsReadStateRequest"
}

func (m *AdsReadStateRequest) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *AdsReadStateRequest) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *AdsReadStateRequest) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func AdsReadStateRequestParse(readBuffer utils.ReadBuffer, commandId CommandId, response bool) (*AdsReadStateRequest, error) {
	if pullErr := readBuffer.PullContext("AdsReadStateRequest"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("AdsReadStateRequest"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &AdsReadStateRequest{
		AdsData: &AdsData{},
	}
	_child.AdsData.Child = _child
	return _child, nil
}

func (m *AdsReadStateRequest) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AdsReadStateRequest"); pushErr != nil {
			return pushErr
		}

		if popErr := writeBuffer.PopContext("AdsReadStateRequest"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *AdsReadStateRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
