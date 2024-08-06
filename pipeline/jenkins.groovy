pipeline {
    agent any
    environment {
        REPO="https://github.com/jkasyan/kbot"
        BRANCH="develop"
    }

    parameters {
        choice(name: 'OS', choices: ['linux', 'darwin', 'windows', 'all'], description: 'Pick OS')
        choice(name: 'ARCH', choices: ['amd64', 'x86_64', 'i686', 'armv7l'], description: 'Pick ARCH')
    }

    stages {
        stage('clone') {
            steps {
                echo "Clone"
                git branch: "${BRANCH}", url: "${REPO}"
            }
        }

        stage('test') {
            steps {
                echo "Tests"
                sh "make test"
            }
        }

        stage('build') {
            steps {
                echo "Build"
                sh "make build"
            }
        }

        stage('image') {
            steps {
                script {
                    echo "Create image"
                    sh "make image"
                }
            }
        }

        stage('push') {
            steps {
                script {
                    docker.withRegistry("", "dockerhub") {
                        sh "make push"
                    }
                }
            }
        }
    }
}