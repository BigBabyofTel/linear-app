package com.example.demo.service;

import com.example.demo.entity.User;
import com.example.demo.repository.UserRepository;
import org.springframework.stereotype.Service;

import java.lang.reflect.Array;
import java.util.List;
import java.util.Optional;

@Service
public class UserService {
    private final UserRepository userRepository;

    public UserService(UserRepository userRepository){
        this.userRepository = userRepository;
    }

   public List<User> findAll() {
       return userRepository.findAll();
   }

   public User findById(Integer id) {
        Optional<User> user = userRepository.findById(id);

        if (user.isPresent()) {
            return user.get();
        } else {
            throw new RuntimeException("User not found");
        }
   }

   public User addUser(User user) {
        return userRepository.save(user);
   }

   public void deleteById(Integer id) {
        User user = findById(id);
        userRepository.delete(user);
   }

   public User updateUser(User userToUpdate) {
       if (userRepository.existsById(userToUpdate.getId())) {
           return userRepository.save(userToUpdate);
       } else {
           throw new RuntimeException("User not found");
       }

   }
}
