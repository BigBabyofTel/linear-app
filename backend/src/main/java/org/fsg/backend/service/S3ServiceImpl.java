package org.fsg.backend.service;

import com.amazonaws.SdkClientException;
import com.amazonaws.regions.Regions;
import com.amazonaws.services.s3.AmazonS3;
import com.amazonaws.services.s3.AmazonS3ClientBuilder;
import com.amazonaws.services.s3.model.ObjectMetadata;
import com.amazonaws.services.s3.model.PutObjectRequest;
import org.springframework.stereotype.Service;

import java.io.IOException;
import java.io.InputStream;
import java.util.UUID;

@Service
public class S3ServiceImpl implements S3Service {
    private final Regions clientRegion = Regions.EU_CENTRAL_1;


    @Override
    public String uploadFile(InputStream inputStream, String type, String originalFilename)
            throws IOException, SdkClientException {
        String fileExtension = originalFilename.substring(originalFilename.lastIndexOf("."));
        String filename = UUID.randomUUID() + fileExtension;

        AmazonS3 s3Client = AmazonS3ClientBuilder.standard()
                .withRegion(clientRegion)
                .build();

        ObjectMetadata metadata = new ObjectMetadata();
        metadata.setContentType(type);
        metadata.setContentLength(inputStream.available());

        String bucketName = "wuhuu-bucket";
        PutObjectRequest req = new PutObjectRequest(bucketName, filename, inputStream, metadata);
        s3Client.putObject(req);

        return "https://s3." + Regions.EU_CENTRAL_1.getName() + ".amazonaws.com/" + "wuhuu-bucket" + "/" + filename;
    }

}