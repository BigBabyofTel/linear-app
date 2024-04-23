package org.fsg.backend.service;

import org.fsg.backend.exceptions.InvalidTokenException;
import org.fsg.backend.exceptions.UserNotFoundException;
import org.fsg.backend.model.User;
import org.springframework.security.core.userdetails.UserDetails;

public interface UserService {
     User findUserByJwtToken(String token) throws InvalidTokenException, UserNotFoundException;
     User findUserByEmail(String email) throws UserNotFoundException;
     UserDetails loadUserByEmail(String email);
}
