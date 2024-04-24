package org.fsg.backend.request;


import jakarta.validation.constraints.NotBlank;
import jakarta.validation.constraints.Size;
import lombok.Data;

@Data
public class NoteRequest {
    @NotBlank(message = "Title is required")
    @Size(min = 3, max = 50, message = "Title must be between 3 and 50 characters")
    private String title;

    @NotBlank(message = "Content is required")
    @Size(min = 10, message = "Content must be at least 10 characters long")
    private String content;
}
