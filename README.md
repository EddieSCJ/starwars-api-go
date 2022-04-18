# <p align="center"> :space_invader: Go Star Wars API :space_invader:</p>

#### <p align="center">  Project Tooling </p>
<div align="center"> 
    <a href="https://app.snyk.io/org/eddiescj/projects" target"_blank">:wolf: Snyk </a>
    <a href="https://sonarcloud.io/project/overview?id=EddieSCJ_starwars-api-go" target"_blank">:detective: SonarCloud </a>
    <a href="https://app.codecov.io/gh/EddieSCJ/starwars-api-go/" target"_blank">:open_umbrella: CodeCov </a>
    <a href="https://golangci-lint.run/" target="_blank">:white_check_mark: GolangCI-Lint </a>

</div>

#### <p align="center">  Description </p>
This api is a simple wrapper for the [Star Wars API](https://swapi.dev/) where you can get information about the characters,
planets, starships, vehicles, species, films, and more with a few extra features where you can handle this data however you 
need.

Please, read the content below to know how to use this api and if is there any doubt, please, contact me.

#### <p align="center">  Using Cloud Tools </p>

* Snyk
    * Just click in the link above and search for starwars-api-go, so you will be able to see the security problems.
* SonarCloud
    * Clicking the link above you will be redirected to the quality analysis of this project.
* CodeCov
    * You can click the link above and see the code coverage details by commit or any type of data or just see the summary in your PR.
* GolangCI-Lint
    * You are able to see it in the actions tab, most specifically in the Lint Check job.

#### <p align="center">  Testing </p>

There are two types of tests which are used in this project, unit and integrated tests.
To run the unit tests you just need to type: `go test ./...` in your terminal and wait to see magic happens.

However, our integrated tests are running with DockerTest to reproduce more realistic scenarios, so, we need some extra commands, nothing scaring, firstly, make sure you have [Docker](https://docs.docker.com/engine/install/) installed in your machine and type `go test -p 1 -tags integration ./...` in your terminal.

If you want to see the benchmarks, just run `go test ./... -bench .`

:grey_question: What exactly is this `-p 1` :grey_question:

The Go Programming language run the test packages in parallel, so, it's possible our docker containers being not ready before some integrations test. To make sure it will not happen, we limit the workers to only one.
