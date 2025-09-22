pipeline {
    agent any

    environment {
        // Define the image name (you can modify it as per your requirements)
        IMAGE_NAME = 'go-web-store'
        REPO_NAME = 'uehara96'
        DOCKER_IMAGE = "${REPO_NAME}/${IMAGE_NAME}:latest"
        EC2_IP = '3.92.135.71'
    }

    options {
        // Set a timeout for the pipeline or specific stages
        timeout(time: 1, unit: 'HOURS')
        // Discard old builds to free up space
        buildDiscarder(logRotator(numToKeepStr: '10', daysToKeepStr: '7'))
    }

    parameters {
        string(name: 'BRANCH_NAME', defaultValue: 'main', description: 'Branch to build and deploy')
        booleanParam(name: 'DEPLOY', defaultValue: true, description: 'Deploy the application after building')
        choice(name: 'ENVIRONMENT', choices: ['development', 'staging', 'production'], description: 'Select deployment environment')
    }

    triggers {
        // Trigger the pipeline on commits to the `main` branch
        pollSCM('H/5 * * * *')
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
                    // Use withCredentials to access the stored Docker Hub credentials
                    withCredentials([usernamePassword(credentialsId: 'docker-login', usernameVariable: 'DOCKER_USERNAME', passwordVariable: 'DOCKER_PASSWORD')]) {
                        // Log in to Docker Hub
                        sh "docker login -u ${DOCKER_USERNAME} -p '${DOCKER_PASSWORD}'"
                        sh "docker push ${DOCKER_IMAGE}"
                    }
                }
            }
        }

        stage('Deploy') {
            when {
                expression { return params.DEPLOY }
            }
            steps {
                script {
                    echo "Deploying with Docker Compose..."
                    sshagent(credentials: ['ec2']) {
                        // Upload files once to reduce redundant SCP commands
                        sh """
                        scp -o StrictHostKeyChecking=no docker-compose.yaml ubuntu@${EC2_IP}:/home/ubuntu
                        ssh -o StrictHostKeyChecking=no ubuntu@${EC2_IP} "docker compose -f /home/ubuntu/docker-compose.yaml down"
                        ssh -o StrictHostKeyChecking=no ubuntu@${EC2_IP} "docker compose -f /home/ubuntu/docker-compose.yaml up -d"
                        """
                    }
                }
            }
        }

        // Parallel stages example
        stage('Test and Deploy') {
            parallel {
                stage('Unit Tests') {
                    steps {
                        script {
                            echo "Running Unit Tests..."
                            // Replace with actual test commands
                            //sh 'npm test'
                        }
                    }
                }
                stage('Integration Tests') {
                    steps {
                        script {
                            echo "Running Integration Tests..."
                            // Replace with actual integration test commands
                            //sh 'npm run integration-tests'
                        }
                    }
                }
            }
        }

        stage('Input') {
            steps {
                script {
                    def userInput = input message: 'Should we proceed with deployment?', parameters: [booleanParam(defaultValue: true, description: 'Proceed?', name: 'Proceed')]
                    if (!userInput) {
                        error 'Deployment aborted by user.'
                    }
                }
            }
        }
    }

    post {
        always {
            echo 'Cleaning up after build...'
            // Cleanup steps, like removing temporary files or Docker images
            //sh 'docker system prune -f'
        }
        success {
            echo 'Pipeline executed successfully!'
        }
        failure {
            echo 'Pipeline failed, check the logs for errors.'
        }
    }
}
