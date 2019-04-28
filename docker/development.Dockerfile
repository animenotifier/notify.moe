# Install development environment
FROM blitzprog/aero

# Download database
RUN git clone --progress --depth=1 https://github.com/animenotifier/database ~/.aero/db/arn

# Download notify.moe dependencies
RUN curl -s -o ~/go.mod https://raw.githubusercontent.com/animenotifier/notify.moe/go/go.mod && \
	curl -s -o ~/go.sum https://raw.githubusercontent.com/animenotifier/notify.moe/go/go.sum && \
	cd ~/ && \
	go mod download && \
	rm ~/go.mod ~/go.sum

# Create empty working directory
WORKDIR /my