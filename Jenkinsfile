pipeline {
    agent any

    environment {
        // Define any global environment variables here if needed
        IMAGE_NAME = "ayee-portal-backend"
    }

    stages {
        stage('Checkout') {
            steps {
                checkout scm
            }
        }

        stage('Build Docker Image') {
            steps {
                script {
                    sh 'docker build -t ${IMAGE_NAME} .'
                }
            }
        }

        stage('Deploy') {
            steps {
                script {
                    // Stop and remove existing container if running (optional, compose handles this usually but good for cleanup)
                    // sh 'docker-compose down || true'
                    
                    // Deploy using docker-compose
                    // This assumes Jenkins has the necessary env vars injected via credentials or global config
                    sh 'docker-compose up -d --build'
                }
            }
        }
    }

    post {
        always {
            cleanWs()
        }
    }
}
