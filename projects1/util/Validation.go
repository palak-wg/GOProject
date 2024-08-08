package util

import (
	"regexp"
)

const (
	minNameLength        = 1
	maxNameLength        = 50
	minDesignationLength = 1
	maxDesignationLength = 50
)

// TakenUserId checks if the UserID is already taken
func TakenUserId(uID string) bool {
	existCred := ReadCredFile()
	_, exists := existCred[uID+"ID"]
	if !exists {
		return false
	}
	return existCred["ID"] == uID
}

// IsValidPhoneNumber checks if the phone number is valid (10 digits)
func IsValidPhoneNumber(phone string) bool {
	return regexp.MustCompile(`^\d{10}$`).MatchString(phone)
}

// IsValidName checks if the name is valid (1 to 50 characters, letters, and spaces only)
func IsValidName(name string) bool {
	if len(name) < minNameLength || len(name) > maxNameLength {
		return false
	}
	return regexp.MustCompile(`^[A-Za-z\s]+$`).MatchString(name)
}

// IsValidDesignation checks if the designation is valid (1 to 50 characters)
func IsValidDesignation(designation string) bool {
	if len(designation) < minDesignationLength || len(designation) > maxDesignationLength {
		return false
	}
	return true
}

// UserExists returns true if user already exists and false if not
func UserExists(cred map[string]string, uID, pass string) bool {
	_, exists := cred[uID+"ID"]
	if !exists {
		return false
	}
	return cred[uID+"ID"] == uID && cred[uID+"Password"] == pass && (cred[uID+"AppStatus"] == "user" || cred[uID+"AppStatus"] == "User")
}
func AdminExists(cred map[string]string, uID, pass string) bool {
	_, exists := cred[uID+"ID"]
	if !exists {
		return false
	}
	return cred[uID+"ID"] == uID && cred[uID+"Password"] == pass && (cred[uID+"AppStatus"] == "admin" || cred[uID+"AppStatus"] == "Admin")
}
