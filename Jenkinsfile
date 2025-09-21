pipeline {
    agent {
        docker {
            image 'docker:19.03.12'
            args '-v /var/run/docker.sock:/var/run/docker.sock'  // Mount Docker socket
        }
    }

    stages {
        stage('Docker Build') {
            steps {
                script {
                    sh 'docker build -t my-app .'
                }
            }
        }
    }
}
