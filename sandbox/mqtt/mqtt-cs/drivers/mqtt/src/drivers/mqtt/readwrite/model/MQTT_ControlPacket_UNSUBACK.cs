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

using System;
using System.Collections.Generic;
using System.Text;
using org.apache.plc4net.api.value;
using org.apache.plc4net.spi.generation;
using org.apache.plc4net.spi.model.values;

// Code generated by code-generation. DO NOT EDIT.

namespace org.apache.plc4net.drivers.mqtt.readwrite.model
{

    public class MQTT_ControlPacket_UNSUBACK : MQTT_ControlPacket
    {

        // Accessors for discriminator values.
        public override MQTT_ControlPacketType PacketType => MQTT_ControlPacketType.UNSUBACK;

        // Properties.
        public byte RemainingLength { get; }
        public ushort PacketIdentifier { get; }
        public uint PropertyLength { get; }
        public MQTT_Property[] Properties { get; }
        public MQTT_ReasonCode[] Results { get; }

        public MQTT_ControlPacket_UNSUBACK(byte remainingLength, ushort packetIdentifier, uint propertyLength, MQTT_Property[] properties, MQTT_ReasonCode[] results)
        {
            RemainingLength = remainingLength;
            PacketIdentifier = packetIdentifier;
            PropertyLength = propertyLength;
            Properties = properties;
            Results = results;
        }

    }
}
