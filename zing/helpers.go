package zing

import (
	"regexp"
	"strings"
)

func GetVariables(action string) []string {
	var vars []string

	// simple angle identifier like <service> (captures the name without brackets)
	reAngle := regexp.MustCompile(`<\s*([a-zA-Z0-9_]+)\s*>`)

	collect := func(matches [][]string, idx int) {
		for _, m := range matches {
			if len(m) <= idx {
				continue
			}
			name := strings.TrimSpace(m[idx])
			if name == "" {
				continue
			}
			vars = append(vars, name)
		}
	}
	collect(reAngle.FindAllStringSubmatch(action, -1), 1)

	return vars
}

func GetFieldsInCommand(action string) []string {
	var vars []string

	// simple angle identifier like <service> (captures the name without brackets)
	reAngle := regexp.MustCompile(`<\s*([a-zA-Z0-9_]+)\s*>`)

	collect := func(matches [][]string, idx int) {
		for _, m := range matches {
			if len(m) <= idx {
				continue
			}
			name := strings.TrimSpace(m[idx])
			if name == "" {
				continue
			}
			vars = append(vars, name)
		}
	}
	collect(reAngle.FindAllStringSubmatch(action, -1), 0)

	return vars
}

func GetVariablesMap(args []string) map[string]any {
	result := map[string]any{}
	if len(args) <= 0 {
		return result
	}

	for _, arg := range args {
		if strings.Contains(arg, "=") {
			argSplit := strings.Split(arg, "=")
			result[argSplit[0]] = argSplit[1]
		}
	}

	return result
}

func GetFieldsMap(action string) map[string]any {
	vars := map[string]any{}

	// simple angle identifier like <service> (captures the name without brackets)
	reAngle := regexp.MustCompile(`<\s*([a-zA-Z0-9_]+)\s*>`)

	collect := func(matches [][]string) {
		idx := 0
		idy := 1
		for _, m := range matches {
			if len(m) <= idx {
				continue
			}
			key := strings.TrimSpace(m[idy])
			if key == "" {
				continue
			}
			name := strings.TrimSpace(m[idx])
			vars[key] = name
		}
	}
	collect(reAngle.FindAllStringSubmatch(action, -1))

	return vars
}
