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

// NetworkRoute is the data-structure of this message
type NetworkRoute struct {
	RouteType                 RouteType
	AdditionalBridgeAddresses []*BridgeAddress
}

// INetworkRoute is the corresponding interface of NetworkRoute
type INetworkRoute interface {
	// GetRouteType returns RouteType (property field)
	GetRouteType() RouteType
	// GetAdditionalBridgeAddresses returns AdditionalBridgeAddresses (property field)
	GetAdditionalBridgeAddresses() []*BridgeAddress
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

func (m *NetworkRoute) GetRouteType() RouteType {
	return m.RouteType
}

func (m *NetworkRoute) GetAdditionalBridgeAddresses() []*BridgeAddress {
	return m.AdditionalBridgeAddresses
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewNetworkRoute factory function for NetworkRoute
func NewNetworkRoute(routeType RouteType, additionalBridgeAddresses []*BridgeAddress) *NetworkRoute {
	return &NetworkRoute{RouteType: routeType, AdditionalBridgeAddresses: additionalBridgeAddresses}
}

func CastNetworkRoute(structType interface{}) *NetworkRoute {
	if casted, ok := structType.(NetworkRoute); ok {
		return &casted
	}
	if casted, ok := structType.(*NetworkRoute); ok {
		return casted
	}
	return nil
}

func (m *NetworkRoute) GetTypeName() string {
	return "NetworkRoute"
}

func (m *NetworkRoute) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *NetworkRoute) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (routeType)
	lengthInBits += 8

	// Array field
	if len(m.AdditionalBridgeAddresses) > 0 {
		for i, element := range m.AdditionalBridgeAddresses {
			last := i == len(m.AdditionalBridgeAddresses)-1
			lengthInBits += element.GetLengthInBitsConditional(last)
		}
	}

	return lengthInBits
}

func (m *NetworkRoute) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func NetworkRouteParse(readBuffer utils.ReadBuffer) (*NetworkRoute, error) {
	if pullErr := readBuffer.PullContext("NetworkRoute"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Simple Field (routeType)
	if pullErr := readBuffer.PullContext("routeType"); pullErr != nil {
		return nil, pullErr
	}
	_routeType, _routeTypeErr := RouteTypeParse(readBuffer)
	if _routeTypeErr != nil {
		return nil, errors.Wrap(_routeTypeErr, "Error parsing 'routeType' field")
	}
	routeType := _routeType
	if closeErr := readBuffer.CloseContext("routeType"); closeErr != nil {
		return nil, closeErr
	}

	// Array field (additionalBridgeAddresses)
	if pullErr := readBuffer.PullContext("additionalBridgeAddresses", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, pullErr
	}
	// Count array
	additionalBridgeAddresses := make([]*BridgeAddress, routeType.AdditionalBridges())
	{
		for curItem := uint16(0); curItem < uint16(routeType.AdditionalBridges()); curItem++ {
			_item, _err := BridgeAddressParse(readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'additionalBridgeAddresses' field")
			}
			additionalBridgeAddresses[curItem] = CastBridgeAddress(_item)
		}
	}
	if closeErr := readBuffer.CloseContext("additionalBridgeAddresses", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("NetworkRoute"); closeErr != nil {
		return nil, closeErr
	}

	// Create the instance
	return NewNetworkRoute(routeType, additionalBridgeAddresses), nil
}

func (m *NetworkRoute) Serialize(writeBuffer utils.WriteBuffer) error {
	if pushErr := writeBuffer.PushContext("NetworkRoute"); pushErr != nil {
		return pushErr
	}

	// Simple Field (routeType)
	if pushErr := writeBuffer.PushContext("routeType"); pushErr != nil {
		return pushErr
	}
	_routeTypeErr := m.RouteType.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("routeType"); popErr != nil {
		return popErr
	}
	if _routeTypeErr != nil {
		return errors.Wrap(_routeTypeErr, "Error serializing 'routeType' field")
	}

	// Array Field (additionalBridgeAddresses)
	if m.AdditionalBridgeAddresses != nil {
		if pushErr := writeBuffer.PushContext("additionalBridgeAddresses", utils.WithRenderAsList(true)); pushErr != nil {
			return pushErr
		}
		for _, _element := range m.AdditionalBridgeAddresses {
			_elementErr := _element.Serialize(writeBuffer)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'additionalBridgeAddresses' field")
			}
		}
		if popErr := writeBuffer.PopContext("additionalBridgeAddresses", utils.WithRenderAsList(true)); popErr != nil {
			return popErr
		}
	}

	if popErr := writeBuffer.PopContext("NetworkRoute"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *NetworkRoute) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
