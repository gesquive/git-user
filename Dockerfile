FROM gesquive/go-builder:latest AS builder

ENV APP=git-user

# This requires that `make release-snapshot` be called first
COPY dist/ /dist/
RUN copy-release
RUN chmod +x /app/git-user

# =============================================================================
FROM gesquive/docker-base:latest
LABEL maintainer="Gus Esquivel <gesquive@gmail.com>"

COPY --from=builder /app/${APP} /app/

ENTRYPOINT ["/app/git-user"]
