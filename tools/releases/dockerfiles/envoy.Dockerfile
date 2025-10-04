<<<<<<< HEAD
FROM debian:13.0@sha256:6d87375016340817ac2391e670971725a9981cfc24e221c47734681ed0f6c0f5 AS envoy
=======
FROM debian:13.1@sha256:fd8f5a1df07b5195613e4b9a0b6a947d3772a151b81975db27d47f093f60c6e6 AS envoy
>>>>>>> b52d7f8b2 (chore(deps): bump debian:13.1 from 833c135 to fd8f5a1 (#662))
ARG ARCH

COPY /build/artifacts-linux-$ARCH/envoy/envoy /envoy
