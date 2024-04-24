package org.fsg.backend.globalExceptions;

import org.fsg.backend.exceptions.InvalidTokenException;
import org.fsg.backend.exceptions.NoteNotFoundException;
import org.fsg.backend.exceptions.UserNotFoundException;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.MethodArgumentNotValidException;
import org.springframework.web.bind.annotation.ExceptionHandler;
import org.springframework.web.bind.annotation.RestControllerAdvice;

import java.util.*;


@RestControllerAdvice
public class GlobalExceptionHandler {

    private ResponseEntity<Map<String, Object>> buildErrorResponse(HttpStatus status, String message, List<Map<String, String>> errors) {
        Map<String, Object> responseBody = new HashMap<>();
        responseBody.put("message", message);
        responseBody.put("errors", errors);
        return ResponseEntity.status(status).body(responseBody);
    }

    @ExceptionHandler(Exception.class)
    public ResponseEntity<Map<String, Object>> handleGeneralException(Exception ex) {
        List<Map<String, String>> errors = new ArrayList<>();
        Map<String, String> error = new HashMap<>();
        error.put("message", ex.getMessage() != null ? ex.getMessage() : "Unexpected error occurred");
        errors.add(error);

        return buildErrorResponse(HttpStatus.INTERNAL_SERVER_ERROR, "Internal server error", errors);
    }

    @ExceptionHandler(MethodArgumentNotValidException.class)
    public ResponseEntity<Map<String, Object>> handleMethodArgumentNotValidException(MethodArgumentNotValidException ex) {
        List<Map<String, String>> errors = new ArrayList<>();
        ex.getBindingResult().getFieldErrors().forEach(fieldError -> {
            Map<String, String> error = new HashMap<>();
            error.put("field", fieldError.getField());
            error.put("errorMessage", fieldError.getDefaultMessage());
            errors.add(error);
        });

        Map<String, Object> responseBody = new HashMap<>();
        responseBody.put("message", "Validation errors occurred");
        responseBody.put("errors", errors);

        return ResponseEntity.status(HttpStatus.BAD_REQUEST).body(responseBody);
    }

    @ExceptionHandler({InvalidTokenException.class, UserNotFoundException.class})
    public ResponseEntity<Map<String, Object>> handleAuthenticationException(Exception ex) {
        List<Map<String, String>> errors = new ArrayList<>();
        Map<String, String> error = new HashMap<>();
        error.put("message", ex.getMessage());
        errors.add(error);

        return buildErrorResponse(HttpStatus.UNAUTHORIZED, "Authentication error", errors);
    }


    @ExceptionHandler(IllegalAccessException.class)
    public ResponseEntity<Map<String, Object>> handleIllegalAccessException(IllegalAccessException ex) {
        List<Map<String, String>> errors = new ArrayList<>();
        Map<String, String> error = new HashMap<>();
        error.put("message", ex.getMessage());
        errors.add(error);

        return buildErrorResponse(HttpStatus.FORBIDDEN, "Access denied", errors);
    }

    @ExceptionHandler(NoteNotFoundException.class)
    public ResponseEntity<Map<String, Object>> handleNoteNotFoundException(NoteNotFoundException ex) {
        return buildErrorResponse(
                HttpStatus.NOT_FOUND,
                "Note not found",
                Collections.singletonList(Collections.singletonMap("message", ex.getMessage()))
        );
    }


}