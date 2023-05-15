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

package schedule

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func CallBackTest(ss interface{}) interface{} {
	return ss
}

func Test_NewTask(t *testing.T) {
	assert := assert.New(t)
	task := NewTask(1)
	now := time.Now()
	err := task.At(now.Format("15:04:05")).Do(CallBackTest, "aaa")
	assert.Nil(err)
	err = task.At(now.Format("15:04:05")).Do(CallBackTest)
	assert.Nil(err)

	ddd := task.GetAt()
	assert.Equal(ddd, now.Format("15:04"))

	eee := task.GetWeekday()
	assert.Equal(eee, time.Weekday(0))

	err = task.Days().Day().Friday().Weeks().Weekday(time.Monday).Wednesday().Tuesday().Thursday().
		Sunday().Saturday().Monday().At(time.Now().Add(1*time.Second).Format("15:04:05")).Do(CallBackTest, "aaa")
	assert.Nil(err)

	var functionNot interface{}
	functionNot = "adffdfa"
	err = task.Hour().Minutes().Minute().Hours().Second().Seconds().Do(functionNot, "df")
	assert.Equal(err, ErrNotAFunction)

	err = task.Hour().Minutes().Minute().Hours().Second().Seconds().DoSafely(functionNot, "df")
	assert.Nil(err)
}
