<<<<<<< HEAD
FROM debian:12.5@sha256:b37bc259c67238d814516548c17ad912f26c3eed48dd9bb54893eafec8739c89 as envoy
=======
FROM debian:13.1@sha256:fd8f5a1df07b5195613e4b9a0b6a947d3772a151b81975db27d47f093f60c6e6 AS envoy
>>>>>>> b52d7f8b2 (chore(deps): bump debian:13.1 from 833c135 to fd8f5a1 (#662))
ARG ARCH

COPY /build/artifacts-linux-$ARCH/envoy/envoy /envoy
