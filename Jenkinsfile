pipeline {
    
    agent {
        docker {
            label 'go'
            image 'docker:29-cli'
            args '-v /var/run/docker.sock:/var/run/docker.sock -u root'
        }    
    }

    environment {
        REGISTRY = "192.168.88.62:8082"
        REPOSITORY = "docker-hosted"
        IMAGE_NAME = "myapp"
        IMAGE_TAG = "${BUILD_NUMBER}"
    }
     
    stages {
        
        stage('Clone') {
            steps {
               git branch: 'main', url: 'https://github.com/sobhan-bagheri/golang-openstack.git'
            }
        }
        
        stage('Build') {
            steps {
               sh 'docker build -t ${IMAGE_NAME}:${IMAGE_TAG} .'
            }
        }

        stage('Tag') {
            steps {
                sh 'docker tag ${IMAGE_NAME}:${IMAGE_TAG} ${REGISTRY}/${REPOSITORY}/${IMAGE_NAME}:${IMAGE_TAG}'
            }
        }

        stage('Login') {
            steps {
                withCredentials([usernamePassword(
                    credentialsId: 'nexus-docker',
                    usernameVariable: 'USERNAME',
                    passwordVariable: 'PASSWORD'
                )]) {
                    sh 'echo $PASSWORD | docker login ${REGISTRY} -u $USERNAME --password-stdin'
                }
            }
        }

        stage('Push') {
            steps {
                sh 'docker push ${REGISTRY}/${REPOSITORY}/${IMAGE_NAME}:${IMAGE_TAG}'
            }
        }
        stage('Deploy to Production') {
    steps {
        sshagent(['production-ssh']) {
            sh '''
            ssh -o StrictHostKeyChecking=no ubuntu@192.168.88.68 "
                docker pull 192.168.88.62:8082/docker-hosted/myapp:${BUILD_NUMBER} &&
                docker stop myapp || true &&
                docker rm myapp || true &&
                docker run -d --name myapp -p 8084:8084 192.168.88.62:8082/docker-hosted/myapp:${BUILD_NUMBER}
            "
            '''
        }
    }
}

    }
}
