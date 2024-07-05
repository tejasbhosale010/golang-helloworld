pipeline {
    agent any

tools {
        // Define the GoLang tool using the configured name in Jenkins Global Tool Configuration
        go 'go1.22.5'
    }
    stages {
        stage('Checkout') {
            steps {
         
                git branch: 'main', url: 'https://github.com/tejasbhosale010/golang-helloworld.git'
            }
        }
        
    stage('Setup Go Environment') {
            steps {
                script {
                    def goHome = tool name: 'go1.22.5', type: 'go'
                    env.PATH = "${goHome}/bin:${env.PATH}"
                }
            }
        }

        stage('Build') {
            steps {
                //sh 'go mod init golang-app'
                sh "go build -o my-go-app -buildvcs=false"
            }
        }
    }
}
