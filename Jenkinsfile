pipeline {
    agent any  // This runs the pipeline on any available agent (or specify a label for a particular node)

    environment {
        // Define the image name (you can modify it as per your requirements)
        IMAGE_NAME = 'go-web-store'
        REPO_NAME = 'uehara96'
        DOCKER_IMAGE = "${REPO_NAME}/${IMAGE_NAME}:latest"
    }

    stages {
        stage('Install Docker') {
            steps {
                script {
                    // Check if Docker is installed
                    def dockerInstalled = sh(script: 'which docker', returnStatus: true)

                    // If Docker is not installed, install it
                    if (dockerInstalled != 0) {
                        echo 'Docker is not installed. Installing...'

                        // Install Docker (for Ubuntu/Debian based systems)
                        sh '''
                        sudo apt-get update
                        sudo apt-get install -y apt-transport-https ca-certificates curl software-properties-common
                        curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
                        sudo add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable"
                        sudo apt-get update
                        sudo apt-get install -y docker-ce
                        '''
                    } else {
                        echo 'Docker is already installed.'
                    }
                }
            }
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
