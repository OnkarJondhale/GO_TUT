package main

import "fmt"

type Course struct {
	Name     string   `json:"name"`
	Price    int      `json:"price"`
	Platform string   `json:"platform"`
	Password string   `json:"-"` // The password will not be present in JSON 
	Tags     []string `json:"tags,omitempty"` // The omitempty will remove the field tags if it is not present
}

func main() {
	// Create a slice of Course
	courses := []Course{
		{"web-dev", 100, "www.edu/web.com", "Pass@123", []string{"web", "frontend", "backend", "database"}},
		{"android-dev", 100, "www.edu/android.com", "Pass@123", []string{"android", "os", "kotlin", "swift"}},
		{"cloud-based-dev", 100, "www.edu/cloud.com", "Pass@123", []string{"aws","azure"}},
		{"machine-learning", 100, "www.edu/ml.com", "Pass@123", []string{"python","data-science"}},
		{"ai-tools", 100, "www.edu/ai.com", "Pass@123", []string{"llama","openai","claude","gpt","gemini"}},
	}

	fmt.Println(courses)
}