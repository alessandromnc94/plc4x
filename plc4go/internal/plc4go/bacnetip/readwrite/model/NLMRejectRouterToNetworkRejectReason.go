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

type NLMRejectRouterToNetworkRejectReason uint8

type INLMRejectRouterToNetworkRejectReason interface {
	Serialize(writeBuffer utils.WriteBuffer) error
}

const (
	NLMRejectRouterToNetworkRejectReason_OTHER                  NLMRejectRouterToNetworkRejectReason = 0
	NLMRejectRouterToNetworkRejectReason_NOT_DIRECTLY_CONNECTED NLMRejectRouterToNetworkRejectReason = 1
	NLMRejectRouterToNetworkRejectReason_BUSY                   NLMRejectRouterToNetworkRejectReason = 2
	NLMRejectRouterToNetworkRejectReason_UNKNOWN_NLMT           NLMRejectRouterToNetworkRejectReason = 3
	NLMRejectRouterToNetworkRejectReason_TOO_LONG               NLMRejectRouterToNetworkRejectReason = 4
	NLMRejectRouterToNetworkRejectReason_SECURITY_ERROR         NLMRejectRouterToNetworkRejectReason = 5
	NLMRejectRouterToNetworkRejectReason_ADDRESSING_ERROR       NLMRejectRouterToNetworkRejectReason = 6
)

var NLMRejectRouterToNetworkRejectReasonValues []NLMRejectRouterToNetworkRejectReason

func init() {
	_ = errors.New
	NLMRejectRouterToNetworkRejectReasonValues = []NLMRejectRouterToNetworkRejectReason{
		NLMRejectRouterToNetworkRejectReason_OTHER,
		NLMRejectRouterToNetworkRejectReason_NOT_DIRECTLY_CONNECTED,
		NLMRejectRouterToNetworkRejectReason_BUSY,
		NLMRejectRouterToNetworkRejectReason_UNKNOWN_NLMT,
		NLMRejectRouterToNetworkRejectReason_TOO_LONG,
		NLMRejectRouterToNetworkRejectReason_SECURITY_ERROR,
		NLMRejectRouterToNetworkRejectReason_ADDRESSING_ERROR,
	}
}

func NLMRejectRouterToNetworkRejectReasonByValue(value uint8) NLMRejectRouterToNetworkRejectReason {
	switch value {
	case 0:
		return NLMRejectRouterToNetworkRejectReason_OTHER
	case 1:
		return NLMRejectRouterToNetworkRejectReason_NOT_DIRECTLY_CONNECTED
	case 2:
		return NLMRejectRouterToNetworkRejectReason_BUSY
	case 3:
		return NLMRejectRouterToNetworkRejectReason_UNKNOWN_NLMT
	case 4:
		return NLMRejectRouterToNetworkRejectReason_TOO_LONG
	case 5:
		return NLMRejectRouterToNetworkRejectReason_SECURITY_ERROR
	case 6:
		return NLMRejectRouterToNetworkRejectReason_ADDRESSING_ERROR
	}
	return 0
}

func NLMRejectRouterToNetworkRejectReasonByName(value string) NLMRejectRouterToNetworkRejectReason {
	switch value {
	case "OTHER":
		return NLMRejectRouterToNetworkRejectReason_OTHER
	case "NOT_DIRECTLY_CONNECTED":
		return NLMRejectRouterToNetworkRejectReason_NOT_DIRECTLY_CONNECTED
	case "BUSY":
		return NLMRejectRouterToNetworkRejectReason_BUSY
	case "UNKNOWN_NLMT":
		return NLMRejectRouterToNetworkRejectReason_UNKNOWN_NLMT
	case "TOO_LONG":
		return NLMRejectRouterToNetworkRejectReason_TOO_LONG
	case "SECURITY_ERROR":
		return NLMRejectRouterToNetworkRejectReason_SECURITY_ERROR
	case "ADDRESSING_ERROR":
		return NLMRejectRouterToNetworkRejectReason_ADDRESSING_ERROR
	}
	return 0
}

func NLMRejectRouterToNetworkRejectReasonKnows(value uint8) bool {
	for _, typeValue := range NLMRejectRouterToNetworkRejectReasonValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastNLMRejectRouterToNetworkRejectReason(structType interface{}) NLMRejectRouterToNetworkRejectReason {
	castFunc := func(typ interface{}) NLMRejectRouterToNetworkRejectReason {
		if sNLMRejectRouterToNetworkRejectReason, ok := typ.(NLMRejectRouterToNetworkRejectReason); ok {
			return sNLMRejectRouterToNetworkRejectReason
		}
		return 0
	}
	return castFunc(structType)
}

func (m NLMRejectRouterToNetworkRejectReason) GetLengthInBits() uint16 {
	return 8
}

func (m NLMRejectRouterToNetworkRejectReason) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func NLMRejectRouterToNetworkRejectReasonParse(readBuffer utils.ReadBuffer) (NLMRejectRouterToNetworkRejectReason, error) {
	val, err := readBuffer.ReadUint8("NLMRejectRouterToNetworkRejectReason", 8)
	if err != nil {
		return 0, nil
	}
	return NLMRejectRouterToNetworkRejectReasonByValue(val), nil
}

func (e NLMRejectRouterToNetworkRejectReason) Serialize(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint8("NLMRejectRouterToNetworkRejectReason", 8, uint8(e), utils.WithAdditionalStringRepresentation(e.name()))
}

func (e NLMRejectRouterToNetworkRejectReason) name() string {
	switch e {
	case NLMRejectRouterToNetworkRejectReason_OTHER:
		return "OTHER"
	case NLMRejectRouterToNetworkRejectReason_NOT_DIRECTLY_CONNECTED:
		return "NOT_DIRECTLY_CONNECTED"
	case NLMRejectRouterToNetworkRejectReason_BUSY:
		return "BUSY"
	case NLMRejectRouterToNetworkRejectReason_UNKNOWN_NLMT:
		return "UNKNOWN_NLMT"
	case NLMRejectRouterToNetworkRejectReason_TOO_LONG:
		return "TOO_LONG"
	case NLMRejectRouterToNetworkRejectReason_SECURITY_ERROR:
		return "SECURITY_ERROR"
	case NLMRejectRouterToNetworkRejectReason_ADDRESSING_ERROR:
		return "ADDRESSING_ERROR"
	}
	return ""
}

func (e NLMRejectRouterToNetworkRejectReason) String() string {
	return e.name()
}
