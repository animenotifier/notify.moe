# Install development environment
FROM blitzprog/aero

# Download database
RUN git clone --progress --depth=1 https://github.com/animenotifier/database ~/.aero/db/arn

# Download notify.moe dependencies
WORKDIR /home/developer
RUN curl -s -o go.mod https://raw.githubusercontent.com/animenotifier/notify.moe/go/go.mod && \
	curl -s -o go.sum https://raw.githubusercontent.com/animenotifier/notify.moe/go/go.sum && \
	go mod download && \
	rm go.*

# Create empty working directory
WORKDIR /my