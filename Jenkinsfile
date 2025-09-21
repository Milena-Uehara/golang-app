pipeline {
    agent {
        docker { image 'node:14' }  // Use a Docker container with Node.js
    }

    stages {
        stage('Checkout') {
            steps {
                checkout scm
            }
        }

        stage('Install Dependencies') {
            steps {
                sh 'npm install'  // Install Node.js dependencies inside the Docker container
            }
        }

        stage('Run Tests') {
            steps {
                sh 'npm test'  // Run tests inside the Docker container
            }
        }

        stage('Build') {
            steps {
                sh 'npm run build'  // Build the app inside the Docker container
            }
        }
    }

    post {
        always {
            echo 'Cleaning up Docker container'
        }
    }
}
