pipeline {
    agent any
    tools {
        go 'go1.20'
    }

    environment {
        // dockerhubId is the credential name created in Jenkins UI
        registryCredentials = 'dockerhubId'
    }
    parameters {
        string(name : 'RECIPIENTS', defaultValue: 'uhabiba005@gmail.com', description: 'Email address of recipient')
    }
    stages {
        stage('Pre-Build') {
            steps {
                echo "Installing dependencies to run go code"
                sh 'go version'

                dir('microservice1') {
                    sh 'go get -u golang.org/x/lint/golint'
                }
                dir('microservice2') {
                    sh 'go get -u golang.org/x/lint/golint'
                }

                echo "Checking docker version"
                sh 'docker version'
            }
        }
        stage('Build') {
            steps {
                echo "Compiling and building"
                dir('microservice1') {
                    sh 'go build .'
                }
                dir('microservice2') {
                    sh 'go build .'
                }
            }
        }
        stage('Test') {
            steps {
                echo "Unit Testing, Vetting, Linting, Formatting for microservice1"
                dir('microservice1') {
                    sh 'go test -cover .'
                    sh 'go vet .'
                    sh 'golint .'
                    sh 'gofmt -s -w .'
                }

                echo "Unit Testing, Vetting, Linting, Formatting for microservice2"
                dir('microservice2') {
                    sh 'go test -cover .'
                    sh 'go vet .'
                    sh 'golint .'
                    sh 'gofmt -s -w .'
                }
            }
        }
        stage('Create and Push Docker Images') {
            steps {
                dir('microservice1') {
                    script{
                        docker.withRegistry('', registryCredentials) {
                            def microservice1Image = docker.build("umehabiba04/microservice1:${env.BUILD_ID}")
                            microservice1Image.push()
                        }
                    }
                }
                dir('microservice2') {
                    script{
                        docker.withRegistry('', registryCredentials) {
                            def microservice2Image = docker.build("umehabiba04/microservice2:${env.BUILD_ID}")
                            microservice2Image.push()
                        }
                    }
                }
            }
        }
    }
    post{
        always {
            script {
                String emailSubject = "Jenkins Build ${currentBuild.currentResult}: ${env.Job_NAME}"
                String emailBody = "Jenkin Build ${currentBuild.result}: ${currentBuild.fullDisplayName}.\nChanges: ${currentBuild.changeSets}.\nBuild URL: ${env.BUILD_URL}"
                if (params.RECIPIENTS) {
                    mail to: "${params.RECIPIENTS}",
                    subject: "${emailSubject}",
                    body: "${emailBody}"
                }
            }
        }
    }
}
