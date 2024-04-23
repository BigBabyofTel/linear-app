package org.fsg.backend.response;


import lombok.Data;

import java.util.Date;

@Data
public class NoteResponse {
    private Long id;
    private String title;
    private String content;
    private Date createdAt;
}
