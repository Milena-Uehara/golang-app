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
                    echo "Pushing Image to DockerHub..."
                    withCredentials([usernamePassword(credentialsId: 'docker-login', passwordVariable: 'PASS', usernameVariable: 'USER')]) {
                        sh "echo $PASS | docker login -u $USER --password-stdin"
                        sh "docker push ${DOCKER_IMAGE}"
                    }
                }
            }
        }
    }
}
