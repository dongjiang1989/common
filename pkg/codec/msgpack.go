/*
Copyright 2023 The KubeService-Stack Authors.

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

package codec

import (
	"github.com/vmihailenco/msgpack/v5"
)

type MSGPack struct{}

func NewMSGPack() Codec {
	return &MSGPack{}
}

func (mc *MSGPack) Marshal(v interface{}) ([]byte, error) {
	return msgpack.Marshal(v)
}

func (mc *MSGPack) Unmarshal(data []byte, v interface{}) error {
	return msgpack.Unmarshal(data, v)
}

func init() {
	Register(MSGPACK, NewMSGPack)
}
