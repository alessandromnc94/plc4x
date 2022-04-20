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

// Code generated by code-generation. DO NOT EDIT.

namespace org.apache.plc4net.drivers.knxnetip.readwrite.model
{

    public enum KnxDatapointMainType
    {
        DPT_UNKNOWN = 0,
        DPT_64_BIT_SET = 1,
        DPT_8_BYTE_UNSIGNED_VALUE = 2,
        DPT_8_BYTE_SIGNED_VALUE = 3,
        DPT_12_BYTE_SIGNED_VALUE = 4,
        DPT_8_BYTE_FLOAT_VALUE = 5,
        DPT_1_BIT = 6,
        DPT_1_BIT_CONTROLLED = 7,
        DPT_3_BIT_CONTROLLED = 8,
        DPT_CHARACTER = 9,
        DPT_8_BIT_UNSIGNED_VALUE = 10,
        DPT_8_BIT_SIGNED_VALUE = 11,
        DPT_2_BYTE_UNSIGNED_VALUE = 12,
        DPT_2_BYTE_SIGNED_VALUE = 13,
        DPT_2_BYTE_FLOAT_VALUE = 14,
        DPT_TIME = 15,
        DPT_DATE = 16,
        DPT_4_BYTE_UNSIGNED_VALUE = 17,
        DPT_4_BYTE_SIGNED_VALUE = 18,
        DPT_4_BYTE_FLOAT_VALUE = 19,
        DPT_ENTRANCE_ACCESS = 20,
        DPT_CHARACTER_STRING = 21,
        DPT_SCENE_NUMBER = 22,
        DPT_SCENE_CONTROL = 23,
        DPT_DATE_TIME = 24,
        DPT_1_BYTE = 25,
        DPT_8_BIT_SET = 26,
        DPT_16_BIT_SET = 27,
        DPT_2_BIT_SET = 28,
        DPT_2_NIBBLE_SET = 29,
        DPT_8_BIT_SET_2 = 30,
        DPT_32_BIT_SET = 31,
        DPT_ELECTRICAL_ENERGY = 32,
        DPT_24_TIMES_CHANNEL_ACTIVATION = 33,
        DPT_16_BIT_UNSIGNED_VALUE_AND_8_BIT_ENUM = 34,
        DPT_8_BIT_UNSIGNED_VALUE_AND_8_BIT_ENUM = 35,
        DPT_DATAPOINT_TYPE_VERSION = 36,
        DPT_ALARM_INFO = 37,
        DPT_3X_2_BYTE_FLOAT_VALUE = 38,
        DPT_SCALING_SPEED = 39,
        DPT_4_1_1_BYTE_COMBINED_INFORMATION = 40,
        DPT_MBUS_ADDRESS = 41,
        DPT_3_BYTE_COLOUR_RGB = 42,
        DPT_LANGUAGE_CODE_ISO_639_1 = 43,
        DPT_SIGNED_VALUE_WITH_CLASSIFICATION_AND_VALIDITY = 44,
        DPT_PRIORITISED_MODE_CONTROL = 45,
        DPT_CONFIGURATION_DIAGNOSTICS_16_BIT = 46,
        DPT_CONFIGURATION_DIAGNOSTICS_8_BIT = 47,
        DPT_POSITIONS = 48,
        DPT_STATUS_32_BIT = 49,
        DPT_STATUS_48_BIT = 50,
        DPT_CONVERTER_STATUS = 51,
        DPT_CONVERTER_TEST_RESULT = 52,
        DPT_BATTERY_INFORMATION = 53,
        DPT_BRIGHTNESS_COLOUR_TEMPERATURE_TRANSITION = 54,
        DPT_STATUS_24_BIT = 55,
        DPT_COLOUR_RGBW = 56,
        DPT_RELATIVE_CONTROL_RGBW = 57,
        DPT_RELATIVE_CONTROL_RGB = 58,
        DPT_F32F32 = 59,
        DPT_F16F16F16F16 = 60,
    }

    public static class KnxDatapointMainTypeInfo
    {

        public static ushort? Number(this KnxDatapointMainType value)
        {
            switch (value)
            {
                case KnxDatapointMainType.DPT_UNKNOWN: { /* '0' */
                    return 0;
                }
                case KnxDatapointMainType.DPT_64_BIT_SET: { /* '1' */
                    return 0;
                }
                case KnxDatapointMainType.DPT_8_BIT_UNSIGNED_VALUE: { /* '10' */
                    return 5;
                }
                case KnxDatapointMainType.DPT_8_BIT_SIGNED_VALUE: { /* '11' */
                    return 6;
                }
                case KnxDatapointMainType.DPT_2_BYTE_UNSIGNED_VALUE: { /* '12' */
                    return 7;
                }
                case KnxDatapointMainType.DPT_2_BYTE_SIGNED_VALUE: { /* '13' */
                    return 8;
                }
                case KnxDatapointMainType.DPT_2_BYTE_FLOAT_VALUE: { /* '14' */
                    return 9;
                }
                case KnxDatapointMainType.DPT_TIME: { /* '15' */
                    return 10;
                }
                case KnxDatapointMainType.DPT_DATE: { /* '16' */
                    return 11;
                }
                case KnxDatapointMainType.DPT_4_BYTE_UNSIGNED_VALUE: { /* '17' */
                    return 12;
                }
                case KnxDatapointMainType.DPT_4_BYTE_SIGNED_VALUE: { /* '18' */
                    return 13;
                }
                case KnxDatapointMainType.DPT_4_BYTE_FLOAT_VALUE: { /* '19' */
                    return 14;
                }
                case KnxDatapointMainType.DPT_8_BYTE_UNSIGNED_VALUE: { /* '2' */
                    return 0;
                }
                case KnxDatapointMainType.DPT_ENTRANCE_ACCESS: { /* '20' */
                    return 15;
                }
                case KnxDatapointMainType.DPT_CHARACTER_STRING: { /* '21' */
                    return 16;
                }
                case KnxDatapointMainType.DPT_SCENE_NUMBER: { /* '22' */
                    return 17;
                }
                case KnxDatapointMainType.DPT_SCENE_CONTROL: { /* '23' */
                    return 18;
                }
                case KnxDatapointMainType.DPT_DATE_TIME: { /* '24' */
                    return 19;
                }
                case KnxDatapointMainType.DPT_1_BYTE: { /* '25' */
                    return 20;
                }
                case KnxDatapointMainType.DPT_8_BIT_SET: { /* '26' */
                    return 21;
                }
                case KnxDatapointMainType.DPT_16_BIT_SET: { /* '27' */
                    return 22;
                }
                case KnxDatapointMainType.DPT_2_BIT_SET: { /* '28' */
                    return 23;
                }
                case KnxDatapointMainType.DPT_2_NIBBLE_SET: { /* '29' */
                    return 25;
                }
                case KnxDatapointMainType.DPT_8_BYTE_SIGNED_VALUE: { /* '3' */
                    return 0;
                }
                case KnxDatapointMainType.DPT_8_BIT_SET_2: { /* '30' */
                    return 26;
                }
                case KnxDatapointMainType.DPT_32_BIT_SET: { /* '31' */
                    return 27;
                }
                case KnxDatapointMainType.DPT_ELECTRICAL_ENERGY: { /* '32' */
                    return 29;
                }
                case KnxDatapointMainType.DPT_24_TIMES_CHANNEL_ACTIVATION: { /* '33' */
                    return 30;
                }
                case KnxDatapointMainType.DPT_16_BIT_UNSIGNED_VALUE_AND_8_BIT_ENUM: { /* '34' */
                    return 206;
                }
                case KnxDatapointMainType.DPT_8_BIT_UNSIGNED_VALUE_AND_8_BIT_ENUM: { /* '35' */
                    return 207;
                }
                case KnxDatapointMainType.DPT_DATAPOINT_TYPE_VERSION: { /* '36' */
                    return 217;
                }
                case KnxDatapointMainType.DPT_ALARM_INFO: { /* '37' */
                    return 219;
                }
                case KnxDatapointMainType.DPT_3X_2_BYTE_FLOAT_VALUE: { /* '38' */
                    return 222;
                }
                case KnxDatapointMainType.DPT_SCALING_SPEED: { /* '39' */
                    return 225;
                }
                case KnxDatapointMainType.DPT_12_BYTE_SIGNED_VALUE: { /* '4' */
                    return 0;
                }
                case KnxDatapointMainType.DPT_4_1_1_BYTE_COMBINED_INFORMATION: { /* '40' */
                    return 229;
                }
                case KnxDatapointMainType.DPT_MBUS_ADDRESS: { /* '41' */
                    return 230;
                }
                case KnxDatapointMainType.DPT_3_BYTE_COLOUR_RGB: { /* '42' */
                    return 232;
                }
                case KnxDatapointMainType.DPT_LANGUAGE_CODE_ISO_639_1: { /* '43' */
                    return 234;
                }
                case KnxDatapointMainType.DPT_SIGNED_VALUE_WITH_CLASSIFICATION_AND_VALIDITY: { /* '44' */
                    return 235;
                }
                case KnxDatapointMainType.DPT_PRIORITISED_MODE_CONTROL: { /* '45' */
                    return 236;
                }
                case KnxDatapointMainType.DPT_CONFIGURATION_DIAGNOSTICS_16_BIT: { /* '46' */
                    return 237;
                }
                case KnxDatapointMainType.DPT_CONFIGURATION_DIAGNOSTICS_8_BIT: { /* '47' */
                    return 238;
                }
                case KnxDatapointMainType.DPT_POSITIONS: { /* '48' */
                    return 240;
                }
                case KnxDatapointMainType.DPT_STATUS_32_BIT: { /* '49' */
                    return 241;
                }
                case KnxDatapointMainType.DPT_8_BYTE_FLOAT_VALUE: { /* '5' */
                    return 0;
                }
                case KnxDatapointMainType.DPT_STATUS_48_BIT: { /* '50' */
                    return 242;
                }
                case KnxDatapointMainType.DPT_CONVERTER_STATUS: { /* '51' */
                    return 244;
                }
                case KnxDatapointMainType.DPT_CONVERTER_TEST_RESULT: { /* '52' */
                    return 245;
                }
                case KnxDatapointMainType.DPT_BATTERY_INFORMATION: { /* '53' */
                    return 246;
                }
                case KnxDatapointMainType.DPT_BRIGHTNESS_COLOUR_TEMPERATURE_TRANSITION: { /* '54' */
                    return 249;
                }
                case KnxDatapointMainType.DPT_STATUS_24_BIT: { /* '55' */
                    return 250;
                }
                case KnxDatapointMainType.DPT_COLOUR_RGBW: { /* '56' */
                    return 251;
                }
                case KnxDatapointMainType.DPT_RELATIVE_CONTROL_RGBW: { /* '57' */
                    return 252;
                }
                case KnxDatapointMainType.DPT_RELATIVE_CONTROL_RGB: { /* '58' */
                    return 254;
                }
                case KnxDatapointMainType.DPT_F32F32: { /* '59' */
                    return 255;
                }
                case KnxDatapointMainType.DPT_1_BIT: { /* '6' */
                    return 1;
                }
                case KnxDatapointMainType.DPT_F16F16F16F16: { /* '60' */
                    return 275;
                }
                case KnxDatapointMainType.DPT_1_BIT_CONTROLLED: { /* '7' */
                    return 2;
                }
                case KnxDatapointMainType.DPT_3_BIT_CONTROLLED: { /* '8' */
                    return 3;
                }
                case KnxDatapointMainType.DPT_CHARACTER: { /* '9' */
                    return 4;
                }
                default: {
                    return 0;
                }
            }
        }

        public static string Name(this KnxDatapointMainType value)
        {
            switch (value)
            {
                case KnxDatapointMainType.DPT_UNKNOWN: { /* '0' */
                    return "Unknown Datapoint Type";
                }
                case KnxDatapointMainType.DPT_64_BIT_SET: { /* '1' */
                    return "Unknown Datapoint Type";
                }
                case KnxDatapointMainType.DPT_8_BIT_UNSIGNED_VALUE: { /* '10' */
                    return "8-bit unsigned value";
                }
                case KnxDatapointMainType.DPT_8_BIT_SIGNED_VALUE: { /* '11' */
                    return "8-bit signed value";
                }
                case KnxDatapointMainType.DPT_2_BYTE_UNSIGNED_VALUE: { /* '12' */
                    return "2-byte unsigned value";
                }
                case KnxDatapointMainType.DPT_2_BYTE_SIGNED_VALUE: { /* '13' */
                    return "2-byte signed value";
                }
                case KnxDatapointMainType.DPT_2_BYTE_FLOAT_VALUE: { /* '14' */
                    return "2-byte float value";
                }
                case KnxDatapointMainType.DPT_TIME: { /* '15' */
                    return "time";
                }
                case KnxDatapointMainType.DPT_DATE: { /* '16' */
                    return "date";
                }
                case KnxDatapointMainType.DPT_4_BYTE_UNSIGNED_VALUE: { /* '17' */
                    return "4-byte unsigned value";
                }
                case KnxDatapointMainType.DPT_4_BYTE_SIGNED_VALUE: { /* '18' */
                    return "4-byte signed value";
                }
                case KnxDatapointMainType.DPT_4_BYTE_FLOAT_VALUE: { /* '19' */
                    return "4-byte float value";
                }
                case KnxDatapointMainType.DPT_8_BYTE_UNSIGNED_VALUE: { /* '2' */
                    return "Unknown Datapoint Type";
                }
                case KnxDatapointMainType.DPT_ENTRANCE_ACCESS: { /* '20' */
                    return "entrance access";
                }
                case KnxDatapointMainType.DPT_CHARACTER_STRING: { /* '21' */
                    return "character string";
                }
                case KnxDatapointMainType.DPT_SCENE_NUMBER: { /* '22' */
                    return "scene number";
                }
                case KnxDatapointMainType.DPT_SCENE_CONTROL: { /* '23' */
                    return "scene control";
                }
                case KnxDatapointMainType.DPT_DATE_TIME: { /* '24' */
                    return "Date Time";
                }
                case KnxDatapointMainType.DPT_1_BYTE: { /* '25' */
                    return "1-byte";
                }
                case KnxDatapointMainType.DPT_8_BIT_SET: { /* '26' */
                    return "8-bit set";
                }
                case KnxDatapointMainType.DPT_16_BIT_SET: { /* '27' */
                    return "16-bit set";
                }
                case KnxDatapointMainType.DPT_2_BIT_SET: { /* '28' */
                    return "2-bit set";
                }
                case KnxDatapointMainType.DPT_2_NIBBLE_SET: { /* '29' */
                    return "2-nibble set";
                }
                case KnxDatapointMainType.DPT_8_BYTE_SIGNED_VALUE: { /* '3' */
                    return "Unknown Datapoint Type";
                }
                case KnxDatapointMainType.DPT_8_BIT_SET_2: { /* '30' */
                    return "8-bit set";
                }
                case KnxDatapointMainType.DPT_32_BIT_SET: { /* '31' */
                    return "32-bit set";
                }
                case KnxDatapointMainType.DPT_ELECTRICAL_ENERGY: { /* '32' */
                    return "electrical energy";
                }
                case KnxDatapointMainType.DPT_24_TIMES_CHANNEL_ACTIVATION: { /* '33' */
                    return "24 times channel activation";
                }
                case KnxDatapointMainType.DPT_16_BIT_UNSIGNED_VALUE_AND_8_BIT_ENUM: { /* '34' */
                    return "16-bit unsigned value & 8-bit enum";
                }
                case KnxDatapointMainType.DPT_8_BIT_UNSIGNED_VALUE_AND_8_BIT_ENUM: { /* '35' */
                    return "8-bit unsigned value & 8-bit enum";
                }
                case KnxDatapointMainType.DPT_DATAPOINT_TYPE_VERSION: { /* '36' */
                    return "datapoint type version";
                }
                case KnxDatapointMainType.DPT_ALARM_INFO: { /* '37' */
                    return "alarm info";
                }
                case KnxDatapointMainType.DPT_3X_2_BYTE_FLOAT_VALUE: { /* '38' */
                    return "3x 2-byte float value";
                }
                case KnxDatapointMainType.DPT_SCALING_SPEED: { /* '39' */
                    return "scaling speed";
                }
                case KnxDatapointMainType.DPT_12_BYTE_SIGNED_VALUE: { /* '4' */
                    return "Unknown Datapoint Type";
                }
                case KnxDatapointMainType.DPT_4_1_1_BYTE_COMBINED_INFORMATION: { /* '40' */
                    return "4-1-1 byte combined information";
                }
                case KnxDatapointMainType.DPT_MBUS_ADDRESS: { /* '41' */
                    return "MBus address";
                }
                case KnxDatapointMainType.DPT_3_BYTE_COLOUR_RGB: { /* '42' */
                    return "3-byte colour RGB";
                }
                case KnxDatapointMainType.DPT_LANGUAGE_CODE_ISO_639_1: { /* '43' */
                    return "language code ISO 639-1";
                }
                case KnxDatapointMainType.DPT_SIGNED_VALUE_WITH_CLASSIFICATION_AND_VALIDITY: { /* '44' */
                    return "Signed value with classification and validity";
                }
                case KnxDatapointMainType.DPT_PRIORITISED_MODE_CONTROL: { /* '45' */
                    return "Prioritised Mode Control";
                }
                case KnxDatapointMainType.DPT_CONFIGURATION_DIAGNOSTICS_16_BIT: { /* '46' */
                    return "configuration/ diagnostics";
                }
                case KnxDatapointMainType.DPT_CONFIGURATION_DIAGNOSTICS_8_BIT: { /* '47' */
                    return "configuration/ diagnostics";
                }
                case KnxDatapointMainType.DPT_POSITIONS: { /* '48' */
                    return "positions";
                }
                case KnxDatapointMainType.DPT_STATUS_32_BIT: { /* '49' */
                    return "status";
                }
                case KnxDatapointMainType.DPT_8_BYTE_FLOAT_VALUE: { /* '5' */
                    return "Unknown Datapoint Type";
                }
                case KnxDatapointMainType.DPT_STATUS_48_BIT: { /* '50' */
                    return "status";
                }
                case KnxDatapointMainType.DPT_CONVERTER_STATUS: { /* '51' */
                    return "Converter Status";
                }
                case KnxDatapointMainType.DPT_CONVERTER_TEST_RESULT: { /* '52' */
                    return "Converter test result";
                }
                case KnxDatapointMainType.DPT_BATTERY_INFORMATION: { /* '53' */
                    return "Battery Information";
                }
                case KnxDatapointMainType.DPT_BRIGHTNESS_COLOUR_TEMPERATURE_TRANSITION: { /* '54' */
                    return "brightness colour temperature transition";
                }
                case KnxDatapointMainType.DPT_STATUS_24_BIT: { /* '55' */
                    return "status";
                }
                case KnxDatapointMainType.DPT_COLOUR_RGBW: { /* '56' */
                    return "Colour RGBW";
                }
                case KnxDatapointMainType.DPT_RELATIVE_CONTROL_RGBW: { /* '57' */
                    return "Relative Control RGBW";
                }
                case KnxDatapointMainType.DPT_RELATIVE_CONTROL_RGB: { /* '58' */
                    return "Relative Control RGB";
                }
                case KnxDatapointMainType.DPT_F32F32: { /* '59' */
                    return "F32F32";
                }
                case KnxDatapointMainType.DPT_1_BIT: { /* '6' */
                    return "1-bit";
                }
                case KnxDatapointMainType.DPT_F16F16F16F16: { /* '60' */
                    return "F16F16F16F16";
                }
                case KnxDatapointMainType.DPT_1_BIT_CONTROLLED: { /* '7' */
                    return "1-bit controlled";
                }
                case KnxDatapointMainType.DPT_3_BIT_CONTROLLED: { /* '8' */
                    return "3-bit controlled";
                }
                case KnxDatapointMainType.DPT_CHARACTER: { /* '9' */
                    return "character";
                }
                default: {
                    return null;
                }
            }
        }

        public static byte? SizeInBits(this KnxDatapointMainType value)
        {
            switch (value)
            {
                case KnxDatapointMainType.DPT_UNKNOWN: { /* '0' */
                    return 0;
                }
                case KnxDatapointMainType.DPT_64_BIT_SET: { /* '1' */
                    return 64;
                }
                case KnxDatapointMainType.DPT_8_BIT_UNSIGNED_VALUE: { /* '10' */
                    return 8;
                }
                case KnxDatapointMainType.DPT_8_BIT_SIGNED_VALUE: { /* '11' */
                    return 8;
                }
                case KnxDatapointMainType.DPT_2_BYTE_UNSIGNED_VALUE: { /* '12' */
                    return 16;
                }
                case KnxDatapointMainType.DPT_2_BYTE_SIGNED_VALUE: { /* '13' */
                    return 16;
                }
                case KnxDatapointMainType.DPT_2_BYTE_FLOAT_VALUE: { /* '14' */
                    return 16;
                }
                case KnxDatapointMainType.DPT_TIME: { /* '15' */
                    return 24;
                }
                case KnxDatapointMainType.DPT_DATE: { /* '16' */
                    return 24;
                }
                case KnxDatapointMainType.DPT_4_BYTE_UNSIGNED_VALUE: { /* '17' */
                    return 32;
                }
                case KnxDatapointMainType.DPT_4_BYTE_SIGNED_VALUE: { /* '18' */
                    return 32;
                }
                case KnxDatapointMainType.DPT_4_BYTE_FLOAT_VALUE: { /* '19' */
                    return 32;
                }
                case KnxDatapointMainType.DPT_8_BYTE_UNSIGNED_VALUE: { /* '2' */
                    return 64;
                }
                case KnxDatapointMainType.DPT_ENTRANCE_ACCESS: { /* '20' */
                    return 32;
                }
                case KnxDatapointMainType.DPT_CHARACTER_STRING: { /* '21' */
                    return 112;
                }
                case KnxDatapointMainType.DPT_SCENE_NUMBER: { /* '22' */
                    return 8;
                }
                case KnxDatapointMainType.DPT_SCENE_CONTROL: { /* '23' */
                    return 8;
                }
                case KnxDatapointMainType.DPT_DATE_TIME: { /* '24' */
                    return 64;
                }
                case KnxDatapointMainType.DPT_1_BYTE: { /* '25' */
                    return 8;
                }
                case KnxDatapointMainType.DPT_8_BIT_SET: { /* '26' */
                    return 8;
                }
                case KnxDatapointMainType.DPT_16_BIT_SET: { /* '27' */
                    return 16;
                }
                case KnxDatapointMainType.DPT_2_BIT_SET: { /* '28' */
                    return 2;
                }
                case KnxDatapointMainType.DPT_2_NIBBLE_SET: { /* '29' */
                    return 8;
                }
                case KnxDatapointMainType.DPT_8_BYTE_SIGNED_VALUE: { /* '3' */
                    return 64;
                }
                case KnxDatapointMainType.DPT_8_BIT_SET_2: { /* '30' */
                    return 8;
                }
                case KnxDatapointMainType.DPT_32_BIT_SET: { /* '31' */
                    return 32;
                }
                case KnxDatapointMainType.DPT_ELECTRICAL_ENERGY: { /* '32' */
                    return 64;
                }
                case KnxDatapointMainType.DPT_24_TIMES_CHANNEL_ACTIVATION: { /* '33' */
                    return 24;
                }
                case KnxDatapointMainType.DPT_16_BIT_UNSIGNED_VALUE_AND_8_BIT_ENUM: { /* '34' */
                    return 24;
                }
                case KnxDatapointMainType.DPT_8_BIT_UNSIGNED_VALUE_AND_8_BIT_ENUM: { /* '35' */
                    return 16;
                }
                case KnxDatapointMainType.DPT_DATAPOINT_TYPE_VERSION: { /* '36' */
                    return 16;
                }
                case KnxDatapointMainType.DPT_ALARM_INFO: { /* '37' */
                    return 48;
                }
                case KnxDatapointMainType.DPT_3X_2_BYTE_FLOAT_VALUE: { /* '38' */
                    return 48;
                }
                case KnxDatapointMainType.DPT_SCALING_SPEED: { /* '39' */
                    return 24;
                }
                case KnxDatapointMainType.DPT_12_BYTE_SIGNED_VALUE: { /* '4' */
                    return 96;
                }
                case KnxDatapointMainType.DPT_4_1_1_BYTE_COMBINED_INFORMATION: { /* '40' */
                    return 48;
                }
                case KnxDatapointMainType.DPT_MBUS_ADDRESS: { /* '41' */
                    return 64;
                }
                case KnxDatapointMainType.DPT_3_BYTE_COLOUR_RGB: { /* '42' */
                    return 24;
                }
                case KnxDatapointMainType.DPT_LANGUAGE_CODE_ISO_639_1: { /* '43' */
                    return 16;
                }
                case KnxDatapointMainType.DPT_SIGNED_VALUE_WITH_CLASSIFICATION_AND_VALIDITY: { /* '44' */
                    return 48;
                }
                case KnxDatapointMainType.DPT_PRIORITISED_MODE_CONTROL: { /* '45' */
                    return 8;
                }
                case KnxDatapointMainType.DPT_CONFIGURATION_DIAGNOSTICS_16_BIT: { /* '46' */
                    return 16;
                }
                case KnxDatapointMainType.DPT_CONFIGURATION_DIAGNOSTICS_8_BIT: { /* '47' */
                    return 8;
                }
                case KnxDatapointMainType.DPT_POSITIONS: { /* '48' */
                    return 24;
                }
                case KnxDatapointMainType.DPT_STATUS_32_BIT: { /* '49' */
                    return 32;
                }
                case KnxDatapointMainType.DPT_8_BYTE_FLOAT_VALUE: { /* '5' */
                    return 64;
                }
                case KnxDatapointMainType.DPT_STATUS_48_BIT: { /* '50' */
                    return 48;
                }
                case KnxDatapointMainType.DPT_CONVERTER_STATUS: { /* '51' */
                    return 16;
                }
                case KnxDatapointMainType.DPT_CONVERTER_TEST_RESULT: { /* '52' */
                    return 48;
                }
                case KnxDatapointMainType.DPT_BATTERY_INFORMATION: { /* '53' */
                    return 16;
                }
                case KnxDatapointMainType.DPT_BRIGHTNESS_COLOUR_TEMPERATURE_TRANSITION: { /* '54' */
                    return 48;
                }
                case KnxDatapointMainType.DPT_STATUS_24_BIT: { /* '55' */
                    return 24;
                }
                case KnxDatapointMainType.DPT_COLOUR_RGBW: { /* '56' */
                    return 48;
                }
                case KnxDatapointMainType.DPT_RELATIVE_CONTROL_RGBW: { /* '57' */
                    return 40;
                }
                case KnxDatapointMainType.DPT_RELATIVE_CONTROL_RGB: { /* '58' */
                    return 24;
                }
                case KnxDatapointMainType.DPT_F32F32: { /* '59' */
                    return 64;
                }
                case KnxDatapointMainType.DPT_1_BIT: { /* '6' */
                    return 1;
                }
                case KnxDatapointMainType.DPT_F16F16F16F16: { /* '60' */
                    return 64;
                }
                case KnxDatapointMainType.DPT_1_BIT_CONTROLLED: { /* '7' */
                    return 2;
                }
                case KnxDatapointMainType.DPT_3_BIT_CONTROLLED: { /* '8' */
                    return 4;
                }
                case KnxDatapointMainType.DPT_CHARACTER: { /* '9' */
                    return 8;
                }
                default: {
                    return 0;
                }
            }
        }
    }

}

