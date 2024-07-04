pipeline {
    agent any

    stages {
        stage('Build') {
            steps {
                // Build GoLang application
                sh 'go build -o myapp .'
            }
        }
        stage('Build Docker Image') {
            environment {
                DOCKER_IMAGE = 'golang/golang:latest'
            }
            steps {
                // Build Docker image
                script {
                    docker.build DOCKER_IMAGE
                }
            }
        }
        stage('Push Docker Image') {
            steps {
                // Push Docker image to registry
                script {
                    docker.withRegistry('https://hub.docker.com', 'DOCKER_CREDITONAL_ID') {
                        dockerImage.push()
                    }
                }
            }
        }
    }
}

