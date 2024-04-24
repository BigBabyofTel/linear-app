package org.fsg.backend.service;

import org.fsg.backend.exceptions.InvalidTokenException;
import org.fsg.backend.model.Note;
import org.fsg.backend.model.User;
import org.fsg.backend.request.NoteRequest;
import org.fsg.backend.response.NoteResponse;

import java.util.List;


public interface NoteService {
    Note createNote(NoteRequest noteRequest, User user);
    NoteResponse createNoteResponse(Note note);
    User validateAndExtractUser(String authToken) throws InvalidTokenException;
    List<Note> findAllNotesByUser(User user);
    Note updateNote(Long id, NoteRequest noteRequest, User user) throws Exception;
    void deleteNote(Long id, User user) throws Exception;
    Note findNoteById(Long id, User user) throws IllegalAccessException;
    List<NoteResponse> createNoteResponseList(List<Note> notes);
}
