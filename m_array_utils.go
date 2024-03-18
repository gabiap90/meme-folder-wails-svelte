package main

func removeString(slice []string, remove string) []string {
	var result []string
	for _, value := range slice {
		if value != remove {
			result = append(result, value)
		}
	}
	return result
}

func addIfNotExists(slice []string, element string) []string {
	for _, value := range slice {
		if value == element {
			return slice
		}
	}

	return append(slice, element)
}
