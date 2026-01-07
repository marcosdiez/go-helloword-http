# HTTP helloworld Container
This is yet another simple golang HTTP hello world server.

It does have some interesting features, which are provided as environmental variables:

* HTTP_PORT the HTTP port this server will listen to. Default is 8080
* SECONDAY_HTTP_PORT a possible, optional HTTP port this server will listen to. Default is None. Useful when your k8s deployments use a different port for healthcheck.
* START_DELAY (int, seconds): How many seconds it should wait before binding to HTTP_PORT. Great for you to test your readiness k8s probes. Default is zero.
* FREEZE_PERCENTAGE (Float between 0 and 100): A chance the server will freeze and NEVER bind to HTTP_PORT. The console output will inform you that. This parameter is evaulated on startup, before START_DELAY so by looking at the logs you will immediatelly know if the server will work or not. Default value is zero, which means is does not freeze.
* MESSAGE (string): a message that will appear in the log of every request. Default is blank, suggested values: "BLUE" and "GREEN"

# URLs
* Docker: https://hub.docker.com/r/marcosdiez/helloworld-http
* GitHub: https://github.com/marcosdiez/go-helloword-http
