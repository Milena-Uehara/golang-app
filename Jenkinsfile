pipeline {
    agent {
        docker { image 'docker:dind' }  // Use a Docker container with Node.js
    }

    stages {
        stage('Checkout') {
            steps {
                checkout scm
            }
        }

        stage('Install Dependencies') {
            steps {
                sh 'docker --version'  // Install Node.js dependencies inside the Docker container
                sh 'git --version'
            }
        }
    }

    post {
        always {
            echo 'Cleaning up Docker container'
        }
    }
}
