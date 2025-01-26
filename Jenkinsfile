pipeline {
    agent any
    environment {
        REPO_URL = 'https://github.com/SchoolManagementSyestem/Backend.git'
        BRANCH = 'main'
    }
    stages {
        stage('Checkout Code') {
            steps {
                // checkout([$class: 'GitSCM',
                //     // branches: [[name: '*/${BRANCH}']],
                //     // userRemoteConfigs: [[url: REPO_URL, credentialsId: 'github_credintial']]
                // ])
            }
        }
        stage('Build Docker Image') {
            steps {
                echo 'Building Docker image...'
                // Add build steps
            }
        }
        stage('Push Docker Image') {
            steps {
                echo 'Pushing Docker image...'
                // Add push steps
            }
        }
        stage('Deploy to Kubernetes') {
            steps {
                echo 'Deploying to Kubernetes...'
                // Add deployment steps
            }
        }
    }
    post {
        always {
            echo 'Pipeline execution completed!'
        }
        failure {
            echo 'Deployment failed. Check logs!'
        }
    }
}
