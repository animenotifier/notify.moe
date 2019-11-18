package autodocs

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Type represents a type in a Go source file.
type Type struct {
	Name       string
	Comment    string
	LineNumber int
}

// Endpoint returns the REST endpoint for that type.
func (typ *Type) Endpoint() string {
	return "/api/" + strings.ToLower(typ.Name)
}

// GitHubLink returns link to display the type in GitHub.
func (typ *Type) GitHubLink() string {
	return fmt.Sprintf("https://github.com/animenotifier/notify.moe/blob/go/arn/%s.go#L%d", typ.Name, typ.LineNumber)
}

// GetTypeDocumentation tries to gather documentation about the given type.
func GetTypeDocumentation(typeName string, filePath string) (*Type, error) {
	typ := &Type{
		Name: typeName,
	}

	file, err := os.Open(filePath)

	if err != nil {
		return typ, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNumber := 0

	var comments []string

	for scanner.Scan() {
		lineNumber++

		line := scanner.Text()
		line = strings.TrimSpace(line)
		isComment := strings.HasPrefix(line, "// ")

		if isComment {
			comment := strings.TrimPrefix(line, "// ")
			comments = append(comments, comment)
			continue
		}

		if strings.HasPrefix(line, "type ") {
			definedTypeName := strings.TrimPrefix(line, "type ")
			space := strings.Index(definedTypeName, " ")
			definedTypeName = definedTypeName[:space]

			if definedTypeName == typeName {
				typ.Comment = strings.Join(comments, " ")
				typ.LineNumber = lineNumber
			}
		}

		if !isComment {
			comments = nil
		}
	}

	if err := scanner.Err(); err != nil {
		return typ, err
	}

	return typ, nil
}
