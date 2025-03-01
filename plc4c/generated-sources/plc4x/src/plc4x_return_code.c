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

#include <string.h>
#include <plc4c/driver_plc4x_static.h>

#include "plc4x_return_code.h"

// Code generated by code-generation. DO NOT EDIT.


// Create an empty NULL-struct
static const plc4c_plc4x_read_write_plc4x_return_code plc4c_plc4x_read_write_plc4x_return_code_null_const;

plc4c_plc4x_read_write_plc4x_return_code plc4c_plc4x_read_write_plc4x_return_code_null() {
  return plc4c_plc4x_read_write_plc4x_return_code_null_const;
}

// Parse function.
plc4c_return_code plc4c_plc4x_read_write_plc4x_return_code_parse(plc4c_spi_read_buffer* readBuffer, plc4c_plc4x_read_write_plc4x_return_code* _message) {
    plc4c_return_code _res = OK;

    _res = plc4c_spi_read_unsigned_byte(readBuffer, 8, (uint8_t*) _message);

    return _res;
}

plc4c_return_code plc4c_plc4x_read_write_plc4x_return_code_serialize(plc4c_spi_write_buffer* writeBuffer, plc4c_plc4x_read_write_plc4x_return_code* _message) {
    plc4c_return_code _res = OK;

    _res = plc4c_spi_write_unsigned_byte(writeBuffer, 8, *_message);

    return _res;
}

plc4c_plc4x_read_write_plc4x_return_code plc4c_plc4x_read_write_plc4x_return_code_value_of(char* value_string) {
    if(strcmp(value_string, "OK") == 0) {
        return plc4c_plc4x_read_write_plc4x_return_code_OK;
    }
    if(strcmp(value_string, "NOT_FOUND") == 0) {
        return plc4c_plc4x_read_write_plc4x_return_code_NOT_FOUND;
    }
    if(strcmp(value_string, "ACCESS_DENIED") == 0) {
        return plc4c_plc4x_read_write_plc4x_return_code_ACCESS_DENIED;
    }
    if(strcmp(value_string, "INVALID_ADDRESS") == 0) {
        return plc4c_plc4x_read_write_plc4x_return_code_INVALID_ADDRESS;
    }
    if(strcmp(value_string, "INVALID_DATATYPE") == 0) {
        return plc4c_plc4x_read_write_plc4x_return_code_INVALID_DATATYPE;
    }
    if(strcmp(value_string, "INVALID_DATA") == 0) {
        return plc4c_plc4x_read_write_plc4x_return_code_INVALID_DATA;
    }
    if(strcmp(value_string, "INTERNAL_ERROR") == 0) {
        return plc4c_plc4x_read_write_plc4x_return_code_INTERNAL_ERROR;
    }
    if(strcmp(value_string, "REMOTE_BUSY") == 0) {
        return plc4c_plc4x_read_write_plc4x_return_code_REMOTE_BUSY;
    }
    if(strcmp(value_string, "REMOTE_ERROR") == 0) {
        return plc4c_plc4x_read_write_plc4x_return_code_REMOTE_ERROR;
    }
    if(strcmp(value_string, "UNSUPPORTED") == 0) {
        return plc4c_plc4x_read_write_plc4x_return_code_UNSUPPORTED;
    }
    if(strcmp(value_string, "RESPONSE_PENDING") == 0) {
        return plc4c_plc4x_read_write_plc4x_return_code_RESPONSE_PENDING;
    }
    return -1;
}

int plc4c_plc4x_read_write_plc4x_return_code_num_values() {
  return 11;
}

plc4c_plc4x_read_write_plc4x_return_code plc4c_plc4x_read_write_plc4x_return_code_value_for_index(int index) {
    switch(index) {
      case 0: {
        return plc4c_plc4x_read_write_plc4x_return_code_OK;
      }
      case 1: {
        return plc4c_plc4x_read_write_plc4x_return_code_NOT_FOUND;
      }
      case 2: {
        return plc4c_plc4x_read_write_plc4x_return_code_ACCESS_DENIED;
      }
      case 3: {
        return plc4c_plc4x_read_write_plc4x_return_code_INVALID_ADDRESS;
      }
      case 4: {
        return plc4c_plc4x_read_write_plc4x_return_code_INVALID_DATATYPE;
      }
      case 5: {
        return plc4c_plc4x_read_write_plc4x_return_code_INVALID_DATA;
      }
      case 6: {
        return plc4c_plc4x_read_write_plc4x_return_code_INTERNAL_ERROR;
      }
      case 7: {
        return plc4c_plc4x_read_write_plc4x_return_code_REMOTE_BUSY;
      }
      case 8: {
        return plc4c_plc4x_read_write_plc4x_return_code_REMOTE_ERROR;
      }
      case 9: {
        return plc4c_plc4x_read_write_plc4x_return_code_UNSUPPORTED;
      }
      case 10: {
        return plc4c_plc4x_read_write_plc4x_return_code_RESPONSE_PENDING;
      }
      default: {
        return -1;
      }
    }
}

uint16_t plc4c_plc4x_read_write_plc4x_return_code_length_in_bytes(plc4c_plc4x_read_write_plc4x_return_code* _message) {
    return plc4c_plc4x_read_write_plc4x_return_code_length_in_bits(_message) / 8;
}

uint16_t plc4c_plc4x_read_write_plc4x_return_code_length_in_bits(plc4c_plc4x_read_write_plc4x_return_code* _message) {
    return 8;
}
