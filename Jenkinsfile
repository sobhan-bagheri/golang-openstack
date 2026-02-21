pipeline {
    agent {

        docker {
            label 'go'
            image 'golang:1.22-alpine'
        }
    }
    environment {
        HOME = "${WORKSPACE}"
        GOCACHE = "${WORKSPACE}/.cache/go-build"
    }

    stages {
        stage('pull') {
            steps {
                git branch: 'main', url: 'https://github.com/sobhan-bagheri/golang-openstack.git/'
            }
        }
        stage('build') {
            steps {
                sh 'go version'
                sh 'go build -o app app.go'
            }
        }
        
        stage('deploy') {
            steps {
                sh './app'
            }
        }
    }
}
