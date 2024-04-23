package org.fsg.backend.utils;

public class AuthUtils {
    public static String extractBearerToken(String authHeader) {
        if (authHeader != null && authHeader.startsWith("Bearer ")) {
            return authHeader.substring(7);
        }

        return null;
    }


}
