// Copyright 2015 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package bigquery

import "golang.org/x/net/context"

// OpenTable creates a handle to an existing BigQuery table. If the table does
// not already exist, subsequent uses of the *Table will fail.
//
// Deprecated: use Dataset.Table instead.
func (c *Client) OpenTable(projectID, datasetID, tableID string) *Table {
	return c.Table(projectID, datasetID, tableID)
}

// CreateTable creates a table in the BigQuery service and returns a handle to it.
//
// Deprecated: use Table.Create instead.
func (c *Client) CreateTable(ctx context.Context, projectID, datasetID, tableID string, options ...CreateTableOption) (*Table, error) {
	t := c.Table(projectID, datasetID, tableID)
	if err := t.Create(ctx, options...); err != nil {
		return nil, err
	}
	return t, nil
}