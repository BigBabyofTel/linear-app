# Build stage
FROM openjdk:17-slim as build

WORKDIR /app

COPY mvnw .
COPY .mvn .mvn
COPY pom.xml .
RUN ./mvnw dependency:go-offline -B

COPY src src
RUN ./mvnw package -DskipTests
RUN mkdir -p target/dependency && (cd target/dependency; jar -xf ../*.jar)

# Final stage
FROM openjdk:17-slim

WORKDIR /app

RUN apt-get update && apt-get install -y netcat

COPY --from=build /app/target/dependency/BOOT-INF/lib /app/lib
COPY --from=build /app/target/dependency/META-INF /app/META-INF
COPY --from=build /app/target/dependency/BOOT-INF/classes /app
COPY wait-for.sh /app
RUN chmod +x /app/wait-for.sh

ENV POSTGRES_URL=""
ENV POSTGRES_USERNAME=""
ENV POSTGRES_PASSWORD=""
ENV AWS_ACCESS_KEY=""
ENV AWS_SECRET_KEY=""

EXPOSE 8080

ENTRYPOINT ["./wait-for.sh", "db:5432", "--", "java", "-cp", "app:app/lib/*", "com.example.demo.WuhuServerApplication"]
