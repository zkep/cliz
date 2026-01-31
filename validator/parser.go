package validator

import (
	"regexp"
	"strconv"
	"strings"
)

// parseValidateTags parses validate tags and creates corresponding validators
func parseValidateTags(validateTags, fieldName string) []Validator {
	var validators []Validator
	var errorMap = make(map[string]string)

	// First pass: process all error tags to populate errorMap
	tags := strings.Split(validateTags, ",")
	for _, tag := range tags {
		tag = strings.TrimSpace(tag)
		if tag == "" {
			continue
		}

		parts := strings.SplitN(tag, "=", 2)
		tagName := strings.TrimSpace(parts[0])
		var tagValue string
		if len(parts) > 1 {
			tagValue = strings.TrimSpace(parts[1])
			tagValue = strings.Trim(tagValue, `"'`)
		}

		if strings.HasPrefix(tagName, "error_") {
			// Extract validator name from error tag
			validatorName := tagName[len("error_"):]
			errorMap[validatorName] = tagValue
		}
	}

	// Second pass: process validator tags
	for _, tag := range tags {
		tag = strings.TrimSpace(tag)
		if tag == "" {
			continue
		}

		parts := strings.SplitN(tag, "=", 2)
		tagName := strings.TrimSpace(parts[0])
		var tagValue string
		if len(parts) > 1 {
			tagValue = strings.TrimSpace(parts[1])
			tagValue = strings.Trim(tagValue, `"'`)
		}

		if strings.HasPrefix(tagName, "error_") {
			continue
		}

		switch tagName {
		case "required":
			var errMsg string
			if val, ok := errorMap["required"]; ok {
				errMsg = val
				delete(errorMap, "required")
			}
			validators = append(validators, &RequiredValidator{FieldName: fieldName, ErrorMessage: errMsg})
		case "range":
			if tagValue != "" {
				parts := strings.Split(tagValue, "-")
				if len(parts) == 2 {
					min, err1 := strconv.ParseFloat(strings.TrimSpace(parts[0]), 64)
					max, err2 := strconv.ParseFloat(strings.TrimSpace(parts[1]), 64)
					if err1 == nil && err2 == nil {
						var errMsg string
						if val, ok := errorMap["range"]; ok {
							errMsg = val
							delete(errorMap, "range")
						}
						validators = append(validators, &RangeValidator{FieldName: fieldName, Min: min, Max: max, ErrorMessage: errMsg})
					}
				}
			}
		case "len":
			if tagValue != "" {
				length, err := strconv.Atoi(tagValue)
				if err == nil {
					var errMsg string
					if val, ok := errorMap["len"]; ok {
						errMsg = val
						delete(errorMap, "len")
					}
					validators = append(validators, &LenValidator{FieldName: fieldName, Length: length, ErrorMessage: errMsg})
				}
			}
		case "pattern":
			if tagValue != "" {
				re, err := regexp.Compile(tagValue)
				if err == nil {
					var errMsg string
					if val, ok := errorMap["pattern"]; ok {
						errMsg = val
						delete(errorMap, "pattern")
					}
					validators = append(validators, &PatternValidator{FieldName: fieldName, Pattern: tagValue, Regexp: re, ErrorMessage: errMsg})
				}
			}
		case "in":
			if tagValue != "" {
				allowed := strings.Split(tagValue, "|")
				var errMsg string
				if val, ok := errorMap["in"]; ok {
					errMsg = val
					delete(errorMap, "in")
				}
				validators = append(validators, &InValidator{FieldName: fieldName, Allowed: allowed, ErrorMessage: errMsg})
			}
		case "eq":
			if tagValue != "" {
				if valInt, err := strconv.Atoi(tagValue); err == nil {
					var errMsg string
					if val, ok := errorMap["eq"]; ok {
						errMsg = val
						delete(errorMap, "eq")
					}
					validators = append(validators, &EqValidator{FieldName: fieldName, Value: valInt, ErrorMessage: errMsg})
				} else if valFloat, err := strconv.ParseFloat(tagValue, 64); err == nil {
					var errMsg string
					if val, ok := errorMap["eq"]; ok {
						errMsg = val
						delete(errorMap, "eq")
					}
					validators = append(validators, &EqValidator{FieldName: fieldName, Value: valFloat, ErrorMessage: errMsg})
				} else if tagValue == "true" || tagValue == "false" {
					valBool, _ := strconv.ParseBool(tagValue)
					var errMsg string
					if val, ok := errorMap["eq"]; ok {
						errMsg = val
						delete(errorMap, "eq")
					}
					validators = append(validators, &EqValidator{FieldName: fieldName, Value: valBool, ErrorMessage: errMsg})
				} else {
					var errMsg string
					if val, ok := errorMap["eq"]; ok {
						errMsg = val
						delete(errorMap, "eq")
					}
					validators = append(validators, &EqValidator{FieldName: fieldName, Value: tagValue, ErrorMessage: errMsg})
				}
			}
		case "gt":
			if tagValue != "" {
				if valFloat, err := strconv.ParseFloat(tagValue, 64); err == nil {
					var errMsg string
					if val, ok := errorMap["gt"]; ok {
						errMsg = val
						delete(errorMap, "gt")
					}
					validators = append(validators, &GtValidator{FieldName: fieldName, Value: valFloat, ErrorMessage: errMsg})
				}
			}
		case "lt":
			if tagValue != "" {
				if valFloat, err := strconv.ParseFloat(tagValue, 64); err == nil {
					var errMsg string
					if val, ok := errorMap["lt"]; ok {
						errMsg = val
						delete(errorMap, "lt")
					}
					validators = append(validators, &LtValidator{FieldName: fieldName, Value: valFloat, ErrorMessage: errMsg})
				}
			}
		case "contains":
			if tagValue != "" {
				var errMsg string
				if val, ok := errorMap["contains"]; ok {
					errMsg = val
					delete(errorMap, "contains")
				}
				validators = append(validators, &ContainsValidator{FieldName: fieldName, Substring: tagValue, ErrorMessage: errMsg})
			}
		case "email":
			var errMsg string
			if val, ok := errorMap["email"]; ok {
				errMsg = val
				delete(errorMap, "email")
			}
			validators = append(validators, &EmailValidator{FieldName: fieldName, ErrorMessage: errMsg})
		case "url":
			var errMsg string
			if val, ok := errorMap["url"]; ok {
				errMsg = val
				delete(errorMap, "url")
			}
			validators = append(validators, &URLValidator{FieldName: fieldName, ErrorMessage: errMsg})
		case "alpha":
			var errMsg string
			if val, ok := errorMap["alpha"]; ok {
				errMsg = val
				delete(errorMap, "alpha")
			}
			validators = append(validators, &AlphaValidator{FieldName: fieldName, ErrorMessage: errMsg})
		case "alphanum":
			var errMsg string
			if val, ok := errorMap["alphanum"]; ok {
				errMsg = val
				delete(errorMap, "alphanum")
			}
			validators = append(validators, &AlphanumValidator{FieldName: fieldName, ErrorMessage: errMsg})
		}
	}
	return validators
}
