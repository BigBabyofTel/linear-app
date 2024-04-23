package org.fsg.backend.service;

import org.fsg.backend.exceptions.InvalidTokenException;
import org.fsg.backend.model.Note;
import org.fsg.backend.model.User;
import org.fsg.backend.repository.NoteRepository;
import org.fsg.backend.request.NoteRequest;
import org.fsg.backend.response.NoteResponse;
import org.fsg.backend.utils.AuthUtils;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import java.util.Date;

@Service
public class NoteServiceImpl implements NoteService {

    private final NoteRepository noteRepository;
    private final UserService userService;

    @Autowired
    public NoteServiceImpl(NoteRepository noteRepository, UserService userService) {
        this.noteRepository = noteRepository;
        this.userService = userService;
    }

    @Override
    public Note createNote(NoteRequest noteRequest, User user) {
        Note newNote = new Note();
        newNote.setTitle(noteRequest.getTitle());
        newNote.setContent(noteRequest.getContent());
        newNote.setCreatedAt(new Date());
        newNote.setUser(user);
        noteRepository.save(newNote);
        return newNote;
    }

    @Override
    public NoteResponse createNoteResponse(Note note) {
        NoteResponse response = new NoteResponse();
        response.setId(note.getId());
        response.setTitle(note.getTitle());
        response.setContent(note.getContent());
        response.setCreatedAt(note.getCreatedAt());
        return response;
    }

    @Override
    public User validateAndExtractUser(String authToken) throws InvalidTokenException {
        String token = AuthUtils.extractBearerToken(authToken);
        if (token == null) {
            throw new InvalidTokenException("Unauthorized: No token provided", null);
        }
        return userService.findUserByJwtToken(token);
    }
}
