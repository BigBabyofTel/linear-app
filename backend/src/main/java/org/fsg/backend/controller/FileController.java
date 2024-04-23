package org.fsg.backend.controller;

import com.amazonaws.regions.Regions;
import org.fsg.backend.service.S3Service;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.multipart.MultipartFile;

import java.util.Map;
import java.util.Objects;

@RestController
@RequestMapping("/api/v1/files")
public class FileController {

    S3Service s3Service;

    @Autowired
    public FileController(S3Service s3Service) {
        this.s3Service = s3Service;
    }

    @PostMapping("/image")
    public ResponseEntity<?> fileUpload(@RequestParam("file") MultipartFile file) {
        try {
            String url = s3Service.uploadFile(file.getInputStream(), file.getContentType(), Objects.requireNonNull(file.getOriginalFilename()));
            return ResponseEntity.ok(Map.of("url", url));
        } catch (Exception e) {
            return ResponseEntity.status(500).body(Map.of("message", "Failed to upload file: " + e.getMessage()));
        }
    }
}