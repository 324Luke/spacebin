/*
 * Copyright 2020 Luke Whrit, Jack Dorland; The Spacebin Authors

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

package structs

// Config is the configuration object
type Config struct {
	Server struct {
		Host              string `koanf:"host"`
		Port              int    `koanf:"port"`
		UseCSP            bool   `koanf:"useCSP"`
		CompresssionLevel int    `koanf:"compressionLevel"`
		Prefork           bool   `koanf:"prefork"`

		Ratelimits struct {
			Requests int `koanf:"requests"`
			Duration int `koanf:"duration"`
		} `koanf:"ratelimits"`
	}

	Documents struct {
		IDLength          int `koanf:"idLength"`
		MaxDocumentLength int `koanf:"maxDocunentLength"`
	} `koanf:"documents"`

	Database struct {
		Dialect       string `koanf:"dialect"`
		ConnectionURI string `koanf:"connection_uri"`
	} `koanf:"database"`
}