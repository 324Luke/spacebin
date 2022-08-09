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

package document

import (
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/orca-group/spirit/internal/pkg/config"
)

// CreateRequest represents a valid body object for the create document request
type CreateRequest struct {
	Content   string
	Extension string
}

// Validate performs validation on the body
func (c CreateRequest) Validate() error {
	/*
	 * This regex matches the file extension for various languages.

	 * Languages including:
	 *	python, javascript, jsx, typescript, tsx, go, kotlin, cpp, sql, csharp,
	 *	c, scala, haskell, shell-session, bash, powershell, php, asm6502, julia,
	 *	objc, perl, crystal, json, yaml, toml, none, rust, ruby, java, markdown,
	 *	markup (HTML, XML, SVG, Atom, RSS, MathML, SSML), css.

	 * For any unsupported formats Plain Text should be used.
	 */
	regex := regexp.MustCompile("^python$|^javascript$|^jsx$|^typescript$|^tsx$|^go$|^kotlin$|^cpp$|^sql$|^csharp$|^c$|^scala$|^haskell$|^shell-session$|^bash$|^powershell$|^php$|^asm6502$|^julia$|^objc$|^perl$|^crystal$|^json$|^yaml$|^toml$|^none$|^rust$|^ruby$|^markup$|^markdown$|^css$|")

	return validation.ValidateStruct(&c,
		validation.Field(
			&c.Content,
			validation.Required,
			// Enforce length to follow what's set in the config
			validation.Length(2, config.Config.Documents.MaxDocumentLength),
		),
		// The purpose of this field is to support client's that perform
		// syntax highlighting and need to know what highlighter to use.
		validation.Field(
			&c.Extension,
			validation.Match(regex),
			validation.Required,
		),
	)
}
