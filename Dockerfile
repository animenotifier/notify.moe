# Install development environment
FROM blitzprog/aero
RUN git clone --progress --verbose --depth=1 https://github.com/animenotifier/database ~/.aero/db/arn

# Expect ~/notify.moe to be mounted as a volume
RUN cd ~/notify.moe && \
	tsc && \
	pack && \
	go build