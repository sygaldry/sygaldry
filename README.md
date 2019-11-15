# sygaldry

## Dependencies
```bash
go mod download
```

## Building Sygaldry
```bash
go test ./...
```

## Using the CLI
```bash
sygaldry build -f runes.yaml
```

## YAML

### Formatting
```yaml
stages:
  build:
  - definition: MavenRune
    mavenome: ~/.m2
    stages: clean install
    version: 3.5.0-jdk-8-slim
  - definition: DockerRune
    version: stable
    command: build
    args: "-t docker.io/sygaldrydemos/springboot:v0.0.2 ."
  publish:
  - definition: DockerRune
    version: stable
    command: push
    dockerConfigPath: ~/.docker/config-sygaldry.json
    args: "docker.io/sygaldrydemos/springboot:v0.0.2"
  deploy:
  - definition: KubectlRune
    command: apply
    version: stable
    args: -f k8s-configs/demo.yaml
definitions:
- "https://raw.githubusercontent.com/sygaldry/sygaldry-runes/master/rune-definitions.yaml"
```

### Anatomy
The following is a valid Sygaldry YAML file:
```yaml
stages:
  creamcheese:
  - definition: DockerRune
    version: stable
    command: build
    args: "-t docker.io/sygaldrydemos/springboot:v0.0.2 ."
definitions:
- "https://raw.githubusercontent.com/sygaldry/sygaldry-runes/master/rune-definitions.yaml"
```
Assuming the name of this file is `myDockerBuild.yaml` we can run the following command without error:
```bash
sygaldry creamcheese -f myDockerBuild.yaml
```

## FAQ

### What is Sygaldry?
Sygaldry is the simplest way to run your [CI/CD](https://www.redhat.com/en/topics/devops/what-is-ci-cd). It takes a simple [YAML](https://en.wikipedia.org/wiki/YAML) file and uses it with underlying [runes](#what-is-a-rune) in order to build out a pipeline for your code. It is written in [Go](https://golang.org/) and ran in command line.

### What is a rune?
Runes are pre-built modules that allow users of [Sygaldry](#what-is-sygaldry) to work with specific services easily.
