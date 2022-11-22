/*
Copyright 2022 The KubeService-Stack Authors.

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

package item

import (
	"time"
)

type LruItem struct {
	Key        interface{}
	Value      interface{}
	Expiration *time.Time
}

// returns boolean value whether this item is expired or not.
func (it *LruItem) IsExpired(now *time.Time) bool {
	if it.Expiration == nil {
		return false
	}
	if now == nil {
		t := time.Now()
		now = &t
	}
	return it.Expiration.Before(*now)
}

func (it *LruItem) Expire() *time.Time {
	return it.Expiration
}
