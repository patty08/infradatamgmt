Surikator
--
Surikator is [middleware](https://en.wikipedia.org/wiki/Middleware) written in Go language, that you run in [CLI](https://en.wikipedia.org/wiki/CLI) to capture events from orchestrators (docker containers)
of operational data (docker events are structured as ID and LABEL). These events are formated, parsed then
send the operational data to Elasticsearch, so it can be visualized with Kibana.


## Documentation and Getting Started
You can find all about Docker from [here](https://www.docker.com/) .

You can find the documentation and getting started guides for each of the Beats
on the [elastic.co site](https://www.elastic.co/guide/):


* [Logstash](https://www.elastic.co/guide/en/logstash/current/index.html)
* [Metricbeat](https://www.elastic.co/guide/en/beats/metricbeat/current/index.html)

### Setting up your GoLang dev environment
Current Go version used for development is Golang 1.8.1.
The location where you clone is important. Please clone under the source
directory of your `GOPATH`. If you don't have `GOPATH` already set, you can
simply set it to your home directory (`export GOPATH=$HOME`).

    $ mkdir -p ${GOPATH}/src/github.com/sebastienmusso
    $ cd ${GOPATH}/src/github.com/sebastienmusso
    $ git clone https://github.com/sebastienmusso/infradatamgmt.git

## Setting up your Docker environment
You can download docker compose file to set up a default configuration for Surikator, from [here](#).

## Contributing
We'd love working with you! You can help make this programme better in many ways:
report issues, help us reproduce issues, fix bugs, add functionality, or even
create your own Beat.

Please start by reading our [CONTRIBUTING](CONTRIBUTING.md) file.

Crédits
--
Send ideas and questions to @Treeptik request features and report bugs using the GitHub Issue Tracker.

To find all Treeptik's github project [here](https://github.com/Treeptik).

Developed by [Treeptik](http://treeptik.fr/)'s team @2017.


