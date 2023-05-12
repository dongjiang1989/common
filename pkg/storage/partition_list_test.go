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

package storage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type fakePartition struct {
	minT      int64
	maxT      int64
	numPoints int
	IsActive  bool

	err error
}

func (f *fakePartition) insertRows(_ []Row) ([]Row, error) {
	return nil, f.err
}

func (f *fakePartition) selectDataPoints(_ string, _ []Label, _, _ int64) ([]*DataPoint, error) {
	return nil, f.err
}

func (f *fakePartition) minTimestamp() int64 {
	return f.minT
}

func (f *fakePartition) maxTimestamp() int64 {
	return f.maxT
}

func (f *fakePartition) size() int {
	return f.numPoints
}

func (f *fakePartition) active() bool {
	return f.IsActive
}

func (f *fakePartition) clean() error {
	return nil
}

func (f *fakePartition) expired() bool {
	return false
}

func Test_partitionList_Remove(t *testing.T) {
	tests := []struct {
		name              string
		partitionList     partitionListImpl
		target            partition
		wantErr           bool
		wantPartitionList partitionListImpl
	}{
		{
			name:          "empty partition",
			partitionList: partitionListImpl{},
			wantErr:       true,
		},
		{
			name: "remove the head node",
			partitionList: func() partitionListImpl {
				second := &partitionNode{
					val: &fakePartition{
						minT: 2,
					},
				}

				first := &partitionNode{
					val: &fakePartition{
						minT: 1,
					},
					next: second,
				}
				return partitionListImpl{
					numPartitions: 2,
					head:          first,
					tail:          second,
				}
			}(),
			target: &fakePartition{
				minT: 1,
			},
			wantPartitionList: partitionListImpl{
				numPartitions: 1,
				head: &partitionNode{
					val: &fakePartition{
						minT: 2,
					},
				},
				tail: &partitionNode{
					val: &fakePartition{
						minT: 2,
					},
				},
			},
		},
		{
			name: "remove the tail node",
			partitionList: func() partitionListImpl {
				second := &partitionNode{
					val: &fakePartition{
						minT: 2,
					},
				}

				first := &partitionNode{
					val: &fakePartition{
						minT: 1,
					},
					next: second,
				}
				return partitionListImpl{
					numPartitions: 2,
					head:          first,
					tail:          second,
				}
			}(),
			target: &fakePartition{
				minT: 2,
			},
			wantPartitionList: partitionListImpl{
				numPartitions: 1,
				head: &partitionNode{
					val: &fakePartition{
						minT: 1,
					},
				},
				tail: &partitionNode{
					val: &fakePartition{
						minT: 1,
					},
				},
			},
		},
		{
			name: "remove the middle node",
			partitionList: func() partitionListImpl {
				third := &partitionNode{
					val: &fakePartition{
						minT: 3,
					},
				}
				second := &partitionNode{
					val: &fakePartition{
						minT: 2,
					},
					next: third,
				}
				first := &partitionNode{
					val: &fakePartition{
						minT: 1,
					},
					next: second,
				}
				return partitionListImpl{
					numPartitions: 3,
					head:          first,
					tail:          third,
				}
			}(),
			target: &fakePartition{
				minT: 2,
			},
			wantPartitionList: partitionListImpl{
				numPartitions: 2,
				head: &partitionNode{
					val: &fakePartition{
						minT: 1,
					},
					next: &partitionNode{
						val: &fakePartition{
							minT: 3,
						},
					},
				},
				tail: &partitionNode{
					val: &fakePartition{
						minT: 3,
					},
				},
			},
		},
		{
			name: "given node not found",
			partitionList: func() partitionListImpl {
				second := &partitionNode{
					val: &fakePartition{
						minT: 2,
					},
				}

				first := &partitionNode{
					val: &fakePartition{
						minT: 1,
					},
					next: second,
				}
				return partitionListImpl{
					numPartitions: 2,
					head:          first,
					tail:          second,
				}
			}(),
			target: &fakePartition{
				minT: 3,
			},
			wantPartitionList: func() partitionListImpl {
				second := &partitionNode{
					val: &fakePartition{
						minT: 2,
					},
				}

				first := &partitionNode{
					val: &fakePartition{
						minT: 1,
					},
					next: second,
				}
				return partitionListImpl{
					numPartitions: 2,
					head:          first,
					tail:          second,
				}
			}(),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.partitionList.remove(tt.target)
			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, tt.wantPartitionList, tt.partitionList)
			tt.partitionList.String()
		})
	}
}

func Test_partitionList_Swap(t *testing.T) {
	tests := []struct {
		name              string
		partitionList     partitionListImpl
		old               partition
		new               partition
		wantErr           bool
		wantPartitionList partitionListImpl
	}{
		{
			name:          "empty partition",
			partitionList: partitionListImpl{},
			wantErr:       true,
		},
		{
			name: "swap the head node",
			partitionList: func() partitionListImpl {
				second := &partitionNode{
					val: &fakePartition{
						minT: 2,
					},
				}

				first := &partitionNode{
					val: &fakePartition{
						minT: 1,
					},
					next: second,
				}
				return partitionListImpl{
					numPartitions: 2,
					head:          first,
					tail:          second,
				}
			}(),
			old: &fakePartition{
				minT: 1,
			},
			new: &fakePartition{
				minT: 100,
			},
			wantPartitionList: partitionListImpl{
				numPartitions: 2,
				head: &partitionNode{
					val: &fakePartition{
						minT: 100,
					},
					next: &partitionNode{
						val: &fakePartition{
							minT: 2,
						},
					},
				},
				tail: &partitionNode{
					val: &fakePartition{
						minT: 2,
					},
				},
			},
		},
		{
			name: "swap the tail node",
			partitionList: func() partitionListImpl {
				second := &partitionNode{
					val: &fakePartition{
						minT: 2,
					},
				}

				first := &partitionNode{
					val: &fakePartition{
						minT: 1,
					},
					next: second,
				}
				return partitionListImpl{
					numPartitions: 2,
					head:          first,
					tail:          second,
				}
			}(),
			old: &fakePartition{
				minT: 2,
			},
			new: &fakePartition{
				minT: 100,
			},
			wantPartitionList: partitionListImpl{
				numPartitions: 2,
				head: &partitionNode{
					val: &fakePartition{
						minT: 1,
					},
					next: &partitionNode{
						val: &fakePartition{
							minT: 100,
						},
					},
				},
				tail: &partitionNode{
					val: &fakePartition{
						minT: 100,
					},
				},
			},
		},
		{
			name: "swap the middle node",
			partitionList: func() partitionListImpl {
				third := &partitionNode{
					val: &fakePartition{
						minT: 3,
					},
				}
				second := &partitionNode{
					val: &fakePartition{
						minT: 2,
					},
					next: third,
				}

				first := &partitionNode{
					val: &fakePartition{
						minT: 1,
					},
					next: second,
				}
				return partitionListImpl{
					numPartitions: 3,
					head:          first,
					tail:          third,
				}
			}(),
			old: &fakePartition{
				minT: 2,
			},
			new: &fakePartition{
				minT: 100,
			},
			wantPartitionList: partitionListImpl{
				numPartitions: 3,
				head: &partitionNode{
					val: &fakePartition{
						minT: 1,
					},
					next: &partitionNode{
						val: &fakePartition{
							minT: 100,
						},
						next: &partitionNode{
							val: &fakePartition{
								minT: 3,
							},
						},
					},
				},
				tail: &partitionNode{
					val: &fakePartition{
						minT: 3,
					},
				},
			},
		},
		{
			name: "given node not found",
			partitionList: func() partitionListImpl {
				second := &partitionNode{
					val: &fakePartition{
						minT: 2,
					},
				}

				first := &partitionNode{
					val: &fakePartition{
						minT: 1,
					},
					next: second,
				}
				return partitionListImpl{
					numPartitions: 2,
					head:          first,
					tail:          second,
				}
			}(),
			old: &fakePartition{
				minT: 100,
			},
			wantPartitionList: partitionListImpl{
				numPartitions: 2,
				head: &partitionNode{
					val: &fakePartition{
						minT: 1,
					},
					next: &partitionNode{
						val: &fakePartition{
							minT: 2,
						},
					},
				},
				tail: &partitionNode{
					val: &fakePartition{
						minT: 2,
					},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.partitionList.swap(tt.old, tt.new)
			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, tt.wantPartitionList, tt.partitionList)
			tt.partitionList.String()
		})
	}
}
