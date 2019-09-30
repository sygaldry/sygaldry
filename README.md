# sygaldry

## Using the CLI
```bash
sygaldry build -f mySygaldryYamlFile.yaml
```

## YAML formatting
```yaml
build:
  SpringbootMavenBuild:
    mavenHome: ~/.m2
  DockerBuild:
    dockerfilePath: docker/Dockerfile
    name: hodor/myKewlApp
publish:
  SpringbootMavenDeploy:
    nexusUrl: https://www.nexus.com
    nexusUsername: hodor
    nexusPassword: h0dor
  DockerPublish:
    registryUrl: https://docker.io
    dockerUsername: hodor
    dockerPassword: h0dor
```

## FAQ

### What is Sygaldry?
Sygaldry is the simplest way to run your [CI/CD](https://www.redhat.com/en/topics/devops/what-is-ci-cd). It takes a simple [YAML](https://en.wikipedia.org/wiki/YAML) file and uses it with underlying [runes](#what-is-a-rune) in order to build out a pipeline for your code. It is written in [Go](https://golang.org/) and ran in command line.

### What is a rune?
Runes are pre-built modules that allow users of [Sygaldry](#what-is-sygaldry) to work with specific services easily.