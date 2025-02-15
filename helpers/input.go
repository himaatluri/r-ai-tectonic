package helpers

import (
	"regexp"
	"strings"
)

func generateDocTypeRegex(docTypes []string) *regexp.Regexp {
	pattern := `(?i)\b(?:` + strings.Join(docTypes, "|") + `)\b`
	return regexp.MustCompile(pattern)
}

func hasDocumentInfo(input string, docTypes []string) (string, bool) {
	docTypeRegex := generateDocTypeRegex(docTypes)
	docPathPattern := `(?:[a-zA-Z]:\\|\\|/)?(?:[\w\s\-\_\.]+[\\/])*[\w\s\-\_]+\.(?:` + strings.Join(docTypes, "|") + `)`
	docPathRegex := regexp.MustCompile(docPathPattern)
	docPathMatch := docPathRegex.FindString(input)
	if docPathMatch != "" {
		return docPathMatch, true
	}
	docTypeMatches := docTypeRegex.FindAllStringIndex(input, -1)
	for _, match := range docTypeMatches {
		if match[0] > 0 && (input[match[0]-1] == '\\' || input[match[0]-1] == '/' || input[match[0]-1] == '.') {
			return input[match[0]:match[1]], true
		}
	}

	return "", false
}

func DetectDocInference(inputPrompt string) (string, bool) {
	docTypes := []string{"csv", "pdf", "docx?", "xlsx?", "pptx?", "txt", "rtf", "html?", "md"}
	return hasDocumentInfo(inputPrompt, docTypes)
}
