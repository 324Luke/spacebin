/*
 * Copyright 2020-2023 Luke Whritenour, Jack Dorland

 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at

 *     http://www.apache.org/licenses/LICENSE-2.0

 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package database

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/orca-group/spirit/internal/config"
)

// Connection holds the current connection to the database
var Connection *sqlx.DB

func migrate() error {
	_, err := Connection.Exec(`
CREATE TABLE IF NOT EXISTS documents (
	id varchar(255) PRIMARY KEY,
	content text NOT NULL,
	created_at timestamp with time zone DEFAULT now(),
	updated_at timestamp with time zone DEFAULT now()
)`)

	return err
}

// Init opens a connection to the database
func Init() (err error) {
	Connection, err = sqlx.Connect("postgres", config.Config.ConnectionURI)

	if err != nil {
		return err
	}

	if err := migrate(); err != nil {
		return err
	}

	return nil
}

func Close() error {
	return nil
}