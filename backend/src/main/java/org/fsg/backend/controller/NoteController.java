package org.fsg.backend.controller;

import org.fsg.backend.exceptions.InvalidTokenException;
import org.fsg.backend.exceptions.UserNotFoundException;
import org.fsg.backend.model.Note;
import org.fsg.backend.model.User;
import org.fsg.backend.request.NoteRequest;
import org.fsg.backend.response.ErrorResponse;
import org.fsg.backend.response.NoteResponse;
import org.fsg.backend.service.NoteService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;


@RestController
@RequestMapping("api/v1/notes")
public class NoteController {

        private final NoteService noteService;

        @Autowired
        public NoteController(NoteService noteService) {
            this.noteService = noteService;
        }


    @PostMapping()
    public ResponseEntity<?> createNoteHandler(
            @RequestBody NoteRequest note,
            @RequestHeader("Authorization") String authToken
    ) {
        try {
            User user = noteService.validateAndExtractUser(authToken);
            Note newNote = noteService.createNote(note, user);
            NoteResponse response = noteService.createNoteResponse(newNote);
            return ResponseEntity.ok(response);
        } catch (InvalidTokenException | UserNotFoundException e) {
            ErrorResponse errorResponse = new ErrorResponse(HttpStatus.UNAUTHORIZED.value(), e.getMessage());
            return ResponseEntity.status(HttpStatus.UNAUTHORIZED).body(errorResponse);
        } catch (Exception e) {
            ErrorResponse errorResponse = new ErrorResponse(HttpStatus.INTERNAL_SERVER_ERROR.value(), "Internal server error");
            return ResponseEntity.status(HttpStatus.INTERNAL_SERVER_ERROR).body(errorResponse);
        }
    }


    //    @DeleteMapping("/{id}")
    //    @PatchMapping("/{id}")
    //    @GetMapping("/all")
    //    @GetMapping("/{id}")
}