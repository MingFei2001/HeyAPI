package handlers

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Note struct {
	ID        int       `json:"id"`
	Note      string    `json:"note"`
	Timestamp time.Time `json:"created_at"`
}

var notes = []Note{}
var nextNoteID = 1

// create a new note
func CreateNoteHandler(w http.ResponseWriter, r *http.Request) {
	// Ensure the request method is POST
	if r.Method != http.MethodPost {
		log.Println("Method not allowed")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse the form data
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Failed to parse form data", http.StatusBadRequest)
		return
	}

	// Get the "note" field from the form data
	noteContent := r.FormValue("note")
	if noteContent == "" {
		http.Error(w, "Note cannot be empty", http.StatusBadRequest)
		return
	}

	// Create a new note
	newNote := Note{
		ID:        nextNoteID,
		Note:      noteContent,
		Timestamp: time.Now(),
	}
	nextNoteID++

	// Add the new note to the list
	notes = append(notes, newNote)

	// Redirect back to the notes page
	http.Redirect(w, r, "/notes", http.StatusSeeOther)
}

// get all notes
func GetNotesHandler(w http.ResponseWriter, r *http.Request) {
	// Ensure the request method is GET
	if r.Method != http.MethodGet {
		log.Println("Method not allowed")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(notes)
}

// grab note by id
func GetNoteHandler(w http.ResponseWriter, r *http.Request) {
	// Ensure the request method is GET
	if r.Method != http.MethodGet {
		log.Println("Method not allowed")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// separate the id from the query parameters
	id := r.URL.Query().Get("id")
	if id == "" {
		log.Println("Id cannot be empty.")
		http.Error(w, "ID parameter cannot be empty.", http.StatusBadRequest)
		return
	}

	// convert id from str to int, if it fails return 400
	idInt, err := strconv.Atoi(id)
	if err != nil {
		log.Printf("Invalid ID: %v", err)
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// loop thru all notes until the id matches, yes i am lazy
	for _, note := range notes {
		if note.ID == idInt {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(note)
			return
		}
	}

	// if all hope is lost, return 404
	log.Println("Note not found")
	http.Error(w, "Note not found", http.StatusNotFound)
}

// delete note by id
func DeleteNoteHandler(w http.ResponseWriter, r *http.Request) {
	// Ensure the request method is DELETE
	if r.Method != http.MethodDelete {
		log.Println("Method not allowed")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id")
	if id == "" {
		log.Println("ID cannot be empty.")
		http.Error(w, "ID parameter cannot be empty.", http.StatusBadRequest)
		return
	}

	// convert id from str to int, if it fails return 400
	idInt, err := strconv.Atoi(id)
	if err != nil {
		log.Printf("Invalid ID: %v", err)
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// make a new array with everything else but the chosen one
	for i, note := range notes {
		if note.ID == idInt {
			notes = append(notes[:i], notes[i+1:]...)
			json.NewEncoder(w).Encode(note)
			return
		}
	}

	log.Printf("Note with ID %d not found", idInt)
	http.Error(w, "Note not found", http.StatusNotFound)
}

func ServeNotesPage(w http.ResponseWriter, r *http.Request) {
	// Ensure the request method is GET
	if r.Method != http.MethodGet {
		log.Println("Method not allowed")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse the notes.html template
	tmpl, err := template.ParseFiles("templates/notes.html")
	if err != nil {
		log.Printf("Error loading template: %v", err)
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}

	// Render the template with the notes data
	err = tmpl.Execute(w, map[string]any{
		"Notes": notes,
	})
	if err != nil {
		log.Printf("Error rendering template: %v", err)
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}
