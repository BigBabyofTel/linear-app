package org.fsg.backend.repository;

import org.fsg.backend.model.Note;
import org.fsg.backend.model.User;
import org.springframework.data.jpa.repository.JpaRepository;

import java.util.List;

public interface NoteRepository extends JpaRepository<Note, Long> {
    List<Note> findByUser(User user);
}
