pipeline {
  agent any

  environment {
    IMAGE = "tpc90/go-ci-app"
  }

  options { timestamps() }

  stages {
    stage('Checkout') {
      steps { checkout scm }
    }

    stage('Build') {
      steps {
        sh 'go mod tidy'
        sh 'go build -v -o go-ci-app ./cmd'
      }
    }

    stage('Test') {
      steps {
        sh 'go test ./... || true'
      }
    }

    stage('Docker Build') {
      steps {
        script {
          def IMAGE_TAG = sh(returnStdout: true, script: 'git rev-parse --short HEAD').trim()
          env.IMAGE_TAG = IMAGE_TAG
          sh "docker build -t ${IMAGE}:${IMAGE_TAG} -t ${IMAGE}:latest ."
        }
      }
    }

    stage('Push to Docker Hub') {
      steps {
        withCredentials([usernamePassword(credentialsId: 'dockerhub-creds',
                                          usernameVariable: 'DOCKER_USER',
                                          passwordVariable: 'DOCKER_PASS')]) {
          sh 'echo "$DOCKER_PASS" | docker login -u "$DOCKER_USER" --password-stdin'
          sh "docker push ${IMAGE}:${env.IMAGE_TAG}"
          sh "docker push ${IMAGE}:latest"
        }
      }
    }

    stage('Deploy') {
      steps {
        sh '''
          docker stop go-ci-app || true
          docker rm go-ci-app || true
          docker run -d --restart unless-stopped --name go-ci-app -p 9090:8080 ${IMAGE}:latest
        '''
      }
    }
  }

  post {
    always { archiveArtifacts artifacts: 'go-ci-app', allowEmptyArchive: true }
  }
}
