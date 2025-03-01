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
package org.apache.plc4x.plugins.codegenerator.language.mspec.model.fields;

import org.apache.plc4x.plugins.codegenerator.types.fields.EnumField;
import org.apache.plc4x.plugins.codegenerator.types.references.EnumTypeReference;
import org.apache.plc4x.plugins.codegenerator.types.terms.Term;

import java.util.Map;
import java.util.Objects;

public class DefaultEnumField extends DefaultTypedNamedField implements EnumField {

    private final String fieldName;

    public DefaultEnumField(Map<String, Term> attributes, EnumTypeReference type, String name, String fieldName) {
        super(attributes, name);
        this.fieldName = Objects.requireNonNull(fieldName);
        this.type = type;
    }

    public String getFieldName() {
        return fieldName;
    }

    @Override
    public String toString() {
        return "DefaultEnumField{" +
            "fieldName='" + fieldName + '\'' +
            "} " + super.toString();
    }

    @Override
    public boolean equals(Object o) {
        if (this == o) return true;
        if (o == null || getClass() != o.getClass()) return false;
        if (!super.equals(o)) return false;
        DefaultEnumField that = (DefaultEnumField) o;
        return Objects.equals(fieldName, that.fieldName);
    }

    @Override
    public int hashCode() {
        return Objects.hash(super.hashCode(), fieldName);
    }
}
