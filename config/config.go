//  Copyright (c) 2015 Couchbase, Inc.
//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
//  except in compliance with the License. You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
//  Unless required by applicable law or agreed to in writing, software distributed under the
//  License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
//  either express or implied. See the License for the specific language governing permissions
//  and limitations under the License.

package config

import (
	// token maps
	_ "github.com/acgshare/bleve/analysis/token_map"

	// fragment formatters
	_ "github.com/acgshare/bleve/search/highlight/fragment_formatters/ansi"
	_ "github.com/acgshare/bleve/search/highlight/fragment_formatters/html"

	// fragmenters
	_ "github.com/acgshare/bleve/search/highlight/fragmenters/simple"

	// highlighters
	_ "github.com/acgshare/bleve/search/highlight/highlighters/ansi"
	_ "github.com/acgshare/bleve/search/highlight/highlighters/html"
	_ "github.com/acgshare/bleve/search/highlight/highlighters/simple"

	// char filters
	_ "github.com/acgshare/bleve/analysis/char_filters/html_char_filter"
	_ "github.com/acgshare/bleve/analysis/char_filters/regexp_char_filter"
	_ "github.com/acgshare/bleve/analysis/char_filters/zero_width_non_joiner"

	// analyzers
	_ "github.com/acgshare/bleve/analysis/analyzers/custom_analyzer"
	_ "github.com/acgshare/bleve/analysis/analyzers/keyword_analyzer"
	_ "github.com/acgshare/bleve/analysis/analyzers/simple_analyzer"
	_ "github.com/acgshare/bleve/analysis/analyzers/standard_analyzer"
	_ "github.com/acgshare/bleve/analysis/analyzers/web"

	// token filters
	_ "github.com/acgshare/bleve/analysis/token_filters/apostrophe_filter"
	_ "github.com/acgshare/bleve/analysis/token_filters/compound"
	_ "github.com/acgshare/bleve/analysis/token_filters/edge_ngram_filter"
	_ "github.com/acgshare/bleve/analysis/token_filters/elision_filter"
	_ "github.com/acgshare/bleve/analysis/token_filters/keyword_marker_filter"
	_ "github.com/acgshare/bleve/analysis/token_filters/length_filter"
	_ "github.com/acgshare/bleve/analysis/token_filters/lower_case_filter"
	_ "github.com/acgshare/bleve/analysis/token_filters/ngram_filter"
	_ "github.com/acgshare/bleve/analysis/token_filters/shingle"
	_ "github.com/acgshare/bleve/analysis/token_filters/stop_tokens_filter"
	_ "github.com/acgshare/bleve/analysis/token_filters/truncate_token_filter"
	_ "github.com/acgshare/bleve/analysis/token_filters/unicode_normalize"

	// tokenizers
	_ "github.com/acgshare/bleve/analysis/tokenizers/exception"
	_ "github.com/acgshare/bleve/analysis/tokenizers/regexp_tokenizer"
	_ "github.com/acgshare/bleve/analysis/tokenizers/single_token"
	_ "github.com/acgshare/bleve/analysis/tokenizers/unicode"
	_ "github.com/acgshare/bleve/analysis/tokenizers/web"
	_ "github.com/acgshare/bleve/analysis/tokenizers/whitespace_tokenizer"

	// date time parsers
	_ "github.com/acgshare/bleve/analysis/datetime_parsers/datetime_optional"
	_ "github.com/acgshare/bleve/analysis/datetime_parsers/flexible_go"

	// languages
	_ "github.com/acgshare/bleve/analysis/language/ar"
	_ "github.com/acgshare/bleve/analysis/language/bg"
	_ "github.com/acgshare/bleve/analysis/language/ca"
	_ "github.com/acgshare/bleve/analysis/language/cjk"
	_ "github.com/acgshare/bleve/analysis/language/ckb"
	_ "github.com/acgshare/bleve/analysis/language/cs"
	_ "github.com/acgshare/bleve/analysis/language/el"
	_ "github.com/acgshare/bleve/analysis/language/en"
	_ "github.com/acgshare/bleve/analysis/language/eu"
	_ "github.com/acgshare/bleve/analysis/language/fa"
	_ "github.com/acgshare/bleve/analysis/language/fr"
	_ "github.com/acgshare/bleve/analysis/language/ga"
	_ "github.com/acgshare/bleve/analysis/language/gl"
	_ "github.com/acgshare/bleve/analysis/language/hi"
	_ "github.com/acgshare/bleve/analysis/language/hy"
	_ "github.com/acgshare/bleve/analysis/language/id"
	_ "github.com/acgshare/bleve/analysis/language/in"
	_ "github.com/acgshare/bleve/analysis/language/it"
	_ "github.com/acgshare/bleve/analysis/language/pt"

	// kv stores
	_ "github.com/acgshare/bleve/index/store/boltdb"
	_ "github.com/acgshare/bleve/index/store/goleveldb"
	_ "github.com/acgshare/bleve/index/store/gtreap"

	// index types
	_ "github.com/acgshare/bleve/index/firestorm"
	_ "github.com/acgshare/bleve/index/upside_down"

	// byte array converters
	_ "github.com/acgshare/bleve/analysis/byte_array_converters/ignore"
	_ "github.com/acgshare/bleve/analysis/byte_array_converters/json"
	_ "github.com/acgshare/bleve/analysis/byte_array_converters/string"
)
