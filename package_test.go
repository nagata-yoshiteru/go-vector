/*
Copyright 2017 Albert Tedja

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package vector

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	assert := assert.New(t)
	v := New(4)
	v[0] = 2.9
	v[1] = 3.1
	v[2] = 3.2
	assert.Equal(v[0], 2.9)
	assert.Equal(v[1], 3.1)
	assert.Equal(v[2], 3.2)
	assert.Equal(v[3], 0.0)
}

func TestNewWithValues(t *testing.T) {
	assert := assert.New(t)
	v := NewWithValues([]float64{0.0, 1.0, 2.0})
	assert.Equal(v[0], 0.0)
	assert.Equal(v[1], 1.0)
	assert.Equal(v[2], 2.0)
}

func TestAdd(t *testing.T) {
	assert := assert.New(t)
	v1 := NewWithValues([]float64{0.0, 1.0, 2.0, 1.0})
	v2 := NewWithValues([]float64{0.0, 1.0, 2.0, 1.0})
	result := Add(v1, v2)
	assert.Equal(result[0], 0.0)
	assert.Equal(result[1], 2.0)
	assert.Equal(result[2], 4.0)
	assert.Equal(result[3], 2.0)

	assert.Equal(v1[1], 1.0)
	assert.Equal(v2[1], 1.0)
}

func TestSubtract(t *testing.T) {
	assert := assert.New(t)
	v1 := NewWithValues([]float64{0.0, 1.0, 2.0, 1.0})
	v2 := NewWithValues([]float64{0.0, 1.0, 2.0, 1.0})
	result := Subtract(v1, v2)
	assert.Equal(result[0], 0.0)
	assert.Equal(result[1], 0.0)
	assert.Equal(result[2], 0.0)
	assert.Equal(result[3], 0.0)

	assert.Equal(v1[1], 1.0)
	assert.Equal(v2[1], 1.0)
}

func TestDot(t *testing.T) {
	assert := assert.New(t)
	v1 := NewWithValues([]float64{0.0, 1.0, 2.0, 1.0})
	v2 := NewWithValues([]float64{0.0, 1.0, 2.0, 1.0})
	result, err := Dot(v1, v2)
	assert.Nil(err)
	assert.Equal(6.0, result)
}

func TestCross(t *testing.T) {
	assert := assert.New(t)
	v1 := NewWithValues([]float64{0.0, 1.0, 2.0})
	v2 := NewWithValues([]float64{0.0, 3.0, 4.0})
	result, err := Cross(v1, v2)
	assert.Nil(err)
	assert.Equal(-2.0, result[0])
	assert.Equal(0.0, result[1])
	assert.Equal(0.0, result[2])

	assert.Equal(1.0, v1[1])
	assert.Equal(3.0, v2[1])

	v3 := NewWithValues([]float64{0.0, 1.0, 2.0, 1.0})
	v4 := NewWithValues([]float64{0.0, 1.0, 2.0})
	_, err = Cross(v3, v4)
	assert.NotNil(err)
}

func TestUnit(t *testing.T) {
	assert := assert.New(t)
	v := NewWithValues([]float64{3.0, 4.0})
	unit := Unit(v)
	assert.InEpsilon(0.6, unit[0], EPSILON)
	assert.InEpsilon(0.8, unit[1], EPSILON)
	assert.Equal(1.0, unit.Magnitude())
}

func TestHadamard(t *testing.T) {
	assert := assert.New(t)
	v1 := NewWithValues([]float64{1.0, 1.0, 2.0})
	v2 := NewWithValues([]float64{0.5, 3.0, 4.0})
	result, err := Hadamard(v1, v2)
	assert.Nil(err)
	assert.Equal(0.5, result[0])
	assert.Equal(3.0, result[1])
	assert.Equal(8.0, result[2])
}
