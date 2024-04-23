package org.fsg.backend.service;

import org.fsg.backend.exceptions.InvalidTokenException;
import org.fsg.backend.model.Note;
import org.fsg.backend.model.User;
import org.fsg.backend.request.NoteRequest;
import org.fsg.backend.response.NoteResponse;


public interface NoteService {
    Note createNote(NoteRequest noteRequest, User user);
    NoteResponse createNoteResponse(Note note);
    User validateAndExtractUser(String authToken) throws InvalidTokenException;
}
