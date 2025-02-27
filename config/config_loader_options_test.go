/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package config

import (
	"testing"
)

import (
	"github.com/stretchr/testify/assert"
)

func TestCheckGenre(t *testing.T) {

	err := checkGenre("abc")
	assert.NotNil(t, err)

	err = checkGenre("json")
	assert.Nil(t, err)
}

func TestFileGenre(t *testing.T) {
	conf := NewLoaderConf(WithPath("../config/testdata/config/properties/application.properties"))
	assert.Equal(t, conf.genre, "properties")
}

func TestRootConfig(t *testing.T) {
	rc := NewRootConfigBuilder().SetApplication(NewApplicationConfigBuilder().SetName("test-app").Build()).Build()
	conf := NewLoaderConf(WithRootConfig(rc))
	assert.Equal(t, conf.rc.Application.Name, "test-app")
}
