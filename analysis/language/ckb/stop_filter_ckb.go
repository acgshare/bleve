//  Copyright (c) 2014 Couchbase, Inc.
//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
//  except in compliance with the License. You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
//  Unless required by applicable law or agreed to in writing, software distributed under the
//  License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
//  either express or implied. See the License for the specific language governing permissions
//  and limitations under the License.

package ckb

import (
	"github.com/acgshare/bleve/analysis"
	"github.com/acgshare/bleve/analysis/token_filters/stop_tokens_filter"
	"github.com/acgshare/bleve/registry"
)

func StopTokenFilterConstructor(config map[string]interface{}, cache *registry.Cache) (analysis.TokenFilter, error) {
	tokenMap, err := cache.TokenMapNamed(StopName)
	if err != nil {
		return nil, err
	}
	return stop_tokens_filter.NewStopTokensFilter(tokenMap), nil
}

func init() {
	registry.RegisterTokenFilter(StopName, StopTokenFilterConstructor)
}
