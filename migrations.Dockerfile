FROM openjdk:8-jre
COPY ./migrations /flyway/sql
WORKDIR /flyway
ENV FLYWAY_VERSION=9.4.0
RUN curl -L https://repo1.maven.org/maven2/org/flywaydb/flyway-commandline/${FLYWAY_VERSION}/flyway-commandline-${FLYWAY_VERSION}.tar.gz -o flyway-commandline-${FLYWAY_VERSION}.tar.gz \
    && tar -xzf flyway-commandline-${FLYWAY_VERSION}.tar.gz --strip-components=1 \
    && rm flyway-commandline-${FLYWAY_VERSION}.tar.gz

CMD sh -c "./flyway -url=jdbc:postgresql://localhost:5432/marine-logistics -user=homestead -password=secret migrate"