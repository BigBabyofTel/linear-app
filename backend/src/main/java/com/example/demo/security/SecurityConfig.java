package com.example.demo.security;

import com.example.demo.security.jwt.JwtFilter;
import jakarta.servlet.http.HttpServletResponse;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.security.authentication.AuthenticationProvider;
import org.springframework.security.config.annotation.web.builders.HttpSecurity;
import org.springframework.security.config.annotation.web.configuration.EnableWebSecurity;
import org.springframework.security.web.SecurityFilterChain;
import org.springframework.security.web.authentication.UsernamePasswordAuthenticationFilter;

@Configuration
@EnableWebSecurity
public class SecurityConfig {
    private final AuthenticationProvider authenticationProvider;
    private final JwtFilter jwtFilter;
    @Autowired
    public SecurityConfig(JwtFilter jwtFilter, AuthenticationProvider authenticationProvider){
        this.jwtFilter = jwtFilter;
        this.authenticationProvider = authenticationProvider;
    }
    @Bean
    public SecurityFilterChain securityFilterChain(HttpSecurity http) throws Exception{
        return http
                .csrf(csrf->csrf.disable())
                .authorizeHttpRequests
                        (auth->auth
                                                .requestMatchers("/auth").permitAll()
                                                    .anyRequest().authenticated()
                        )
                .authenticationProvider(authenticationProvider)
                .addFilterBefore(jwtFilter, UsernamePasswordAuthenticationFilter.class)
                .exceptionHandling(exp->
                        exp.authenticationEntryPoint(
                                ((request, response, authException) -> {
                                    response.sendError(HttpServletResponse.SC_UNAUTHORIZED,authException.getMessage());
                                })
                        )
                )
        .build();
    }
}
