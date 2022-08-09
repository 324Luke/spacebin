/*
 * Copyright 2020-2022 Luke Whrit, Jack Dorland

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

package main

import (
	"fmt"
	"log"

	"github.com/coral-dev/spirit/internal/app"
	"github.com/coral-dev/spirit/internal/pkg/config"
	"github.com/coral-dev/spirit/internal/pkg/database"
	"github.com/coral-dev/spirit/internal/pkg/document"
)

func init() {
	// Load config
	if err := config.Load(); err != nil {
		log.Fatalf("Couldn't load configuration file: %v", err)
	}

	// Start server and initialize database
	database.Init()

	// Start expire document cron job
	document.ExpireDocument().Start()
}

func main() {
	app := app.Start()
	address := fmt.Sprintf("%s:%d", config.Config.Server.Host, config.Config.Server.Port)

	log.Fatal(app.Listen(address))
}
