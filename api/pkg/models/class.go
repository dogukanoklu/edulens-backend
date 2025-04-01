package models

// Database 

type Class struct {
	ID string `json:"id"`
	Level string `json:"level"`
	Branch string `json:"branch"`
}

// Response

type ResClasses struct {
	ID string `json:"id"`
	Level string `json:"level"`
	Branch string `json:"branch"`
}