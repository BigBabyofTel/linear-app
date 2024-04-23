package org.fsg.backend.controller;


import org.fsg.backend.model.Note;
import org.fsg.backend.model.User;
import org.fsg.backend.repository.NoteRepository;
import org.fsg.backend.request.NoteRequest;
import org.fsg.backend.response.NoteResponse;
import org.fsg.backend.service.UserService;
import org.fsg.backend.utils.AuthUtils;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.util.Date;

@RestController
@RequestMapping("api/v1/notes")
public class NoteController {

        private final NoteRepository noteRepository;
        private final UserService userService;

        @Autowired
        public NoteController(NoteRepository noteRepository, UserService userService) {
            this.noteRepository = noteRepository;
            this.userService = userService;
        }



    @PostMapping()
    public ResponseEntity<NoteResponse> createNoteHandler(
            @RequestBody NoteRequest note,
            @RequestHeader("Authorization") String authToken
    ) throws Exception {

        String token = AuthUtils.extractBearerToken(authToken);
        if (token == null) {
            return ResponseEntity.status(HttpStatus.UNAUTHORIZED).build();
        }
        User user = userService.findUserByJwtToken(token);

        Note newNote = new Note();
        newNote.setTitle(note.getTitle());
        newNote.setContent(note.getContent());
        newNote.setCreatedAt(new Date());
        newNote.setUser(user);
        noteRepository.save(newNote);

        NoteResponse response = new NoteResponse();
        response.setId(newNote.getId());
        response.setTitle(newNote.getTitle());
        response.setContent(newNote.getContent());
        response.setCreatedAt(newNote.getCreatedAt());

        return ResponseEntity.ok(response);
    }

    @GetMapping("/hello")
    public String hello() {
        return "Hello World!";
    }
    //    @DeleteMapping("/{id}")
    //    @PatchMapping("/{id}")
    //    @GetMapping("/all")
    //    @GetMapping("/{id}")
}
