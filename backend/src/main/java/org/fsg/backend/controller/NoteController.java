package org.fsg.backend.controller;

import jakarta.validation.Valid;
import org.fsg.backend.model.Note;
import org.fsg.backend.model.User;
import org.fsg.backend.request.NoteRequest;
import org.fsg.backend.response.NoteResponse;
import org.fsg.backend.service.NoteService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@RestController
@RequestMapping("api/v1/notes")
public class NoteController {

    private final NoteService noteService;

    @Autowired
    public NoteController(NoteService noteService) {
        this.noteService = noteService;
    }

    @PostMapping()
    public ResponseEntity<NoteResponse> createNoteHandler(
            @Valid
            @RequestBody NoteRequest note,
            @RequestHeader("Authorization")
            String authToken) {
        User user = noteService.validateAndExtractUser(authToken);
        Note newNote = noteService.createNote(note, user);
        NoteResponse response = noteService.createNoteResponse(newNote);
        return ResponseEntity.ok(response);
    }

    @DeleteMapping("/{id}")
    public ResponseEntity<?> deleteNoteHandler(
            @PathVariable Long id,
            @RequestHeader("Authorization") String authToken
    ) throws Exception {
        User user = noteService.validateAndExtractUser(authToken);
        noteService.deleteNote(id, user);
        return ResponseEntity.ok().build();
    }

    @PatchMapping("/{id}")
    public ResponseEntity<NoteResponse> updateNoteHandler(
            @Valid
            @PathVariable Long id,
            @RequestBody NoteRequest noteRequest,
            @RequestHeader("Authorization") String authToken
    ) throws Exception {
        User user = noteService.validateAndExtractUser(authToken);
        Note updatedNote = noteService.updateNote(id, noteRequest, user);
        NoteResponse response = noteService.createNoteResponse(updatedNote);
        return ResponseEntity.ok(response);
    }

    @GetMapping("/all")
    public ResponseEntity<List<NoteResponse>> getAllNotesHandler(
            @RequestHeader("Authorization") String authToken
    ) {
        User user = noteService.validateAndExtractUser(authToken);
        List<Note> notes = noteService.findAllNotesByUser(user);

        List<NoteResponse> response = noteService.createNoteResponseList(notes);
        return ResponseEntity.ok(response);
    }

    @GetMapping("/{id}")
    public ResponseEntity<NoteResponse> getNoteByIdHandler(
            @PathVariable Long id,
            @RequestHeader("Authorization") String authToken
    ) throws IllegalAccessException {
        User user = noteService.validateAndExtractUser(authToken);
        Note note = noteService.findNoteById(id, user);
        NoteResponse response = noteService.createNoteResponse(note);
        return ResponseEntity.ok(response);
    }
}

