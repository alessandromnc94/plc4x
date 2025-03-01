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

#ifndef PLC4C_PLC4X_READ_WRITE_PLC4X_REQUEST_TYPE_H_
#define PLC4C_PLC4X_READ_WRITE_PLC4X_REQUEST_TYPE_H_

#include <stdbool.h>
#include <stdint.h>
#include <plc4c/spi/read_buffer.h>
#include <plc4c/spi/write_buffer.h>

// Code generated by code-generation. DO NOT EDIT.


enum plc4c_plc4x_read_write_plc4x_request_type {
  plc4c_plc4x_read_write_plc4x_request_type_CONNECT_REQUEST = 0x01,
  plc4c_plc4x_read_write_plc4x_request_type_CONNECT_RESPONSE = 0x02,
  plc4c_plc4x_read_write_plc4x_request_type_DISCONNECT_REQUEST = 0x03,
  plc4c_plc4x_read_write_plc4x_request_type_DISCONNECT_RESPONSE = 0x04,
  plc4c_plc4x_read_write_plc4x_request_type_READ_REQUEST = 0x05,
  plc4c_plc4x_read_write_plc4x_request_type_READ_RESPONSE = 0x06,
  plc4c_plc4x_read_write_plc4x_request_type_WRITE_REQUEST = 0x07,
  plc4c_plc4x_read_write_plc4x_request_type_WRITE_RESPONSE = 0x08,
  plc4c_plc4x_read_write_plc4x_request_type_SUBSCRIPTION_REQUEST = 0x09,
  plc4c_plc4x_read_write_plc4x_request_type_SUBSCRIPTION_RESPONSE = 0x0A,
  plc4c_plc4x_read_write_plc4x_request_type_UNSUBSCRIPTION_REQUEST = 0x0B,
  plc4c_plc4x_read_write_plc4x_request_type_UNSUBSCRIPTION_RESPONSE = 0x0C
};
typedef enum plc4c_plc4x_read_write_plc4x_request_type plc4c_plc4x_read_write_plc4x_request_type;

// Get an empty NULL-struct
plc4c_plc4x_read_write_plc4x_request_type plc4c_plc4x_read_write_plc4x_request_type_null();

plc4c_return_code plc4c_plc4x_read_write_plc4x_request_type_parse(plc4c_spi_read_buffer* readBuffer, plc4c_plc4x_read_write_plc4x_request_type* message);

plc4c_return_code plc4c_plc4x_read_write_plc4x_request_type_serialize(plc4c_spi_write_buffer* writeBuffer, plc4c_plc4x_read_write_plc4x_request_type* message);

plc4c_plc4x_read_write_plc4x_request_type plc4c_plc4x_read_write_plc4x_request_type_value_of(char* value_string);

int plc4c_plc4x_read_write_plc4x_request_type_num_values();

plc4c_plc4x_read_write_plc4x_request_type plc4c_plc4x_read_write_plc4x_request_type_value_for_index(int index);

uint16_t plc4c_plc4x_read_write_plc4x_request_type_length_in_bytes(plc4c_plc4x_read_write_plc4x_request_type* message);

uint16_t plc4c_plc4x_read_write_plc4x_request_type_length_in_bits(plc4c_plc4x_read_write_plc4x_request_type* message);


#endif  // PLC4C_PLC4X_READ_WRITE_PLC4X_REQUEST_TYPE_H_
