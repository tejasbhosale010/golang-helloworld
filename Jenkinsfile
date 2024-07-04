{\rtf1\ansi\ansicpg1252\cocoartf2761
\cocoatextscaling0\cocoaplatform0{\fonttbl\f0\fswiss\fcharset0 Helvetica;}
{\colortbl;\red255\green255\blue255;}
{\*\expandedcolortbl;;}
\paperw11900\paperh16840\margl1440\margr1440\vieww11520\viewh8400\viewkind0
\pard\tx720\tx1440\tx2160\tx2880\tx3600\tx4320\tx5040\tx5760\tx6480\tx7200\tx7920\tx8640\pardirnatural\partightenfactor0

\f0\fs24 \cf0 pipeline \{\
    agent any\
\
    environment \{\
        SONARQUBE_SCANNER_HOME = tool name: 'SonarQubeScanner', type: 'hudson.plugins.sonar.SonarRunnerInstallation'\
        SONAR_HOST_URL = 'http://your-sonarqube-server'\
        SONAR_AUTH_TOKEN = credentials('sonarqube-token')\
        DOCKER_IMAGE = 'your-docker-image'\
    \}\
\
    stages \{\
        stage('Checkout') \{\
            steps \{\
                git 'https://your-repo-url.git'\
            \}\
        \}\
\
        stage('Build') \{\
            steps \{\
                sh 'go build -o myapp'\
            \}\
        \}\
\
        stage('SonarQube Analysis') \{\
            steps \{\
                withSonarQubeEnv('SonarQube') \{\
                    sh """\
                    $\{env.SONARQUBE_SCANNER_HOME\}/bin/sonar-scanner \\\
                    -Dsonar.projectKey=myapp \\\
                    -Dsonar.sources=. \\\
                    -Dsonar.host.url=$\{env.SONAR_HOST_URL\} \\\
                    -Dsonar.login=$\{env.SONAR_AUTH_TOKEN\}\
                    """\
                \}\
            \}\
        \}\
\
        stage('Run Tests') \{\
            steps \{\
                sh 'go test -coverprofile=coverage.out ./...'\
            \}\
        \}\
\
        stage('Docker Build') \{\
            steps \{\
                script \{\
                    docker.build("$\{env.DOCKER_IMAGE\}")\
                \}\
            \}\
        \}\
\
        stage('Deploy with Ansible') \{\
            steps \{\
                ansiblePlaybook credentialsId: 'ansible-vault-password', inventory: 'path/to/your/inventory', playbook: 'path/to/your/playbook.yml'\
            \}\
        \}\
    \}\
\
    post \{\
        always \{\
            cleanWs()\
        \}\
    \}\
\}\
}