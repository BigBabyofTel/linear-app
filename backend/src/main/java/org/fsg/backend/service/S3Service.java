package org.fsg.backend.service;

import com.amazonaws.SdkClientException;

import java.io.IOException;
import java.io.InputStream;

public interface S3Service {
     String uploadFile(InputStream inputStream, String type, String originalFilename)
            throws IOException, SdkClientException;
}
