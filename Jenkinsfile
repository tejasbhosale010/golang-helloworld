pipeline {
    agent any

    environment {
        SONARQUBE_SCANNER_HOME = tool name: 'SonarQubeScanner', type: 'hudson.plugins.sonar.SonarRunnerInstallation'
        SONAR_HOST_URL = 'http://your-sonarqube-server'
        SONAR_AUTH_TOKEN = credentials('sonarqube-token')
        DOCKER_IMAGE = 'your-docker-image'
    }

    stages {
        stage('Checkout') {
            steps {
                git 'https://github.com/tejasbhosale010/golang-helloworld.git'
            }
        }

        stage('Build') {
            steps {
                sh 'go build -o myapp'
            }
        }

        stage('SonarQube Analysis') {
            steps {
                withSonarQubeEnv('SonarQube') {
                    sh """
                    /usr/local/sonar \
                    -Dsonar.organization=demo-tejas \
                    -Dsonar.projectKey=demo-tejas_hellogo-lang \
                    -Dsonar.sources=. \
                    -Dsonar.host.url=https://sonarcloud.io
                    """
                }
            }
        }

        stage('Run Tests') {
            steps {
                sh 'go test -coverprofile=coverage.out ./...'
            }
        }

        stage('Docker Build') {
            steps {
                script {
                    docker.build("${env.DOCKER_IMAGE}")
                }
            }
        }

        stage('Deploy with Ansible') {
            steps {
                ansiblePlaybook credentialsId: 'ansible-vault-password', inventory: 'path/to/your/inventory', playbook: 'path/to/your/playbook.yml'
            }
        }
    }

    post {
        always {
            cleanWs()
        }
    }
}

