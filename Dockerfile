FROM scratch
MAINTAINER Jake Dahn <jake@markupisart.com>
ADD bin/echo /echo

ENTRYPOINT ["/echo"]
