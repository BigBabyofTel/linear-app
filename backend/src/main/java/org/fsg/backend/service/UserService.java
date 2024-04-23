package org.fsg.backend.service;

import org.fsg.backend.model.User;
import org.springframework.security.core.userdetails.UserDetails;

public interface UserService {
     User findUserByJwtToken(String token) throws Exception;
     User findUserByEmail(String email) throws Exception;
     UserDetails loadUserByEmail(String email);
}
