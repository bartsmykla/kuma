<<<<<<< HEAD
FROM debian:12.9@sha256:35286826a88dc879b4f438b645ba574a55a14187b483d09213a024dc0c0a64ed AS envoy
=======
FROM debian:13.1@sha256:fd8f5a1df07b5195613e4b9a0b6a947d3772a151b81975db27d47f093f60c6e6 AS envoy
>>>>>>> b52d7f8b2 (chore(deps): bump debian:13.1 from 833c135 to fd8f5a1 (#662))
ARG ARCH

COPY /build/artifacts-linux-$ARCH/envoy/envoy /envoy
