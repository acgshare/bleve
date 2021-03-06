//  Copyright (c) 2015 Couchbase, Inc.
//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
//  except in compliance with the License. You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
//  Unless required by applicable law or agreed to in writing, software distributed under the
//  License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
//  either express or implied. See the License for the specific language governing permissions
//  and limitations under the License.

package firestorm

import (
	"encoding/json"
	"sync/atomic"
)

type indexStat struct {
	f                                 *Firestorm
	updates, deletes, batches, errors uint64
	analysisTime, indexTime           uint64
}

func (i *indexStat) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{}
	m["updates"] = atomic.LoadUint64(&i.updates)
	m["deletes"] = atomic.LoadUint64(&i.deletes)
	m["batches"] = atomic.LoadUint64(&i.batches)
	m["errors"] = atomic.LoadUint64(&i.errors)
	m["analysis_time"] = atomic.LoadUint64(&i.analysisTime)
	m["index_time"] = atomic.LoadUint64(&i.indexTime)
	m["lookup_queue_len"] = len(i.f.lookuper.workChan)
	return json.Marshal(m)
}
