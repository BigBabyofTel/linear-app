package org.fsg.backend.service;


import io.jsonwebtoken.JwtException;
import org.fsg.backend.config.JwtProvider;
import org.fsg.backend.exceptions.InvalidTokenException;
import org.fsg.backend.exceptions.UserNotFoundException;
import org.fsg.backend.model.User;
import org.fsg.backend.repository.UserRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.security.core.userdetails.UserDetails;
import org.springframework.security.core.userdetails.UsernameNotFoundException;
import org.springframework.stereotype.Service;

import java.util.ArrayList;

@Service
public class UserServiceImpl implements UserService {
    private final UserRepository userRepository;
    private final JwtProvider jwtProvider;

    @Autowired
    public UserServiceImpl(UserRepository userRepository, JwtProvider jwtProvider) {
        this.userRepository = userRepository;
        this.jwtProvider = jwtProvider;
    }

    @Override
    public User findUserByJwtToken(String token) throws InvalidTokenException, UserNotFoundException {
        try {
            String email = jwtProvider.getEmailFromJwtToken(token);
            User user = userRepository.findByEmail(email);
            if (user == null) {
                throw new UserNotFoundException("User not found with email: " + email);
            }
            return user;
        } catch (JwtException | IllegalArgumentException ex) {
            throw new InvalidTokenException("Invalid JWT token", ex);
        }
    }

    @Override
    public User findUserByEmail(String email) throws UserNotFoundException {
        User user = userRepository.findByEmail(email);
        if (user == null) {
            throw new UserNotFoundException("User not found with email: " + email);
        }
        return user;
    }

    @Override
    public UserDetails loadUserByEmail(String email) {
        User user = userRepository.findByEmail(email);
        if (user == null) {
            throw new UsernameNotFoundException("User not found with email: " + email);
        }

        return new org.springframework.security.core.userdetails.User(
                user.getEmail(), user.getPassword(), new ArrayList<>()
        );
    }
}

