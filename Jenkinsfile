pipeline {
    agent any

    environment {
        // Define the image name (you can modify it as per your requirements)
        IMAGE_NAME = 'go-web-store'
        REPO_NAME = 'uehara96'
        DOCKER_IMAGE = "${REPO_NAME}/${IMAGE_NAME}:latest"
    }

    stages {
        stage('Checkout') {
            steps {
                // Checkout the source code from the repository
                checkout scm
            }
        }

        stage('Build') {
            steps {
                script {
                    // Build the Docker image
                    sh '''
                    set -e
                    docker build -t ${DOCKER_IMAGE} -f docker/Dockerfile .
                    '''
                }
            }
        }

        stage('Push Image') {
            steps {
                script {
                    // Push the image to a Docker registry
                    sh '''
                    echo $DOCKER_PASSWORD | docker login -u $DOCKER_USERNAME --password-stdin
                    docker push ${DOCKER_IMAGE}
                    '''
                }
            }
        }

        stage('Deploy') {
            steps {
                // Add your deployment steps here
                echo 'Deploying the application...'
            }
        }
    }

    post {
        always {
            // Clean up Docker images after the pipeline finishes
            sh 'docker system prune -f'
        }
    }
}
