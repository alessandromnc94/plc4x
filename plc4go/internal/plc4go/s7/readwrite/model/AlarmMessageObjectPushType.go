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
const AlarmMessageObjectPushType_VARIABLESPEC uint8 = 0x12

// AlarmMessageObjectPushType is the data-structure of this message
type AlarmMessageObjectPushType struct {
	LengthSpec       uint8
	SyntaxId         SyntaxIdType
	NumberOfValues   uint8
	EventId          uint32
	EventState       *State
	LocalState       *State
	AckStateGoing    *State
	AckStateComing   *State
	AssociatedValues []*AssociatedValueType
}

// IAlarmMessageObjectPushType is the corresponding interface of AlarmMessageObjectPushType
type IAlarmMessageObjectPushType interface {
	// GetLengthSpec returns LengthSpec (property field)
	GetLengthSpec() uint8
	// GetSyntaxId returns SyntaxId (property field)
	GetSyntaxId() SyntaxIdType
	// GetNumberOfValues returns NumberOfValues (property field)
	GetNumberOfValues() uint8
	// GetEventId returns EventId (property field)
	GetEventId() uint32
	// GetEventState returns EventState (property field)
	GetEventState() *State
	// GetLocalState returns LocalState (property field)
	GetLocalState() *State
	// GetAckStateGoing returns AckStateGoing (property field)
	GetAckStateGoing() *State
	// GetAckStateComing returns AckStateComing (property field)
	GetAckStateComing() *State
	// GetAssociatedValues returns AssociatedValues (property field)
	GetAssociatedValues() []*AssociatedValueType
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

func (m *AlarmMessageObjectPushType) GetLengthSpec() uint8 {
	return m.LengthSpec
}

func (m *AlarmMessageObjectPushType) GetSyntaxId() SyntaxIdType {
	return m.SyntaxId
}

func (m *AlarmMessageObjectPushType) GetNumberOfValues() uint8 {
	return m.NumberOfValues
}

func (m *AlarmMessageObjectPushType) GetEventId() uint32 {
	return m.EventId
}

func (m *AlarmMessageObjectPushType) GetEventState() *State {
	return m.EventState
}

func (m *AlarmMessageObjectPushType) GetLocalState() *State {
	return m.LocalState
}

func (m *AlarmMessageObjectPushType) GetAckStateGoing() *State {
	return m.AckStateGoing
}

func (m *AlarmMessageObjectPushType) GetAckStateComing() *State {
	return m.AckStateComing
}

func (m *AlarmMessageObjectPushType) GetAssociatedValues() []*AssociatedValueType {
	return m.AssociatedValues
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for const fields.
///////////////////////

func (m *AlarmMessageObjectPushType) GetVariableSpec() uint8 {
	return AlarmMessageObjectPushType_VARIABLESPEC
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewAlarmMessageObjectPushType factory function for AlarmMessageObjectPushType
func NewAlarmMessageObjectPushType(lengthSpec uint8, syntaxId SyntaxIdType, numberOfValues uint8, eventId uint32, eventState *State, localState *State, ackStateGoing *State, ackStateComing *State, AssociatedValues []*AssociatedValueType) *AlarmMessageObjectPushType {
	return &AlarmMessageObjectPushType{LengthSpec: lengthSpec, SyntaxId: syntaxId, NumberOfValues: numberOfValues, EventId: eventId, EventState: eventState, LocalState: localState, AckStateGoing: ackStateGoing, AckStateComing: ackStateComing, AssociatedValues: AssociatedValues}
}

func CastAlarmMessageObjectPushType(structType interface{}) *AlarmMessageObjectPushType {
	if casted, ok := structType.(AlarmMessageObjectPushType); ok {
		return &casted
	}
	if casted, ok := structType.(*AlarmMessageObjectPushType); ok {
		return casted
	}
	return nil
}

func (m *AlarmMessageObjectPushType) GetTypeName() string {
	return "AlarmMessageObjectPushType"
}

func (m *AlarmMessageObjectPushType) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *AlarmMessageObjectPushType) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Const Field (variableSpec)
	lengthInBits += 8

	// Simple field (lengthSpec)
	lengthInBits += 8

	// Simple field (syntaxId)
	lengthInBits += 8

	// Simple field (numberOfValues)
	lengthInBits += 8

	// Simple field (eventId)
	lengthInBits += 32

	// Simple field (eventState)
	lengthInBits += m.EventState.GetLengthInBits()

	// Simple field (localState)
	lengthInBits += m.LocalState.GetLengthInBits()

	// Simple field (ackStateGoing)
	lengthInBits += m.AckStateGoing.GetLengthInBits()

	// Simple field (ackStateComing)
	lengthInBits += m.AckStateComing.GetLengthInBits()

	// Array field
	if len(m.AssociatedValues) > 0 {
		for i, element := range m.AssociatedValues {
			last := i == len(m.AssociatedValues)-1
			lengthInBits += element.GetLengthInBitsConditional(last)
		}
	}

	return lengthInBits
}

func (m *AlarmMessageObjectPushType) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func AlarmMessageObjectPushTypeParse(readBuffer utils.ReadBuffer) (*AlarmMessageObjectPushType, error) {
	if pullErr := readBuffer.PullContext("AlarmMessageObjectPushType"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Const Field (variableSpec)
	variableSpec, _variableSpecErr := readBuffer.ReadUint8("variableSpec", 8)
	if _variableSpecErr != nil {
		return nil, errors.Wrap(_variableSpecErr, "Error parsing 'variableSpec' field")
	}
	if variableSpec != AlarmMessageObjectPushType_VARIABLESPEC {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", AlarmMessageObjectPushType_VARIABLESPEC) + " but got " + fmt.Sprintf("%d", variableSpec))
	}

	// Simple Field (lengthSpec)
	_lengthSpec, _lengthSpecErr := readBuffer.ReadUint8("lengthSpec", 8)
	if _lengthSpecErr != nil {
		return nil, errors.Wrap(_lengthSpecErr, "Error parsing 'lengthSpec' field")
	}
	lengthSpec := _lengthSpec

	// Simple Field (syntaxId)
	if pullErr := readBuffer.PullContext("syntaxId"); pullErr != nil {
		return nil, pullErr
	}
	_syntaxId, _syntaxIdErr := SyntaxIdTypeParse(readBuffer)
	if _syntaxIdErr != nil {
		return nil, errors.Wrap(_syntaxIdErr, "Error parsing 'syntaxId' field")
	}
	syntaxId := _syntaxId
	if closeErr := readBuffer.CloseContext("syntaxId"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (numberOfValues)
	_numberOfValues, _numberOfValuesErr := readBuffer.ReadUint8("numberOfValues", 8)
	if _numberOfValuesErr != nil {
		return nil, errors.Wrap(_numberOfValuesErr, "Error parsing 'numberOfValues' field")
	}
	numberOfValues := _numberOfValues

	// Simple Field (eventId)
	_eventId, _eventIdErr := readBuffer.ReadUint32("eventId", 32)
	if _eventIdErr != nil {
		return nil, errors.Wrap(_eventIdErr, "Error parsing 'eventId' field")
	}
	eventId := _eventId

	// Simple Field (eventState)
	if pullErr := readBuffer.PullContext("eventState"); pullErr != nil {
		return nil, pullErr
	}
	_eventState, _eventStateErr := StateParse(readBuffer)
	if _eventStateErr != nil {
		return nil, errors.Wrap(_eventStateErr, "Error parsing 'eventState' field")
	}
	eventState := CastState(_eventState)
	if closeErr := readBuffer.CloseContext("eventState"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (localState)
	if pullErr := readBuffer.PullContext("localState"); pullErr != nil {
		return nil, pullErr
	}
	_localState, _localStateErr := StateParse(readBuffer)
	if _localStateErr != nil {
		return nil, errors.Wrap(_localStateErr, "Error parsing 'localState' field")
	}
	localState := CastState(_localState)
	if closeErr := readBuffer.CloseContext("localState"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (ackStateGoing)
	if pullErr := readBuffer.PullContext("ackStateGoing"); pullErr != nil {
		return nil, pullErr
	}
	_ackStateGoing, _ackStateGoingErr := StateParse(readBuffer)
	if _ackStateGoingErr != nil {
		return nil, errors.Wrap(_ackStateGoingErr, "Error parsing 'ackStateGoing' field")
	}
	ackStateGoing := CastState(_ackStateGoing)
	if closeErr := readBuffer.CloseContext("ackStateGoing"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (ackStateComing)
	if pullErr := readBuffer.PullContext("ackStateComing"); pullErr != nil {
		return nil, pullErr
	}
	_ackStateComing, _ackStateComingErr := StateParse(readBuffer)
	if _ackStateComingErr != nil {
		return nil, errors.Wrap(_ackStateComingErr, "Error parsing 'ackStateComing' field")
	}
	ackStateComing := CastState(_ackStateComing)
	if closeErr := readBuffer.CloseContext("ackStateComing"); closeErr != nil {
		return nil, closeErr
	}

	// Array field (AssociatedValues)
	if pullErr := readBuffer.PullContext("AssociatedValues", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, pullErr
	}
	// Count array
	AssociatedValues := make([]*AssociatedValueType, numberOfValues)
	{
		for curItem := uint16(0); curItem < uint16(numberOfValues); curItem++ {
			_item, _err := AssociatedValueTypeParse(readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'AssociatedValues' field")
			}
			AssociatedValues[curItem] = CastAssociatedValueType(_item)
		}
	}
	if closeErr := readBuffer.CloseContext("AssociatedValues", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("AlarmMessageObjectPushType"); closeErr != nil {
		return nil, closeErr
	}

	// Create the instance
	return NewAlarmMessageObjectPushType(lengthSpec, syntaxId, numberOfValues, eventId, eventState, localState, ackStateGoing, ackStateComing, AssociatedValues), nil
}

func (m *AlarmMessageObjectPushType) Serialize(writeBuffer utils.WriteBuffer) error {
	if pushErr := writeBuffer.PushContext("AlarmMessageObjectPushType"); pushErr != nil {
		return pushErr
	}

	// Const Field (variableSpec)
	_variableSpecErr := writeBuffer.WriteUint8("variableSpec", 8, 0x12)
	if _variableSpecErr != nil {
		return errors.Wrap(_variableSpecErr, "Error serializing 'variableSpec' field")
	}

	// Simple Field (lengthSpec)
	lengthSpec := uint8(m.LengthSpec)
	_lengthSpecErr := writeBuffer.WriteUint8("lengthSpec", 8, (lengthSpec))
	if _lengthSpecErr != nil {
		return errors.Wrap(_lengthSpecErr, "Error serializing 'lengthSpec' field")
	}

	// Simple Field (syntaxId)
	if pushErr := writeBuffer.PushContext("syntaxId"); pushErr != nil {
		return pushErr
	}
	_syntaxIdErr := m.SyntaxId.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("syntaxId"); popErr != nil {
		return popErr
	}
	if _syntaxIdErr != nil {
		return errors.Wrap(_syntaxIdErr, "Error serializing 'syntaxId' field")
	}

	// Simple Field (numberOfValues)
	numberOfValues := uint8(m.NumberOfValues)
	_numberOfValuesErr := writeBuffer.WriteUint8("numberOfValues", 8, (numberOfValues))
	if _numberOfValuesErr != nil {
		return errors.Wrap(_numberOfValuesErr, "Error serializing 'numberOfValues' field")
	}

	// Simple Field (eventId)
	eventId := uint32(m.EventId)
	_eventIdErr := writeBuffer.WriteUint32("eventId", 32, (eventId))
	if _eventIdErr != nil {
		return errors.Wrap(_eventIdErr, "Error serializing 'eventId' field")
	}

	// Simple Field (eventState)
	if pushErr := writeBuffer.PushContext("eventState"); pushErr != nil {
		return pushErr
	}
	_eventStateErr := m.EventState.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("eventState"); popErr != nil {
		return popErr
	}
	if _eventStateErr != nil {
		return errors.Wrap(_eventStateErr, "Error serializing 'eventState' field")
	}

	// Simple Field (localState)
	if pushErr := writeBuffer.PushContext("localState"); pushErr != nil {
		return pushErr
	}
	_localStateErr := m.LocalState.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("localState"); popErr != nil {
		return popErr
	}
	if _localStateErr != nil {
		return errors.Wrap(_localStateErr, "Error serializing 'localState' field")
	}

	// Simple Field (ackStateGoing)
	if pushErr := writeBuffer.PushContext("ackStateGoing"); pushErr != nil {
		return pushErr
	}
	_ackStateGoingErr := m.AckStateGoing.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("ackStateGoing"); popErr != nil {
		return popErr
	}
	if _ackStateGoingErr != nil {
		return errors.Wrap(_ackStateGoingErr, "Error serializing 'ackStateGoing' field")
	}

	// Simple Field (ackStateComing)
	if pushErr := writeBuffer.PushContext("ackStateComing"); pushErr != nil {
		return pushErr
	}
	_ackStateComingErr := m.AckStateComing.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("ackStateComing"); popErr != nil {
		return popErr
	}
	if _ackStateComingErr != nil {
		return errors.Wrap(_ackStateComingErr, "Error serializing 'ackStateComing' field")
	}

	// Array Field (AssociatedValues)
	if m.AssociatedValues != nil {
		if pushErr := writeBuffer.PushContext("AssociatedValues", utils.WithRenderAsList(true)); pushErr != nil {
			return pushErr
		}
		for _, _element := range m.AssociatedValues {
			_elementErr := _element.Serialize(writeBuffer)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'AssociatedValues' field")
			}
		}
		if popErr := writeBuffer.PopContext("AssociatedValues", utils.WithRenderAsList(true)); popErr != nil {
			return popErr
		}
	}

	if popErr := writeBuffer.PopContext("AlarmMessageObjectPushType"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *AlarmMessageObjectPushType) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
