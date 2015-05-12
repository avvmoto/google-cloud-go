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

import bq "google.golang.org/api/bigquery/v2"

// A Table is a reference to a BigQuery table.
type Table struct {
	projectID string
	datasetID string
	tableID   string

	CreateDisposition CreateDisposition
	WriteDisposition  WriteDisposition
}

// CreateDisposition specifies the circumstances under which destination table will be created.
// The default is CreateIfNeeded.
type CreateDisposition string

const (
	// The table will be created if it does not already exist.  Tables are created atomically on successful completion of a job.
	CreateIfNeeded CreateDisposition = "CREATE_IF_NEEDED"

	// The table must already exist and will not be automatically created.
	CreateNever CreateDisposition = "CREATE_NEVER"
)

// WriteDisposition specifies how existing data in a destination table is treated.
// The default is WriteAppend.
type WriteDisposition string

const (
	// Data will be appended to any existing data in the destination table.
	// Data is appended atomically on successful completion of a job.
	WriteAppend WriteDisposition = "WRITE_APPEND"

	// Existing data in the destination table will be overwritten.
	// Data is overwritten atomically on successful completion of a job.
	WriteTruncate WriteDisposition = "WRITE_TRUNCATE"

	// Writes will fail if the destination table already contains data.
	WriteEmpty WriteDisposition = "WRITE_EMPTY"
)

func (t *Table) implementsDestination() {
}

// OpenTable constructs a reference to a BigQuery table.
func (c *Client) OpenTable(datasetID, tableID string) *Table {
	return &Table{
		projectID: c.projectID,
		datasetID: datasetID,
		tableID:   tableID,
	}
}

func (t *Table) customizeLoadDst(conf *bq.JobConfigurationLoad) {
	conf.DestinationTable = &bq.TableReference{
		ProjectId: t.projectID,
		DatasetId: t.datasetID,
		TableId:   t.tableID,
	}
	conf.CreateDisposition = string(t.CreateDisposition)
	conf.WriteDisposition = string(t.WriteDisposition)
}
