<<<<<<< HEAD
FROM debian:12.7@sha256:27586f4609433f2f49a9157405b473c62c3cb28a581c413393975b4e8496d0ab AS envoy
=======
FROM debian:13.1@sha256:fd8f5a1df07b5195613e4b9a0b6a947d3772a151b81975db27d47f093f60c6e6 AS envoy
>>>>>>> b52d7f8b2 (chore(deps): bump debian:13.1 from 833c135 to fd8f5a1 (#662))
ARG ARCH

COPY /build/artifacts-linux-$ARCH/envoy/envoy /envoy
