FROM maven:latest as builder
WORKDIR /app/
COPY ./ /app/
RUN mvn -v
RUN mvn package
# EXPOSE 8000
# CMD ["java", "-jar", "target/app.jar"]

FROM openjdk:latest  
WORKDIR /app/
COPY --from=builder /app/target/ /app/
EXPOSE 8000
CMD ["java", "-jar", "app.jar"]

