// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package elasticsearch

import (
	"github.com/elastic/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "elasticsearch", asset.ModuleFieldsPri, AssetElasticsearch); err != nil {
		panic(err)
	}
}

// AssetElasticsearch returns asset data.
// This is the base64 encoded zlib format compressed contents of module/elasticsearch.
func AssetElasticsearch() string {
	return "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"
}
