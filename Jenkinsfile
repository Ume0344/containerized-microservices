pipeline {
    agent any
    tools {
        go 'go1.20'
    }

    parameters {
        string(name : 'RECIPIENTS', defaultValue: 'uhabiba005@gmail.com', description: 'Email address of recipient')
    }
    stages {
        stage('Pre-Build') {
            steps {
                echo "Installing dependencies to run go code"
                sh 'go version'
                sh 'go get -u golang.org/x/lint/golint'
            }

        }
        stage('build') {
            steps {
                echo "Compiling and building"
                sh 'go build'
            } 
        }
        stage('Test') {
            steps {
                echo "Linting"
                sh 'golint ./...'

                echo "Formatting"
                sh 'gofmt -s -w .'

                echo "Vetting"
                sh 'go vet ./...'

                echo "Unit Testing"
                sh 'go test -cover ./...'
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
