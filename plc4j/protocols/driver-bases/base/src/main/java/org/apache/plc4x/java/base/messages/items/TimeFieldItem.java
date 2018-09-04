/*
Licensed to the Apache Software Foundation (ASF) under one
or more contributor license agreements.  See the NOTICE file
distributed with this work for additional information
regarding copyright ownership.  The ASF licenses this file
to you under the Apache License, Version 2.0 (the
"License"); you may not use this file except in compliance
with the License.  You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing,
software distributed under the License is distributed on an
"AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
KIND, either express or implied.  See the License for the
specific language governing permissions and limitations
under the License.
*/
package org.apache.plc4x.java.base.messages.items;

import java.time.LocalDate;
import java.time.LocalDateTime;
import java.time.LocalTime;

public class TimeFieldItem extends FieldItem<LocalDateTime> {

    public TimeFieldItem(LocalDateTime[] values) {
        super(values);
    }

    @Override
    public boolean isValidTime(int index) {
        return (getValue(index) != null);
    }

    @Override
    public LocalTime getTime(int index) {
        if(isValidTime(index)) {
            return getValue(index).toLocalTime();
        }
        return null;
    }

    @Override
    public boolean isValidDate(int index) {
        return (getValue(index) != null);
    }

    @Override
    public LocalDate getDate(int index) {
        if(isValidTime(index)) {
            return getValue(index).toLocalDate();
        }
        return null;
    }

    @Override
    public boolean isValidDateTime(int index) {
        return (getValue(index) != null);
    }

    @Override
    public LocalDateTime getDateTime(int index) {
        if(isValidTime(index)) {
            return getValue(index);
        }
        return null;
    }

}

