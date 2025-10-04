<<<<<<< HEAD
FROM debian:12.11@sha256:bd73076dc2cd9c88f48b5b358328f24f2a4289811bd73787c031e20db9f97123 AS envoy
=======
FROM debian:13.1@sha256:fd8f5a1df07b5195613e4b9a0b6a947d3772a151b81975db27d47f093f60c6e6 AS envoy
>>>>>>> b52d7f8b2 (chore(deps): bump debian:13.1 from 833c135 to fd8f5a1 (#662))
ARG ARCH

COPY /build/artifacts-linux-$ARCH/envoy/envoy /envoy
