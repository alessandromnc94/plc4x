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

#include <stdio.h>
#include <plc4c/spi/evaluation_helper.h>
#include "alarm_message_object_ack_type.h"

// Code generated by code-generation. DO NOT EDIT.


// Constant values.
static const uint8_t PLC4C_S7_READ_WRITE_ALARM_MESSAGE_OBJECT_ACK_TYPE_LENGTH_const = 0x08;
uint8_t PLC4C_S7_READ_WRITE_ALARM_MESSAGE_OBJECT_ACK_TYPE_LENGTH() {
  return PLC4C_S7_READ_WRITE_ALARM_MESSAGE_OBJECT_ACK_TYPE_LENGTH_const;
}
static const uint8_t PLC4C_S7_READ_WRITE_ALARM_MESSAGE_OBJECT_ACK_TYPE_VARIABLE_SPEC_const = 0x12;
uint8_t PLC4C_S7_READ_WRITE_ALARM_MESSAGE_OBJECT_ACK_TYPE_VARIABLE_SPEC() {
  return PLC4C_S7_READ_WRITE_ALARM_MESSAGE_OBJECT_ACK_TYPE_VARIABLE_SPEC_const;
}

// Parse function.
plc4c_return_code plc4c_s7_read_write_alarm_message_object_ack_type_parse(plc4c_spi_read_buffer* readBuffer, plc4c_s7_read_write_alarm_message_object_ack_type** _message) {
  uint16_t startPos = plc4c_spi_read_get_pos(readBuffer);
  plc4c_return_code _res = OK;

  // Allocate enough memory to contain this data structure.
  (*_message) = malloc(sizeof(plc4c_s7_read_write_alarm_message_object_ack_type));
  if(*_message == NULL) {
    return NO_MEMORY;
  }

  // Const Field (variableSpec)
  uint8_t variableSpec = 0;
  _res = plc4c_spi_read_unsigned_byte(readBuffer, 8, (uint8_t*) &variableSpec);
  if(_res != OK) {
    return _res;
  }
  if(variableSpec != PLC4C_S7_READ_WRITE_ALARM_MESSAGE_OBJECT_ACK_TYPE_VARIABLE_SPEC()) {
    return PARSE_ERROR;
    // throw new ParseException("Expected constant value " + PLC4C_S7_READ_WRITE_ALARM_MESSAGE_OBJECT_ACK_TYPE_VARIABLE_SPEC + " but got " + variableSpec);
  }

  // Const Field (length)
  uint8_t length = 0;
  _res = plc4c_spi_read_unsigned_byte(readBuffer, 8, (uint8_t*) &length);
  if(_res != OK) {
    return _res;
  }
  if(length != PLC4C_S7_READ_WRITE_ALARM_MESSAGE_OBJECT_ACK_TYPE_LENGTH()) {
    return PARSE_ERROR;
    // throw new ParseException("Expected constant value " + PLC4C_S7_READ_WRITE_ALARM_MESSAGE_OBJECT_ACK_TYPE_LENGTH + " but got " + length);
  }

  // Simple Field (syntaxId)
  plc4c_s7_read_write_syntax_id_type syntaxId;
  _res = plc4c_s7_read_write_syntax_id_type_parse(readBuffer, (void*) &syntaxId);
  if(_res != OK) {
    return _res;
  }
  (*_message)->syntax_id = syntaxId;

  // Simple Field (numberOfValues)
  uint8_t numberOfValues = 0;
  _res = plc4c_spi_read_unsigned_byte(readBuffer, 8, (uint8_t*) &numberOfValues);
  if(_res != OK) {
    return _res;
  }
  (*_message)->number_of_values = numberOfValues;

  // Simple Field (eventId)
  uint32_t eventId = 0;
  _res = plc4c_spi_read_unsigned_int(readBuffer, 32, (uint32_t*) &eventId);
  if(_res != OK) {
    return _res;
  }
  (*_message)->event_id = eventId;

  // Simple Field (ackStateGoing)
  plc4c_s7_read_write_state* ackStateGoing;
  _res = plc4c_s7_read_write_state_parse(readBuffer, (void*) &ackStateGoing);
  if(_res != OK) {
    return _res;
  }
  (*_message)->ack_state_going = ackStateGoing;

  // Simple Field (ackStateComing)
  plc4c_s7_read_write_state* ackStateComing;
  _res = plc4c_s7_read_write_state_parse(readBuffer, (void*) &ackStateComing);
  if(_res != OK) {
    return _res;
  }
  (*_message)->ack_state_coming = ackStateComing;

  return OK;
}

plc4c_return_code plc4c_s7_read_write_alarm_message_object_ack_type_serialize(plc4c_spi_write_buffer* writeBuffer, plc4c_s7_read_write_alarm_message_object_ack_type* _message) {
  plc4c_return_code _res = OK;

  // Const Field (variableSpec)
  uint8_t variableSpec = PLC4C_S7_READ_WRITE_ALARM_MESSAGE_OBJECT_ACK_TYPE_VARIABLE_SPEC();
  plc4c_spi_write_unsigned_byte(writeBuffer, 8, variableSpec);

  // Const Field (length)
  uint8_t length = PLC4C_S7_READ_WRITE_ALARM_MESSAGE_OBJECT_ACK_TYPE_LENGTH();
  plc4c_spi_write_unsigned_byte(writeBuffer, 8, length);

  // Simple Field (syntaxId)
  plc4c_s7_read_write_syntax_id_type syntaxId = _message->syntax_id;
  _res = plc4c_s7_read_write_syntax_id_type_serialize(writeBuffer, &syntaxId);
  if(_res != OK) {
    return _res;
  }

  // Simple Field (numberOfValues)
  uint8_t numberOfValues = _message->number_of_values;
  _res = plc4c_spi_write_unsigned_byte(writeBuffer, 8, numberOfValues);
  if(_res != OK) {
    return _res;
  }

  // Simple Field (eventId)
  uint32_t eventId = _message->event_id;
  _res = plc4c_spi_write_unsigned_int(writeBuffer, 32, eventId);
  if(_res != OK) {
    return _res;
  }

  // Simple Field (ackStateGoing)
  plc4c_s7_read_write_state* ackStateGoing = _message->ack_state_going;
  _res = plc4c_s7_read_write_state_serialize(writeBuffer, ackStateGoing);
  if(_res != OK) {
    return _res;
  }

  // Simple Field (ackStateComing)
  plc4c_s7_read_write_state* ackStateComing = _message->ack_state_coming;
  _res = plc4c_s7_read_write_state_serialize(writeBuffer, ackStateComing);
  if(_res != OK) {
    return _res;
  }

  return OK;
}

uint16_t plc4c_s7_read_write_alarm_message_object_ack_type_length_in_bytes(plc4c_s7_read_write_alarm_message_object_ack_type* _message) {
  return plc4c_s7_read_write_alarm_message_object_ack_type_length_in_bits(_message) / 8;
}

uint16_t plc4c_s7_read_write_alarm_message_object_ack_type_length_in_bits(plc4c_s7_read_write_alarm_message_object_ack_type* _message) {
  uint16_t lengthInBits = 0;

  // Const Field (variableSpec)
  lengthInBits += 8;

  // Const Field (length)
  lengthInBits += 8;

  // Simple field (syntaxId)
  lengthInBits += plc4c_s7_read_write_syntax_id_type_length_in_bits(&_message->syntax_id);

  // Simple field (numberOfValues)
  lengthInBits += 8;

  // Simple field (eventId)
  lengthInBits += 32;

  // Simple field (ackStateGoing)
  lengthInBits += plc4c_s7_read_write_state_length_in_bits(_message->ack_state_going);

  // Simple field (ackStateComing)
  lengthInBits += plc4c_s7_read_write_state_length_in_bits(_message->ack_state_coming);

  return lengthInBits;
}

