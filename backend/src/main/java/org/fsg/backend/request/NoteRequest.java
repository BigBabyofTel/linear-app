package org.fsg.backend.request;


import lombok.Data;

@Data
public class NoteRequest {
    private String title;
    private String content;
}
